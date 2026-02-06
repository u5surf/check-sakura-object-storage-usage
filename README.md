# check-sakura-object-storage-usage

## Description
Check free space of sakura object storage bucket.

さくらのオブジェクトストレージにおけるバケット容量のチェックをするmackerel check pluginです.

## Synopsis
```
export SAKURA_ACCESS_TOKEN=01234567-89ab-cdef-0123-456789abcdef
export SAKURA_ACCESS_TOKEN_SECRET=XXXXXXXXXXXXXX
check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo
```

## Installation

まずはインストールします.
[リリースページ](https://github.com/u5surf/check-sakura-object-storage-usage/releases/) からダウンロード
もしくは, mkr コマンドにて以下を実施してください.
```
mkr plugin install u5surf/check-sakura-object-storage-usage
```

次に、以下のようにコマンドを実行します.


```
export SAKURA_ACCESS_TOKEN=01234567-89ab-cdef-0123-456789abcdef
export SAKURA_ACCESS_TOKEN_SECRET=XXXXXXXXXXXXXX
./check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo
Sakura Object Storage Usage OK: usage: site:isk01, bucket:foo, current free: 29.575990%
```
環境変数 `SAKURA_ACCESS_TOKEN`と`SAKURA_ACCESS_TOKEN_SECRET`は, [APIキー | さくらのクラウド マニュアル](https://manual.sakura.ad.jp/cloud/api/apikey.html) を参考に入手します.
APIキーを発行する際に, サービスへのアクセス権で「オブジェクトストレージ」を選択してください.

`site`は, 石狩第1サイトならば、`isk01` を指定してください.
`bucket`は, 監視したいバケット名を指定してください.

## Setting for mackerel-agent

mackerel check pluginとして、以下のように設定をすれば利用可能です.

```
[plugin.checks.objectstorage-free]
command = ["check-sakura-object-storage-usage", "--warning", "10%", "--critical", "5%", "--site", "isk01", "--bucket", "foo"]
env = { SAKURA_ACCESS_TOKEN = "01234567-89ab-cdef-0123-456789abcdef", SAKURA_ACCESS_TOKEN_SECRET = "XXXXXXXXXXXXXX" }
check_interval = 60
timeout_seconds = 60
max_check_attempts = 2
```

### Options

```
  -w, --warning=N%                  Exit with WARNING status if less than N% of bucket storage is free
  -c, --critical=N%                 Exit with CRITICAL status if less than N% of bucket storage is free
  -s, --site=STRING                 Choose a site where monitored bucket
  -b, --Bucket=STRING               Choose a monitoring bucket
```

## For more information
- `check-sakura-object-storage-usage -h` を実行することで、コマンドラインオプションを取得できます.
- このプラグインは, さくらのオブジェクトストレージ APIを使用しています。詳しいAPIドキュメントはこちらを参照ください. [さくらのオブジェクトストレージ APIドキュメント](https://manual.sakura.ad.jp/api/cloud/objectstorage/)
