# drone-telegram

[![Docker Stars](https://img.shields.io/docker/stars/jessynt/drone-telegram.svg)](https://hub.docker.com/r/jessynt/drone-telegram/)
[![Docker Pulls](https://img.shields.io/docker/pulls/jessynt/drone-telegram.svg)](https://hub.docker.com/r/jessynt/drone-telegram/)
[![Go Doc](https://godoc.org/github.com/jessynt/drone-telegram?status.svg)](http://godoc.org/github.com/jessynt/drone-telegram)
[![Go Report](https://goreportcard.com/badge/github.com/jessynt/drone-telegram)](https://goreportcard.com/report/github.com/jessynt/drone-telegram)
![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)

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

Example configuration using secrets:

```
pipeline:
  telegram:
    image: jessynt/drone-telegram
    secrets: [telegram_access_token, telegram_chat_id]
```