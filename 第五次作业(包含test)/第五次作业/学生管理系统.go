package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Student struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

var students map[string]Student

func init() {
	students = make(map[string]Student)
}
func add(ctx context.Context, c *app.RequestContext) {
	var newStudent Student
	if err := c.Bind(&newStudent); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"message": "Invalid request data"})
		return
	}
	students[newStudent.Id] = newStudent
	c.JSON(consts.StatusOK, utils.H{"message": "添加成功"})
}

func profile(ctx context.Context, c *app.RequestContext) {
	var prostudent Student
	if err := c.Bind(&prostudent); err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"message": "Invalid request data"})
		return
	}
	if _, ok := students[prostudent.Id]; ok {
		c.JSON(consts.StatusOK, utils.H{"message": "修改成功"})
	} else {
		c.JSON(consts.StatusNotFound, utils.H{"message": "Invalid student"})
	}
}

func search(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	if student, ok := students[id]; ok {
		c.JSON(consts.StatusOK, student)
	} else {
		c.JSON(consts.StatusNotFound, utils.H{"message": "没有该学生"})
	}
}
func main() {
	h := server.Default()
	h.POST("/add", add)
	h.POST("/profile", profile)
	h.GET("/search", search)
	h.Spin()
}
