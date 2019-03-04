package bank

import (
	"EvaluationServer/msql"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//UpDataQues 更新题库
func UpDataQues() {

	//打开文件
	file, err := os.OpenFile("../bank/ques.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	//关闭文件
	defer file.Close()

	//清空数据库
	msql.Empyt()

	var ques msql.QuesData
	reader := bufio.NewReader(file) //带缓冲区的读写
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		if err != nil {
			log.Println("题库已更新")
			break
		}
		ques.Ques = str

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.Ans = str

		str, err = reader.ReadString('\n')
		str = strings.TrimSpace(str)
		ques.A = str

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

		//fmt.Println(ques)
		msql.InsertQues(ques)
	}

}

//GenerateRandomNumber 生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int32) []int {
	//范围检查
	if end < start || (end-start) < int(count) {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)

	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < int(count) {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
