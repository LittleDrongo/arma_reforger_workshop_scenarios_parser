package main

import (
	"fmt"
	"log"
	"strings"
	"workshop_parser/model/worshop"

	"github.com/LittleDrongo/fmn-lib/console/cmd"
)

func main() {
	quickPrintSample("https://reforger.armaplatform.com/workshop/7777777777777771-PodvalPvP-PostSovietEra/scenarios")
	quickPrintSample("https://reforger.armaplatform.com/workshop/7777777777777777-Podval_PVP_Vanilla/scenarios")
	quickPrintSample("https://reforger.armaplatform.com/workshop/5F0ECCEE7B9AA8CE-LobbyMissionsRHS/scenarios")
	quickPrintSample("https://reforger.armaplatform.com/workshop/5EAF1CB52EF38F0D-ReforgerLobbyMissions/scenarios")

	cmd.Input("Нажмите Enter чтобы закрыть приложение.")
}

func quickPrintSample(url string) {
	addon, err := worshop.GetAddonsDataFromWorkshop(url)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	addon.PrintScenariosCompact()
}

func printMissionsSample(url string) {
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
