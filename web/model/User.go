package model

import (
	"fmt"
	"goTest/web/utils"
)

//User  结构体
type User struct {
	ID            int
	Username      string
	Password      string
	Email         string
}

//AddUser1 添加User的方法1
func (user *User) AddUser1() error{
	//1.写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("预编译出现异常:",err)
		return err
	}

	//3.执行
	_, err2 := inStmt.Exec("admin","123456","wujinfan@100tal.com")
	if err2 != nil{
		fmt.Println("预编译出现异常:",err)
		return err2
	}

	return nil
}

//AddUser2 添加User的方法2
func (user *User) AddUser2() error{
	//1.写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	//2.预编译

	//3.执行
	_, err := utils.Db.Exec(sqlStr,"admin","666666","wujinfan@100tal.com")
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}
	return nil
}
