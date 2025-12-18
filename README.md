# check-sakura-object-storage-usage

## Description
Check free space of sakura object storage bucket.

さくらのオブジェクトストレージにおけるバケット容量のチェックをするmackerel pluginです.

## Synopsis
```
export SAKURA_API_ACCESS_TOKEN=01234567-89ab-cdef-0123-456789abcdef
export SAKURA_API_ACCESS_TOKEN_SECRET=XXXXXXXXXXXXXX
check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo
```

## Installation

まずはインストールします.
```
go install github.com/u5surf/check-sakura-object-storage-usage@latest
```

次に、以下のようにコマンドを実行します.


```
export SAKURA_API_ACCESS_TOKEN=01234567-89ab-cdef-0123-456789abcdef
export SAKURA_API_ACCESS_TOKEN_SECRET=XXXXXXXXXXXXXX
./check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo
Sakura Object Storage Usage OK: usage: site:isk01, bucket:foo, current free: 29.575990%
```
環境変数 `SAKURA_API_ACCESS_TOKEN`と`SAKURA_API_ACCESS_TOKEN_SECRET`は, [さくらのオブジェクトストレージ APIドキュメント-基本的な使い方-APIキーの発行](https://manual.sakura.ad.jp/api/cloud/objectstorage/#section/%E5%9F%BA%E6%9C%AC%E7%9A%84%E3%81%AA%E4%BD%BF%E3%81%84%E6%96%B9/API) を参考に入手します.
`site`は, 石狩第1サイトならば、`isk01` を指定してください.
`bucket`は, 監視したいバケット名を指定してください.

## Setting for mackerel-agent

If there are no problems in the execution result, add a setting in mackerel-agent.conf .

```
[plugin.checks.objectstorage-free]
command = ["check-sakura-object-storage-usage", "--warning", "10%", "--critical", "5%", "--site", "isk01", "--bucket", "foo"]
env = { SAKURA_API_ACCESS_TOKEN = "01234567-89ab-cdef-0123-456789abcdef", SAKURA_API_ACCESS_TOKEN_SECRET = "XXXXXXXXXXXXXX" }
check_interval = 60
timeout_seconds = 60
max_check_attempts = 2
```

## Usage
### Options

```
  -w, --warning=N%                  Exit with WARNING status if less than N% of bucket storage is free
  -c, --critical=N%                 Exit with CRITICAL status if less than N% of bucket storage is free
  -s, --site=STRING                 Choose a site where monitored bucket
  -b, --Bucket=STRING               Choose a monitoring bucket
```

## For more information
Please refer to the following.
- Execute `check-sakura-object-storage-usage -h` and you can get command line options.
