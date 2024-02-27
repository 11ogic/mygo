package algo

import (
	"fmt"
	"reflect"
	"time"
)

func TimeSpent(fn interface{}, params ...interface{}) []reflect.Value {
	function := reflect.ValueOf(fn)
	if function.Kind() != reflect.Func {
		panic("expected a function")
	}

	inner := make([]reflect.Value, len(params))
	for i, val := range params {
		inner[i] = reflect.ValueOf(val)
	}

	start := time.Now()
	result := function.Call(inner)

	elapsed := time.Since(start).Seconds()
	fmt.Printf("elapsed: %.10f result: %v \n", elapsed, result[0])

	return result
}
