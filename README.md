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
go get github.com/Tvative/Go-Log
```

### Usage

Here's how you can use Go Log in your Go code:

```go
package main

import goLog "github.com/Tvative/Go-Log"

var logData *goLog.LogData

func main() {
    // Create a LogData instance

    logData = &goLog.LogData{}

    // Log a default message to the terminal

    logData.LogP("This is a normal message")

    // Log a warning message to the terminal with colors

    logData.WarningPC("This is a colored message")

    // Log a fatal message to a file

    logData.FatalF("This message is logged to a file")
}
```

Make sure to import the `github.com/Tvative/Go-Log` package and create a LogData instance to use
the provided logging functions

## Documentation

For detailed documentation, check the [package](https://pkg.go.dev/github.com/Tvative/Go-Log) for this project

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
