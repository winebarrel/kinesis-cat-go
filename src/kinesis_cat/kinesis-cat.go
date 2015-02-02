package kinesis_cat

import (
	"encoding/json"
	"github.com/AdRoll/goamz/aws"
	"github.com/AdRoll/goamz/kinesis"
	"time"
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

func NewKinesisCat(params *KinesisCatParams) (kinesisCat *KinesisCat, err error) {
	region := aws.GetRegion(params.regionName)
	auth, err := aws.GetAuth(params.accessKey, params.secretKey, "", time.Time{})
	client := kinesis.New(auth, region)

	kinesisCat = &KinesisCat{
		client:       client,
		streamName:   params.streamName,
		partitionKey: params.partitionKey,
	}

	return
}
