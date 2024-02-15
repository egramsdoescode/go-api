package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/ethandoes-code/go-api/internal/handlers"
    log "github.com/sirupsen/logrus"
)

func main() {
    // Prints logging for files
    log.SetReportCaller(true)

    // Setting up API
    var r *chi.Mux = chi.NewRouter()

    //
    handlers.Handler(r)

    fmt.Println("Starting GO API service...")
    
    fmt.Println(`
         ______     ______        ______     ______   __    
        /\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
        \ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
         \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
          \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ 
    `)
    
    // Starts server
    err := http.ListenAndServe("localhost:8000", r)
    if err != nil {
        log.Error(err)
    }


}
