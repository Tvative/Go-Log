// Warning Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// Warning logs a message to the terminal with warning formatting

func (logData LogInstance) Warning(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageWarning, jsonContent, messageContent...)
}

// WarningC logs a message to the terminal with warning formatting

func (logData LogInstance) WarningC(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, true, MessageWarning, jsonContent, messageContent...)
}

// FWarning logs a warning message to the log file

func (logData LogInstance) FWarning(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageWarning, jsonContent, messageContent...)
}
