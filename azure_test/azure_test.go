package azure_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/writeameer/cloudlogger-go/azure"
)

func TestAzureBlobStorage(t *testing.T) {

	args := &azure.WriterArgs{
		AccountName:   os.Getenv("AZURE_STORAGE_ACCOUNTNAME"),
		ContainerName: os.Getenv("AZURE_STORAGE_CONTAINERNAME"),
		AccountKey:    os.Getenv("AZURE_ACCOUNT_KEY"),
		FileName:      "test.log",
	}

	w := azure.NewWriter(args)

	for i := 1; i < 5; i++ {
		_, err := fmt.Fprintf(w, "%d. This is a test from cloudlogger !\n", i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
