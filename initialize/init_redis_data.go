package initialize

import (
	"context"
	"log"
	"strconv"

	"github.com/chienduynguyen1702/go-redis/models"
	"github.com/redis/go-redis/v9"
)

var FoodItems = []models.FoodItem{
	{
		ID:       1,
		Name:     "Apple",
		Quantity: 58,
	},
	{
		ID:       2,
		Name:     "Banana",
		Quantity: 30,
	},
	{
		ID:       3,
		Name:     "Orange",
		Quantity: 13,
	},
	{
		ID:       4,
		Name:     "Pineapple",
		Quantity: 80,
	},
	{
		ID:       5,
		Name:     "Gum",
		Quantity: 22,
	},
	{
		ID:       6,
		Name:     "Chocolate",
		Quantity: 14,
	},
}

func InitSampleData(redisClient *redis.Client, ctx context.Context) {
	for _, item := range FoodItems {
		key := "food_item:" + strconv.Itoa(item.ID)
		err := redisClient.HSet(ctx, key, "name", item.Name, "quantity", item.Quantity).Err()
		if err != nil {
			log.Fatal("Error initializing food items:", err)

		}
	}
}
