package kinesis_cat

import (
	"code.google.com/p/go-uuid/uuid"
	"flag"
	"github.com/vaughan0/go-ini"
	"log"
	"os"
)

type KinesisCatParams struct {
	accessKey    string
	secretKey    string
	regionName   string
	streamName   string
	partitionKey string
}

func loadProfile(profile string, credentialsPath string, accessKey *string, secretKey *string) {
	credentialsFile, err := ini.LoadFile(credentialsPath)

	if err != nil {
		log.Fatal(err)
	}

	credentials, ok := credentialsFile[profile]

	if !ok {
		log.Fatalf("The config profile (%s) could not be found", profile)
	}

	*accessKey = credentials["aws_access_key_id"]
	*secretKey = credentials["aws_secret_access_key"]
}

func ParseFlag() (params *KinesisCatParams) {
	params = &KinesisCatParams{}
	var profile, credentialsPath string

	flag.StringVar(&params.accessKey, "accesskey", "", "AWS access key id credential")
	flag.StringVar(&params.secretKey, "secretkey", "", "AWS secret access key credential")
	flag.StringVar(&profile, "profile", "", "The credentials file's profile")
	flag.StringVar(&credentialsPath, "credentials-path", os.Getenv("HOME")+"/.aws/credentials", "The credentials file's path")
	flag.StringVar(&params.regionName, "region", os.Getenv("AWS_REGION"), "The name of the stream")
	flag.StringVar(&params.streamName, "stream", "", "The default AWS region")
	flag.StringVar(&params.partitionKey, "partitionkey", uuid.New(), "Determines which shard")
	flag.Parse()

	if profile != "" {
		loadProfile(profile, credentialsPath, &params.accessKey, &params.secretKey)
	}

	if params.streamName == "" {
		log.Fatal("argument -stream is required")
	}

	return
}
