package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func base1() {
	scoremap := make(map[string]int, 8)
	scoremap["张三"] = 90
	scoremap["李四"] = 100
	fmt.Println(scoremap)         //map[张三:90 李四:100]
	fmt.Println(scoremap["张三"])   //90
	fmt.Printf("类型：%T", scoremap) //类型：map[string]int
}

// map也支持在声明的时候填充元素
func base2() {
	scorem := map[string]int{
		"章": 80,
		"里": 60,
	}
	fmt.Println(scorem)
}

// 判断某个键值是否存在
// value,ok:=map[key]
func vaok() {
	scorema := make(map[string]int, 8)
	scorema["zhang"] = 50
	scorema["li"] = 30
	scorema["qi"] = 20
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scorema["zhang"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人！")
	}
}

// map的遍历
func rmap() {
	s := make(map[string]int, 8)
	s["李"] = 20
	s["期"] = 20
	s["是"] = 50
	for k, v := range s {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range s {
		fmt.Println(k)
	}
}

// 使用DELETE函数删除键值
// delete(map,key)
func demap() {
	scoremap := make(map[string]int, 8)
	scoremap["张三"] = 90
	scoremap["李四"] = 100
	scoremap["王五"] = 220
	delete(scoremap, "王五")
	for k, v := range scoremap {
		fmt.Println(k, v)
		// 张三 90
		// 李四 100
	}

}

func randmap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoremap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100)
		scoremap[key] = value
	}
	fmt.Println(scoremap)
	var keys = make([]string, 0, 20)
	for key := range scoremap {
		keys = append(keys, key)
	}

	var ints = make([]int, 0, 20)
	for _, va := range scoremap {
		ints = append(ints, va)
	}
	//对切片进行排序
	sort.Strings(keys)
	sort.Ints(ints)
	for _, k := range keys {
		fmt.Println(k, scoremap[k])
	}
	//目前看来，只能通过键去拿值
}

// map 类型的切片
func slmap() {
	var sliceMap = make([]map[string]int, 3)
	for in, va := range sliceMap {
		fmt.Printf("%d,%v\n", in, va)
		// 0,map[]
		// 1,map[]
		// 2,map[]
	}
	sliceMap[0] = make(map[string]int, 10)
	sliceMap[0]["张三"] = 20
	sliceMap[0]["刘"] = 100
	sliceMap[0]["李"] = 100
	fmt.Println(sliceMap) //[map[刘:100 张三:20 李:100] map[] map[]]
}

// 切片类型的map
func mapsl() {
	var mapSlice = make(map[string][]string)
	// for in, va := range mapSlice {
	// 	fmt.Printf("%v,%v\n", in, va)
	// } 错误的，因为这的map不是切片类型， 切片在map里面，无法索引
	fmt.Println(mapSlice)
	key := "中国"
	value, ok := mapSlice[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	mapSlice[key] = value
	fmt.Println(mapSlice)
}
func main() {
	mapsl()
}
