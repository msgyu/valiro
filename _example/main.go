package main

import (
	"encoding/json"
	"fmt"

	"github.com/msgyu/golang-validation/interceptor"
)

type User struct {
	Username string `json:"username" validate:"required,min=8,max=10" label:"ユーザー名"`
	Age      int    `json:"age" validate:"gte=0,lte=130" label:"年齢"`
	Email    string `json:"email" validate:"required,email,max=10" label:"メールアドレス"`
	Repeat   int    `json:"repeat" validate:"eqfield=Age" label:"年齢の確認"`
}

func main() {
	interceptor.Init()

	user := User{
		Username: "12333333333333333333333333333",
		Age:      25,
		Email:    "11111111111111111111",
		Repeat:   25,
	}
	err := interceptor.Validate(user)
	if err != nil {
		mapError := interceptor.CreateFieldErrorMap(err)
		m, err := json.Marshal(mapError)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(m))
	}
}
