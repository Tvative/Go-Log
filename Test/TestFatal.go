// Fatal Message Manual Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

func TestFatal() {
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
