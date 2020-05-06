package main

import "fmt"

type Stu struct {
	Id    int
	Name  string
	Age   int
	Score int
}

func newStu(id int, name string, age int, score int) *Stu {
	return &Stu{
		id,
		name,
		age,
		score,
	}
}

type StuList struct {
	StuList []*Stu
}

func newStuList() *StuList {
	return &StuList{
		StuList: make([]*Stu, 0, 200),
	}
}

//添加
func (s *StuList) create(stu *Stu) {
	s.StuList = append(s.StuList, stu)
}

func (s *StuList) update(stu *Stu) {
	for k, v := range s.StuList {
		if v.Id == stu.Id {
			s.StuList[k] = stu
			return
		}
	}
	fmt.Printf("未找到学生 %v", stu.Name)
}

func (s *StuList) delete(stu *Stu) {
	for k, v := range s.StuList {
		if v.Id == stu.Id {
			s.StuList = append(s.StuList[:k], s.StuList[k+1:]...)
			return
		}
	}
	fmt.Printf("未找到学生 %v", stu.Name)
}

func (s *StuList) list() {
	for _, v := range s.StuList {
		fmt.Printf("id:%d 姓名:%s 年龄:%d 分数:%d\n", v.Id, v.Name, v.Age, v.Score)
	}
}
