package main

import (
	"flag"
	"fmt"

	"github.com/go-redis/redis"
)

var serverAddress *string = flag.String("server", "localhost:6379", "The Server Host")
var password *string = flag.String("password", "", "Password to Connect to Redis With.")
var databaseId *int = flag.Int("database", 0, "The Value to Set where Key is")
var key *string = flag.String("key", "", "The Key to Update")
var value *string = flag.String("value", "", "The Value to Set where Key is")

func main() {
	flag.Parse()

	fmt.Println(" Server Address: ", *serverAddress)
	fmt.Println("       Database: ", *databaseId)
	fmt.Println("        Set Key: ", *key)
	fmt.Println("      Set Value: ", *value)

	client := redis.NewClient(&redis.Options{
		Addr:     *serverAddress,
		Password: *password,   // no password set
		DB:       *databaseId, // use default DB
	})
	defer client.Close()

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("         Status: ", "Failed.")
		fmt.Println(err)
	} else {
		err := client.Set(*key, *value, 0).Err()
		if err != nil {
			fmt.Println("         Status: ", "Failed.")
			panic(err)
		} else {
			fmt.Println("         Status: ", "Good to Go!")
		}
	}
}
