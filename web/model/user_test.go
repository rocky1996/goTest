package model

import (
	"fmt"
	"testing"
)

func TestAddUser1(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &User{}

	//调用添加用户的方法
	user.AddUser1()
	user.AddUser2()
}
