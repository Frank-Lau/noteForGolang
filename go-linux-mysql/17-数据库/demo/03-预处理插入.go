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

	// 准备待插入数据库的数据
	str := [][]string{{"6", "司马懿"}, {"7", "赵云"}, {"8", "孙悟空"}}

	// 预处理 sql语句
	stmt, err := db.Prepare(`insert into st7 values(?, ?)`)
	if err != nil {
		fmt.Println("sql.Open err:", err)
		return
	}
	// 循环从str取数据写入数据库
	for _, data := range str {
		stmt.Exec(data[0], data[1])
	}

	fmt.Println("insert over!")
}
