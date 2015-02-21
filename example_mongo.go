package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"strconv"
)

type Person struct {
	Name  string
	Phone string
}

func AddUsersToMongoDB(count int, session *mgo.Session) {
	c := session.DB("test").C("people")

	for i := 0; i < count; i++ {
		_ = c.Insert(&Person{"Ale" + strconv.Itoa(i), "+55 53 8116 9639"})
	}

}
