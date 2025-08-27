package model

import (
  "time"
)

type User struct {
  ID  uint `json:"id" gorm:"primaryKey"`
  Phone string `json:"phone" gorm:"uniqueIndex"`
  CreatedAt time.Time `json:"createdat"`
}

type MinUser struct {
  Phone string `json:"phone"`
}

type SignUser struct {
  Phone string `json:"phone"`
  Code string `json:"code"`
}
