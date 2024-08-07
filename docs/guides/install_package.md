# Install package

The original example package installation is available in the [Schwitzd](https://github.com/Schwitzd/IaC-HomeRouter/blob/main/container_backend.tf) repository.

## Example
```terraform
resource "null_resource" "download_container_npk" {
  provisioner "local-exec" {
    command = <<EOT
      chmod +x ./helper/download_routeros_packages.sh
      ./helper/download_routeros_packages.sh ${local.system_architecture} "${local.system_version}" "container"
    EOT
  }
}

resource "null_resource" "upload_container_npk" {
  provisioner "local-exec" {
    command = "scp -i ${local.router_ssh_key} \"/tmp/routeros_packages/${local.container_npk_name}\" ${local.router_user}@${var.router_ip}:/${local.container_npk_name}"
  }

  depends_on = [ null_resource.download_container_npk ]
}

resource "null_resource" "install_container_npk" {
  provisioner "local-exec" {
      command = <<EOT
        ssh -i ${local.router_ssh_key} ${local.router_user}@${var.router_ip} '/system reboot'; sleep 3
        until ssh -i ${local.router_ssh_key} -o ConnectTimeout=2 ${local.router_user}@${var.router_ip} ':put True' 2> /dev/null
        do
          echo "Waiting for router to reboot and become available..."
          sleep 10
        done
      EOT
  }

  depends_on = [ null_resource.upload_container_npk ]
}
```

```shell
#!/bin/bash

# Input parameters
ARCHITECTURE_NAME=$1
VERSION=$2
PACKAGE_NAME_PREFIX=$3

# Define the base URL and package format
BASE_URL="https://download.mikrotik.com/routeros"
PACKAGE_FORMAT="all_packages-${ARCHITECTURE_NAME}-${VERSION}.zip"

# Construct the full URL
FULL_URL="${BASE_URL}/${VERSION}/${PACKAGE_FORMAT}"

# Define the download and extraction paths
DOWNLOAD_PATH="/tmp/${PACKAGE_FORMAT}"
EXTRACT_PATH="/tmp/routeros_packages"

# Download the package
echo "Downloading package from: ${FULL_URL}"
curl -o "${DOWNLOAD_PATH}" "${FULL_URL}"

# Verify download
if [ $? -ne 0 ]; then
  echo "Failed to download the package."
  exit 1
fi

# Create the extraction directory
mkdir -p "${EXTRACT_PATH}"

# List all files in the ZIP archive and filter by the PACKAGE_NAME_PREFIX
echo "Finding package that starts with: ${PACKAGE_NAME_PREFIX}"
MATCHED_FILES=$(unzip -l "${DOWNLOAD_PATH}" | awk '{print $4}' | grep "^${PACKAGE_NAME_PREFIX}")

# Check if any files were matched
if [ -z "$MATCHED_FILES" ]; then
  echo "No files found starting with '${PACKAGE_NAME_PREFIX}'."
  exit 1
fi

# Extract matched files
for FILE in $MATCHED_FILES; do
  echo "Extracting: ${FILE}"
  unzip -jo "${DOWNLOAD_PATH}" "${FILE}" -d "${EXTRACT_PATH}"

  if [ $? -ne 0 ]; then
    echo "Failed to extract: ${FILE}"
    exit 1
  fi
done

echo "Extraction completed successfully in: ${EXTRACT_PATH}"
```