package main

import "fmt"

type student struct {
	Name  string
	ID    int
	Age   int
	Class string
}

// 学生的构造函数
func newStudent(name, class string, id int, age int) *student {
	return &student{
		Name:  name,
		ID:    id,
		Age:   age,
		Class: class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allstudents []*student
}

// studentMgr类型的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allstudents: make([]*student, 0, 100),
	}
}

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	   函数体
//	}
//
// 1.添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allstudents = append(s.allstudents, newStu)
}

// 2.编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allstudents {
		if newStu.ID == v.ID {
			s.allstudents[i] = newStu
			return
		}
	}
	fmt.Printf("输入的学生有误，没有找到学号为%d的学生", newStu.ID)
}

// 3.查看学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allstudents {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 班级：%s\n", v.ID, v.Name, v.Age, v.Class)
	}
}
