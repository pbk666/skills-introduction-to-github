package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"testing"
)

func TestAdd(t *testing.T) {
	body := `{Id:"1",Name:"alice",Gender:"male",}`
	method := "POST"
	path := "/add"
	c := ut.CreateUtRequestContext(method, path, &ut.Body{
		Body: bytes.NewBufferString(body),
		Len:  len(body),
	})
	add(context.Background(), c)
	assert.DeepEqual(t, c.Response.StatusCode(), consts.StatusOK)
	expectedMsg := []byte(`{"message":"添加成功"}`)
	assert.DeepEqual(t, c.Response.Body(), expectedMsg)
}
func TestProfile(t *testing.T) {
	initStudent := Student{
		Id:     "1",
		Name:   "bob",
		Gender: "male",
	}
	students[initStudent.Id] = initStudent
	body := `{id:"1",name:"bob",gender:"male",}`
	method := "POST"
	path := "/profile"
	c := ut.CreateUtRequestContext(method, path, &ut.Body{
		Body: bytes.NewBufferString(body),
	})
	profile(context.Background(), c)
	assert.DeepEqual(t, c.Response.StatusCode(), consts.StatusOK)
	exceptMsg := []byte(`{"message":"修改成功"}`)
	assert.DeepEqual(t, c.Response.Body(), exceptMsg)
}
func TestSearch(t *testing.T) {
	searchstudent := Student{
		Id:     "1",
		Name:   "bob",
		Gender: "male",
	}
	students[searchstudent.Id] = searchstudent
	method := "GET"
	path := "/search?id=1"
	c := ut.CreateUtRequestContext(method, path, nil)
	search(context.Background(), c)
	assert.DeepEqual(t, c.Response.StatusCode(), consts.StatusOK)
	exceptMsg, _ := json.Marshal(searchstudent)
	assert.DeepEqual(t, c.Response.Body(), exceptMsg)
}
