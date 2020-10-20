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

/*	// 单行插入

	// 组织sql语句
	sql := `insert into st7 values(1, 'Tom');`

	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("Exec err:", err)
		return
	}
	// 获取影响的行数
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	// 显示sql语句执行结果
	fmt.Printf("insert ok! %d rows Affected.\n", n)*/

	// 多行插入
	// 组织sql语句
	sql := `insert into st7 values(2, 'Tom'), (3, 'kitty'),(4, '李白'),(5, '曹操');`

	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("Exec err:", err)
		return
	}
	// 获取影响的行数
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err:", err)
		return
	}
	// 显示sql语句执行结果
	fmt.Printf("insert ok! %d rows Affected.\n", n)
}
