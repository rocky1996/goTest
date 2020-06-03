package model

import (
	"fmt"
	"testing"
)

//如果函数名不是以Test开头,默认不执行
//func TestAddUser1(t *testing.T) {
//	fmt.Println("测试添加用户")
//	user := &User{}
//
//	//调用添加用户的方法
//	user.AddUser1()
//	user.AddUser2()
//}

func TestGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := &User{
		ID:2,
	}

	u,_ := user.GetUserByID()
	fmt.Println("得到的User的信息是:",u)
}