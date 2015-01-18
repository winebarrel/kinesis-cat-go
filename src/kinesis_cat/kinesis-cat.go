package kinesis_cat

import (
	"encoding/json"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/kinesis"
)

type KinesisCat struct {
	client       *kinesis.Kinesis
	streamName   string
	partitionKey string
}

func (p *KinesisCat) PutJSON(jsonArray []interface{}) (err error) {
	for _, v := range jsonArray {
		var data []byte
		data, err = json.Marshal(v)

		if err == nil {
			_, err = p.client.PutRecord(p.streamName, p.partitionKey, data, "", "")
		}

		if err != nil {
			return
		}
	}

	return
}

func NewKinesisCat(params *KinesisCatParams) *KinesisCat {
	region := aws.GetRegion(params.regionName)

	auth := aws.Auth{
		AccessKey: params.accessKey,
		SecretKey: params.secretKey,
	}

	client := kinesis.New(auth, region)

	return &KinesisCat{
		client:       client,
		streamName:   params.streamName,
		partitionKey: params.partitionKey,
	}
}
