# kinesis-cat-go
Amazon Kinesis cli for put JSON data.

[![Build Status](https://travis-ci.org/winebarrel/kinesis-cat-go.svg?branch=master)](https://travis-ci.org/winebarrel/kinesis-cat-go)

## Installation

```sh
make build
make install
```

## Usage

```
Usage of kinesis-cat:
  -accesskey="": AWS access key id credential
  -credentials-path="/Users/sugawara/.aws/credentials": The credentials file's path
  -partitionkey="714855c5-acfb-4af5-a7da-62ac67cb9f19": Determines which shard
  -profile="": The credentials file's profile
  -region="ap-northeast-1": The name of the stream
  -secretkey="": AWS secret access key credential
  -stream="": The default AWS region```
```

## Example

```sh
echo '{"key":"val"}' | kinesis-cat -stream my-stream
```
