package azure_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/writeameer/cloudlogger-go/azure"
)

func TestLogtoAzureBlobStorage(t *testing.T) {

	accountName := os.Getenv("AZURE_STORAGE_ACCOUNTNAME")
	containerName := os.Getenv("AZURE_STORAGE_CONTAINERNAME")
	accountKey := os.Getenv("AZURE_ACCOUNT_KEY")
	fileName := "test.log"

	w := azure.NewAzureWriter(accountName, containerName, accountKey, fileName)

	fmt.Fprintf(w, "This is a test from cloudlogger !")
}
