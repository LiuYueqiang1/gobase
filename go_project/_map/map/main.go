package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//map[KeyType]ValueType

func basic() { //初始化map
	scoremap := make(map[string]int, 8)
	scoremap["张三"] = 90
	scoremap["李四"] = 80
	fmt.Println(scoremap)        //map[张三:90 李四:80]
	fmt.Println(scoremap["张三"])  //90
	fmt.Printf("%T\n", scoremap) //map[string]int
}
func basic2() { //初始化map
	userinfo := map[string]string{
		"username": "大帅哥",
		"password": "帅煞",
	}
	fmt.Println(userinfo)
}

// 判断 map中键是否存在的特殊写法，格式
// value, ok := map[key]
func okmap() {
	scmap := map[string]int{
		"土": 20,
		"够": 30,
	}
	v, ok := scmap["土"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人！")
	}

}

// 写法二make构造 value，ok
func okmap2() {
	smap := make(map[string]int)
	smap["小明"] = 90
	smap["李三"] = 80
	value, ok := smap["小明"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}
}

// forrange遍历map
func formap() {
	mapname := map[string]int{
		"刘": 100,
		"张": 60,
		"李": 80,
	}
	for k, v := range mapname {
		fmt.Println(k, v)
	}
	for k2 := range mapname { //只遍历k值
		fmt.Println(k2)
	}
}

// delete(map,key)
func delmap() {
	score := make(map[string]int)
	score["张三"] = 50
	score["里斯"] = 60
	score["帅额"] = 80
	delete(score, "张三")
	for key, value := range score {
		fmt.Println(key, value)
	}
}

// 按照指定顺序遍历map
func randmap() {
	var scoreMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("st%2d", i)
		value := rand.Intn(100) //每次循环生成0-100的随机整数
		scoreMap[key] = value
	}
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}

	//对这些进行排序
	var skey = make([]string, 0, 100)

	for k := range scoreMap {
		skey = append(skey, k)
	}
	sort.Strings(skey)

	for _, va := range skey {

		fmt.Println(va, scoreMap[va])
	}
}

// 切片中的元素为map
func msilce() {
	var mapsilec = make([]map[string]string, 3)

	mapsilec[0] = make(map[string]string, 3)
	mapsilec[0]["name"] = "小王子"
	mapsilec[0]["name"] = "小王子"
	mapsilec[0]["password"] = "123456"
	mapsilec[0]["address"] = "沙河"

	for _, v := range mapsilec {
		fmt.Println(v)
	}
}

// map中值为切片类型
func slmap() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}

func test() {
	var a []string
	var b = []string{}
	var c = make([]string, 3)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
}
func main() {
	test()
}
