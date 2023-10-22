// Log Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// A Log logs a message to the terminal with normal formatting
func (logInstance *LogInstance) Log(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(false, true, false, MessageNormal, jsonContent, messageContent...)
}

// A FLog logs a message to the log file
func (logInstance *LogInstance) FLog(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(true, false, false, MessageNormal, jsonContent, messageContent...)
}
