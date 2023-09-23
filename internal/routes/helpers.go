package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// const (
// 	catFactUrl = "https://catfact.ninja/fact?max_length=1000"
// )

type CatFact struct {
	Fact   string `json:"fact"`
	Length int64  `json:"length"`
}

type ChatMessage struct {
	Ip      string
	Name    string
	Message string
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

func GetChatMessages() []ChatMessage {
	chatMessages := []ChatMessage{}

	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return chatMessages
	}

	db.AutoMigrate(&ChatMessage{})
	// db.Create(&ChatMessages{
	// 	Ip:      "127.0.0.1",
	// 	Name:    "Athithyan",
	// 	Message: "Hello sqlite3",
	// })

	db.Find(&chatMessages)

	return chatMessages
}

func CreateChatMessage(chatMessage *ChatMessage) error {
	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	db.Create(chatMessage)
	return nil
}
