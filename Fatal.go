// Fatal Log Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// FatalP logs a message to the terminal with fatal formatting

func (logData LogData) FatalP(jsonContent string, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageFatal, jsonContent, messageContent...)
}

// FatalPC logs a message to the terminal with fatal formatting

func (logData LogData) FatalPC(jsonContent string, messageContent ...any) {
	logData.printOutPut(false, true, true, MessageFatal, jsonContent, messageContent...)
}

// FatalF logs a fatal message to the log file

func (logData LogData) FatalF(jsonContent string, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageFatal, jsonContent, messageContent...)
}
