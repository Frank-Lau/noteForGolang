package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)


func main()  {

	// 组织连接数据库的资源信息
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/mydb2")
	if err != nil {
		fmt.Println("sql.Open err:", err)
		return
	}
	defer db.Close();

	// 测试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping err:", err)
		return
	}

	fmt.Println("连接数据库成功！")

}
