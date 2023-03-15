package azure

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/appendblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

// Writer is an implementation of io.Writer to write to Azure Blob Storage

type BlobStorageWriter struct {
	ctx context.Context
	w   *appendblob.Client
}

type BlobStorageWriterArgs struct {
	// The name of the Azure storage account. Please specify without the FQDN.
	// For e.g. if your storage account is https://myaccount.blob.core.windows.net, then just  pass in 'myaccount'
	AccountName string

	// The name of an existing blob container you have created in the storage account where you want to write to. For e.g. mycontainer
	ContainerName string

	// The Azure storage account key that allows you write access to your blob container
	AccountKey string

	// The name of the file you want to create/append in your blob container
	FileName string
}

func NewBlobStorageWriter(args *BlobStorageWriterArgs) *BlobStorageWriter {

	// Setup context
	ctx := context.Background()

	// Generate container url
	containerURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", args.AccountName, args.ContainerName)

	// Create credential
	credential, err := azblob.NewSharedKeyCredential(args.AccountName, args.AccountKey)
	if err != nil {
		log.Printf("Could not create shared key for account: %s\n", args.AccountName)
		log.Println(err.Error())
	}
	// Create a container client with creds and url
	containerClient, err := container.NewClientWithSharedKeyCredential(containerURL, credential, nil)
	if err != nil {
		log.Printf("Could not create container client key for url: %s\n", containerURL)
		log.Println(err.Error())
	}
	// Create appendClient from container client
	appendClient := containerClient.NewAppendBlobClient(args.FileName)

	// Create file if not exists
	_, err = appendClient.Create(ctx, nil)
	if err != nil {
		log.Println("Error creating 0-size append blob")
		log.Fatal(err.Error())
	}

	return &BlobStorageWriter{
		w:   appendClient,
		ctx: ctx,
	}

}
func (e BlobStorageWriter) Write(data []byte) (int, error) {

	tmStamp := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	message := fmt.Sprintf("%s  %s", tmStamp, string(data))

	body := streaming.NopCloser(strings.NewReader(message))

	_, err := e.w.AppendBlock(e.ctx, body, nil)
	if err != nil {
		log.Println("Error appending blob:" + err.Error())
		return 0, handleError(err)
	}

	return len([]byte(message)), nil
}

func handleError(err error) error {
	if err != nil {
		log.Println(err.Error())
	}

	return err
}
