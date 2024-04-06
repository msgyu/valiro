package interceptor

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
	// タグのlabelを使ってフィールド名を登録
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate関数をジェネリクスを使って定義
func Validate[T any](v T) error {
	return validate.Struct(v)
}

// FieldErrorMap はフィールド名をキーとするエラーメッセージのマップです。
type FieldErrorMap map[string][]string

func (fem FieldErrorMap) String() string {
	var messages []string
	for field, msgs := range fem {
		messages = append(messages, fmt.Sprintf("%s フィールドエラー: %s", field, strings.Join(msgs, "; ")))
	}
	return strings.Join(messages, "\n")
}

func CreateFieldErrorMap(err error) FieldErrorMap {
	errorMessages := make(FieldErrorMap)
	if err == nil {
		return errorMessages
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			message := CreateErrorMessage(fieldError)
			errorMessages[field] = append(errorMessages[field], message)
		}
	}

	return errorMessages
}

// カスタムエラーメッセージを作成する関数
func CreateErrorMessage(fe validator.FieldError) string {
	var message string
	switch fe.Tag() {
	case "required":
		message = "必須項目です"
	case "min":
		minValue := fe.Param()
		message = fmt.Sprintf("最低でも%s文字以上である必要があります", minValue)
	case "max":
		maxValue := fe.Param()
		message = fmt.Sprintf("%s文字以下である必要があります", maxValue)
	case "gte":
		message = "0以上の値である必要があります"
	case "lte":
		message = "130以下の値である必要があります"
	case "eqfield":
		message = "年齢と一致する値である必要があります"
	case "email":
		message = "有効なメールアドレス形式である必要があります"
	default:
		message = fmt.Sprintf("%sバリデーションエラーが発生しました", fe.Tag())
	}
	return message
}
