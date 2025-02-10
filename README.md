# Golang Project

## Setup

### Initialize a new Go module:
```sh
go mod init <module-name>
```
### Install dependencies:
```sh
go mod tidy
```
## Running Tests with Statement Coverage
```sh
cd <package-name>
```
### Run tests with coverage and generate a coverage report:
```sh
go test -cover -coverprofile="coverage.out"
```
### View the coverage details in the terminal:
```sh
go tool cover -func="coverage.out"
```
### Generate an HTML report for better visualization:
```sh
go tool cover -html="coverage.out" -o coverage.html
```

## Running Tests with Branch Coverage
```sh
cd <package-name>
```
### Installation
#### With go1.17 or later:
```sh
go install github.com/rillig/gobco@latest
```
#### With go1.16:
```sh
go get github.com/rillig/gobco
```
### Run tests with branch coverage
```sh
gobco
```
