package main

import (
	"fmt"
	"log"

	"github.com/barbucatalinn/collatz-conjecture/app"
)

func main() {
	application, err := app.New("./data/sample.csv")
	if err != nil {
		log.Fatal(err)
	}

	result := application.CalculateCollatzConjectureSteps()
	fmt.Println("result", result)
}
