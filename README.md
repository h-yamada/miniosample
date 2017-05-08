# MINIO via aws-sdk-go

## MINIO

### MINIOとは

AWS S3 APIとの互換性を持ったオープンソースのオブジェクトストレージです。

### Docker Container

AccessKeyとSecretKeyを指定しないと毎回ランダムに発行されます。
なお、利用したBucketを永続的に利用したい場合はマウント先を指定して下さい。

```
docker pull minio/minio

docker run -p 9000:9000 \
 -e "MINIO_ACCESS_KEY=AccessKey" \
 -e "MINIO_SECRET_KEY=SecretKey" \
 minio/minio server /export
```

## GO言語サンプル

### ビルド

コンパイルオプションに応じて環境毎に定義したconfigでビルドします。
localでMINIO、それ以外はAWS利用を想定した作りにしてます。

```
go build -tags=local
```
### Usage

```
$ ./miniosample -help
Usage of ./miniosample:
  -bucket string
    	S3 Bucket
  -file string
    	local file
  -key string
    	S3 file
  -type string
    	command type d:Download u:Upload
```

### Example

* アップロード

```
./miniosample -type=u -bucket=sample -file=./Freeza.jpg -key=Freeza.jpg
```

* ダウンロード

```
./miniosample -type=d -bucket=sample -file=~/Downloads/Freeza.jpg -key=Freeza.jpg
```
