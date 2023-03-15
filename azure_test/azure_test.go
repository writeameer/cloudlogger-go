package azure_test

import (
	"fmt"
	"github.com/writeameer/cloudlogger-go/azure"
	"os"
	"testing"
)

func TestAzureBlobStorage(t *testing.T) {

	fmt.Printf("The azure account is:%s \n", os.Getenv("AZURE_STORAGE_ACCOUNTNAME"))

	args := &azure.BlobStorageWriterArgs{
		AccountName:   os.Getenv("AZURE_STORAGE_ACCOUNTNAME"),
		ContainerName: os.Getenv("AZURE_STORAGE_CONTAINERNAME"),
		AccountKey:    os.Getenv("AZURE_ACCOUNT_KEY"),
		FileName:      "test.log",
	}

	w := azure.NewBlobStorageWriter(args)

	for i := 1; i < 5; i++ {
		_, err := fmt.Fprintf(w, "%d. This is a test from cloudlogger !\n", i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
