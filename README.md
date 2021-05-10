# Using Go-Swagger to Build an API

[Overview](#overview)  
[Quick Start](#quick-start)  
[Generated Code](#generated-code)  
[ID Cloud Integration](#id-cloud-integration)

***

## Overview

This accelerator is a starting point for building a new API using [go-swagger](https://github.com/go-swagger/go-swagger).

It includes:

- Starter swagger spec
- Middleware for OAuth 2.0 authentication
- Sample endpoint handlers
- Configuration loaded from environment variables
- Makefile for common tasks

## Quick Start

The general process to build out your API is:

1. Define API spec in the `/swagger` folder
2. Run `make generate` to generate docs and server
3. Implement endpoint handlers
4. Run local server with `make run_local`

You can now view API docs at http://127.0.0.1:8080/docs and send sample requests:

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

| Folder                         | Purpose                                      |
| ------------------------------ | -------------------------------------------- |
| `docs_gen`                     | Documentation (Swagger UI or Redoc)          |
| `swagger_gen/cmd/main.go`      | Entrypoint to start server                   |
| `swagger_gen/models/*`         | Structs for entities including validation    |
| `swagger_gen/server/*`         | Other code to configure and run http server  |
| `swagger_gen/server/restapi/*` | Request/response parameters used in handlers |

## ID Cloud Integration

After you've gotten the skeleton API running, you can:

- Uncomment AM token introspection in `pkg/auth/authenticator.go`
- Update the swagger spec and regenerate the API
- Organize new handlers into packages and hook them up in `pkg/handler/handler.go`