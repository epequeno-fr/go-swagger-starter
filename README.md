# Using Go-Swagger to Build an API

[Overview](#overview)  
[Quick Start](#quick-start)  
[Generated Code](#generated-code)  
[ID Cloud Integration](#id-cloud-integration)  
[Improvements](#improvements)

***

## Overview

This accelerator is a starting point for building a new API using [go-swagger](https://github.com/go-swagger/go-swagger). The goals are:

- Avoid recreating the wheel for basic API plumbing
- Focus on defining the API spec and implementing endpoint handlers
- Encourage consistency with our API implementations

This starter includes:

- Sample swagger spec
- Sample endpoint handlers
- Middleware for OAuth 2.0 authentication with AM
- Configuration loaded from environment variables
- Makefile for common tasks

## Quick Start

### Prerequisites

1. Install [go-swagger](https://github.com/go-swagger/go-swagger) if you haven't already
2. Set your kube context to a valid customer environment (required for authentication with AM)

### Creating a New API

1. Copy this starter into the monorepo
2. Define the API spec in the `/swagger` folder
3. Run `make generate` to generate docs and server code
4. Implement your endpoint handlers
5. Run a local server with `make run_local` to test

You can now view API docs at http://127.0.0.1:8080/docs. Obtain an access token from your customer environment and
use it to send requests, for example:

```bash
curl http://127.0.0.1:8080/ponies/foo \
  --header 'accept-api-version: resource=1.0' \
  --header 'authorization: Bearer xxx' \
  | jq

curl -X PUT http://127.0.0.1:8080/ponies/foo \
  --header 'accept-api-version: resource=1.0' \
  --header 'authorization: Bearer xxx' \
  --header 'content-type: application/json' \
  --data '{"color":"Black","name":"Joe"}' \
  | jq
```

## Generated Code

These files are generated, and you don't necessarily need to know what they do.  They should not be edited by hand.

| Folder                         | Purpose                                      |
| ------------------------------ | -------------------------------------------- |
| `docs_gen`                     | Documentation (Swagger UI or Redoc)          |
| `swagger_gen/cmd/main.go`      | Entrypoint to start server                   |
| `swagger_gen/models/*`         | Structs for entities including validation    |
| `swagger_gen/server/*`         | Other code to configure and run http server  |
| `swagger_gen/server/restapi/*` | Request/response parameters used in handlers |


**Note**: The file `swagger_gen/server/configure_*.go` is generated, but only once.  You can customize to configure authentication, middleware, etc. Some of that customization has been done in this starter.

## ID Cloud Integration

After you've gotten the skeleton API running, you can:

- change package name from `example.ponies.com` to something appropriate for ID Cloud like: `github.com/ForgeCloud/go/services/...` (see [PR #4221](https://github.com/ForgeCloud/saas/pull/4221/files) for an example)
  - double check `./swagger_gen/server/configure_ponies.go` has also been updated for package changes since this file is not automatically updated.
- remove the `go.mod` and `go.sum` files from the starter root directory.
- generate bazel files as documented in the saas repo README
- Uncomment AM token introspection in `pkg/auth/authenticator.go`
- Update the swagger spec and regenerate the API
- Organize new handlers into packages and hook them up in `pkg/handler/handler.go`

## Improvements

- Add CodeFresh stage template 
- Add kustomize template