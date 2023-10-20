// Unit Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

import goLog "github.com/Tvative/Go-Log/Source"

var logData *goLog.LogData

func main() {
	logData = &goLog.LogData{}

	var fileDestination = "Test/_log.log"
	logData.Initialize(fileDestination)

	testLog()
}

func testLog() {
}
