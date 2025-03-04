package worshop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Asset struct {
	ID                   string     `json:"id"`
	Name                 string     `json:"name"`
	Type                 string     `json:"type"`
	Summary              string     `json:"summary"`
	Description          string     `json:"description"`
	License              string     `json:"license"`
	LicenseText          string     `json:"licenseText"`
	Unlisted             bool       `json:"unlisted"`
	Private              bool       `json:"private"`
	Blocked              bool       `json:"blocked"`
	RatingCount          int        `json:"ratingCount"`
	SubscriberCount      int        `json:"subscriberCount"`
	CurrentVersionNumber string     `json:"currentVersionNumber"`
	Scenarios            []Scenario `json:"scenarios"`
}

type Scenario struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ScenarioID  string `json:"gameId"`
	GameMode    string `json:"gameMode"`
	PlayerCount int    `json:"playerCount"`
	AuthorName  string `json:"authorName"`
}

type PageProps struct {
	Asset     Asset      `json:"asset"`
	Scenarios []Scenario `json:"scenarios"`
}

type AddonData struct {
	Props struct {
		PageProps PageProps `json:"pageProps"`
	} `json:"props"`
}

// Необходимо использовать ссылку на сценарий, к примеру: https://reforger.armaplatform.com/workshop/7777777777777771-PodvalPvP-PostSovietEra/scenarios
func GetAddonsDataFromWorkshop(url string) (AddonData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return AddonData{}, fmt.Errorf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return AddonData{}, fmt.Errorf("Ошибка при парсинге HTML: %v", err)
	}

	var jsonData string
	doc.Find("script#__NEXT_DATA__").Each(func(i int, s *goquery.Selection) {
		jsonData = s.Text()
	})

	if jsonData == "" {
		return AddonData{}, fmt.Errorf("Не удалось найти JSON данные в теге <script>")
	}

	var data AddonData
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return AddonData{}, fmt.Errorf("Ошибка при парсинге JSON: %v", err)
	}

	return data, nil
}

func (a *AddonData) PrintInfo() {
	asset := a.Props.PageProps.Asset
	fmt.Println("Информация о модификации:")
	fmt.Printf("ID: %s\n", asset.ID)
	fmt.Printf("Name: %s\n", asset.Name)
	fmt.Printf("Type: %s\n", asset.Type)
	fmt.Printf("Summary: %s\n", asset.Summary)
	fmt.Printf("Description: %s\n", asset.Description)
	fmt.Printf("License: %s\n", asset.License)
	fmt.Printf("Unlisted: %v\n", asset.Unlisted)
	fmt.Printf("Private: %v\n", asset.Private)
	fmt.Printf("Blocked: %v\n", asset.Blocked)
	fmt.Printf("RatingCount: %d\n", asset.RatingCount)
	fmt.Printf("SubscriberCount: %d\n", asset.SubscriberCount)
	fmt.Printf("CurrentVersionNumber: %s\n", asset.CurrentVersionNumber)
}

func (a *AddonData) PrintScenarios() {
	for _, scenario := range a.Props.PageProps.Asset.Scenarios {
		fmt.Printf("Название: %s\n", scenario.Name)
		fmt.Printf("Описание: %s\n", scenario.Description)
		fmt.Printf("Режим игры: %s\n", scenario.GameMode)
		fmt.Printf("Количество игроков: %d\n", scenario.PlayerCount)
		fmt.Println(strings.Repeat("-", 40))
	}
}

func (a *AddonData) PrintScenariosCompact() {
	for _, scenario := range a.Props.PageProps.Asset.Scenarios {

		name := scenario.Name
		// if len(scenario.Name) > 42 {
		// 	name = strings.Join([]string{scenario.Name[:41], "..."}, "")
		// }

		fmt.Printf("│ %-4v │ %-50s │ %-80s │ %-40s │ %-40s\n",
			scenario.PlayerCount,
			name,
			scenario.ScenarioID,
			a.Props.PageProps.Asset.Name,
			scenario.AuthorName,
		)
	}
}
