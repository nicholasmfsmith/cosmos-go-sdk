# Contributing to Cosmos Go SDK

👍🎉 Thank you for taking the time to contribute! 🎉👍

The following are a set of guidelines for contributing to the **Cosmos Go SDK** and its packages. We certainly appreciate your input as this is an effort to offer the community a needed Go SDK to enhance our developer experience.

Disclaimer: This repository is not maintained by Microsoft. We share no affiliation to Microsoft. 

## TODO: Steps for creating Issues

## Steps for creating Pull Requests

NOTE: Pull Requests without Unit Testing will *not* be accepted. 

### Branching

Please adhere to the branching conventions below.

```bash
$ git checkout -b "[type of PR]/[package(s) affected]/[topic addressed]"

Examples:
$ git commit -m "feature/rest/post-request-set-optional-headers"
$ git commit -m "bugfix/collection/error-handling-missing"
$ git commit -m "docs/database/added-package-usage-example"
```

### Commit Message

Please send us a Pull Request with a clear commit message and list of what you are accomplishing. 

```bash
$ git commit -m "[Description of PR changes]"

Example:
$ git commit -m "Added Rest Package ability to set optional headers for POST request"
```

## External Documentation

We are closely following [Azure Cosmos DB: Rest API Reference](https://docs.microsoft.com/en-us/rest/api/cosmos-db/) for implementation - including nomenclature (ex. Collections vs Container).

## TODO: Coding Conventions

### Testing

[Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/)

## TODO: Code of Conduct

## TODO: Community and behavioral expectations