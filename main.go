package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// create session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "tteshima",
		SharedConfigState: session.SharedConfigEnable,
	}))

	filename := "test.txt"
	file, err := os.Open("./" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("teshima-test-bucket"),
		Key:    aws.String("/" + filename),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")

}
