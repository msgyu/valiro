package valiro

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/msgyu/valiro/field_error_map"
)

var validater *validator.Validate

func Init() {
	validater = validator.New()
	// タグのlabelを使ってフィールド名を登録
	// RegisterTagNameFunc は、StructFields の代替名を取得する関数を登録する。
	// これにより、バリデーションエラーメッセージにおいて、構造体フィールドの JSON タグの名前を使用することができます。
	validater.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate関数をジェネリクスを使って定義
func Validate[T any](v T) field_error_map.FieldErrorMap {
	var validateErrMap field_error_map.FieldErrorMap
	validateErr := validater.Struct(v)
	if validateErr != nil {
		validateErrMap = field_error_map.CreateFieldErrorMap(validateErr)
	}
	return validateErrMap
}
