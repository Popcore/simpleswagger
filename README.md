# Simple Swagger API
This application is a simple HTTP server documented using Swagger.

Its purpose is to demonstrate some of possible comment annotations and features that are available to turn go code into a fully compliant swagger 2.0 spec.


## Dependencies
The Simple Swagger API requires the following tools:

### Golang
To install and get started with the Go programming language see the [official documentation]( https://golang.org/doc/install?download=go1.5.windows-amd64.msi2).

### Go Swagger
The easier way to install `go swagger` is to install it from source. Once Golang is installed it can be done by running
```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

For alternative ways to install the package please see:
https://goswagger.io/#docs


## Run the application
The application can be run by issuing the following command inside the application root folder
```
go run main.go
```

The app runs at `http://localhost:3001` and the only available endpoint is `http://localhost:3001/items`.

Items can be filtered by specifing their id as query argument (e.g `http://localhost:3001/items?q=1`).


## Swagger Specs and UI

Swagger specs can be generated running
```
swagger generate spec -o ./swagger.json
```
This command will create a `swagger.json` specification document used by swagger to generate API docs.
The API documentation can be visualized through a user interface running
```
swagger serve swagger.json -F redoc
```

The Makefile included in this repository can be used to issue the commands above
```
make doc_spec # generates spec

make doc_ui   # runs the user interface
```

## Learn More
More information about GoSwagger can be found at:
https://goswagger.io/#docs

To lean more about the  Swagger project please visit:
https://swagger.io/
