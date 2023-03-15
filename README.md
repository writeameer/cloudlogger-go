# Overview

An io.Writer implementation for cloud APIs. Use`fmt.Printf` write/log directly to cloud !



## Logging to Azure Blob Storage

Create the io.Writer:

```go
accountName := os.Getenv("AZURE_STORAGE_ACCOUNTNAME")
containerName := os.Getenv("AZURE_STORAGE_CONTAINERNAME")
accountKey := os.Getenv("AZURE_ACCOUNT_KEY")
fileName := "test.log"

// Create the writer
w := azure.NewAzureWriter(accountName, containerName, accountKey, fileName)
```

Start writing to it using `fmt.Prtinf` !

```go
// Creates a file caled `test.log` in the Azure storage account define earlier
fmt.Fprintf(w, "This is a test from cloudlogger !")
```