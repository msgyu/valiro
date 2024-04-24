package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/msgyu/golang-validation/interceptor"
)

type User struct {
	Username string `json:"username" validate:"required,min=8,max=10" label:"ユーザー名"`
	Age      int    `json:"age" validate:"gte=0,lte=130" label:"年齢"`
	Email    string `json:"email" validate:"required,email,max=10" label:"メールアドレス"`
	Repeat   int    `json:"repeat" validate:"eqfield=Age" label:"年齢の確認"`
}

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		user    User
		wantErr bool
	}{
		"valid user": {
			user: User{
				Username: "validUser008",
				Age:      30,
				Email:    "valid@example.com",
				Repeat:   30,
			},
			wantErr: false,
		},
		"invalid email": {
			user: User{
				Username: "validUser123",
				Age:      25,
				Email:    "invalid", // 無効なメールアドレスでエラーが発生
				Repeat:   25,
			},
			wantErr: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			err := interceptor.Validate(c.user)
			if (err != nil) != c.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}

func TestCreateFieldErrorMap(t *testing.T) {
	cases := map[string]struct {
		model  User
		output interceptor.FieldErrorMap
	}{
		"no error": {
			model: User{
				Username: "validUser008",
				Age:      30,
				Email:    "",
				Repeat:   30,
			},
			output: interceptor.FieldErrorMap{},
		},
		"required fields": {
			model: User{},
			output: interceptor.FieldErrorMap{
				"Username": []string{"必須項目です"},
				"Email":    []string{"必須項目です"},
			},
		},
		"invalid email": {},
	}

	for name, c := range cases {
		t.Parallel()
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := interceptor.Validate(c.model)
			output := interceptor.CreateFieldErrorMap(err)
			if diff := cmp.Diff(output, c.output); diff != "" {
				t.Errorf("ValidateAndPrintErrors() = %v, want %v", output, c.output)
			}
		})
	}
}
