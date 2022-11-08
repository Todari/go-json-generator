package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Metadata struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Image       string `json:"image"`
	// Attributes  []Attributes `json:"attributes"`
	// Properties  Properties   `json:"properties"`
}

type Attributes struct {
	Trait_type string `json:"trait_type"`
	Value      string `json:"value"`
}

type Properties struct {
	Files []Files `json:"files"`
}

type Files struct {
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

func main() {

	for i := 1; i < 10001; i++ {
		var result Metadata
		// result.Attributes = make([]Attributes, 11)
		// result.Properties.Files = make([]Files, 1)

		result.Symbol = "WCREW"
		result.Description = "Join the Whowho Crew!\nThis is an NFT project with Whowho Crew.\nExplore the sea of virtual assets with Whowho Crew!"
		// result.Attributes[0].Trait_type = "legend"
		// result.Attributes[0].Value = strings.Split(traitArr[0], "#")[0]
		// result.Attributes[1].Trait_type = "background"
		// result.Attributes[1].Value = strings.Split(traitArr[1], "#")[0]
		// result.Attributes[2].Trait_type = "body"
		// result.Attributes[2].Value = strings.Split(traitArr[2], "#")[0]
		// result.Attributes[3].Trait_type = "outfit"
		// result.Attributes[3].Value = strings.Split(traitArr[3], "#")[0]
		// result.Attributes[4].Trait_type = "backpack"
		// result.Attributes[4].Value = strings.Split(traitArr[4], "#")[0]
		// result.Attributes[5].Trait_type = "pet"
		// result.Attributes[5].Value = strings.Split(traitArr[5], "#")[0]
		// result.Attributes[6].Trait_type = "hat"
		// result.Attributes[6].Value = strings.Split(traitArr[6], "#")[0]
		// result.Attributes[7].Trait_type = "eye"
		// result.Attributes[7].Value = strings.Split(traitArr[7], "#")[0]
		// result.Attributes[8].Trait_type = "mouth"
		// result.Attributes[8].Value = strings.Split(traitArr[8], "#")[0]
		// result.Attributes[9].Trait_type = "ring"
		// result.Attributes[9].Value = strings.Split(traitArr[9], "#")[0]
		// result.Attributes[10].Trait_type = "rarity"
		// result.Attributes[10].Value = strings.Split(traitArr[10], "#")[0]

		// result.Properties.Files[0].Type = "image/png"
		result.Name = "WhowhoCrew#" + strconv.Itoa(i)
		result.Image = "https://kvada.vpay.co.kr/nft/Whowhocrew/" + strconv.Itoa(i) + "/image.png"
		// result.Properties.Files[0].Uri = strconv.Itoa(i) + ".png"
		doc, _ := json.Marshal(result)
		os.Mkdir("./nft/Whowhocrew/"+strconv.Itoa(i)+"/", 0777)

		err2 := ioutil.WriteFile("./nft/Whowhocrew/"+strconv.Itoa(i)+"/meta.json", doc, os.FileMode(0644))

		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}
