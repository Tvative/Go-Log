// Log Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// LogP logs a message to the terminal with normal formatting

func (logData LogData) LogP(jsonContent string, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageNormal, jsonContent, messageContent...)
}

// LogF logs a message to the log file

func (logData LogData) LogF(jsonContent string, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageNormal, jsonContent, messageContent...)
}
