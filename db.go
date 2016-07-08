package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type USER struct {
	Name   string
	LineID string
}

func testDB() {
	session, err := mgo.Dial("mongodb://heroku_h1g317z7:k5c6qmc4so8glsjvb68m659m26@ds017205.mlab.com:17205")
	c := session.DB("heroku_h1g317z7").C("_User")
	err = c.Insert(&USER{"Ale", "+55 53 8116 9639"},
		&USER{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := USER{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.LineID)
}
