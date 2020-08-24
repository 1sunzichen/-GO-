package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
 
	fmt.Println("MyFavorite is",rand.Intn(10))
}