// Warning Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// WarningP logs a message to the terminal with warning formatting

func (logData LogData) WarningP(messageContent string) {
	logData.printOutPut(false, true, false, MessageWarning, messageContent)
}

// WarningPC logs a message to the terminal with warning formatting

func (logData LogData) WarningPC(messageContent string) {
	logData.printOutPut(false, true, true, MessageWarning, messageContent)
}

// WarningF logs a warning message to the log file

func (logData LogData) WarningF(messageContent string) {
	logData.printOutPut(true, false, false, MessageWarning, messageContent)
}
