package models


import (
  "time"
  "errors"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type Personagem struct{
  // a decidir
  ID int
  Nome string
  Habilidade1 string
  Habilidade2 string
  Posicao string
  Elemento string
  Created time.Time
  Expires time.Time
}