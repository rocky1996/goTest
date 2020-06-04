package dao

import (
	"goTest/bookstore/model"
	"goTest/bookstore/utils"
)

//CheckUserNameAndPassword，根据用户名，密码返回数据库的记录
func CheckUserNameAndPassword(username string,password string) (*model.User, error)  {
	//sql语句
	sqlStr := "select id,username,password,email from users where username=? and password=?"
	//执行
	row := utils.Db.QueryRow(sqlStr,username,password)
	user := &model.User{}

	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user,nil
}

//检查用户名
func CheckUserName(username string) (*model.User, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where username=?"
	//执行
	row := utils.Db.QueryRow(sqlStr,username)
	user := &model.User{}

	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user,nil
}

//保存用户的
func SaveUser(username string,password string,email string) (error) {
	//sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_,err := utils.Db.Exec(sqlStr,username,password,email)
	if err != nil{
		return err
	}
	return nil
}
