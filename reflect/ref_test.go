package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c *Cal) GetSub(name string) {
	fmt.Printf("%v 完成了减法运算，%v - %v = %v", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func TestRef(t *testing.T) {
	cal := Cal{}
	cVal := reflect.ValueOf(&cal)
	cTyp := reflect.TypeOf(&cal)
	for i := 0; i < cVal.Elem().NumField(); i++ {
		fmt.Printf("name: %v, value: %v \n", cTyp.Elem().Field(i).Name, cVal.Elem().Field(i))
	}
	cVal.Elem().FieldByName("Num1").SetInt(8)
	cVal.Elem().FieldByName("Num2").SetInt(5)

	cVal.MethodByName("GetSub").Call([]reflect.Value{
		reflect.ValueOf("Tom"),
	})
}
