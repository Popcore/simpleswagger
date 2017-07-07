#Simple Swagger

## Generate Spec
swagger generate spec -o ./swagger.json

## Generate Client
swagger generate client

## Generate and run UI
swagger serve swagger.json -F redoc
