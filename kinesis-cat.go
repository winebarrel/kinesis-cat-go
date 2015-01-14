package main

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/kinesis"
	"io/ioutil"
	"log"
	"os"
)

func parseFlag(accessKey *string, secretKey *string, region *aws.Region, streamName *string, partitionKey *string) {
	var regionName string

	flag.StringVar(accessKey, "accesskey", "", "AWS access key id credential")
	flag.StringVar(secretKey, "secretkey", "", "AWS secret access key credential")
	flag.StringVar(&regionName, "region", os.Getenv("AWS_REGION"), "The name of the stream")
	flag.StringVar(streamName, "stream", "", "The default AWS region")
	flag.StringVar(partitionKey, "partitionkey", uuid.New(), "Determines which shard")
	flag.Parse()

	if *accessKey == "" {
		*accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	}

	if *secretKey == "" {
		*secretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}

	if *streamName == "" {
		log.Fatal("argument -stream is required")
	}

	*region = aws.GetRegion(regionName)
}

func readJSON() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func parseJSON(src []byte) (jsonArray []interface{}) {
	var data interface{}
	dec := json.NewDecoder(bytes.NewReader(src))
	dec.Decode(&data)

	switch data.(type) {
	case []interface{}:
		jsonArray = data.([]interface{})
	case map[string]interface{}:
		jsonArray = []interface{}{data.(map[string]interface{})}
	default:
		fmt.Fprintf(os.Stderr, "invalid JSON: %v", string(src))
	}

	return
}

func putRecord(client *kinesis.Kinesis, streamName *string, partitionKey *string, data []byte) {
	_, err := client.PutRecord(*streamName, *partitionKey, data, "", "")

	if err != nil {
		log.Fatal(err)
	}
}

func putJSON(client *kinesis.Kinesis, streamName *string, partitionKey *string) {
	src := readJSON()
	jsonArray := parseJSON(src)

	for _, v := range jsonArray {
		data, err := json.Marshal(v)

		if err != nil {
			log.Fatal(err)
		}

		putRecord(client, streamName, partitionKey, data)
	}
}

func init() {
	log.SetFlags(0)
}

func main() {
	var accessKey, secretKey, streamName, partitionKey string
	var region aws.Region

	parseFlag(&accessKey, &secretKey, &region, &streamName, &partitionKey)

	auth := aws.Auth{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	client := kinesis.New(auth, region)
	putJSON(client, &streamName, &partitionKey)
}
