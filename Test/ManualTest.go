// Unit Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

import goLog "github.com/Tvative/Go-Log"

var logData *goLog.LogData

func main() {
	logData = &goLog.LogData{}

	var fileDestination = "Test/_log.log"
	logData.Initialize(fileDestination)

	testLog()

	// Ignore ..
	//
	// testFatal()

	testWarning()
}

func testLog() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logData.Log(nil, "Sample log message 01", "111")
	logData.Log(jsonString, "Sample log message 02")
	logData.Log(nil, "Sample log message 03")
	logData.Log(jsonString, "Sample log message 04")

	logData.FLog(nil, "Sample log message 01")
	logData.FLog(jsonString, "Sample log message 02")
	logData.FLog(nil, "Sample log message 03")
	logData.FLog(jsonString, "Sample log message 04")
}

func testFatal() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logData.Fatal(nil, "Sample fatal log message 01")
	logData.Fatal(jsonString, "Sample fatal log message 02")
	logData.Fatal(nil, "Sample fatal log message 03")
	logData.Fatal(jsonString, "Sample fatal log message 04")

	logData.FatalC(nil, "Sample fatal log message 01")
	logData.FatalC(jsonString, "Sample fatal log message 02")
	logData.FatalC(nil, "Sample fatal log message 03")
	logData.FatalC(jsonString, "Sample fatal log message 04")

	logData.FFatal(nil, "Sample fatal log message 01")
	logData.FFatal(jsonString, "Sample fatal log message 02")
	logData.FFatal(nil, "Sample fatal log message 03")
	logData.FFatal(jsonString, "Sample fatal log message 04")
}

func testWarning() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logData.Warning(nil, "Sample warning log message 01")
	logData.Warning(jsonString, "Sample warning log message 02")
	logData.Warning(nil, "Sample warning log message 03")
	logData.Warning(jsonString, "Sample warning log message 04")

	logData.WarningC(nil, "Sample warning log message 01")
	logData.WarningC(jsonString, "Sample warning log message 02")
	logData.WarningC(nil, "Sample warning log message 03")
	logData.WarningC(jsonString, "Sample warning log message 04")

	logData.FWarning(nil, "Sample warning log message 01")
	logData.FWarning(jsonString, "Sample warning log message 02")
	logData.FWarning(nil, "Sample warning log message 03")
	logData.FWarning(jsonString, "Sample warning log message 04")
}
