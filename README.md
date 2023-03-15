# Overview

An io.Writer implementation for cloud APIs. Use`fmt.Printf` write/log directly to cloud !



## Logging to Azure Blob Storage

Create the io.Writer:

```go
// Define Azure args
args := &azure.AzureWriterArgs{
    AccountName:   os.Getenv("AZURE_STORAGE_ACCOUNTNAME"),
    ContainerName: os.Getenv("AZURE_STORAGE_CONTAINERNAME"),
    AccountKey:    os.Getenv("AZURE_ACCOUNT_KEY"),
    FileName:      "test.log",
}

// Create Azure writer
w := azure.NewAzureWriter(args)
```

Start writing to it  !

```go
// Creates a file called `test.log` in the Azure storage account define earlier
fmt.Fprintf(w, "This is a test from cloudlogger !")
```