package controllers

import (
	"github.com/gofiber/fiber/v2"

	"log"
	"root/database"
	"root/models"
)

var (
	items = []models.Item{
		{ItemName: "Pizza", Amount: 12.49},
		{ItemName: "Ghorme sabzi", Amount: 6.95},
		{ItemName: "Ab gosht", Amount: 17.99},
	}
)

func UserGetAllFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserGetOrders(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserCreateOrder(c *fiber.Ctx) error {
	// Create Items
	for index := range items {
		database.DB.Create(&items[index])
	}
	order := models.Order{Status: "pending"}
	database.DB.Create(&order)
	item1 := models.OrderItem{OrderID: order.ID, ItemID: items[0].ID, Quantity: 1}
	item2 := models.OrderItem{OrderID: order.ID, ItemID: items[1].ID, Quantity: 4}
	database.DB.Create(&item1)
	database.DB.Create(&item2)

	// Query with joins
	rows, err := database.DB.Table("orders").Where("orders.id = ? and status = ?", order.ID, "pending").
		Joins("Join order_items on order_items.order_id = orders.id").
		Joins("Join items on items.id = order_items.id").
		Select("orders.id, orders.status, order_items.order_id, order_items.item_id, order_items.quantity" +
			", items.item_name, items.amount").Rows()
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	// Values to load into
	newOrder := &models.Order{}
	newOrder.OrderItems = make([]models.OrderItem, 0)

	for rows.Next() {
		orderItem := models.OrderItem{}
		item := models.Item{}
		err = rows.Scan(&newOrder.ID, &newOrder.Status, &orderItem.OrderID, &orderItem.ItemID, &orderItem.Quantity, &item.ItemName, &item.Amount)
		if err != nil {
			log.Panic(err)
		}
		orderItem.Item = item
		newOrder.OrderItems = append(newOrder.OrderItems, orderItem)
	}
	return c.JSON(order)
}

func UserGetHistory(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserGetFavorites(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserCreateComment(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}
