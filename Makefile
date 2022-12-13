VERSION=$(shell git describe --tags --abbrev=0)

all: compile checksum clean

test:
	/usr/bin/go test -timeout 30s github.com/vaerh/terraform-provider-routeros/routeros

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