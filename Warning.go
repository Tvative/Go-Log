// Warning Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// Warning logs a message to the terminal with warning formatting

func (logData LogData) Warning(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageWarning, jsonContent, messageContent...)
}

// WarningC logs a message to the terminal with warning formatting

func (logData LogData) WarningC(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, true, MessageWarning, jsonContent, messageContent...)
}

// FWarning logs a warning message to the log file

func (logData LogData) FWarning(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageWarning, jsonContent, messageContent...)
}
