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