# drone-telegram

[![Go Doc](https://godoc.org/github.com/jessynt/drone-telegram?status.svg)](http://godoc.org/github.com/jessynt/drone-telegram)
[![Go Report](https://goreportcard.com/badge/github.com/jessynt/drone-telegram)](https://goreportcard.com/report/github.com/jessynt/drone-telegram)

Drone plugin for sending telegram notifications. 

## Build

Build the binary with the following commands:

```
make all
```

## Usage

```
pipeline:
  telegram:
    image: jessynt/drone-telegram
    access_token: your access token
    chat_id: telegram chat id, you can obtain it from @chatid_echo_bot
```