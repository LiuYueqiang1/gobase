package main

import "fmt"

//1.空接口可以作为函数的参数
//2.空接口可以作为map的value

func main() {
	var x interface{} //定义了一个空接口变量x
	x = "Da"
	fmt.Println(x)
	x = 20
	fmt.Println(x)
	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "纳扎"
	m["age"] = 16
	m["city"] = "北京"
	m["hobby"] = []string{"篮球", "足球", "乒乓球"}
	fmt.Println(m)

	ret, ok := x.(string)
	if !ok {
		fmt.Println("不是string类型")
	} else {
		fmt.Println("是字符串类型", ret)
	}

	//使用switch判断
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型,为%v\n", v)
	case int:
		fmt.Printf("是整型,为%v\n", v)
	case bool:
		fmt.Printf("是bool类型,为%v\n", v)
	default:
		fmt.Println("类型未知")
	}
}
