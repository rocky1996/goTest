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
	_, err2 := inStmt.Exec("admin","123456","rocky@qq.com")
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
	_, err := utils.Db.Exec(sqlStr,"admin","666666","rocky@qq.com")
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return err
	}
	return nil
}

//获取一条记录
func (user *User) GetUserByID() (*User,error){
	//sql语句
	sqlStr := "select id,username,password,email from users where id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr,user.ID)
	//声明
	var id int
	var username string
	var password string
	var email string

	err := row.Scan(&id,&username,&password,&email)
	if err != nil{
		return nil,err
	}

	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u,err
}

//查询全部的记录
func (user *User) GetUsers() ([]*User,error) {
	//sql语句
	sqlStr := "select id,username,password,email from users"

	//执行
	rows,err := utils.Db.Query(sqlStr)
	if err != nil{
		return nil,err
	}
	//创建User切片
	var users []*User
	for rows.Next(){
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id,&username,&password,&email)
		if err != nil{
			return nil,err
		}

		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users,nil
}