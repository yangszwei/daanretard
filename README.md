# 靠北大安4.0 (daanretard)

![build](https://github.com/yangszwei/daanretard/workflows/Go%20Test/badge.svg)
[![](https://img.shields.io/badge/status-active%20development-blue)](https://img.shields.io/badge/status-under%20development-blue)
[![](https://tokei.rs/b1/github/yangszwei/daanretard)](https://github.com/yangszwei/daanretard)

靠北大安4.0投稿系統

## Prerequisite

- MySQL database (refer to [https://github.com/go-gorm/mysql](https://github.com/go-gorm/mysql))
- Facebook Graph API App (refer to [https://developers.facebook.com/docs/graph-api/reference/application/](https://developers.facebook.com/docs/graph-api/reference/application/))

## Build

- Prerequisite

```
Yarn
Go 1.15+
Make
```

- Run `make build`
- A file named "daanretard" should be created

> NOTE: You may want to rename daanretard to daanretard.exe on Windows system

## Install

- Create data folder if you want to create folder with custom permission (default is 0755)
- Run `./daanretard` or `./daanretard.exe`

## Config

You may create one in advance or after starting the app

```env
ADDR=
SECRET=
DATA_PATH=
FB_PAGE_ID=
FB_GRAPH_APP_ID=
FB_GRAPH_APP_SECRET=
DB_DSN=
```
