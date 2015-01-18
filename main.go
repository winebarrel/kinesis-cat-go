package main

import (
	"io/ioutil"
	"kinesis_cat"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func readStdin() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func reatJSON() []interface{} {
	src := readStdin()
	json, err := kinesis_cat.ParseJSON(src)

	if err != nil {
		log.Fatal(err)
	}

	return json
}

func putJSON(client *kinesis_cat.KinesisCat, json []interface{}) {
	err := client.PutJSON(json)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	params := kinesis_cat.ParseFlag()
	client := kinesis_cat.NewKinesisCat(params)
	json := reatJSON()
	putJSON(client, json)
}
