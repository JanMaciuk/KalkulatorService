package main

import "github/JahoPL/CalculatorService/client/client"

func main() {

	address := "localhost:50051"

	c := client.NewClient(address)

	c.Calculator(2, 31)
}
