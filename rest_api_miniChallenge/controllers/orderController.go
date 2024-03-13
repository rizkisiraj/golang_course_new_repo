package controllers

import (
	"fmt"
	"net/http"
	"rest_api_mini_challenge/database"
	"rest_api_mini_challenge/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	if err := ctx.BindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"data":    newOrder,
		})
		return
	}

	err := db.Create(&newOrder).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"data":    newOrder,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    newOrder,
		"message": "succeed create order",
	})

}

func GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func GetOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	err := db.Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})

}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"data":    order,
		})
		return
	}

	err := db.Model(&models.Order{}).Where("id = ?", id).Updates(order).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"data":    order,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "succeed create order",
	})

}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order

	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	err := db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "succesfully deleting order",
	})
}
