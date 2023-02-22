package main

import "fmt"

// 因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。
type Person struct {
	name   string
	age    int8
	dreams []string
}

//func (p *Person) SetDreams(dreams []string) {
//	p.dreams = dreams
//}

// 正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
func main() {
	p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // [吃饭 睡觉 打豆豆]
	fmt.Println(data)      //[吃饭 不睡觉 打豆豆]
}
