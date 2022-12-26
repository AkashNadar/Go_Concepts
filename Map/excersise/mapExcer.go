package main

import "fmt"

const (
	ONLINE      = 0
	OFFLINE     = 1
	MAINTENANCE = 2
	RETIRED     = 3
)

func printStatus(serverStat map[string]int) {
	statCount := make(map[int]int)
	for _, stat := range serverStat {
		statCount[stat] += 1
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println(statCount[ONLINE], "no of user online")
	fmt.Println(statCount[OFFLINE], "no of user offline")
	fmt.Println(statCount[MAINTENANCE], "no of user maintenance")
	fmt.Println(statCount[RETIRED], "no of user retired")
}

func main() {
	servers := []string{"darkster", "aiur", "omicron", "corona", "base24"}

	status := make(map[string]int)

	for _, server := range servers {
		status[server] = ONLINE
	}

	for key, val := range status {
		println("server :", key, "status :", val)
	}
	printStatus(status)
	fmt.Println("1111111111111111111111111111111111111")

	status["darkster"] = RETIRED
	status["aiur"] = OFFLINE
	for key, val := range status {
		println("server :", key, "status :", val)
	}
	printStatus(status)

	fmt.Println("22222222222222222222222222222222222222")
	status["darkster"] = MAINTENANCE
	status["aiur"] = MAINTENANCE
	status["omicron"] = MAINTENANCE
	status["corona"] = MAINTENANCE
	status["base24"] = MAINTENANCE
	for key, val := range status {
		println("server :", key, "status :", val)
	}
	printStatus(status)

}
