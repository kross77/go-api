package main

import (
	"fmt"
	"net/http"
	"runtime"
	"reflect"
	
)

func main() {
	http.HandleFunc("/", hello)
	//bind := fmt.Sprintf("%s:%s", "localhost", 8713)
	//fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe("localhost:8713", nil)
	if err != nil {
		panic(err)
	}
	//mysql.ConnectToMySQL();
}

func hello(res http.ResponseWriter, req *http.Request) {
	n := Name{"Alexander", "Krasouski", "asdfasdfasd"}
	fmt.Fprintf(res, "hello "+n.getName()+", "+mysql.InsertRecord()+", world from %s", runtime.Version())
}

type Name struct {
	First string `mysql:"First"`
	Second string `mysql:"Secound"`
	Third string `mysql:"third"`
}

func (b *Name) getName() string {
	return b.First + " " + b.Second;
}

func (b *Name) getField() string{
	retS := "name: "
	s := reflect.ValueOf(b).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		retS += fmt.Sprintf("\n%d: %s %s = %v %s", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface(), typeOfT.Field(i).Tag.Get("mysql"))
	}
	return retS;
}
