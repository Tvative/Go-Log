// Warning Message Generation
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package GoLog

// Warning logs a message to the terminal with warning formatting
func (logInstance *LogInstance) Warning(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(false, true, false, MessageWarning, jsonContent, messageContent...)
}

// WarningC logs a message to the terminal with warning formatting
func (logInstance *LogInstance) WarningC(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(false, true, true, MessageWarning, jsonContent, messageContent...)
}

// FWarning logs a warning message to the log file
func (logInstance *LogInstance) FWarning(jsonContent map[string]interface{}, messageContent ...interface{}) {
	logInstance.printOutPut(true, false, false, MessageWarning, jsonContent, messageContent...)
}
