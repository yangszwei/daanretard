# 靠北大安4.0 (daanretard)

![build](https://github.com/yangszwei/daanretard/workflows/Go%20Test/badge.svg)
[![](https://tokei.rs/b1/github/yangszwei/daanretard)](https://github.com/yangszwei/daanretard)

靠北大安4.0投稿系統

## Prerequisite

- go1.15
- mariaDB or any database supported by [https://github.com/go-gorm/mysql](https://github.com/go-gorm/mysql)
- facebook graph api app

## Install

(at project root)

- Create .env file:

```.env
ADDR=
SECRET=
FB_APP_ID=
FB_APP_SECRET=
DB_DSN= 
```

- Run `go run daanretard/cmd/daanretard`
