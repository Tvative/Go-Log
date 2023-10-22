// Log Message Manual Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

func TestLog() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logInstance.Log(nil, "Sample log message 01", " 111")
	logInstance.Log(jsonString, "Sample log message 02")
	logInstance.Log(nil, "Sample log message 03")
	logInstance.Log(jsonString, "Sample log message 04")

	logInstance.FLog(nil, "Sample log message 01")
	logInstance.FLog(jsonString, "Sample log message 02")
	logInstance.FLog(nil, "Sample log message 03")
	logInstance.FLog(jsonString, "Sample log message 04")
}
