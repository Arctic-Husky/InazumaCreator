package main

import (
  "database/sql"
  "net/http"
  "log"
  "flag"
  "os"
  
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  personagens *mysql.PersonagemModel
}

func main() {
  addr := flag.String("addr", ":4000", "Porta da Rede")

  dsn := flag.String("dsn",
                     "1rJx8vlHTM:Nx2KvlJg58@tcp(remotemysql.com)/1rJx8vlHTM?parseTime=true", 
                     "MySql DSN")

  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil{
    errorLog.Fatal(err)
  }
  defer db.Close()
  
  app := &application{
    errorLog:errorLog,
    infoLog:infoLog,
  }
  
  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }

  app.infoLog.Printf("Inicializando o servidor na porta: %s\n", *addr)
  err = srv.ListenAndServe()
  app.errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)
  if err!= nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err
  }
  return db, nil
}
