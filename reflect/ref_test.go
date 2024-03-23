package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type myInt int8

func test(i interface{}) {
	rVal := reflect.ValueOf(i)
	rTyp := reflect.TypeOf(i)
	//num := rVal.Int()
	//typ := rTyp.Kind()
	fmt.Printf("type: %T, value: %v \n", rVal, rVal)
	fmt.Printf("type: %T, value: %v \n", rTyp, rTyp)
}

func TestRef(t *testing.T) {
	var num myInt = 1
	test(num)
}
