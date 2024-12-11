VERSION=$(shell git describe --tags --abbrev=0)

.PHONY: docs debug

all: docs tfformat compile checksum clean

test:
	go test -timeout 30s github.com/terraform-routeros/terraform-provider-routeros

docs:
	go generate
	# !!! GNU Sed
	find docs -type f -exec sed -i -E '/^.*__[[:alpha:]_]+__/d' {} \;

tfformat:
	terraform fmt -recursive examples/

debug:
	go build -gcflags="all=-N -l" -o terraform-provider-routeros_${VERSION} main.go

compile:
	mkdir -p pkg
	echo "Removing previously built packages"
	rm -rf pkg/*
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_arm.zip terraform-provider-routeros_${VERSION}
	
	GOOS=linux GOARCH=arm64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_arm64.zip terraform-provider-routeros_${VERSION}

	GOOS=linux GOARCH=386 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_386.zip terraform-provider-routeros_${VERSION}

	GOOS=linux GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_amd64.zip terraform-provider-routeros_${VERSION}

	GOOS=windows GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_windows_amd64.zip terraform-provider-routeros_${VERSION}

	GOOS=windows GOARCH=386 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_windows_386.zip terraform-provider-routeros_${VERSION}

	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_darwin_amd64.zip terraform-provider-routeros_${VERSION}

	GOOS=darwin GOARCH=arm64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_darwin_arm64.zip terraform-provider-routeros_${VERSION}

checksum:
	cd pkg && sha256sum *.zip > terraform-provider-routeros_${VERSION}_SHA256SUMS

clean:
	rm terraform-provider-routeros_${VERSION}