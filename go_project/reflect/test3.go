package main

import (
	"fmt"
	"reflect"
)

//通过反射修改变量的值
//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
//而反射中使用专有的Elem()方法来获取指针对应的值

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}

}
func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(300)
	}
}

func main() {
	var a int64 = 100
	//reflectSetValue(a)
	fmt.Println(a)
	reflectValue2(&a)
	fmt.Println(a)

	//IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	// *int类型空指针
	var b *int
	fmt.Println("var a *int IsNil", reflect.ValueOf(b).IsNil())
	//nil
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	//实例化一个匿名结构体
	var c = struct{}{}
	//尝试从结构体中查找“abc”字段
	fmt.Println("不存在的结构体成员", reflect.ValueOf(c).FieldByName("abc").IsNil())
	//尝试从结构体中查找“abc”方法
	fmt.Println("不存在的结构体成员", reflect.ValueOf(c).MethodByName("abc").IsNil())
	//map
	d := map[string]int{}
	//尝试从map中寻找一个不存在的键
	fmt.Println("map中不存在的键", reflect.ValueOf(d).MapIndex(reflect.ValueOf("纳扎")).IsValid())
}
