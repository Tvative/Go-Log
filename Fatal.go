// Fatal Log Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// FatalP logs a message to the terminal with fatal formatting

func (logData LogData) FatalP(messageContent ...any) {
	logData.printOutPut(false, true, false, MessageFatal, messageContent...)
}

// FatalPC logs a message to the terminal with fatal formatting

func (logData LogData) FatalPC(messageContent ...any) {
	logData.printOutPut(false, true, true, MessageFatal, messageContent...)
}

// FatalF logs a fatal message to the log file

func (logData LogData) FatalF(messageContent ...any) {
	logData.printOutPut(true, false, false, MessageFatal, messageContent...)
}
