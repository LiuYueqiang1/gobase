package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON序列化
// 键/值对组合中的键名写在前面并用双引号""包裹，
// 使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔
type student struct {
	ID   int
	Name string
}
type class struct {
	Title   string `json:"title"`
	Stuents []student
}

func newStudent(name string, id int) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {

	c1 := class{
		Title:   "火箭班",
		Stuents: make([]student, 0, 8),
	}
	for i := 0; i < 10; i++ {
		tmpstu := newStudent(fmt.Sprintf("%02d", i), i)
		c1.Stuents = append(c1.Stuents, tmpstu)
	}
	fmt.Printf("%#v\n", c1)
	//JSON序列化：Go语言中的数据 ->JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal is failed, err", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//JSON反序列化：JSON格式的字符串->Go语言中的数据
	jsonStr := `{"Title":"火箭班","Stuents":[{"ID":0,"Name":"00"},{"ID":1,"Name":"01"},{"ID":2,"Name":"02"},{"ID":3,"Name":"03"}]}`

	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json marshal is failed, err", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
