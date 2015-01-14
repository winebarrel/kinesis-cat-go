# kinesis-cat-go
Amazon Kinesis cli for put JSON data.

## Build

```sh
go get github.com/crowdmob/goamz/kinesis
go get code.google.com/p/go-uuid/uuid
go build kinesis-cat.go
```

## Usage

```
Usage of kinesis-cat:
  -accesskey="": AWS access key id credential
  -partitionkey="8c07bece-d278-4121-aac4-46f0b12706ea": Determines which shard
  -region="ap-northeast-1": The name of the stream
  -secretkey="": AWS secret access key credential
  -stream="": The default AWS region
```

## Example

```sh
echo '{"key":"val"}' | kinesis-cat -stream my-stream
```
