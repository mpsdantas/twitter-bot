package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"bot/client"
)

func main() {
	data, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		log.Fatalln("could not read credentials", err)
	}

	var cred *client.Credentials

	if err := json.Unmarshal(data, &cred); err != nil {
		log.Fatalln("could not unmarshal credentials", err)
	}

	c := client.New(cred)

	res, err := c.CreateTwitter(&client.CreateTwitterRequest{
		Text: fmt.Sprintf("Olá, eu sou um twitter criado via api em %s", time.Now().Format("02/01/2006 15:01:05")),
	})
	if err != nil {
		log.Fatalln("could not create twitter", err)
	}

	log.Println("twitter created successfully", res.Data.ID)
}
