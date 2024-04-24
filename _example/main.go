package main

import (
	"encoding/json"
	"fmt"

	"github.com/msgyu/valiro"
)

type User struct {
	Username string `json:"username" validate:"required,min=8,max=10" label:"ユーザー名"`
	Age      int    `json:"age" validate:"gte=0,lte=130" label:"年齢"`
	Email    string `json:"email" validate:"required,email,max=10" label:"メールアドレス"`
	Repeat   int    `json:"repeat" validate:"eqfield=Age" label:"年齢の確認"`
}

func example() {
	valiro.Init()

	user := User{
		Username: "12333333333333333333333333333",
		Age:      25,
		Email:    "11111111111111111111",
		Repeat:   25,
	}

	validateErrMap := valiro.Validate(user)
	if len(validateErrMap) > 0 {
		m, err := json.Marshal(validateErrMap)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(m))
	}
}
