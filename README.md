# Go Lang Learning Echo Framework

## Description
This is a simple project to learn the Echo Framework in Go Lang.

## Installation
1. Clone the repository
2. Run the following command to install the dependencies
```bash
go mod tidy
```

## Commands

makefile commands:
- `make tidy` - Run go mod tidy
- `make watch` - Run the server in watch mode
- `make build` - Build the project
- `make compose-runs` - Run the project using docker-compose
- `make compose-down` - Stop the project using docker-compose
- `make compose-logs` - Show the logs of the project using docker-compose
- `make compose-restart` - Restart the project using docker-compose
- `help` - Show the help message