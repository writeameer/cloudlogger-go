package azure_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/writeameer/cloudlogger-go/azure"
)

func TestAzureBlobStorage(t *testing.T) {

	accountName := os.Getenv("AZURE_STORAGE_ACCOUNTNAME")
	containerName := os.Getenv("AZURE_STORAGE_CONTAINERNAME")
	accountKey := os.Getenv("AZURE_ACCOUNT_KEY")
	fileName := "test.log"

	w := azure.NewAzureWriter(accountName, containerName, accountKey, fileName)

	for i := 1; i < 5; i++ {
		_, err := fmt.Fprintf(w, "%d. This is a test from cloudlogger !\n", i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}
