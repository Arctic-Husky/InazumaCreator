package main

import(
  "net/http"
)

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
  
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/personagem", app.showPersonagem)
  mux.HandleFunc("/personagem/create", app.createPersonagem)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))

  return mux
}