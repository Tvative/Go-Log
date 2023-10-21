// Warning Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// WarningP logs a message to the terminal with warning formatting

func (logData LogData) WarningP(jsonContent string, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageWarning, jsonContent, messageContent...)
}

// WarningPC logs a message to the terminal with warning formatting

func (logData LogData) WarningPC(jsonContent string, messageContent ...any) {
	logData.printOutPut(false, true, true, MessageWarning, jsonContent, messageContent...)
}

// WarningF logs a warning message to the log file

func (logData LogData) WarningF(jsonContent string, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageWarning, jsonContent, messageContent...)
}
