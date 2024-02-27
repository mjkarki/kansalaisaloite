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

type Arguments struct {
	help bool
	all  bool
	text string
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

func parseArguments() Arguments {
	var args Arguments
	args.help = false
	args.all = false
	args.text = ""

	if len(os.Args) < 2 {
		args.help = true
	}

	for _, arg := range os.Args[1:] {
		switch l_arg := strings.ToLower(arg); l_arg {
		case "-h":
			args.help = true
		case "--help":
			args.help = true
		case "-a":
			args.all = true
		case "--all":
			args.all = true
		default:
			args.text += l_arg + " "
		}
	}
	args.text = strings.TrimRight(args.text, " ")

	return args
}

func main() {
	var data []byte
	var aloitteet []Initiative = []Initiative{}
	var kaikki = false
	var offset = 0
	var nimi string

	args := parseArguments()

	if args.help {
		fmt.Println("Käyttö:")
		fmt.Println("  " + os.Args[0] + ` [-h | --help | -a | --all] ["hakuteksti" ...]`)
		fmt.Println()
		fmt.Println(`  "hakuteksti"  etsii hakutekstiä aloitteiden otsikoista`)
		fmt.Println(`  -h | --help   tulostaa tämän ohjeen`)
		fmt.Println(`  -a | --all    tulostaa kaikki kansalaisaloitteet`)
		os.Exit(0)
	}

	if args.all {
		fmt.Println("Listataan kaikki aloitteet:")
		kaikki = true
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
		if kaikki {
			fmt.Println("- " + nimi)
		} else {
			nimi_l := strings.ToLower(nimi)
			if strings.Contains(nimi_l, args.text) {
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
