package service

import (
	"context"
	"math/big"
  "crypto/rand"

	"github.com/shimodii/dekamond-otp-challenge/repository"
)

func GenCode(length uint) string {
  code := ""
  for i := 0; i< int(length) ; i++ {
    n, err := rand.Int(rand.Reader, big.NewInt(10))
    if err != nil {
      panic(err)
    }
    code += n.String()
  }

  return code
}

func ValidateCode(code string, phone string) bool {
  ourCode := repository.RedisDB.Get(context.Background(), phone).String()
  
  if ourCode == code {
    return true
  }
  return false
}
