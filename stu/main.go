package main

import (
	"fmt"
	"os"
)

const (
	CREATE = 1
	UPDATE = 2
	DELETE = 3
	LIST   = 4
	EXIT   = 0
)

var (
	welcome = "欢迎使用学生信息管理系统\n\n"
	menu    = map[int]string{
		CREATE: "添加学生信息",
		UPDATE: "修改学生信息",
		DELETE: "删除学生信息",
		LIST:   "查看学生信息",
		EXIT:   "退出系统",
	}
	action int
)

func main() {

	stuList := newStuList()

	for {
		//菜单
		showMenu()

		switch action {
		case CREATE:
			stu := inputStuInfo()
			stuList.create(stu)
		case UPDATE:
			stu := inputStuInfo()
			stuList.update(stu)
		case DELETE:
			stu := inputStuInfo()
			stuList.delete(stu)
		case LIST:
			stuList.list()
		case EXIT:
			os.Exit(0)
		}
	}
}

func showMenu() {
	//欢迎语
	fmt.Print(welcome)
	//菜单
	for k, v := range menu {
		fmt.Printf("%d:%s\n", k, v)
	}
	//选择菜单
	fmt.Print("输入操作编号:")

	_, _ = fmt.Scanf("%d\n", &action)
}

func inputStuInfo() *Stu {
	var (
		id    int
		name  string
		age   int
		score int
	)
	fmt.Print("输入id:")
	_, _ = fmt.Scanf("%d\n", &id)
	fmt.Print("输入姓名:")
	_, _ = fmt.Scanf("%s\n", &name)
	fmt.Print("输入年龄:")
	_, _ = fmt.Scanf("%d\n", &age)
	fmt.Print("输入分数:")
	_, _ = fmt.Scanf("%d\n", &score)

	return newStu(id, name, age, score)
}
