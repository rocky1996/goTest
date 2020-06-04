package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao里面的函数")
	t.Run("验证用户名或密码",testLogin)
	t.Run("验证用户名",testRegist)
	t.Run("保存用户",testSave)
}

func testLogin(t *testing.T)  {
	user,_ := CheckUserNameAndPassword("admin","123456")
	fmt.Println("获取的用户信息是",user)
}

func testRegist(t *testing.T)  {
	user,_ := CheckUserName("admin")
	fmt.Println("获取的用户信息是",user)
}

func testSave(t *testing.T)  {
	SaveUser("admin2","123456","dede@1010fal.com")
}