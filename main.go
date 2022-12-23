package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
type food struct {
	ID int 
	Name string 
	Category string 
	Description string
	UrlImage string
}

type message struct {
	Content string 
}

var msg_err = message{
	Content: "No data found",
}

func ConectDB()(conection *sql.DB)  {
	Driver := "mysql"
	User := "root"
	Password := "root"
	NameDB := "api_food"

	conection,err := sql.Open(Driver,User+":"+Password+"@tcp(127.0.0.1:8889)/"+NameDB) 
	if err != nil {
		panic(err.Error())
	}
	return conection
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/foods",getFoods).Methods("GET")
	router.HandleFunc("/food/{id}",foodById).Methods("GET")
	router.HandleFunc("/foodByCategory/{category}",foodByCategory).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000",router))
}

func getFoods(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	conectionDB := ConectDB()
	getAllFood,err := conectionDB.Query("SELECT id,name,category,description,url_image FROM foods WHERE status = 1");
	if err!= nil {
		panic(err.Error())
	}
	var foods = food{}
	arrFood := []food{}

	for getAllFood.Next() {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var id int
		var name,category,description,url_image string
		err = getAllFood.Scan(&id,&name,&category,&description,&url_image)
		if err!= nil {
			panic(err.Error())
		}
		foods.ID = id
		foods.Name = name
		foods.Category = category
		foods.Description = description
		foods.UrlImage = url_image
		arrFood = append(arrFood, foods)
	}
	if  len(arrFood) >0 {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(arrFood)
		return
	}else{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg_err)
	}
}

func foodByCategory(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	foodCategory := strings.ToValidUTF8(vars["category"],"")
	conectionDB := ConectDB()
	getFood,err := conectionDB.Query("SELECT id,name,category,description,url_image FROM foods WHERE status = 1 and category = ?",foodCategory);
	if err!= nil {
		panic(err.Error())
	}
	var foodFounded = food{}
	arrFood := []food{}
	count := 0
	for getFood.Next() {
		var id int
		var name,category,description,url_image string
		err = getFood.Scan(&id,&name,&category,&description,&url_image)
		if err!= nil {
			panic(err.Error())
		}
		foodFounded.ID = id
		foodFounded.Name = name
		foodFounded.Category = category
		foodFounded.Description = description
		foodFounded.UrlImage = url_image
		count +=1
		arrFood = append(arrFood, foodFounded)
	}
	if  count>0 {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(arrFood)
		return
	}else{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg_err)
	}
}


func foodById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	foodID,err := strconv.Atoi(vars["id"])
	if err!=nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"Invalid ID")
		return
	}

	conectionDB := ConectDB()
	getFood,err := conectionDB.Query("SELECT id,name,category,description,url_image FROM foods WHERE status = 1 and id = ?",foodID);
	if err!= nil {
		panic(err.Error())
	}
	var foodFounded = food{}
	count := 0
	for getFood.Next() {
		var id int
		var name,category,description,url_image string
		err = getFood.Scan(&id,&name,&category,&description,&url_image)
		if err!= nil {
			panic(err.Error())
		}
		foodFounded.ID = id
		foodFounded.Name = name
		foodFounded.Category = category
		foodFounded.Description = description
		foodFounded.UrlImage = url_image
		count +=1
	}
	if  count>0 {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foodFounded)
		return
	}else{
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg_err)
	}
}
