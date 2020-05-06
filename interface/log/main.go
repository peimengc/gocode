package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(log string)
}

type FileLog struct {
	filename string
}

type ConsoleLog struct {
}

func (c *ConsoleLog) Info(s string) {
	fmt.Print(s)
}

func (fl *FileLog) Info(s string) {
	//打开文件,不存在创建
	f, err := os.OpenFile(fl.filename, os.O_CREATE|os.O_APPEND, 0664)

	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	defer f.Close()
	//写入字符串
	_, err = f.WriteString(s)

	if err != nil {
		fmt.Println("文件写入失败")
		return
	}
}

func main() {
	var log Logger
	log = &ConsoleLog{}
	log.Info("写入日志\n")

	log = &FileLog{filename: "./log.txt"}
	log.Info("写入日志\n")
}
