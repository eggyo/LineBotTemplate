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
	session, err := mgo.Dial("mongodb://heroku_7swm6cvp:kshu6a3cjl8ilpe1l3pdi0llq4@ds017195.mlab.com:17195/heroku_7swm6cvp")
	c := session.DB("heroku_7swm6cvp").C("_User")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
