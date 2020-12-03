![goreleaser](https://github.com/arisetransfer/arise/workflows/goreleaser/badge.svg)
[![GoDoc](https://godoc.org/github.com/arisetransfer/arise?status.svg)](https://godoc.org/github.com/arisetranfer/arise)

# Arise File Transfer
Transfer file between two devices using gRPC streams.

## Installation
Using Go
```bash
go get -u github.com/arisetransfer/arise
```
Or Download the binary from releases and add to your path

## Setting up the config file
Create the config file
```bash
mkdir -p $HOME/.arise/ && touch $HOME/.arise/config.toml
```
and add IP and Port of the Server
Example `config.toml`
```toml
# Configuration file for arise relay and port

ip = "127.0.0.1"

port = "6969"
```

## Usage

### To send a File
```bash
arise send filename
```

### To Receive a File
```bash
arise receive unique_code
```

## Setting Up The Server

### Using Docker
```bash
docker pull ghcr.io/arisetransfer/arise:latest
docker run -d -p 6969:6969 ghcr.io/arisetransfer/arise:latest
```

### Using CLI
```bash
arise relay
```
This will listen on port 6969

### Help
```bash
NAME:
   arise - Transfer file between two devices

USAGE:
   arise [global options] command [command options] [arguments...]

COMMANDS:
   send     Send the file over relay
   relay    Start an arise relay on port default(6969)
   receive  Receive file using the unique code
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## License
MIT
