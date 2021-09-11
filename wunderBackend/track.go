package main

import (
	// "fmt"

	"fmt"
	"net/http"
	"strconv"
)

func track(w http.ResponseWriter, r *http.Request) {

	fmt.Println("track")
	r.ParseForm()
	var ab=1
	ab,_=strconv.Atoi(r.FormValue("sec"))
	fmt.Println(ab)
	if c!=0 {
		if(!a[ab]){count+=1;}
		a[ab] = true

	}
	prog(w,r)
}
