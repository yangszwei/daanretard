# 靠北大安4.0 (daanretard)

![build](https://github.com/yangszwei/daanretard/workflows/Go%20Test/badge.svg)
[![](https://img.shields.io/badge/status-active%20development-blue)](https://img.shields.io/badge/status-under%20development-blue)
[![](https://tokei.rs/b1/github/yangszwei/daanretard)](https://github.com/yangszwei/daanretard)

靠北大安4.0投稿系統

## Prerequisite

- MySQL database (refer to [https://github.com/go-gorm/mysql](https://github.com/go-gorm/mysql))
- Facebook Graph API App (refer to [https://developers.facebook.com/docs/graph-api/reference/application/](https://developers.facebook.com/docs/graph-api/reference/application/))

## Install

(at project root)

- Create .env file:

```.env
ADDR=
DATA=
SECRET=
FB_APP_ID=
FB_APP_SECRET=
DB_DSN= 
```

- Run `make build`
- Run `./daanretard`