package main

import (
	"fmt"

	"study/feature1"
	"study/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hellow, main and Git")
	feature1.Feature1()

	simple_connection.ChekConnection()
}
