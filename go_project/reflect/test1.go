package main

import (
	"fmt"
	"reflect"
)

//	func 函数名(参数)(返回值){
//	   函数体
//	}
//
// func 函数名(参数 参数类型) 无返回值
// 使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v,type:%v\n", v.Name(), v.Kind())
}

// 类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，
// 而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
type myInt int64
type person struct {
	name string
	age  int
}

func main() {
	var a float32 = 3.14
	reflectType(a) //type:float32,type:float32
	var b int64 = 100
	reflectType(b) //type:int64,type:int64
	var c myInt = 50
	reflectType(c) //type:myInt,type:int64
	var d *float32
	reflectType(d) //type:,type:ptr

	var p = person{
		name: "alen",
		age:  16,
	}
	reflectType(p) //type:person,type:struct
	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
}
