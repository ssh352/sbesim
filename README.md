# sbesim

sbesim is a simulation client/broker using SBE protocol

## Getting Started
```
$ make help
all -- install dependencies and build binary
build -- build binary
install -- install dependencies
clean -- clean up dependencies and binaries
fmt -- auto format file
mockgen -- generate mock package
test -- run unit test
run -- start service
```
### Prerequisites
1. Golang 1.8
2. Dep

## Running the tests

sbesim uses golang test framework and simply run below
```
make test
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* the schema code is generated via cme tool