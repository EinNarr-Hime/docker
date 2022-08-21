## Go lang イメージ

### ソースコードの変更

- ./app がコンテナ内の/go/src/app にマウントされるので
  ローカルから./app 以下に対する変更がコンテナ側にも反映される。

### Go 実行手順

```shell
# コンテナの起動(バックグラウンド)
docker-compose up -d

# goコンテナに対してhello.goを実行
docker-compose exec [コンテナ名] go run hello.go
```
