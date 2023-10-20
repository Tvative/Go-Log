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
}

func testLog() {
	logData.LogP("Sample log message 01")
	logData.LogP("Sample log message 02")
	logData.LogP("Sample log message 03")
	logData.LogP("Sample log message 04")

	logData.LogF("Sample log message 01")
	logData.LogF("Sample log message 02")
	logData.LogF("Sample log message 03")
	logData.LogF("Sample log message 04")
}
