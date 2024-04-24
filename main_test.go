package valiro

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/msgyu/valiro/field_error_map"
)

type User struct {
	Name   string `json:"name" validate:"required,min=3,max=10" label:"ユーザー名"`
	Age    int    `json:"age" validate:"gte=0,lte=130" label:"年齢"`
	Email  string `json:"email" validate:"required,email" label:"メールアドレス"`
	Repeat int    `json:"repeat" validate:"eqfield=Age" label:"年齢の確認"`
}

func TestCreateFieldErrorMap(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		model    User
		expected field_error_map.FieldErrorMap
	}{
		"no_error": {
			model: User{
				Name:   "tanaka",
				Age:    30,
				Email:  "test@example.com",
				Repeat: 30,
			},
			expected: nil,
		},
		"required_fields": {
			model: User{},
			expected: field_error_map.FieldErrorMap{
				"name":  []string{"必須項目です"},
				"email": []string{"必須項目です"},
			},
		},
		"max_error": {
			model: User{
				Name:   "validUser008",
				Age:    30,
				Email:  "test@example.com",
				Repeat: 30,
			},
			expected: field_error_map.FieldErrorMap{
				"name": []string{"10文字以下である必要があります"},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			Init()
			output := Validate(c.model)
			if diff := cmp.Diff(c.expected, output); diff != "" {
				t.Errorf("result mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
