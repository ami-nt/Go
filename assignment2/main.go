package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/AminaAmangeldi15/Go/assignment2/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Register(u User){

	result, err := db.Exec("insert into users (name, surname, login, password) values ($1, $2, $3, $4)", u.Name, u.Surname, u.Login, u.Password)
	if err != nil{
		panic(err)
	}
	// fmt.Println(result.LastInsertId())  
    fmt.Println("You registered a new user")
	result.LastInsertId()
}

func Authorize(a Authorization){
	result, err := db.Exec("select * from users where login = $1 and password = $2", a.Login, a.Password)
	if err != nil{
		panic(err)
	}
	fmt.Println("You autorized!")
	result.LastInsertId()
}

func AddItem(i Item){
	result,err := db.Exec("insert into items (name, price, rating) values ($1, $2, $3)", i.Name, i.Price, i.Rating)
    if err != nil{
        panic(err)
    }
	fmt.Println("You added a new item")
	result.LastInsertId()
}

func SearchItemByName(name string){
	rows, err := db.Query("select * from items where name like $1", name)
    if err != nil {
        panic(err)
    }
    
    items := []Item{}
     
    for rows.Next(){
        i := Item{}
        err := rows.Scan(&i.Name, &i.Price, &i.Rating)
        if err != nil{
            fmt.Println(err)
            continue
        }
        items  = append(items , i)
    }

	fmt.Println("Searching item is:")
    for _, i := range items {
        fmt.Println(i.Name, i.Price, i.Rating)
    }
}

func FilterByPrice(from, to int){
	rows, err := db.Query("select * from items where price between $1 and $2", from, to)
    if err != nil {
        panic(err)
    }
    
    items := []Item{}
     
    for rows.Next(){
        i := Item{}
        err := rows.Scan(&i.Name, &i.Price, &i.Rating)
        if err != nil{
            fmt.Println(err)
            continue
        }
        items  = append(items , i)
    }

	if len(items) == 0{
		fmt.Printf("Searching items, which price is from %s to %s do not exist...\n", strconv.Itoa(from), strconv.Itoa(to))
	} else {
		fmt.Printf("Searching items, which price is from %s to %s are:\n", strconv.Itoa(from), strconv.Itoa(to))
		for _, i := range items {
			fmt.Println(i.Name, i.Price, i.Rating)
		}
	}
}

func FilterByRating(from, to int){
	rows, err := db.Query("select * from items where rating between $1 and $2", from, to)
    if err != nil {
        panic(err)
    }
    
    items := []Item{}
     
    for rows.Next(){
        i := Item{}
        err := rows.Scan(&i.Name, &i.Price, &i.Rating)
        if err != nil{
            fmt.Println(err)
            continue
        }
        items  = append(items , i)
    }

	if len(items) == 0{
		fmt.Printf("Searching items, which rating is from %s to %s do not exist...\n", strconv.Itoa(from), strconv.Itoa(to))
	} else {
		fmt.Printf("Searching items, which rating is from %s to %s are:\n", strconv.Itoa(from), strconv.Itoa(to))
		for _, i := range items {
			fmt.Println(i.Name, i.Price, i.Rating)
		}
	}
}

func Rate(r Rating){
	result,err := db.Exec("insert into rating (item_name, rating) values ($1, $2)", r.Item_name, r.Rating)
	if err != nil{
		panic(err)
	}
	result.LastInsertId()
	rows, err := db.Query("select * from rating where item_name like $1", r.Item_name)
	if err != nil {
		panic(err)
	}
	
	ratings := []Rating{}
		
	for rows.Next(){
		r := Rating{}
		err := rows.Scan(&r.Item_name, &r.Rating)
		if err != nil{
			fmt.Println(err)
			continue
		}
		ratings = append(ratings, r)
	}
	var sum float64
	for _, r := range ratings{
		sum += r.Rating
	}
	rate := sum / float64(len(ratings ))
	row1, err := db.Query("update items set rating = $1 where name = $2", rate, r.Item_name)
	if err != nil{
		panic(err)
	}
	row1.Close()
}
func main(){
	connStr := "user=postgres password=9792amina dbname=golang sslmode=disable"
    DB, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer DB.Close()

	db = DB

	// result, err := db.Exec("insert into users (name, surname, login, password) values ('Amina', 'Amangeldi', 'login', 'passwprd')")
	// if err != nil{
	// 	panic(err)
	// }
	// u := User{"a", "a", "l1", "p1"}
	// // register(u)
	// Register(u)
	// Authorize(Authorization{"ff", "d"})
	// AddItem(Item{"Iphone 12", 15000, 0})
	// SearchItemByName("Iphone 12")
	// Rate(Rating{"Iphone 12", 4})
	// Rate(Rating{"Iphone 12", 3})
	// Rate(Rating{"Iphone 12", 2})
	SearchItemByName("Iphone 12")
	// FilterByPrice(10000, 200000)
	FilterByRating(2, 4)

}