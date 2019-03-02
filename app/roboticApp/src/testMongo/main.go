package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name string "bson: `name`"
	Age  int    "bson: `age`"
}

func main() {
	var mydb = getDB()
	mydb.Login("root", "Server201901")
	c := mydb.C("student")
	err := c.Insert(&User{Name: "Chris", Age: 23})
	if err != nil {
		panic(err)
	}

	fmt.Println("Login Successfully")
	result := User{}
	err = c.Find(bson.M{"name": "Chris"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Age:", result.Age)
}

func getDB() *mgo.Database {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("admin")
	return db
}
