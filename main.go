package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const URL string = "https://www.kansalaisaloite.fi/api/v1/initiatives?limit=50&orderBy=createdNewest&offset="

type InitiativeStruct interface {
	[]Initiative | InitiativeInfo
}

func getData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Tried to fetch data from URL"+url+":", err)
	}
	return data
}

func getInitiativeStructFromJSON[T InitiativeStruct](data []byte) T {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal("Tried to convert JSON to structured data:", err)
	}
	return result
}

func main() {
	var data []byte
	var aloitteet []Initiative = []Initiative{}
	var listall = false
	var offset = 0
	var nimi string

	if len(os.Args) < 2 {
		fmt.Println("Käyttö:")
		fmt.Println("    " + os.Args[0] + ` "hakuteksti"`)
		fmt.Println("  tai")
		fmt.Println("    " + os.Args[0] + ` -a`)
		os.Exit(0)
	}

	search := os.Args[1]

	if search == "-a" {
		fmt.Println("Listataan kaikki aloitteet:")
		listall = true
	}

	data = getData(fmt.Sprintf("%s%d", URL, offset))
	for !bytes.Equal(data, []byte{91, 93}) {
		a := getInitiativeStructFromJSON[[]Initiative](data)
		aloitteet = append(aloitteet, a...)
		offset += 50
		data = getData(fmt.Sprintf("%s%d", URL, offset))
	}

	for _, initiative := range aloitteet {
		if initiative.Name["fi"] == nil {
			nimi = initiative.Name["sv"].(string)
		} else {
			nimi = initiative.Name["fi"].(string)
		}
		if listall {
			fmt.Println("- " + nimi)
		} else {
			nimi_l := strings.ToLower(nimi)
			if strings.Contains(nimi_l, search) {
				url := initiative.ID
				data = getData(url)
				initiativeStruct := getInitiativeStructFromJSON[InitiativeInfo](data)
				fmt.Println("- " + nimi + ":")
				fmt.Println("  Kannatusilmoituksia:", initiativeStruct.SupportCount)
				fmt.Println("  Keräys päättyy:     ", initiativeStruct.EndDate)
			}
		}
	}
}
