package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/stat/query_overview", data)
	router.HandleFunc("/stat/query_overview/{dataId}", data)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func data(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataId := vars["dataId"]
	dbgInfs := []dataStruct{
		//dataStruct{"debug", `dataId: "test.txt" Not Found`, "Cynhard"},
		dataStruct{dataId, "data string", "mageek"},
	}
	fmt.Fprintln(w, dbgInfs)
}

type dataStruct struct {
Level  string `json:"level,omitempty"` // Level解析为level,忽略空值
Msg    string `json:"message"`                     // Msg解析为message
author string `json:"-"`
}

