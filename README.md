# Brewlink
Link tools installed with brew to a sourceable path

Created for use with the TSL HPC. It is designed to be used in cunjunction with `brew` and `create-testing-wrapper` to make it easy and quick to install tools onto our HPC.

Example usage:

```
brew install go
brewlink go-1.5.3
create-software-testing-wrapper go-1.5.3
```

