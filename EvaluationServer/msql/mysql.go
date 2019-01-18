/* 数据库
create database evaluation;

create table user(
	userid int AUTO_INCREMENT NOT NULL PRIMARY KEY,
	username varchar(20) NOT NULL unique,
	password varchar(20) NOT NULL
)charset=utf8 engine InnoDB;


UPDATE user set password='132456' WHERE username='iii';

*/

package msql

import (
	"database/sql"
	"log"
	"strings"

	// "github.com/go-sql-driver/mysql" 打开数据库,前者是驱动名，所以要导入
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

// UserData 用户数据类型
type UserData struct {
	Userid   int
	Username string
	Password string
}

// DB 数据库连接池
var DB *sql.DB

func init() {

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
		log.Panicln("opon database fail")
		return
	}
	log.Println("database connnect success")
}

// InsertUser 插入数据
func InsertUser(user UserData) bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("tx fail")
		return false
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`username`, `password`) VALUES (?, ?)")
	if err != nil {
		log.Println("Prepare fail")
		return false
	}

	//将参数传递到sql语句中并且执
	if _, err := stmt.Exec(user.Username, user.Password); err != nil {
		log.Println("Exec fail")
		return false
	}

	//将事务提交
	tx.Commit()

	//获得上一个插入自增的id
	//log.Println(res.LastInsertId())
	return true
}

// UpdateUser 更新数据
func UpdateUser(name string, passwd string) bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	//准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET password = ? WHERE username = ?")
	if err != nil {
		log.Println("Prepare fail")
		return false
	}

	//设置参数以及执行sql语句
	if _, err = stmt.Exec(passwd, name); err != nil {
		log.Println("Exec fail")
		return false
	}

	//提交事务
	tx.Commit()

	return true
}

// DeleteData 删除用户
func DeleteData(username string) bool {

	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		log.Println("tx fail")
	}

	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		log.Println("Prepare fail")
		return false
	}

	//设置参数以及执行sql语句
	if _, err = stmt.Exec(username); err != nil {
		log.Println("Exec fail")
		return false
	}

	//提交事务
	tx.Commit()

	return true
}

// SelectUserByname 读取数据
func SelectUserByname(name string) (UserData, bool) {
	var User UserData
	err := DB.QueryRow("SELECT * FROM user WHERE username = ?", name).Scan(&User.Userid, &User.Username, &User.Password)
	if err != nil {
		log.Println(name, "用户不存在")
		return User, false
	}
	return User, true
}

// IsUserByname 查看用户是否存在
func IsUserByname(name string) bool {
	if err := DB.QueryRow("SELECT userid FROM user WHERE username = ?", name); err != nil {
		return false
	}
	return true
}

// IsPasswdTrueByUsername 通过用户名查询密码是否正确
func IsPasswdTrueByUsername(name string, passwd string) bool {
	var slqpasswd string
	err := DB.QueryRow("SELECT password FROM user WHERE username = ?", name).Scan(&slqpasswd)
	if err != nil {
		return false
	}
	return passwd == slqpasswd
}
