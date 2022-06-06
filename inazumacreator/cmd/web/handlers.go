package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
  "time"

  
  "github.com/Arctic-Husky/InazumaCreator/pkg/models"
)

type personagem struct{
  Modelo string
  Nome string
  Habilidade1 string
  Habilidade2 string
  Posicao string
  Elemento string
  Expires time.Time
}

func(app *application) welcome(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  } 
  
  files := []string{
    "./ui/html/welcome.tmpl.html",
    "./ui/html/footer.tmpl.html",
    "./ui/html/topPrincipal.tmpl.html",
    "./ui/html/base.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  

  personagens, err := app.personagens.Latest()
  if err != nil{
    app.serverError(rw,err)
    return
  }

  // MUDAR AQUI
  files := []string{
    "./ui/html/home.tmpl.html",
    "./ui/html/footer.tmpl.html",
    "./ui/html/top.tmpl.html",
    "./ui/html/base.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    // app.errorLog.Println(err.Error())
    // http.Error(rw, "Internal Error",500)
    return
  }
  err = ts.Execute(rw, personagens)
  if err != nil{
    app.serverError(rw, err)
    return
  }

  // personagens, err := app.personagens.Latest()
  // if err != nil{
  //   app.serverError(rw,err)
  //   return
  // }
  // for _,s := range personagens{
  //   fmt.Fprintf(rw, "%v \n",s)
  // }
}

//http://localhost:4000/snippet?id=123
func(app *application) showPersonagem(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    // http.NotFound(rw, r)
    return
  }

  s, err := app.personagens.Get(id)
  if err == models.ErrNoRecord{
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }


  files := []string{
    "./ui/html/show.tmpl.html",
    "./ui/html/footer.tmpl.html",
    "./ui/html/top.tmpl.html",
    "./ui/html/base.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, s)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  //fmt.Fprintf(rw, "%v",s)
}

func (app *application) createPersonagem(rw http.ResponseWriter, r *http.Request){

  pers := personagem{
    Modelo: r.URL.Query().Get("modelo"),
    Nome: r.URL.Query().Get("fname"),
    Habilidade1: r.URL.Query().Get("check1"),
    Habilidade2: r.URL.Query().Get("check2"),
    Posicao: r.URL.Query().Get("posicao"),
    Elemento: r.URL.Query().Get("elemento"),
  }

  
  
  if len(pers.Modelo) > 0{
    id,err := app.personagens.Insert(pers.Modelo, pers.Nome, pers.Habilidade1, pers.Habilidade2, pers.Posicao, pers.Elemento, "7")
  if err != nil{
    app.serverError(rw,err)
    return
  }

  http.Redirect(rw,r,fmt.Sprintf("/personagem?id=%d",id), http.StatusSeeOther)
  }

  
  
  files := []string{
    "./ui/html/creator.tmpl.html",
    "./ui/html/footer.tmpl.html",
    "./ui/html/top.tmpl.html",
    "./ui/html/base.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, nil)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}