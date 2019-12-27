package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"gopkg.in/guregu/null.v3"
	"os"
)

type Client struct {
	Id      string     `csv:"client_id"`
	Name    string     `csv:"client_name"`
	Balance null.Float `csv:"client_balance"`
	NotUsed string     `csv:"-"`
}

func main() {
	clientsFile, err := os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	for _, client := range clients {
		if client.Balance.Valid {
			fmt.Println("Hello", client.Name)
		}
	}
}
