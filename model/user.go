package model

import (
  "time"
)

type User struct {
  ID  uint `json:"id" gorm:"primaryKey"`
  Phone string `json:"phone" gorm:"uniqueIndex"`
  CreatedAt time.Time `json:"createdat"`
}
