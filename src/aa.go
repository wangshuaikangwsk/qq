package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	daba,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/wsktest")
	checkErr(err)

	//添加数据
	//stmt,err :=daba.Prepare("INSERT INTO music set name=?,year=?,singerId=?")
	//checkErr(err)
	//
	//res,err :=stmt.Exec("稻香",5000,1)
	//checkErr(err)
	//
	//fmt.Println(res)

	//修改数据
	//stmt,err :=daba.Prepare("UPDATE  music set name=?,year=? where id=?")
	//checkErr(err)
	//
	//res,err :=stmt.Exec("南方姑娘",2000,4)
	//checkErr(err)
	//
	//fmt.Println(res)


	//删除数据
	stmt,err :=daba.Prepare("DELETE FROM music where id=?")
	checkErr(err)

	id,err :=stmt.Exec(3)
	checkErr(err)

	fmt.Println(id)




	//查询全部数据
	rows,err := daba.Query("SELECT m.*,s.* FROM music m LEFT JOIN singer s ON m.singerId = s.sid")
	checkErr(err)

	for rows.Next(){
		var id int
		var name string
		var year int
		var singerId int
		var sid int
		var sname string

		err := rows.Scan(&id,&name,&year,&singerId,&sid,&sname)
		checkErr(err)
		fmt.Printf("id是:%d,name是:%s,year是:%d,singerId是:%d,sid是:%d,sname是:%s\n",id,name,year,singerId,sid,sname)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
