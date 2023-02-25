package main

import (
	"fmt"
	"reflect"
)

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //拿到值对应的类型
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is Float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is Float64,value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float64 = 3.14
	var b int64 = 30
	reflectValue(a)
	reflectValue(b)

	var c float64 = 3.14
	d := reflect.TypeOf(c)
	fmt.Printf("%v\n", d) //float64
	fmt.Printf("%T\n", d) //*reflect.rtype

	var e float64 = 3.14
	f := reflect.ValueOf(e)
	fmt.Printf("%T\n", f) //reflect.Value
}
