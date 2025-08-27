package service

import (
  "time"
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

// GPT HELPED ME IN THIS PART :D
func CanRequestOTP(phone string) (bool, error) {
	key := "otp:count:" + phone

	// increment request count
	count, err := repository.RedisDB.Incr(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		// first time → set 10 min expiry
		repository.RedisDB.Expire(context.Background(), key, 10*time.Minute)
	}

	if count > 3 {
		return false, nil // limit exceeded
	}

	return true, nil
}
