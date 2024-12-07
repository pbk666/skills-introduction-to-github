package main

import (
	"context"
	"database/sql"
	"decleration/sqlop"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

var db *sql.DB

func init() {
	dsn := "root:hj2005691@tcp(127.0.0.1:3306)/my_db_01"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

var students map[int]Student

func init() {
	students = make(map[int]Student)
}
func add(ctx context.Context, c *app.RequestContext) {
	var newStudent Student
	if err := c.Bind(&newStudent); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"message": "Invalid request data"})
		return
	}
	sqlop.Addsql(newStudent.Name, newStudent.Gender)
	c.JSON(consts.StatusOK, utils.H{"message": "添加成功"})
}

func profile(ctx context.Context, c *app.RequestContext) {
	var proStudent Student
	if err := c.Bind(&proStudent); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"message": "类型转变失败"})
		return
	}
	sqlop.Profilesql(proStudent.Gender, proStudent.Id)
	if _, ok := students[proStudent.Id]; !ok {
		c.JSON(consts.StatusOK, utils.H{"message": "修改成功"})
	} else {
		c.JSON(consts.StatusNotFound, utils.H{"message": "Invalid student"})
	}
}

func search(ctx context.Context, c *app.RequestContext) {
	var s Student
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	row := db.QueryRow("select id,name,gender from my_db_01.students where id=?;", idInt)
	err := row.Scan(&s.Id, &s.Name, &s.Gender)
	if err == nil {
		c.JSON(consts.StatusOK, s)
		sqlop.SearchSql(idInt)
	} else {
		c.JSON(consts.StatusNotFound, utils.H{"message": "没有该学生"})
	}
}

func delete(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	res, err := db.Exec("delete from my_db_01.students where id=?;", id)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, utils.H{"message": "删除失败"})
		return
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		c.JSON(consts.StatusNotFound, utils.H{"message": "没有找到要删除的id"})
		return
	}
	c.JSON(consts.StatusOK, utils.H{"message": "删除成功"})
}

func main() {
	sqlop.InitDB()
	h := server.Default()
	h.POST("/add", add)
	h.POST("/profile", profile)
	h.GET("/search", search)
	h.DELETE("/delete", delete)
	h.Spin()
}
