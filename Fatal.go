// Fatal Log Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// Fatal logs a message to the terminal with fatal formatting
func (logInstance *LogInstance) Fatal(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(false, true, false, MessageFatal, jsonContent, messageContent...)
}

// FatalC logs a message to the terminal with fatal formatting
func (logInstance *LogInstance) FatalC(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(false, true, true, MessageFatal, jsonContent, messageContent...)
}

// FFatal logs a fatal message to the log file
func (logInstance *LogInstance) FFatal(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(true, false, false, MessageFatal, jsonContent, messageContent...)
}
