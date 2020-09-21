package loggingapi

import (
	"reflect"
)

func testLoggerMethods(logger Logger) {
	inputParams := []interface{}{"logTag", "message: %s", "haha"}
	input := make([]reflect.Value, len(inputParams))
	for i, param := range inputParams {
		input[i] = reflect.ValueOf(param)
	}

	testLogger := reflect.ValueOf(logger)
	for i := 0; i < testLogger.NumMethod(); i++ {
		method := testLogger.Method(i)
		method.Call(input)
	}
}
