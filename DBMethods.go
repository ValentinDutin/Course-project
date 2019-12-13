package main

import (
	"./DataBase"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)
func getProducts(c echo.Context) error {
	productsCollection := DataBase.ProductsDB.C("Products")
	query := bson.M{}
	products := []DataBase.Product{}
	productsCollection.Find(query).All(&products)
	var sb strings.Builder
	for _, item := range products{
		sb.WriteString("category: " + item.Category + ", company: " + item.Company + ", model: " + item.Model + " Description: " + item.Description + ", price: " + fmt.Sprintf("%f", item.Price) + ", amount: " + strconv.Itoa(int(item.Amount)) + "\n")
	}
	return c.String(http.StatusOK, sb.String())
}
func signUp(c echo.Context) error {
	usersCollection := DataBase.ProductsDB.C("Users")
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	query := bson.M{
		"login" : login,
	}
	var users DataBase.User
	usersCollection.Find(query).One(&users)
	if login == "" || password == ""{
		return c.String(http.StatusOK, "Data entry error")
	}
	if users.Login == login{
		return c.String(http.StatusOK, "Error, user with such login already exists")
	}
	unew := &DataBase.User{Id:bson.NewObjectId(), Login:login, Password:password}
	err := usersCollection.Insert(unew)
	if err != nil{
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "Registration completed succesfully. Welcome, " + login)
}
func signIn(c echo.Context) error {
	usersCollection := DataBase.ProductsDB.C("Users")
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	query := bson.M{
		"login" : login,
		"password" : password,
	}
	var users DataBase.User
	usersCollection.Find(query).One(&users)
	if users.Login == login && users.Password == password {
		return c.String(http.StatusOK, "Welcome back, " + users.Login)
	}
	return c.String(http.StatusOK, "Incorrect login or password")
}

func getUsers(c echo.Context) error {
	usersCollection := DataBase.ProductsDB.C("Users")
	query := bson.M{}
	users := []DataBase.User{}
	usersCollection.Find(query).All(&users)
	var sb strings.Builder
	for _, item := range users{
		sb.WriteString(item.Login + "\n")
	}
	return c.String(http.StatusOK, sb.String())
}


func filterProducts(c echo.Context) error{
	query := []string{c.QueryParam("category"), c.QueryParam("company")}
	var sb strings.Builder
	for _, item := range intersection(query){
		sb.WriteString("category: " + item.Category + ", company: " + item.Company + ", model: " + item.Model + " Description: " + item.Description + ", price: " + fmt.Sprintf("%f", item.Price) + ", amount: " + strconv.Itoa(int(item.Amount)) + "\n")
	}
	return c.String(http.StatusOK, sb.String())
}

func intersection(query []string) []DataBase.Product{
	productsCollection := DataBase.ProductsDB.C("Products")
	products := []DataBase.Product{}
	productsCollection.Find(bson.M{}).All(&products)
	intersectionMap := make(map[DataBase.Product]int)
	for _, item := range products{
		intersectionMap[item] = 0
	}
	var isSingleQuery bool
	if query[0] == "" || query[1] == ""{
		isSingleQuery = true
	}
	for _, item := range products{
		if item.Category == query[0] {
			intersectionMap[item] += 1
		}
		if item.Company == query[1] {
			intersectionMap[item] += 1
		}
	}
	filterResult := []DataBase.Product{}
	for key, value := range intersectionMap{
		if isSingleQuery == true && value == 1{
			filterResult = append(filterResult, key)
		} else if isSingleQuery == false && value > 1 {
			filterResult = append(filterResult, key)
		}
	}
	return filterResult
}

func main() {
	fmt.Println("Hello")
	DataBase.InitDataBase()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main page")
	})
	e.GET("/products", getProducts)
	e.GET("/signUp", signUp)
	e.GET("/signIn", signIn)
	e.GET("/filter", filterProducts)
	defer DataBase.Session.Close()
	e.Logger.Fatal(e.Start(":27017"))
}
