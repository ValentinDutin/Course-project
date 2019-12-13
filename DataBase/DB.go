package DataBase

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var ProductsDB *mgo.Database
var Session *mgo.Session
type User struct{
	Id       bson.ObjectId `bson:"_id"`
	Login    string        `bson:"login"`
	Password string        `bson:"password"`
	CurMoney    int        `bson:"curMoney"`
}
type Product struct{
	Id 		bson.ObjectId `bson:"_id"`
	Category string		  `bson:"category"`
	Company  string 	  `bson:"company"`
	Model 	 string 	  `bson:"model"`
	Description string    `bson:"description"`
	Price 	 float32      `bson:"price"`
	Amount 	 uint		  `bson:"amount"`
}
func InitDataBase() {
	Session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	ProductsDB = Session.DB("DataBase")
	usersCollection := Session.DB("DataBase").C("Users")
	usersSlice := []*User{&User{Id: bson.NewObjectId(), Login: "user1", Password: "user1password", CurMoney: 10},
		&User{Id: bson.NewObjectId(), Login: "user2", Password: "user2password", CurMoney: 1101},
		&User{Id: bson.NewObjectId(), Login: "user3", Password: "user3password", CurMoney: 2222},
	}
	err = usersCollection.Insert(usersSlice)
	if err != nil{
		fmt.Println(err)
	}
	productsCollection := Session.DB("DataBase").C("Products")
	productsSlice := []*Product{&Product{Id: bson.NewObjectId(), Category: "headphones", Company: "Marshall", Model: "MID BT", Description: "color: black", Price: 300, Amount: 3},
		&Product{Id: bson.NewObjectId(), Category: "headphones", Company: "Apple", Model: "airpods pro", Description: "color: white, ANC", Price:630, Amount: 1},
		&Product{Id: bson.NewObjectId(), Category: "phone", Company: "Xiaomi", Model: "Redmi Note 8", Description: "color: blue, storage: 64GB", Price: 317, Amount: 300},
	}
	err = productsCollection.Insert(productsSlice)
	if err != nil{
		fmt.Println(err)
	}
	defer Session.Close()
}
