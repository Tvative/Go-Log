# Go Log

A Simple Logging Library for Go Language

## Introduction

Go Log is a Go library for flexible and customizable logging. It
provides an easy way to log messages to both the terminal and log files.
You can also add color to your terminal log messages for better readability

## Features

- Log messages to the terminal with or without colors
- Log messages to a file
- Flexible and customizable log output
- Simple API for logging messages

## Getting Started

### Installation

To use Go Log in your Go project, you need to install it using Go modules.
You can add it to your project with the following command:

```bash
go get github.com/Tvative/Package-Go-Log
```

### Usage

Here's how you can use Go Log in your Go code:

```go
package main

import "github.com/Tvative/Package-Go-Log"

var logInstance *GoLog.LogInstance

func main() {
	var fileDestination = "Test/_log.log"
	logInstance = GoLog.Initialize(fileDestination)

	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": 1.5,
	}

	logInstance.Warning(jsonString, "Sample warning log message")
}
```

Make sure to import the `github.com/Tvative/Package-Go-Log` package and create a LogData instance to use
the provided logging functions

## Documentation

For detailed documentation, check the [package](https://pkg.go.dev/github.com/Tvative/Package-Go-Log) for this project

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
