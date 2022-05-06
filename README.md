# xk6-telegram

[![Build xk6 telegram](https://github.com/Maksimall89/xk6-telegram/actions/workflows/go.yml/badge.svg)](https://github.com/Maksimall89/xk6-telegram/actions/workflows/go.yml)

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Contents
* [Build](#build)
* [Usage](#usage)
* [Configure](#configure)
    + [Debug mode](#debug-mode)
    + [Create telegram bot](#create-telegram-bot)


## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. [Install](https://github.com/grafana/xk6) `xk6`:

  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  CGO_ENABLED=1
  xk6 build --with github.com/maksimall89/xk6-telegram@latest
  ```

If you use Windows:

```shell
set CGO_ENABLED=1
xk6 build master --with github.com/maksimall89/xk6-telegram@latest
```

## Usage

```javascript
import http from "k6/http";
import telegram from "k6/x/telegram";

const conn = telegram.connect(`${__ENV.TOKEN}`, false);
const chatID = 123456789;

export function setup() {
    telegram.send(conn, chatID, "Starting test");
    telegram.sendReplay(conn, chatID, 1, "Replay 1 message from the chat");
}

export default function () {
    http.get('http://test.k6.io');
}

export function teardown() {
    telegram.send(conn, chatID, "Result <b>test</b> with HTML tag");
    telegram.shareImage(conn, chatID, "http://k6.io/logo.png");
    telegram.uploadImagePath(conn, chatID, "/home/k6/logo.png");
    telegram.getUpdate(conn, chatID);
}

```

Result output:

```shell
$ ./k6 run example.js

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)


running (00m00.8s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.7s/10m0s  1/1 iters, 1 per VU

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=683.33ms min=683.33ms med=683.33ms max=683.33ms p(90)=683.33ms p(95)=683.33ms
     iterations...........: 1   1.217797/s
     
```

## Configure

### Debug mode

if you need, you can debug mode on:

```javascript
import telegram from "k6/x/telegram";

const conn = telegram.connect(`${__ENV.TOKEN}`, true);
```

### Create telegram bot

Before using this extension, you need to create a bot in telegram use [@BotFather](https://t.me/BotFather) after you need set environment variable `TOKEN`.
