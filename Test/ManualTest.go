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
	logData.LogP("Sample log message 01", "111")
	logData.LogP("Sample log message 02")
	logData.LogP("Sample log message 03")
	logData.LogP("Sample log message 04")

	logData.LogF("Sample log message 01")
	logData.LogF("Sample log message 02")
	logData.LogF("Sample log message 03")
	logData.LogF("Sample log message 04")
}

func testFatal() {
	logData.FatalP("Sample fatal log message 01")
	logData.FatalP("Sample fatal log message 02")
	logData.FatalP("Sample fatal log message 03")
	logData.FatalP("Sample fatal log message 04")

	logData.FatalPC("Sample fatal log message 01")
	logData.FatalPC("Sample fatal log message 02")
	logData.FatalPC("Sample fatal log message 03")
	logData.FatalPC("Sample fatal log message 04")

	logData.FatalF("Sample fatal log message 01")
	logData.FatalF("Sample fatal log message 02")
	logData.FatalF("Sample fatal log message 03")
	logData.FatalF("Sample fatal log message 04")
}

func testWarning() {
	logData.WarningP("Sample warning log message 01")
	logData.WarningP("Sample warning log message 02")
	logData.WarningP("Sample warning log message 03")
	logData.WarningP("Sample warning log message 04")

	logData.WarningPC("Sample warning log message 01")
	logData.WarningPC("Sample warning log message 02")
	logData.WarningPC("Sample warning log message 03")
	logData.WarningPC("Sample warning log message 04")

	logData.WarningF("Sample warning log message 01")
	logData.WarningF("Sample warning log message 02")
	logData.WarningF("Sample warning log message 03")
	logData.WarningF("Sample warning log message 04")
}
