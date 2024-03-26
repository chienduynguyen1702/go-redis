package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/chienduynguyen1702/go-redis/initialize"
	"github.com/chienduynguyen1702/go-redis/models"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Load environment variables and context
	initialize.LoadEnvVarFile()
	ctx := context.Background()

	// Create a new redis client
	redisClient := initialize.NewRedisClient(ctx)
	// create sample data
	if os.Getenv("FIRST_TIME") == "true" {
		initialize.InitSampleData(redisClient, ctx)
	}
	// NumberOfItem := len(initialize.FoodItems)
	var sampleFoodItemList = initialize.FoodItems
	menu := -1

	// create a menu cli to interact with cli user
MENU_LOOP:
	for {
		fmt.Println("=====================================")
		fmt.Println("1. Display all item in store")
		fmt.Println("2. Regenerate items")
		fmt.Println("3. Launch a sale")
		fmt.Println("0. Exit")
		fmt.Printf("Enter option: ")
		fmt.Scanf("%d", &menu)
		fmt.Println("=====================================")
		switch menu {
		case 1:
			// Display remaining quantities of food items
			displayRemainingQuantities(redisClient, ctx, sampleFoodItemList)
		case 2:
			initialize.InitSampleData(redisClient, ctx)
			fmt.Println("Regenerate items successfully")
		case 3:
			launchSale(redisClient, ctx, sampleFoodItemList)
		case 0:
			break MENU_LOOP
		default:
			fmt.Println("Invalid choice !! Try again")
		}
		// Consume newline character from input buffer
		fmt.Scanln()
	}
}

// Process orders received from the orderChannel
func processOrders(redisClient *redis.Client, ctx context.Context, foodItemList []models.FoodItem, orderChannel <-chan int) {
	for orderID := range orderChannel {
		buyRandomFoodItem(redisClient, ctx, orderID, foodItemList)
	}
}

// Generate orders and send them to the orderChannel
func generateOrders(orderChannel chan<- int) {
	for i := 0; i < 100; i++ {
		orderChannel <- i
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond) // Simulate random order interval
	}
	close(orderChannel)
}

// Launch a random sale
func launchSale(redisClient *redis.Client, ctx context.Context, foodItemList []models.FoodItem) {
	// Start order processing goroutine

	// Create a channel to handle orders
	orderChannel := make(chan int)
	// Generate orders in separate goroutine
	go generateOrders(orderChannel)
	// New go routine to process orders
	go processOrders(redisClient, ctx, foodItemList, orderChannel)
	// Wait for purchases to complete
	time.Sleep(3 * time.Second)
	fmt.Println("done: time.Sleep is 3 seconds")
	fmt.Println("")
	displayRemainingQuantities(redisClient, ctx, foodItemList)
}

// Simulate a order buying a random food item
func buyRandomFoodItem(redisClient *redis.Client, ctx context.Context, order int, foodItemList []models.FoodItem) {
	// Randomly select a food item
	itemID := rand.Intn(len(foodItemList)) + 1
	key := "food_item:" + strconv.Itoa(itemID)

	// Check if item is available
	quantity, err := redisClient.HGet(ctx, key, "quantity").Int()
	if err != nil {
		log.Println("Error checking quantity of food item:", err)
		return
	}

	if quantity <= 0 {
		log.Printf("Order %3d failed %15s: Out of stock\n", order, foodItemList[itemID-1].Name)
		return
	}

	// Simulate purchase
	err = redisClient.HIncrBy(ctx, key, "quantity", -1).Err()
	if err != nil {
		log.Printf("Order %3d failed decrement quantity %15s: %v\n", order, foodItemList[itemID-1].Name, err)
		return
	}

	log.Printf("Order %3d bought %15s: Successfully\n", order, foodItemList[itemID-1].Name)
}

// Display remaining quantities of food items
func displayRemainingQuantities(redisClient *redis.Client, ctx context.Context, foodItemList []models.FoodItem) {
	fmt.Println("Remaining quantities of food items:")
	fmt.Printf("%5s|%15s|%10s\n", "ID ", "Name     ", "Quantity")
	length := len(foodItemList)
	for i := 0; i < length; i++ {

		key := "food_item:" + strconv.Itoa(foodItemList[i].ID)
		currentQuantity, err := redisClient.HGet(ctx, key, "quantity").Int()
		if err != nil {
			log.Println("Error getting quantity of food item:", err)
			continue
		}

		fmt.Printf("%5d|%15s|%7d\n", foodItemList[i].ID, foodItemList[i].Name, currentQuantity)
	}
}
