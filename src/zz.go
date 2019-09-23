package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,err := sql.Open("mysql","root:root@(127.0.0.1:3306)/wsktest")
	checkerr(err)

	//添加数据
	stmt,err := db.Prepare("INSERT INTO music SET name=?,year=?,singerId=?")
	checkerr(err)

	res,err := stmt.Exec("说好不哭",2019,1)
	checkerr(err)

	fmt.Println(res)

	//修改数据
	stme,err := db.Prepare("UPDATE MUSIC SET NAME = ?,YEAR=?,SINGERID = ? WHERE ID = ?")
	checkerr(err)

	rese,err := stme.Exec("李香兰",1999,2,2)
	checkerr(err)

	id,err := rese.RowsAffected()
	checkerr(err)

	fmt.Println(id)

	//删除数据
	stmr,err := db.Prepare("DELETE FROM music where id = ?")
	checkerr(err)

	resq,err := stmr.Exec(9)
	checkerr(err)

	id01,err := resq.RowsAffected()
	checkerr(err)

	fmt.Println(id01)

	//查询全部数据
	rows,err := db.Query("select * from music m left join singer s on m.singerId = s.sid")
	checkerr(err)

	for rows.Next(){
		var id int
		var name string
		var year int
		var singerId int
		var sid int
		var sname string
		err := rows.Scan(&id,&name,&year,&singerId,&sid,&sname)
		checkerr(err)
		fmt.Printf("id是%d,name是%s,year是%d,singerId是%d,sid是%d,sname是%s\n",id,name,year,singerId,sid,sname)
	}
}

func checkerr(err error) {
	if err != nil{
		panic(err)
	}
}
