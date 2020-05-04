package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var prompt string = "Enter a file path. "


type Person struct {
  fname string
	lname string }

var pSlice = make([]Person, 0)


func main(){
	
	
	fmt.Print("%s", prompt)
	var filename string
	fmt.Scan(&filename)
	dat, _ := ioutil.ReadFile(filename)
	dat_s := string(dat)
	f_string := strings.Replace(dat_s, string('\n'), " ", -1)
	string_splits := strings.Split(string(f_string), " ")
	//We have one long string with newline characters revmoved,
	//Now each word is in a separate cell of a slice


	//Now, we take the slice of strings and turn it into a slice
	//of person structs...
	for i:=0; i < len(string_splits) - 1; i += 2 {
		pSlice = append(pSlice, Person{fname: string(string_splits[i]),
							lname: string(string_splits[i+1])})
	}

	//and now we can print it
	for _, v := range pSlice {
       fmt.Printf("%s ", v.fname)
       fmt.Printf("%s\n", v.lname)
  }

}