package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

  
  "github.com/Arctic-Husky/InazumaCreator/pkg/models"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  } 

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

  fmt.Fprintf(rw, "%v",s)
}

func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){
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
    "./ui/html/base.tmpl.html",
    "./ui/html/footer.tmpl.html",
    "./ui/html/top.tmpl.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    // app.errorLog.Println(err.Error())
    // http.Error(rw, "Internal Error",500)
    return
  }
  err = ts.Execute(rw, s)
  if err != nil{
    app.serverError(rw, err)
    return
  }

  fmt.Fprintf(rw, "%v",s)
  
  //fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func (app *application) createPersonagem(rw http.ResponseWriter, r *http.Request){
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
  // nome := "Teste"
  // hab1 := "god hand"
  // hab2 := "back tornado"
  // posicao := "FW"
  // elemento := "Terra"
  // expires := "7"

  // id, err := app.personagens.Insert(nome,hab1,hab2,posicao,elemento,expires)
  // if err != nil{
  //   app.serverError(rw,err)
  //   return
  // }

  // http.Redirect(rw,r,fmt.Sprintf("/personagem?id=%d",id), http.StatusSeeOther)
}

// func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
//   if r.Method != "POST"{
//     rw.Header().Set("Allow","POST")
//     app.clientError(rw, http.StatusMethodNotAllowed)
//     // http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
//     return
//   }

//   title := "ALMOÇO DE HJ"
//   content := "RICKS BUERGUE RICKS BURGIER"
//   expires := "7"

//   id, err := app.personagens.Insert(title,content,expires)
//   if err != nil{
//     app.serverError(rw,err)
//     return
//   }

//   http.Redirect(rw,r,fmt.Sprintf("/snippet:id=%d",id), http.StatusSeeOther)
// }