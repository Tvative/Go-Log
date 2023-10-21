// Initialize Package
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// LogData is a struct that holds information about logging

type LogData struct {
	logDestination *os.File // logDestination is the file where the log will be written
}

// ColorDefault represents the ANSI escape sequence for resetting the text color to the default
// ColorRed represents the ANSI escape sequence for setting text color to red
// ColorGreen represents the ANSI escape sequence for setting text color to green
// ColorYellow represents the ANSI escape sequence for setting text color to yellow

const (
	ColorDefault string = "\x1b[0;0m"
	ColorRed     string = "\x1b[31;1m"
	ColorYellow  string = "\x1b[33;1m"
)

// MessageNormal represents a normal message identifier.
// MessageFatal represents a fatal error message identifier.
// MessageWarning represents a warning message identifier.

const (
	MessageNormal  string = " [ INFO ] "
	MessageFatal   string = " [ ERRO ] "
	MessageWarning string = " [ WARN ] "
)

// Initialize initializes the log data with the provided file destination
// It opens the file specified by fileDestination and prepares it for writing
// If the file cannot be opened, it returns false along with an error message
//
// Parameters:
// 	fileDestination: The path to the log file.
//
// Returns:
// 	bool: true if the file was successfully opened, false otherwise

func (logData *LogData) Initialize(fileDestination string) (bool, string) {
	fileDescriptor, openError := os.Create(fileDestination)

	if openError != nil {
		return false, openError.Error()
	}

	logData.logDestination = fileDescriptor
	return true, ""
}

// Print writes the log message to the specified output destinations
//
// Parameters:
// 	logData: The LogData instance
// 	needFileOutput: If true, the message is written to the log file
// 	needTerminalOutput: If true, the message is displayed on the terminal
// 	needTerminalColoredOutput: If true, the terminal output is colored
// 	messageType: The message type to use for terminal output type
// 	messageContent: The content of the log message

func (logData *LogData) printOutPut(needFileOutput bool, needTerminalOutput bool,
	needTerminalColoredOutput bool, messageType string,
	jsonContent map[string]interface{}, messageContent ...any) {
	var messagePrefix string

	// Generate message prefix

	getTime := time.Now()
	generateLongTime := getTime.Format("2006-01-02 15:04:05")
	generatedTimeMillSeconds := getTime.Nanosecond() / 1e6
	generatedTimeNanoSeconds := getTime.Nanosecond()

	var generatedTime string = generateLongTime + ":" +
		strconv.Itoa(generatedTimeMillSeconds) + ":" +
		strconv.Itoa(generatedTimeNanoSeconds)

	messagePrefix = generatedTime + messageType

	// Print to the file

	if needFileOutput {
		fmt.Fprint(logData.logDestination, messagePrefix)
		fmt.Fprint(logData.logDestination, messageContent...)

		if jsonContent != nil {
			logData.generateJSON(true, false, jsonContent)
		}

		fmt.Fprintln(logData.logDestination)
	}

	// Print to the terminal

	if needTerminalOutput && needTerminalColoredOutput {
		var colorCode string

		switch messageType {
		case MessageNormal:
			colorCode = ColorDefault

		case MessageFatal:
			colorCode = ColorRed

		case MessageWarning:
			colorCode = ColorYellow
		}

		fmt.Print(colorCode, messagePrefix)
		fmt.Print(messageContent...)

		if jsonContent != nil {
			logData.generateJSON(false, true, jsonContent)
		}

		fmt.Println(ColorDefault)
	} else if needTerminalOutput {
		fmt.Print(messagePrefix)
		fmt.Print(messageContent...)

		if jsonContent != nil {
			logData.generateJSON(false, true, jsonContent)
		}

		fmt.Println()
	}

	// Exit if fatal

	if messageType == MessageFatal {
		os.Exit(1)
	}
}

// Generate and print JSON content
//
// Parameters:
// 	needFileOutPut: If true, the message is written to the log file
// 	needTerminalOutput: If true, the message is displayed on the terminal
// 	jsonData: The JSON content of the message

func (logData *LogData) generateJSON(needFileOutPut bool, needTerminalOutput bool,
	jsonData map[string]interface{}) {

	if needFileOutPut {
		fmt.Fprint(logData.logDestination, " [")
	}

	if needTerminalOutput {
		fmt.Print(" [")
	}

	for jsonKey, jsonValue := range jsonData {
		if needFileOutPut {
			fmt.Fprint(logData.logDestination, " (", jsonKey, ": ", jsonValue, ")")
		}

		if needTerminalOutput {
			fmt.Print(" (", jsonKey, ": ", jsonValue, ")")
		}
	}

	if needFileOutPut {
		fmt.Fprint(logData.logDestination, " ]")
	}

	if needTerminalOutput {
		fmt.Print(" ]")
	}
}
