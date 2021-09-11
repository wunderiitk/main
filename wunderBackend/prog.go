package main

import (
	// "fmt"

	"fmt"
	"net/http"
	"strconv"
)

func prog(w http.ResponseWriter, r *http.Request) {

	coo := &http.Cookie{
		Name: "prog",
		Value: strconv.Itoa(count),
	}
	fmt.Println("prog")
	http.SetCookie(w,coo)
	fmt.Println(count)

}
