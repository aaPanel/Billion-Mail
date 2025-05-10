# !/bin/sh

# How to use:
# cd /opt/BillionMail
# docker run -d --name p-g-alpine --hostname p-g-alpine --restart=always -v ./core:/opt/core -v ./Dockerfiles/core/repositories:/etc/apk/repositories alpine:3.20 tail -f /dev/null
# docker exec -it p-g-alpine sh
# cd /opt/core
# sh go-build.sh

set -e  # Exit immediately when encountering an error

PLATFORMS=${1:-"all"}
if [[ "$PLATFORMS" != "all" && "$PLATFORMS" != "x86" && "$PLATFORMS" != "arm" ]]; then
  echo -e "Packaging architecture parameters are incorrect, please enter all, x86, arm"
  echo -e "Example: sh go-build.sh all"
  exit 1
fi

# Check and install Go (if not exist or version is lower than 1.22)
if ! command -v go &>/dev/null; then
    echo "Error: Go is not installed, it is being installed automatically..."
    apk update && apk add go file
elif [ "$(go version | awk '{print $3}' | sed 's/go//')" \< "1.22" ]; then
    echo "Error: Go version is lower than 1.22, please upgrade manually."
    exit 1
fi

# Check and install the file (if not exists)
if ! command -v file &>/dev/null; then
    echo "Warning: file not installed, it is being installed automatically..."
    apk update && apk add file
fi

# Verify installation results
echo "Current environment:"
echo "- Go version: $(go version 2>/dev/null || echo 'Not installed')"
echo "- file version: $(file --version 2>/dev/null || echo 'Not installed')"

#go env -w GOPROXY=https://goproxy.io,direct

if [ ! -d "/opt/core" ]; then
    echo "/opt/core not found"
    exit 1
fi
cd /opt/core

if [ ! -f "main.go" ]; then
    echo "main.go not found"
    exit 1
fi

if [[ "$PLATFORMS" == "all" || "$PLATFORMS" == "x86" ]]; then

    # amd64 x86_64
    echo "build start amd64 : x86_64"
    rm -f billionmail-x86_64
    export GOOS=linux
    export GOARCH=amd64
    go build -ldflags="-s -w" -o billionmail-x86_64 main.go
    if [ ! -f "billionmail-x86_64" ]; then
        echo "build x86_64 failed"
        exit 1
    fi

    check_file=$(file billionmail-x86_64 | grep "x86-64,")
    if [ -z "${check_file}" ];then
        echo "billionmail-x86_64 is not an arm64 file, package failed.";
        exit 0;
    fi
fi

if [[ "$PLATFORMS" == "all" || "$PLATFORMS" == "arm" ]]; then
    # arm64 aarch64
    echo "build start arm64 : aarch64"
    rm -f billionmail-aarch64
    export GOOS=linux
    export GOARCH=arm64
    go build -ldflags="-s -w" -o billionmail-aarch64 main.go
    if [ ! -f "billionmail-aarch64" ]; then
        echo "build aarch64 failed"
        exit 1
    fi
    check_file=$(file billionmail-aarch64 | grep "ARM aarch64,")
    if [ -z "${check_file}" ];then
        echo "billionmail-aarch64 is not an arm64 file, package failed.";
        exit 0;
    fi
fi

ls -al billionmail-*