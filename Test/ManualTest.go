// Unit Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

import goLog "github.com/Tvative/Go-Log"

var logInstance *goLog.LogInstance

func main() {
	var fileDestination = "Test/_log.log"
	logInstance = goLog.Initialize(fileDestination)

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

	logInstance.Log(nil, "Sample log message 01", "111")
	logInstance.Log(jsonString, "Sample log message 02")
	logInstance.Log(nil, "Sample log message 03")
	logInstance.Log(jsonString, "Sample log message 04")

	logInstance.FLog(nil, "Sample log message 01")
	logInstance.FLog(jsonString, "Sample log message 02")
	logInstance.FLog(nil, "Sample log message 03")
	logInstance.FLog(jsonString, "Sample log message 04")
}

func testFatal() {
	jsonString := map[string]interface{}{
		"key_01": "a",
		"key_02": 1,
		"key_03": "B",
	}

	logInstance.Fatal(nil, "Sample fatal log message 01")
	logInstance.Fatal(jsonString, "Sample fatal log message 02")
	logInstance.Fatal(nil, "Sample fatal log message 03")
	logInstance.Fatal(jsonString, "Sample fatal log message 04")

	logInstance.FatalC(nil, "Sample fatal log message 01")
	logInstance.FatalC(jsonString, "Sample fatal log message 02")
	logInstance.FatalC(nil, "Sample fatal log message 03")
	logInstance.FatalC(jsonString, "Sample fatal log message 04")

	logInstance.FFatal(nil, "Sample fatal log message 01")
	logInstance.FFatal(jsonString, "Sample fatal log message 02")
	logInstance.FFatal(nil, "Sample fatal log message 03")
	logInstance.FFatal(jsonString, "Sample fatal log message 04")
}

func testWarning() {
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
