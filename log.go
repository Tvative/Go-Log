// Go Log Package
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

// Package golog provides a logging mechanism for managing log instances
//
// Go Log is a Go library for flexible and customizable logging. It provides
// an easy way to log messages to both the terminal and log files. You can
// also add color to your terminal log messages for better readability
package golog

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Log format types
const (
	DefaultFormat int = iota // DefaultFormat generates log message in default format
	JsonFormat               // JsonFormat generates log message in JSON format
)

// Colors for log messages
const (
	DefaultColor string = "\033[0;0m"  // DefaultColor print log messages in default color
	RedColor     string = "\033[1;31m" // RedColor print log messages in red color
	GreenColor   string = "\033[1;32m" // GreenColor print log messages in green color
	YellowColor  string = "\033[1;33m" // YellowColor print log messages in yellow color
	BlueColor    string = "\033[1;34m" // BlueColor print log messages in blue color
	MagentaColor string = "\033[1;35m" // MagentaColor print log messages in magenta color
	CyanColor    string = "\033[1;36m" // CyanColor print log messages in cyan color
)

// Log message prefix types
const (
	NormalLog  int = iota // NormalLog is the prefix for default log
	SuccessLog            // SuccessLog is the prefix for success log
	ErrorLog              // ErrorLog is the prefix for error log
	WarningLog            // WarningLog is the prefix for warning log
	DebugLog              // DebugLog is the prefix for debug log
	InfoLog               // InfoLog is the prefix for information log
)

// Instance is a struct that holds base information about logging
type Instance struct {
	Destination    *os.File // Destination hold the destination file for logging
	FileFormat     int      // FileFormat hold the file format for logging
	TerminalFormat int      // TerminalFormat hold the terminal format for logging
}

// Initialize initializes a new log instance
//
// Initialize creates a new log instance with default values
// It sets the destination file for logging to stdout,
// the file format for logging to default format,
// and the terminal format for logging to default format.
//
// It returns a pointer to the new log instance
func Initialize() *Instance {
	instance := Instance{
		Destination:    nil,
		FileFormat:     DefaultFormat,
		TerminalFormat: DefaultFormat,
	}

	return &instance
}

// SetFile sets the destination file for logging
func (instance *Instance) SetFile(filePath string) {
	descriptor, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	instance.Destination = descriptor
}

// SetFileFormat sets the file format for logging
func (instance *Instance) SetFileFormat(format int) {
	instance.FileFormat = format
}

// SetTerminalFormat sets the terminal format for logging
func (instance *Instance) SetTerminalFormat(format int) {
	instance.TerminalFormat = format
}

// generateLogTime generates a log time in the format "2006-01-02 15:04:05:000000"
func generateLogTime() string {
	getTime := time.Now()
	generateLongTime := getTime.Format("2006-01-02 15:04:05")
	generatedTimeMillSeconds := getTime.Nanosecond() / 1e6
	generatedTimeNanoSeconds := getTime.Nanosecond()

	return generateLongTime + ":" +
		strconv.Itoa(generatedTimeMillSeconds) + ":" +
		strconv.Itoa(generatedTimeNanoSeconds)
}

// printToTerminal prints the log to the terminal
func printToTerminal(content ...interface{}) {
	length := len(content)

	for i, arg := range content {
		if _, err := fmt.Print(arg); err != nil {
			return
		}

		if i < length-1 {
			if _, err := fmt.Print(" "); err != nil {
				return
			}
		}
	}

	fmt.Println()
}

// printToFile prints the log to the file
func printToFile(destination *os.File, content ...interface{}) {
	length := len(content)

	for i, arg := range content {
		if _, err := fmt.Fprint(destination, arg); err != nil {
			return
		}

		if i < length-1 {
			if _, err := fmt.Fprint(destination, " "); err != nil {
				return
			}
		}
	}

	if _, err := fmt.Fprintln(destination); err != nil {
		return
	}
}

// returnPrefix returns the prefix for the default log message based on the prefix type
func returnDefaultPrefix(prefix int, isFile bool) string {
	var (
		NormalLogString  string
		SuccessLogString string
		ErrorLogString   string
		WarningLogString string
		DebugLogString   string
		InfoLogString    string
	)

	if isFile == true {
		NormalLogString = "[ LOG ]"
		SuccessLogString = "[ SUC ]"
		ErrorLogString = "[ ERR ]"
		WarningLogString = "[ WRN ]"
		DebugLogString = "[ DBG ]"
		InfoLogString = "[ INF ]"
	} else {
		NormalLogString = CyanColor + "[ LOG ]" + DefaultColor
		SuccessLogString = GreenColor + "[ SUC ]" + DefaultColor
		ErrorLogString = RedColor + "[ ERR ]" + DefaultColor
		WarningLogString = YellowColor + "[ WRN ]" + DefaultColor
		DebugLogString = MagentaColor + "[ DBG ]" + DefaultColor
		InfoLogString = BlueColor + "[ INF ]" + DefaultColor
	}

	switch prefix {
	case NormalLog:
		return NormalLogString
	case SuccessLog:
		return SuccessLogString
	case ErrorLog:
		return ErrorLogString
	case WarningLog:
		return WarningLogString
	case DebugLog:
		return DebugLogString
	case InfoLog:
		return InfoLogString
	default:
		return ""
	}
}

// returnPrefix returns the prefix for the json log message based on the prefix type
func returnJsonPrefix(prefix int) string {
	switch prefix {
	case NormalLog:
		return "Log"
	case SuccessLog:
		return "Success"
	case ErrorLog:
		return "Error"
	case WarningLog:
		return "Warning"
	case DebugLog:
		return "Debug"
	case InfoLog:
		return "Information"
	default:
		return ""
	}
}

// parseJson converts a map to a JSON string
func convertMapToString(jsonContent map[string]interface{}) string {
	if jsonContent == nil {
		return ""
	}

	jsonData, err := json.Marshal(jsonContent)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

// printLog prints the log to the terminal and file
func (instance *Instance) printLog(logPrefix int, jsonContent map[string]interface{},
	messageContent ...interface{}) {
	var jsonMessage = make(map[string]interface{})
	logTime := generateLogTime()

	var defaultMessage interface{}
	defaultMessage = fmt.Sprint(messageContent...)
	if defaultMessage == "<nil>" {
		defaultMessage = nil
	}

	var defaultJson interface{}
	if jsonContent != nil {
		defaultJson = convertMapToString(jsonContent)
	} else {
		defaultJson = ""
	}

	if instance.FileFormat == JsonFormat || instance.TerminalFormat == JsonFormat {
		jsonMessage["Level"] = returnJsonPrefix(logPrefix)
		jsonMessage["Time"] = logTime

		if messageContent != nil {
			jsonMessage["Message"] = defaultMessage
		} else {
			jsonMessage["Message"] = nil
		}

		if jsonContent != nil {
			jsonMessage["Additional"] = jsonContent
		} else {
			jsonMessage["Additional"] = nil
		}
	}

	if instance.Destination != nil {
		// Line-Delimited JSON method is used here

		if instance.FileFormat == JsonFormat {
			printToFile(instance.Destination, convertMapToString(jsonMessage))
		} else if instance.FileFormat == DefaultFormat {
			if defaultMessage == nil {
				printToFile(instance.Destination, logTime, returnDefaultPrefix(logPrefix, true),
					defaultJson)
			} else {
				printToFile(instance.Destination, logTime, returnDefaultPrefix(logPrefix, true),
					defaultMessage, defaultJson)
			}
		}
	}

	if instance.TerminalFormat == JsonFormat {
		printToTerminal(convertMapToString(jsonMessage))
	} else if instance.TerminalFormat == DefaultFormat {
		if defaultMessage == nil {
			printToTerminal(logTime, returnDefaultPrefix(logPrefix, false),
				defaultJson)
		} else {
			printToTerminal(logTime, returnDefaultPrefix(logPrefix, false),
				defaultMessage, defaultJson)
		}
	}
}

// Log prints the default log to the terminal and file
func (instance *Instance) Log(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(NormalLog, jsonContent, messageContent...)
}

// Success prints the success log to the terminal and file
func (instance *Instance) Success(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(SuccessLog, jsonContent, messageContent...)
}

// Error prints the error log to the terminal and file
func (instance *Instance) Error(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(ErrorLog, jsonContent, messageContent...)
}

// Warning prints the warning log to the terminal and file
func (instance *Instance) Warning(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(WarningLog, jsonContent, messageContent...)
}

// Debug prints the debug log to the terminal and file
func (instance *Instance) Debug(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(DebugLog, jsonContent, messageContent...)
}

// Information prints the information log to the terminal and file
func (instance *Instance) Information(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(InfoLog, jsonContent, messageContent...)
}

// Fatal prints the error log to the terminal and file and exits the program with status code 1
func (instance *Instance) Fatal(jsonContent map[string]interface{}, messageContent ...interface{}) {
	instance.printLog(ErrorLog, jsonContent, messageContent...)
	os.Exit(1)
}
