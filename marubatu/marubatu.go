package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	http.HandleFunc("/form", form1)
	http.HandleFunc("/reset", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Vals struct {
	V1  string
	V2  string
	V3  string
	V4  string
	V5  string
	V6  string
	V7  string
	V8  string
	V9  string
	Msg string
}

var gameBoard []string = []string{"", "", "", "", "", "", "", "", ""}
var check bool = true
var count int = 0
var marubatu string

func index(w http.ResponseWriter, r *http.Request) {
	vals := new(Vals)
	if check {
		marubatu = "〇"
		vals.Msg = "×の手番"
		check = false
	} else {
		marubatu = "×"
		vals.Msg = "〇の手番"
		check = true
	}
	gameBoard[0] = ""
	gameBoard[1] = ""
	gameBoard[2] = ""
	gameBoard[3] = ""
	gameBoard[4] = ""
	gameBoard[5] = ""
	gameBoard[6] = ""
	gameBoard[7] = ""
	gameBoard[8] = ""
	vals.V1 = gameBoard[0]
	vals.V2 = gameBoard[1]
	vals.V3 = gameBoard[2]
	vals.V4 = gameBoard[3]
	vals.V5 = gameBoard[4]
	vals.V6 = gameBoard[5]
	vals.V7 = gameBoard[6]
	vals.V8 = gameBoard[7]
	vals.V9 = gameBoard[8]
	count = 0
	t, _ := template.ParseFiles("temp.html")
	t.Execute(w, vals)

}

func form1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp.html")
	vals := new(Vals)
	id := r.FormValue("id")

	if check {
		marubatu = "〇"
		vals.Msg = "×の手番"
		check = false
	} else {
		marubatu = "×"
		vals.Msg = "〇の手番"
		check = true
	}
	count++

	if id == "1" {
		gameBoard[0] = marubatu
	} else if id == "2" {
		gameBoard[1] = marubatu
	} else if id == "3" {
		gameBoard[2] = marubatu
	} else if id == "4" {
		gameBoard[3] = marubatu
	} else if id == "5" {
		gameBoard[4] = marubatu
	} else if id == "6" {
		gameBoard[5] = marubatu
	} else if id == "7" {
		gameBoard[6] = marubatu
	} else if id == "8" {
		gameBoard[7] = marubatu
	} else if id == "9" {
		gameBoard[8] = marubatu
	}

	if gameBoard[0] == "〇" && gameBoard[1] == "〇" && gameBoard[2] == "〇" || gameBoard[3] == "〇" && gameBoard[4] == "〇" && gameBoard[5] == "〇" || gameBoard[6] == "〇" && gameBoard[7] == "〇" && gameBoard[8] == "〇" {
		vals.Msg = "まるの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	} else if gameBoard[0] == "〇" && gameBoard[3] == "〇" && gameBoard[6] == "〇" || gameBoard[1] == "〇" && gameBoard[4] == "〇" && gameBoard[7] == "〇" || gameBoard[2] == "〇" && gameBoard[5] == "〇" && gameBoard[8] == "〇" {
		vals.Msg = "まるの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	} else if gameBoard[0] == "〇" && gameBoard[4] == "〇" && gameBoard[8] == "〇" || gameBoard[2] == "〇" && gameBoard[6] == "〇" && gameBoard[4] == "〇" {
		vals.Msg = "まるの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	} else if gameBoard[0] == "×" && gameBoard[1] == "×" && gameBoard[2] == "×" || gameBoard[3] == "×" && gameBoard[4] == "×" && gameBoard[5] == "×" || gameBoard[6] == "×" && gameBoard[7] == "×" && gameBoard[8] == "×" {
		vals.Msg = "ばつの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	} else if gameBoard[0] == "×" && gameBoard[3] == "×" && gameBoard[6] == "×" || gameBoard[1] == "×" && gameBoard[4] == "×" && gameBoard[7] == "×" || gameBoard[2] == "×" && gameBoard[5] == "×" && gameBoard[8] == "×" {
		vals.Msg = "ばつの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	} else if gameBoard[0] == "×" && gameBoard[4] == "×" && gameBoard[8] == "×" || gameBoard[2] == "×" && gameBoard[6] == "×" && gameBoard[4] == "×" {
		vals.Msg = "ばつの勝利"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	}
	if count == 9 {
		vals.Msg = "引き分け"
		gameBoard[0] = ""
		gameBoard[1] = ""
		gameBoard[2] = ""
		gameBoard[3] = ""
		gameBoard[4] = ""
		gameBoard[5] = ""
		gameBoard[6] = ""
		gameBoard[7] = ""
		gameBoard[8] = ""
		count = 0
	}
	vals.V1 = gameBoard[0]
	vals.V2 = gameBoard[1]
	vals.V3 = gameBoard[2]
	vals.V4 = gameBoard[3]
	vals.V5 = gameBoard[4]
	vals.V6 = gameBoard[5]
	vals.V7 = gameBoard[6]
	vals.V8 = gameBoard[7]
	vals.V9 = gameBoard[8]

	t.Execute(w, vals)
}
