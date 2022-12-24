package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	gcpClient  *storage.Client
	bucketName string
	ctx        context.Context
	wg         sync.WaitGroup
)

func init() {
	ctx = context.Background()
	credentialsFile, err := os.ReadFile("./credentials.json")
	if err != nil {
		panic(err)
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credentialsFile))
	if err != nil {
		panic(err)
	}
	bucketName = "goexpert-bucket-example"
	gcpClient = client
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
	defer gcpClient.Close()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, bucketName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		<-uploadControl // libera o canal
		errorFileUpload <- filename
		return
	}
	objectHandle := gcpClient.Bucket(bucketName).Object(filename)
	writer := objectHandle.NewWriter(ctx)
	defer writer.Close()
	_, err = io.Copy(writer, f)
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", completeFileName, err)
		<-uploadControl // libera o canal
		errorFileUpload <- filename
		return
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // libera o canal
}
