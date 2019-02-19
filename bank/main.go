package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dbinit()
	var ques QuesData

	//打开文件
	file, err := os.OpenFile("ques.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	//关闭文件
	defer file.Close()

	reader := bufio.NewReader(file) //带缓冲区的读写
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		//str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取结束")
			break
		}
		ques.Ques = str

		str, err = reader.ReadString('\n')
		//TODO

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.A = str[:]

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.B = str

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.C = str

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.D = str

		str, err = reader.ReadString('\n')

		fmt.Println(ques)
		InsertUser(ques)
	}
}
