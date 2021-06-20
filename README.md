# Golang api client kata

## Intro

Gustaw Fit - 20-06-2021 (gustaw.fit@gmail.com)

## Prerequisites

This app assumes you are running it under a LINUX-like system.
We are assuming basic golang infra, makefile and docker were already
installed.

First thing to do is to modify your ~/.bashprofile to include:

```bash
export GOPATH=$HOME/go
export GOBIN=$HOME/go/bin
export PATH=$PATH:$HOME/go/bin
```

This is to ensure you have the right go env variables set.

Please execute after editing:
```bash
source ~/.bash_profile
```

Afterwards please ensure you have installed:
- golang linter
```bash
go get -u golang.org/x/lint/golint
```

## Installation

This library is currently not publicly released.

## Usage

Construct a new api client, then use the various services on the client to access different parts of the API. For example:

```go
client := clientapi.CreateClient()

// Fetch the health status
account, err := client.Health.Check()
```

For more - see [Example](example.go) file.

## FAQ

##### How to run the build and verification process?

By command line (runs in your devbox/laptop/computer):
```bash
make all-locally
```

You can review the specific commands in [Makefile](Makefile) file.

By executing docker-compose:
```bash
docker-compose up
```

## License

See [LICENSE](LICENSE) file.
