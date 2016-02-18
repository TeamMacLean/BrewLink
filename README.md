# Brewlink
Link tools installed with brew to a sourceable path

Created for use with the TSL HPC. It is designed to be used in cunjunction with `brew` and `create-testing-wrapper` to make it easy and quick to install tools onto our HPC.

Example usage:

```
brew install go
brewlink go-1.5.3
create-software-testing-wrapper go-1.5.3
```

Installation on our HPC:
```
source go-1.5.3
mkdir ~/go
export GOPATH=~/go
mkdir -p /tsl/software/testing/brewlink/default/
git clone https://github.com/TeamMacLean/BrewLink.git /tsl/software/testing/brewlink/default/x86_64
cd /tsl/software/testing/brewlink/default/x86_64
go get -d
mkdir bin
cp .brewlink.json bin/
cd bin
go build ../brewlink.go
create-software-testing-wrapper brewlink-default
```
