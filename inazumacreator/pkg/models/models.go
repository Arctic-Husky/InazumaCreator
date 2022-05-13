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
  habilidade1 string
  habilidade2 string
  posicao string
  elemento string
  Created time.Time
  Expires time.Time
}