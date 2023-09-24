package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/AthithyanR/portfolio/internal/database"
)

// const (
// 	catFactUrl = "https://catfact.ninja/fact?max_length=1000"
// )

type CatFact struct {
	Fact   string `json:"fact"`
	Length int64  `json:"length"`
}

func GetTechIKnow() string {
	techIKnow := []string{
		"React/Nextjs",
		"Nodejs",
		"Sql/NoSql databases",
		"CI/CD",
		"AWS",
		"linux",
	}

	techIKnowLen := len(techIKnow)

	if techIKnowLen == 1 {
		return techIKnow[0]
	}

	lastIdx := techIKnowLen - 1

	techIKnowString := strings.Join(techIKnow[:lastIdx], ", ")
	techIKnowString += fmt.Sprintf(" and %s", techIKnow[lastIdx])

	return techIKnowString
}

func GetCatFact() string {
	// catFactReq, err := http.Get(catFactUrl)
	// if err != nil {
	// 	log.Println(err)
	// 	return ""
	// }

	// defer catFactReq.Body.Close()

	// catFact := CatFact{}
	// err = json.NewDecoder(catFactReq.Body).Decode(&catFact)
	// if err != nil {
	// 	log.Println(err)
	// 	return ""
	// }

	// return catFact.Fact

	buff, err := os.ReadFile("static/cat-facts.json")
	if err != nil {
		log.Println(err)
		return ""
	}

	catFacts := []CatFact{}

	err = json.Unmarshal(buff, &catFacts)
	if err != nil {
		log.Println(err)
		return ""
	}

	return catFacts[rand.Intn(len(catFacts))].Fact
}

func GetChatMessages() []database.ChatMessage {
	chatMessages := []database.ChatMessage{}

	database.DB.Find(&chatMessages)

	return chatMessages
}

func CreateChatMessage(chatMessage *database.ChatMessage) error {
	database.DB.Create(chatMessage)
	return nil
}
