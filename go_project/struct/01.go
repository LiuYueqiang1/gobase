package main

import "fmt"

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

//	 结构体的定义
//		type 类型名 struct {
//		   字段名 字段类型
//		   字段名 字段类型
//		   …
//		}

type person struct {
	name, city string
	age        int
}

func main() {
	var a NewInt = 10
	var b MyInt = 5

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
	fmt.Println(a)                  //10
	fmt.Println(b)                  //5

	var p1 person
	p1.name = "job"
	p1.city = "beijing"
	p1.age = 18
	fmt.Println(p1)         //{job beijing 18}
	fmt.Printf("%#v\n", p1) //main.person{name:"job", city:"beijing", age:18}

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user) //main.person{name:"job", city:"beijing", age:18}
}
