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
	go logData.Print(false, true, false, MessageNormal, messageContent)
}

// LogPC logs a message to the terminal with colored formatting

func (logData LogData) LogPC(messageContent string) {
	go logData.Print(false, true, true, MessageNormal, messageContent)
}

// LogF logs a message to the log file

func (logData LogData) LogF(messageContent string) {
	go logData.Print(true, false, false, MessageNormal, messageContent)
}
