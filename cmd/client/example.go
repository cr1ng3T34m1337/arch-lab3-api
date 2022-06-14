package main

import (
	"fmt"
	"log"

	"github.com/cr1ng3T34m1337/arch-lab3-api/sdk"
	"github.com/cr1ng3T34m1337/arch-lab3-api/tools"
)

func main() {
	config, err := tools.ParseConfig("config.json")
	if err != nil {
		log.Fatal("Cannot read config: ", err)
	}

	base := fmt.Sprintf("http://%s:%d", config.Api.Addr, config.Api.Port)
	client := &sdk.Client{BaseUrl: base}

	fmt.Println("=== Scenario 1 ===")
	balancers, err := client.ListBalancers()
	if err != nil {
		log.Fatal("Cannot list balancers: ", err)
	}
	fmt.Print("Balancers: ")
	fmt.Println(balancers)

	fmt.Println("=== Scenario 2 ===")
	usedMachines := balancers[0].UsedMachines
	err = client.SetWorkingStatus(usedMachines[len(usedMachines)-1], "no")
	if err != nil {
		log.Fatal("Cannot set working status: ", err)
	}
	fmt.Println("Working status has been set")
	balancers, err = client.ListBalancers()
	if err != nil {
		log.Fatal("Cannot list balancers: ", err)
	}
	fmt.Print("New balancers list: ")
	fmt.Println(balancers)
}
