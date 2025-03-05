package main

import (
	"fmt"
	"log"
	"strings"
	"workshop_parser/model/worshop"
)

func main() {
	printMissions("https://reforger.armaplatform.com/workshop/7777777777777771-PodvalPvP-PostSovietEra/scenarios")
	// quickPrint("https://reforger.armaplatform.com/workshop/7777777777777771-PodvalPvP-PostSovietEra/scenarios")
	// quickPrint("https://reforger.armaplatform.com/workshop/7777777777777777-Podval_PVP_Vanilla/scenarios")
	// quickPrint("https://reforger.armaplatform.com/workshop/5F0ECCEE7B9AA8CE-LobbyMissionsRHS/scenarios")
	// quickPrint("https://reforger.armaplatform.com/workshop/5EAF1CB52EF38F0D-ReforgerLobbyMissions/scenarios")
}

func quickPrint(url string) {
	addon, err := worshop.GetAddonsDataFromWorkshop(url)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	addon.PrintScenariosCompact()
	fmt.Println("Всего миссий: ", addon.GetCountMissions())
	fmt.Println()
}

func printMissions(url string) {
	addon, err := worshop.GetAddonsDataFromWorkshop(url)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	addon.PrintInfo()
	fmt.Println(strings.Repeat("-", 40))
	addon.PrintScenarios()
	fmt.Println("Всего миссий: ", addon.GetCountMissions())
	fmt.Println()
}
