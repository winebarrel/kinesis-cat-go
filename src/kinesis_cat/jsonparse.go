package kinesis_cat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ParseJSON(src []byte) (jsonArray []interface{}, err error) {
	var data interface{}
	dec := json.NewDecoder(bytes.NewReader(src))
	dec.Decode(&data)

	switch data.(type) {
	case []interface{}:
		jsonArray = data.([]interface{})
	case map[string]interface{}:
		jsonArray = []interface{}{data.(map[string]interface{})}
	default:
		err = fmt.Errorf("invalid JSON: %v", string(src))
	}

	return
}
