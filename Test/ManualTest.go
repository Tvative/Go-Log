// Manual Test
//
// Copyright (c) 2023 Tvative
// All Rights Reserved
//
// Use of this source code is governed by
// certain licenses found in the LICENSE file

package main

import (
	"fmt"

	GoLog "github.com/Tvative/Package-Go-Log"
)

var logInstance *GoLog.LogInstance

func main() {
	var fileDestination = "Test/_log.log"
	logInstance, _ = GoLog.Initialize(fileDestination)

	if logInstance.ReturnFile() == nil {
		fmt.Println("File is not available")
	}

	TestLog()

	// Ignore ..
	//
	// TestFatal()

	TestWarning()
}
