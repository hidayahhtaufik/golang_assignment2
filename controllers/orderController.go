package controllers

import (
	"assignment2/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type orderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *orderController {
	return &orderController{
		db: db,
	}
}

func (in *orderController) GetOrder(c *gin.Context) {
	var (
		order []models.Order
	)

	err := in.db.Preload("Item").Find(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &order,
	})
}

func (in *orderController) CreateOrder(c *gin.Context) {
	var order models.Order

	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Create(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": &order,
	})

}

func (in *orderController) DeleteOrderByID(c *gin.Context) {
	var (
		order models.Order
	)

	id := c.Param("id")

	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success !",
	})
}

func (in *orderController) UpdateOrderByID(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
	)
	id := c.Param("id")
	err := in.db.First(&order, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newOrder)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Model(&order).Updates(newOrder).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "update data success !",
		"data": newOrder,
	})
}
