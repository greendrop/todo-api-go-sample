# todo-api-go-sample

GO言語を使用したAPIのサンプルです。

## 使用言語・ライブラリ

- Golang
- dep
- godo
- Echo
- GORM

## 環境構築アプリケーション

- Docker
- Docker Compose
- direnv
- git

## セットアップ

```shell
$ git clone git@github.com:greendrop/todo-api-go-sample.git
$ cd todo-api-go-sample
$ vi .envrc
$ direnv allow
$ docker-compose pull
$ docker-compose build
$ docker-compose up mysql
$ docker-compose exec mysql bash
# mysql -u root -p mysql
create database todo;
exit
# exit
$ docker-compose run --rm app bash
$ go get -u github.com/golang/dep/cmd/dep
$ go get -u gopkg.in/godo.v2/cmd/godo
$ dep ensure
$ docker-compose up
```

### .envrc

```
export USER_ID=`id -u`
export GROUP_ID=`id -g`
```
