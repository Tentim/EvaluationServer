/* 数据库 与用户同一个数据库
create database evaluation;

create table question(
	question_id int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	question varchar(50) NOT NULL unique,
	answer_A varchar(20) NOT NULL,
	answer_B varchar(20) NOT NULL,
	answer_C varchar(20) NOT NULL,
	answer_D varchar(20) NOT NULL
)charset=utf8 engine InnoDB;

insert into question (question, answer_A, answer_B, answer_C, answer_D)
value ( "在以太网中，是根据______地址来区分不同的设备的2.",
		"面向连接的协议",
		"面向非连接的协议",
		"都是传输层协议",
		"以上均不对" );
*/

package main

import (
	// "github.com/go-sql-driver/mysql" 打开数据库,前者是驱动名，所以要导入
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	username = "root"
	passwd   = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "evaluation"
)

// QuesData 用户数据类型
type QuesData struct {
	Quesid int
	Ans    string
	Ques   string
	A      string
	B      string
	C      string
	D      string
}

// DB 数据库连接池
var DB *sql.DB

//初始化
func dbinit() {
	//设置日志格式
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds)

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{username, ":", passwd, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)

	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	//验证连接
	if err := DB.Ping(); err != nil {
		log.Panicln("opon database fail", err)
		return
	}
	log.Println("database connnect success")
}

// InsertUser 插入数据
func InsertUser(ques QuesData) bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return false
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO question (`answer`, `question`, `answer_A`, `answer_B`, `answer_C`, `answer_D`) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare fail", err)
		return false
	}

	//将参数传递到sql语句中并且执
	if _, err := stmt.Exec(ques.Ans, ques.Ques, ques.A, ques.B, ques.C, ques.D); err != nil {
		log.Println("Exec fail", err)
		return false
	}

	//将事务提交
	tx.Commit()

	//获得上一个插入自增的id
	//log.Println(res.LastInsertId())
	return true
}

//Empyt 清空数据库
func Empyt() bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("tx fail", err)
		return false
	}

	//准备sql语句
	stmt, err := tx.Prepare("truncate table question")
	if err != nil {
		log.Println("Prepare fail", err)
		return false
	}

	//将参数传递到sql语句中并且执行
	if _, err := stmt.Exec(); err != nil {
		log.Println("Exec fail", err)
		return false
	}

	//将事务提交
	tx.Commit()

	log.Println("数据表已清空")
	return true
}
