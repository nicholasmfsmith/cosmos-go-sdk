# Cosmos Go SDK

**DISCLAIMER: THIS PROJECT IS STILL IN THE VERY EARLY PHASES AND IS NOT YET FUNCTIONAL**

[![Build Status](https://travis-ci.com/nicholasmfsmith/cosmos-go-sdk.svg?branch=master)](https://travis-ci.com/nicholasmfsmith/cosmos-go-sdk)

## Purpose
- Create a developer friendly Azure Cosmos Go SDK that enables the use of chained resource-driven components. We hope to achieve something similar to the developer experience provided by the JavaScript Cosmos SDK (while maintaining idiomatic Go)

Example:
```go
// Intended interaction for a Document Read
client := client.New("url", "key")
client.Database("dbName").Collection("collName").Document("docName").Read()
```

## Note
- We nor this repository are affiliated with Microsoft. This repository simply serves to make the developer experience with the Azure Cosmos REST API more seamless. Happy coding! :)

## Contributors
- [Nicholas Smith](https://github.com/nicholasmfsmith)