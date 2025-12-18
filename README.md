# check-sakura-object-storage-usage

## Description
Check free space of sakura object storage bucket.

さくらのオブジェクトストレージにおけるバケット容量のチェックをするmackerel pluginです.

## Synopsis
```
check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo --api_key=daed-beaf-foo-bar --api_secret=XXXXXXXXXXXXXX
```

## Installation

まずはインストールします.
```
go install github.com/u5surf/check-sakura-object-storage-usage@latest
```

次に、以下のようにコマンドを実行します.


```
./check-sakura-object-storage-usage --warning=10% --critical=5% --site=isk01 --bucket=foo --api_key=daed-beaf-foo-bar --api_secret=XXXXXXXXXXXXXX
Sakura Object Storage Usage OK: usage: site:isk01, bucket:foo, current free: 29.575990%
```
`api_key`と`api_secret`は, [さくらのオブジェクトストレージ APIドキュメント-基本的な使い方-APIキーの発行](https://manual.sakura.ad.jp/api/cloud/objectstorage/#section/%E5%9F%BA%E6%9C%AC%E7%9A%84%E3%81%AA%E4%BD%BF%E3%81%84%E6%96%B9/API) を参考に入手します.
`site`は, 石狩第1サイトならば、`isk01` を指定してください.
`bucket`は, 監視したいバケット名を指定してください.

## Setting for mackerel-agent

If there are no problems in the execution result, add a setting in mackerel-agent.conf .

```
[plugin.checks.objectstorage-free]
command = ["check-sakura-object-storage-usage", "--warning", "10%", "--critical", "5%", "--site", "isk01", "--bucket", "foo", "--api_key", "daed-beaf-foo-bar", "--api_secret", "XXXXXXXXXXXXXX" ]
check_interval = 60
```

## Usage
### Options

```
  -w, --warning=N%                  Exit with WARNING status if less than N% of bucket storage is free
  -c, --critical=N%                 Exit with CRITICAL status if less than N% of bucket storage is free
  -s, --site=STRING                 Choose a site where monitored bucket
  -b, --Bucket=STRING               Choose a monitoring bucket
  -k, --api_key=STRING              Sakura ObjectStorage API key
  -S, --api_secret=STRING           Sakura ObjectStorage API secreet
```

## For more information
Please refer to the following.
- Execute `check-sakura-object-storage-usage -h` and you can get command line options.
