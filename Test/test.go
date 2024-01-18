package main

import (
	"github.com/Tvative/Package-Go-Log/v2"
)

var instance *golog.Instance

func main() {
	instance = golog.Initialize()

	instance.SetFile("Test/test.log")
	instance.SetFileFormat(golog.JsonFormat)
	instance.SetTerminalFormat(golog.JsonFormat)

	var moreJson = map[string]interface{}{
		"Sample": "Content",
	}

	var testJson = map[string]interface{}{
		"Sample": "Content",
		"More":   moreJson,
	}

	instance.Log(testJson, "Default Log")
	instance.Success(testJson, "Success Log")
	instance.Error(testJson, "Error Log")
	instance.Warning(testJson, "Warning Log")
	instance.Debug(testJson, "Debug Log")
	instance.Information(testJson, "Information Log")
	// instance.Fatal(testJson, "Fatal Log")

	instance.Log(testJson, nil)
	instance.Log(nil, "Default Log")
}
