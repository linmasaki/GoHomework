package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/linmasaki/gohomework/internal/model"
)

func main() {
	router := gin.Default()
	router.GET("/roles", get)
	router.GET("/role/:id", getById)
	router.POST("/role", post)
	router.PUT("/role", put)
	router.DELETE("/role/:id", delete)
	router.Run(":8080")
}

// 取得全部資料
func get(c *gin.Context) {
	c.JSON(http.StatusOK, model.Data)
}

// 取得單一筆資料
func getById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.String(400, "Parameter Error!")
		return
	}

	var result model.Role
	for _, role := range model.Data {
		if role.ID == uint(id) {
			result = role
			break
		}
	}

	c.JSON(http.StatusOK, result)
}

// 新增資料
func post(c *gin.Context) {
	var data RoleVM
	roleString := json.NewDecoder(c.Request.Body)
	roleString.Decode(&data)

	newRole := model.Role{
		ID:      data.ID,
		Name:    data.Name,
		Summary: data.Summary,
	}
	model.Data = append(model.Data, newRole)

	c.JSON(http.StatusOK, model.Data)
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func put(c *gin.Context) {
	var data RoleVM
	roleString := json.NewDecoder(c.Request.Body)
	err := roleString.Decode(&data)
	if err != nil {
		c.String(400, "Parameter Error!")
		return
	}

	var indexToUpdate int
	dataLength := len(model.Data)
	for i := 0; i < dataLength; i++ {
		if model.Data[i].ID == uint(data.ID) {
			model.Data[i].Name = data.Name
			model.Data[i].Summary = data.Summary
			indexToUpdate = i
			break
		}
	}

	c.JSON(http.StatusOK, model.Data[indexToUpdate])
}

// 刪除資料
func delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.String(400, "Parameter Error!")
		return
	}

	var indexToDelete int
	for i, role := range model.Data {
		if role.ID == uint(id) {
			indexToDelete = i
			break
		}
	}

	dataLength := len(model.Data)
	switch {
	case indexToDelete == 0:
		model.Data = model.Data[1:]

	case indexToDelete == dataLength-1:
		model.Data = model.Data[:indexToDelete]

	default:
		model.Data = append(model.Data[:indexToDelete], model.Data[1+indexToDelete:]...)

	}
	c.JSON(http.StatusOK, model.Data)
}
