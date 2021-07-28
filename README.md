# expression_evaluator

REST API that evaluates expressions

# Building and running the application

Open a command prompt and navigate to the project folder.

Windows command line:

```console
$ go build -o expression_eval.exe .\cmd\web\expression_eval\.
$ expression_eval.exe
```

Bash:

```console
$ go build -o expression_eval ./cmd/web/expression_eval/.
$ ./expression_eval
```

# Testing the application

From the project root folder:

```console
$ go test ./...
```

# Checking test coverage

From the project root folder:

```console
$ go test -coverprofile=coverage.out ./...
$ go tool cover -html=coverage.out
```

# Design decisions

## Folder structure

I've opted to use a file structure commonly used in Go projects. This folder structure should be a good starting point for Domain-Driven Design.

- cmd/web/expression_eval
  - This file structure allows for multiple commands to built with the same packages (e.g. a cli app, a back end REST API etc)
  - I've added routes.go here to allow routes to be configured per command
- pkg - contains project packages. A different convention could be to name this folder "internal" for internal packages
- pkg/handlers - contains http handlers linked to the routes in cmd/web/expression_eval/routes.go

## Design

This is a REST API that accepts JSON for both input and output. I've tried to use appropriate http response codes. Some expressions involve a combination of int and float64 values. Where I need to cast values I casted them to float64 to preserve result accuracy.

As mentioned above, the folder structure should be a good starting point for Domain-Driven Design. I believe it is important to keep the SOLID design principles in mind when designing applications as it reduces the risk bugs when changing / refactoring code. The folder structure used and e.g. the various sets of expressions that have been placed in separate struct methods brings about separation of concern.

## Frameworks

I've opted to not use any 3rd party libraries for the current requirements. For more complex routing I would consider using e.g. the chi router and libraries to assist with security.
