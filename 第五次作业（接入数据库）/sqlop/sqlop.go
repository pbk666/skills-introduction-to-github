package sqlop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id     int
	name   string
	gender string
}

var db2 *sql.DB

// 初始化
func InitDB() (err error) {
	dsn := "root:hj2005691@tcp(127.0.0.1:3306)/my_db_01"
	db2, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = db2.Ping()
	if err != nil {
		panic(err.Error())
	}
	return
}

// search 查询学生
func SearchSql(n int) {
	var s1 Student
	row := db2.QueryRow("select id,name,gender from my_db_01.students where id=?;", n)
	row.Scan(&s1.id, &s1.name, &s1.gender)
	fmt.Println(s1)
}

// add 添加学生
func Addsql(name string, gender string) {
	res, err := db2.Exec("insert into students(name,gender) values(?,?);", name, gender)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("insert id:", id)
}

// profile 修改数据
func Profilesql(gender string, id int) {
	res, err := db2.Exec("update students set gender=? where id=?;", gender, id)
	if err != nil {
		panic(err.Error())
	}
	n, _ := res.RowsAffected()
	fmt.Printf("更新第%d行数据", n)
}

// 打印数据库
func Mysql() {
	rows, err := db2.Query("SELECT * FROM students")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var s Student
		err := rows.Scan(&s.id, &s.name, &s.gender)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(s)
	}
}

