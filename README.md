# GoSlackBot

## 概要
goでslackBotを作る

## 準備
- パッケージインストール
```
$ go get -u github.com/nlopes/slack
```
- password.goのさくせい
```
package main

func getApiToken()(string) {
    return "YOUR-API-TOKEN"
}
```

- 実行
```
$ go run *.go
```


ref
https://kuroeveryday.blogspot.jp/2017/10/tutorial-slack-rtm-api-with-golang.html