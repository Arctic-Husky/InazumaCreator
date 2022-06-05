package mysql

import(
  "database/sql"
  "github.com/Arctic-Husky/InazumaCreator/pkg/models"
)

type PersonagemModel struct{
  DB *sql.DB
}

func(m *PersonagemModel)Insert(nome,habilidade1,habilidade2,posicao,elemento,expires string)(int,error){
  stmt := `INSERT INTO Personagens (nome, habilidade1, habilidade2, posicao, elemento,created, expires)
  VALUES(?,?,?,?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY))`

  result,err := m.DB.Exec(stmt, nome, habilidade1, habilidade2, posicao, elemento, expires)
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
  stmt := `SELECT id, nome, habilidade1, habilidade2, posicao, elemento, created, expires FROM Personagens WHERE expires > UTC_TIMESTAMP() AND id = ?`

  row := m.DB.QueryRow(stmt, id)

  s := &models.Personagem{}

  err := row.Scan(&s.ID, &s.Nome, &s.Habilidade1, &s.Habilidade2, &s.Posicao, &s.Elemento, &s.Created,&s.Expires) // <--

  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  } else if err != nil{
    return nil, err
  }

  return s, nil
}

func(m *PersonagemModel)Latest()([]*models.Personagem, error){
  stmt := `SELECT id, nome, habilidade1, habilidade2, posicao, elemento, created, expires FROM Personagens WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

  rows , err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  personagens := []*models.Personagem{}
  for rows.Next(){
    s := &models.Personagem{}
    err = rows.Scan(&s.ID, &s.Nome, &s.Habilidade1, &s.Habilidade2, &s.Posicao, &s.Elemento, &s.Created,&s.Expires)// <--
    if err != nil{
      return nil,err
    }
    personagens = append(personagens,s)
  }
  err = rows.Err()
  if err != nil{
    return nil,err
  }
  return personagens,nil
}