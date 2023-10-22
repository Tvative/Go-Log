// Fatal Log Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// Fatal logs a message to the terminal with fatal formatting

func (logData *LogInstance) Fatal(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, false, MessageFatal, jsonContent, messageContent...)
}

// FatalC logs a message to the terminal with fatal formatting

func (logData *LogInstance) FatalC(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(false, true, true, MessageFatal, jsonContent, messageContent...)
}

// FFatal logs a fatal message to the log file

func (logData *LogInstance) FFatal(jsonContent map[string]interface{}, messageContent ...any) {
	logData.printOutPut(true, false, false, MessageFatal, jsonContent, messageContent...)
}
