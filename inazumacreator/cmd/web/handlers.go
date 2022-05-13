package main

//go run cmd/web/*
import (
	"fmt"
	//"html/template"
	"net/http"
	"strconv"
  
  "github.com/Arctic-Husky/Aprendendo-GO/pkg/models"
)

func(app *application) home(rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  } 

  // files := []string{
  //   "./ui/html/home.page.tmpl.html",
  //   "./ui/html/base.layout.tmpl.html",
  //   "./ui/html/footer.partial.tmpl.html",
  // }
  // ts, err := template.ParseFiles(files...)
  // if err != nil{
  //   app.serverError(rw, err)
  //   // app.errorLog.Println(err.Error())
  //   // http.Error(rw, "Internal Error",500)
  //   return
  // }
  // err = ts.Execute(rw, nil)
  // if err != nil{
  //   app.serverError(rw, err)
  //   return
  // }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw,err)
    return
  }
  for _,s := range snippets{
    fmt.Fprintf(rw, "%v \n",s)
  }
}

//http://localhost:4000/snippet?id=123
func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    // http.NotFound(rw, r)
    return
  }

  s, err := app.snippets.Get(id)
  if err == models.ErrNoRecord{
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }

  fmt.Fprintf(rw, "%v",s)
  
  //fmt.Fprintf(rw, "Exibir o Snippet de ID: %d", id)
}

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    // http.Error(rw, "Metodo não permitido", http.StatusMethodNotAllowed)
    return
  }

  title := "ALMOÇO DE HJ"
  content := "RICKS BUERGUE RICKS BURGIER"
  expires := "7"

  id, err := app.snippets.Insert(title,content,expires)
  if err != nil{
    app.serverError(rw,err)
    return
  }

  http.Redirect(rw,r,fmt.Sprintf("/snippet:id=%d",id), http.StatusSeeOther)
}