package main

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "github.com/go-redis/redis"
)

// unused redis db var
var client *redis.Client

type Comment struct {
    Author string
    Comment string
}

// main page data struct
type MainPage struct {
    Title string
    Subtitle string
    JumboColor string
    JumboBg string
    BodyColor string
    Comments []Comment
}

// global main page var since db currently unimplemented (bad)
var p MainPage = MainPage{
    Title: "Bad Discord", 
    Subtitle: "\"It's just worse\"",
    JumboColor: "whitesmoke",
    JumboBg: "dimgrey",
    BodyColor: "grey",
    Comments: []Comment{},
}

// sets up main page for get request
func mainPageGetHandler(w http.ResponseWriter, r *http.Request) {
    // some code for using redis database
    //comments, _ := client.LRange("comments", 0, 10).Result()
    t, _ := template.ParseFiles("templates/main.html")
    t.Execute(w,p)
}

// gathers discussion info and pushes data to html via template for post request
func mainPagePostHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    c := Comment {
        Author: r.PostForm.Get("author"),
        Comment: r.PostForm.Get("comment"),
    } 
    p.Comments = append([]Comment{c}, p.Comments...)
    if len(p.Comments) > 10 {
        p.Comments = p.Comments[:len(p.Comments) - 1]
    }
    t, _ := template.ParseFiles("templates/main.html")
    t.Execute(w,p)
}

func main() {
    //more code for redis database, left it in for implementation &
    //  avoiding an unused package error
    client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    //using gorilla mux router
    r := mux.NewRouter()
    r.HandleFunc("/", mainPageGetHandler).Methods("GET")
    r.HandleFunc("/", mainPagePostHandler).Methods("Post")
    http.Handle("/", r)
    http.ListenAndServe(":80", nil)
}