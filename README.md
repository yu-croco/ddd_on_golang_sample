# DDD on Golang
**基本的にまだ開発中**

- [yu-croco/ddd_on_scala](https://github.com/yu-croco/ddd_on_scala) のGolangバージョン
- Golang（Gin）を使い、なんちゃってモンハンの世界をDomain-Driven Designで実装している

## 技術スタック
- Golang: v1.15
- Go modules
- Gin: v1.6.3
- Gorm: v1.9.16
- Docker: 19.03.12
- docker-compose: 1.26.2

## 構成

```
├── Dockerfile
├── README.md
├── app
│   ├── adapter
│   ├── domain
│   ├── infrastructure
│   └── usecase
├── bin
│   ├── build.sh
│   ├── main
│   └── tidy.sh
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

## セットアップ
- このレポジトリをgit cloneする
- `docker-compose up`でAPIサーバーとDBが起動する

## 実装してみての考察
WIP

（Scalaなどとは別パラダイムの言語である）Golangを用いてDDDをどこまでできたのかをまとめる予定

### 意味のあると感じたこと
- 階層/パッケージ分けにより責任の所在をある程度まとめられる
  - DDDというよりはオニオンアーキテクチャなどの観点かもしれないが..

### 辛み
- Genericを使った共通処理が作れないため記述量が増える
    - implicit classが恋しい
    - そもそも言語パラダイム的に兼ね備えていないので文句を言う方が不適切な気もするが...
- パッケージ構成に慣れが必要
    - `import cycle not allowed` で怒られる...

## 参考
- [Clean Architecture in Go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
- [Practical Persistence in Go: Organising Database Access](https://www.alexedwards.net/blog/organising-database-access)
- [【必須科目 DI】DIの仕組みをGoで実装して理解する](https://qiita.com/yoshinori_hisakawa/items/a944115eb77ed9247794)
- [pospomeのサーバサイドアーキテクチャ3](https://booth.pm/ja/items/1578182)
- [GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】](https://qiita.com/gold-kou/items/45a95d61d253184b0f33)
- [Gin と GORM で作るオレオレ構成 API](https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4#controller-%E3%81%AB-action-%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%99%E3%82%8B)