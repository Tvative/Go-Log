// Warning Message Manual Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

func TestWarning() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logInstance.Warning(nil, "Sample warning log message 01")
	logInstance.Warning(jsonString, "Sample warning log message 02")
	logInstance.Warning(nil, "Sample warning log message 03")
	logInstance.Warning(jsonString, "Sample warning log message 04")

	logInstance.WarningC(nil, "Sample warning log message 01")
	logInstance.WarningC(jsonString, "Sample warning log message 02")
	logInstance.WarningC(nil, "Sample warning log message 03")
	logInstance.WarningC(jsonString, "Sample warning log message 04")

	logInstance.FWarning(nil, "Sample warning log message 01")
	logInstance.FWarning(jsonString, "Sample warning log message 02")
	logInstance.FWarning(nil, "Sample warning log message 03")
	logInstance.FWarning(jsonString, "Sample warning log message 04")
}
