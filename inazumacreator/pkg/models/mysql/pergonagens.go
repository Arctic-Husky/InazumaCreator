package mysql

import(
  "database/sql"
  "github.com/Arctic-Husky/InazumaCreator/pkg/models"
)

type PersonagemModel struct{
  DB *sql.DB
}

func(m *PersonagemModel)Insert(title,content,expires string)(int,error){
  stmt := ``

  result,err := m.DB.Exec(stmt, title, content, expires)
  if err != nil{
    return 0, err
  }
  id,err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  
  return int(id), nil
}

func(m *PersonagemModel)Get(id int)(*models.Personagem, error){
  stmt := ``

  row := m.DB.QueryRow(stmt, id)

  s := &models.Personagem{}

  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) // <--

  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  } else if err != nil{
    return nil, err
  }

  return s, nil
}

func(m *PersonagemModel)Latest()([]*models.Personagem, error){
  stmt := ``

  rows , err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  snippets := []*models.Personagem{}
  for rows.Next(){
    s := &models.Personagem{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) // <--
    if err != nil{
      return nil,err
    }
    snippets = append(snippets,s)
  }
  err = rows.Err()
  if err != nil{
    return nil,err
  }
  return snippets,nil
}