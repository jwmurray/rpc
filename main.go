package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

// returned error indicates taht call will not later return a value
func (a *API) GetByName(name string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == name {
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	// needle := GetByName(item.Title)
	// needle.Body = item.Body
	// return item

	var changed Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = database[idx]
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func (a *API) GetDB(item Item, reply *[]Item) error {
	*reply = database
	return nil
}

func log_fatal_on_error(err_string string, err error) {
	if err != nil {
		log.Fatal(err_string, err)
	}
}

func main() {

	var api = new(API)

	err := rpc.Register(api)
	log_fatal_on_error("error registering API", err)

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	log_fatal_on_error("listener error: ", err)

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	log_fatal_on_error("error serving: ", err)

	fmt.Println("Initial database ", database)
	// a := Createitem("first", "first Body")
	// b := Createitem("second", "second Body")
	// c := Createitem("third", "third Body")
	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("database ", database)
	// DeleteItem(b)
	// fmt.Println("database ", database)
	// EditItem(Createitem("first", "new first Body"))
	// fmt.Println("database ", database)

	// x := GetByName("first")
	// y := GetByName("fourth")
	// fmt.Println(x,y)
	// fmt.Println("database ", database)
}
