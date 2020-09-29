package main

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "github.com/go-redis/redis"
)

var client *redis.Client

type Comment struct {
    Author string
    Comment string
}

type MainPage struct {
    Title string
    Subtitle string
    JumboColor string
    JumboBg string
    BodyColor string
    Comments []Comment
}

var p MainPage = MainPage{
    Title: "Bad Discord", 
    Subtitle: "\"It's just worse\"",
    JumboColor: "whitesmoke",
    JumboBg: "dimgrey",
    BodyColor: "grey",
    Comments: []Comment{},
}

func mainPageGetHandler(w http.ResponseWriter, r *http.Request) {
    //comments, _ := client.LRange("comments", 0, 10).Result()
    t, _ := template.ParseFiles("templates/main.html")
    t.Execute(w,p)
}

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
    client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    r := mux.NewRouter()
    r.HandleFunc("/", mainPageGetHandler).Methods("GET")
    r.HandleFunc("/", mainPagePostHandler).Methods("Post")
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}