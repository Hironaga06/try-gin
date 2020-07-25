# try_gin

## 準備

### Docker for Macのインストール
https://docs.docker.com/docker-for-mac/install/

### Goのインストール
1.14系のインストール(goenv, anyenv等は使用しない)  
環境変数に GO111MODULE=on を追加

### rebirthのインストール
```bash
$ go get github.com/goccy/rebirth/cmd/rebirth
```

### 起動準備(初回のみ)
```bash
$ docker-compose build
```

### 起動
コンテナ生成、立ち上げ
```bash
$ docker-compose up -d
```
ホットリローダの起動
```bash
$ rebirth
```

### テーブルのマイグレーションファイル作成
```bash
$ script/create_migrarion_file.sh migrations/schema {テーブル名}
```