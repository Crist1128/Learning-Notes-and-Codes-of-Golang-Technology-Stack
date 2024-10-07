/**
 * @File : main.go
 * @Description : 请填写文件描述
 * @Author : 请填写作者的真实姓名
 * @Date : 2024-09-19
 */
package main

import "fmt"

type Employee interface {
	Work()
}

type Manager interface {
	Employee
	Manage()
}

type Developer struct {
}

type TeamLeader struct {
}

func (d Developer) Work() {
	//TODO implement me
	fmt.Println("Developer is working...")
}

func (t TeamLeader) Work() {
	//TODO implement me
	fmt.Println("TeamLeader is working...")
}

func (t TeamLeader) Manage() {
	//TODO implement me
	fmt.Println("TeamLeader is managing the team...")
}

func AssignWork(i interface{}) {
	if e, ok := i.(Employee); ok {
		e.Work()
	}
	if m, ok := i.(Manager); ok {
		m.Manage()
	}

}

func main() {
	d := Developer{}
	t := TeamLeader{}
	e := d
	m := t
	AssignWork(e)
	AssignWork(m)
}
