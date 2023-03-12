# Parser

## Directories

### helper

utilityパッケージ

### str
文字列操作系のパッケージ

### xml

ログファイルのxmlをparseするパッケージ

#### gameinfo.go

対局情報についての処理

#### player.go

対局プレイヤーについての処理

#### xml.go

ログファイル全体の処理

##### hai

牌情報の処理

##### init

対局前の情報の処理

- 手牌
- 親
- 点数
- 乱数シード

##### operation

対局時のプレイヤーの操作の処理

- 副露
- 捨牌
- 自摸牌
- リーチ
- 流局
- 和了
