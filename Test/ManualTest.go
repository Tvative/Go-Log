// Manual Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

import "github.com/Tvative/Go-Log"

var logInstance *GoLog.LogInstance

func main() {
	var fileDestination = "Test/_log.log"
	logInstance = GoLog.Initialize(fileDestination)

	TestLog()

	// Ignore ..
	//
	// TestFatal()

	TestWarning()
}
