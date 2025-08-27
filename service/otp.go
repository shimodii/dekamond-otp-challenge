package service

import (
	"context"
	"crypto/rand"
	//"fmt"

	"math/big"

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
  ourCode, _ := repository.RedisDB.Get(context.Background(), phone).Result()
  //fmt.Println("our code:", ourCode)
  //fmt.Println("your code:", code)

  if ourCode == code {
    return true
  }
  return false
}
