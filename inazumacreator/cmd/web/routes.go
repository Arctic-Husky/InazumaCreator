package main

import(
  "net/http"
)

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
  
  mux.HandleFunc("/", app.welcome)
  mux.HandleFunc("/personagem", app.showPersonagem)
  mux.HandleFunc("/personagem/create", app.createPersonagem)
  mux.HandleFunc("/lista", app.home)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static",fileServer))

  return mux
}