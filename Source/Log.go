// Log Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// LogP logs a message to the terminal with normal formatting

func (logData LogData) LogP(messageContent string) {
	logData.Print(false, true, false, MessageNormal, messageContent)
}

// LogF logs a message to the log file

func (logData LogData) LogF(messageContent string) {
	logData.Print(true, false, false, MessageNormal, messageContent)
}
