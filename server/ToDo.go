package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type Task struct {
    Id string `json:"id"`
    Desc string `json:"desc"`
    Done bool `json:"done"`
}

var tasks map[string] *Task

func GetToDo(rw http.ResponseWriter, req * http.Request) {

    vars := mux.Vars(req)
    task := tasks[vars["id"]]

    js, err := json.Marshal(task)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(rw, string(js))
}

func UpdateToDo(rw http.ResponseWriter, req * http.Request) {

    vars := mux.Vars(req)

    task:= tasks[vars["id"]]

    dec:= json.NewDecoder(req.Body)
    err:= dec.Decode( & task)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    task.Id = vars["id"]

    retjs, err:= json.Marshal(task)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(rw, string(retjs))
}

func DeleteToDo(rw http.ResponseWriter, req * http.Request) {

    vars := mux.Vars(req)
    delete(tasks, vars["id"])
    fmt.Fprint(rw, "{status : 'success'}")
}

func AddToDo(rw http.ResponseWriter, req * http.Request) {

    task:= new(Task)

    dec:= json.NewDecoder(req.Body)
    err:= dec.Decode( & task)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    tasks[task.Id] = task

    retjs, err:= json.Marshal(task)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(rw, string(retjs))
}

func GetToDos(rw http.ResponseWriter, req * http.Request) {

    v := make([]*Task, 0, len(tasks))

    for  _, value := range tasks {
       v = append(v, value)
    }
    js, err:= json.Marshal(v)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(rw, string(js))
}

func main() {

    var port = 9001
    router:= mux.NewRouter()
    tasks = make(map[string] *Task)
    router.HandleFunc("/tasks", GetToDos).Methods("GET")
    router.HandleFunc("/tasks", AddToDo).Methods("POST")
    router.HandleFunc("/tasks/{id}", GetToDo).Methods("GET")
    router.HandleFunc("/tasks/{id}", UpdateToDo).Methods("PUT")
    router.HandleFunc("/tasks/{id}", DeleteToDo).Methods("DELETE")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../")))
    
    fmt.Println("Listening on port", port)
    http.ListenAndServe("localhost:" + strconv.Itoa(port), router)
}