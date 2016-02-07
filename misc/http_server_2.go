package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "log"
)

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

var users []User

func UserHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if req.Method == "GET" {
        fmt.Fprint(w, "<ul>")
        for _, user := range users {
            fmt.Fprint(w, fmt.Sprintf("<li>%s</li>", user.Name))
        }
        fmt.Fprint(w, "</ul>")
    } else if req.Method == "POST" {
        var user User

        decoder := json.NewDecoder(req.Body)
        err := decoder.Decode(&user)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)

        w.WriteHeader(http.StatusCreated)

    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/users", UserHandler)

    http.ListenAndServe(":5000", nil)
}
