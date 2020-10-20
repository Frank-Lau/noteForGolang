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

/*	// 查询单行
	row := db.QueryRow(`select * from st7 where id=5`)
	// 定义变量
	var id, name string

	row.Scan(&id, &name)

	fmt.Println(id, "-", name)*/

/*	// 多行查询
	rows, err := db.Query(`select * from st7 where id >=3`)
	if err != nil {
		fmt.Println("Query err:", err)
		return
	}
	var id, name string
	// 将游标 下移，取一条数据，不能 next 时候，结束循环
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, "-", name)
	}*/

	// 预处理查询
	stmt, _:= db.Prepare(`select * from st7 where id >= ?`)

	rows, _:= stmt.Query(3)

	var id, name string

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, "-", name)
	}
}