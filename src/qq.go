package main
//
//import (
//	"database/sql"
//	"fmt"
//	_"github.com/go-sql-driver/mysql"
//)
//
//func main() {
//	db,err := sql.Open("mysql","root:root@/wsktest")
//	checkerr(err)
//
//	//添加数据
//	//stem,err := db.Prepare("INSERT INTO music set name=?,year=?,singerId=?")
//	//checkerr(err)
//	//
//	//rest,err :=stem.Exec("Joy",3000,1)
//	//checkerr(err)
//	//
//	//la,err :=rest.LastInsertId()
//	//checkerr(err)
//	//
//	//fmt.Println(la)
//
//	//修改数据
//	//stem,err := db.Prepare("update music set name=?,singerId=? where id=?")
//	//checkerr(err)
//	//
//	//ex,err :=stem.Exec("双节棍",1,4)
//	//checkerr(err)
//	//
//	//aff,err :=ex.RowsAffected()
//	//checkerr(err)
//	//
//	//fmt.Println(aff)
//
//	//删除数据
//	stem,err := db.Prepare("delete from music where id=?")
//	checkerr(err)
//
//	ex,err :=stem.Exec(1)
//	checkerr(err)
//
//	aff,err :=ex.RowsAffected()
//	checkerr(err)
//
//	fmt.Println(aff)
//
//
//	//全部数据
//	rows,err := db.Query("select m.*,s.* from music m left join siger s on m.singerId = s.sid")
//	checkerr(err)
//
//	for rows.Next(){
//		var id int
//		var name string
//		var year int
//		var singerId int
//		var sid int
//		var sname string
//		err :=rows.Scan(&id,&name,&year,&singerId,&sid,&sname)
//		checkerr(err)
//		//fmt.Println(id)
//		//fmt.Println(name)
//		//fmt.Println(year)
//		//fmt.Println(singerId)
//		//fmt.Println(sid)
//		//fmt.Println(sname)
//		fmt.Printf("%d,%s,%d,%d,%d,%s\n",id,name,year,singerId,sid,sname)
//	}
//
//}
//
//func checkerr(err error) {
//	if err != nil {
//		p
//	}
//}