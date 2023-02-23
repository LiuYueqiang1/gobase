package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")
}
func getInput() *student {
	var (
		Name  string
		ID    int
		Age   int
		Class string
	)
	fmt.Println("请输入学员的名字")
	fmt.Scanf("%s\n", &Name)
	fmt.Println("请输入学员的学号")
	fmt.Scanf("%d\n", &ID)
	fmt.Println("请输入学员的年龄")
	fmt.Scanf("%d\n", &Age)
	fmt.Println("请输入学员的班级")
	fmt.Scanf("%s\n", &Class)
	stu := newStudent(Name, Class, ID, Age)
	return stu
}

func main() {
	//打印系统菜单
	sm := newStudentMgr()
	//等待用户要执行的选项
	for {
		showMenu()
		fmt.Printf("请输入你要执行的选项:")
		var input int
		fmt.Scanf("%d\n", &input)
		fmt.Printf("%d\n", input)

		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("输出无效")
		}
	}
}
