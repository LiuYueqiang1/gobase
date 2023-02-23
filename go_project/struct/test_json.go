package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	ID   int
}
type Class struct {
	Title    string `json:"title"`
	Students []*Student
}

func main() {
	c1 := &Class{
		Title:    "101班",
		Students: make([]*Student, 0, 100),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:   i,
			Name: fmt.Sprintf("stu%02d", i),
		}
		c1.Students = append(c1.Students, stu)
	}
	fmt.Println(*c1)
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &Class{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c2)
}
{101班 [0xc000092060 0xc000092078 0xc000092090 0xc0000920a8 0xc0000920c0 0xc0000
920d8 0xc0000920f0 0xc000092108 0xc000092120 0xc000092138]}

json:{"Title":"101班","Students":[{"Name":"stu00","ID":0},{"Name":"stu01","ID":1
},{"Name":"stu02","ID":2},{"Name":"stu03","ID":3},{"Name":"stu04","ID":4},{"Name
":"stu05","ID":5},{"Name":"stu06","ID":6},{"Name":"stu07","ID":7},{"Name":"stu08
","ID":8},{"Name":"stu09","ID":9}]}

&main.Class{Title:"101", Students:[]*main.Student{(*main.Student)(0xc0000922a0),
(*main.Student)(0xc0000922b8), (*main.Student)(0xc0000922d0), (*main.Student)(0
xc0000922e8), (*main.Student)(0xc000092318), (*main.Student)(0xc000092330), (*ma
in.Student)(0xc000092360), (*main.Student)(0xc000092378), (*main.Student)(0xc000
092390), (*main.Student)(0xc0000923c0)}}