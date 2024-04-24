package main

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/msgyu/golang-validation/field_error_map"
)

var validater *validator.Validate

func Init() {
	validater = validator.New()
	// タグのlabelを使ってフィールド名を登録
	// RegisterTagNameFunc は、StructFields の代替名を取得する関数を登録する。
	// 例えば、通常の Go フィールド名ではなく、構造体の JSON 表現に指定された名前を使用する場合などです：
	validater.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate関数をジェネリクスを使って定義
func Validate[T any](v T) (field_error_map.FieldErrorMap, error) {
	var validateErrMap field_error_map.FieldErrorMap
	err := validater.Struct(v)
	if err != nil {
		validateErrMap = field_error_map.CreateFieldErrorMap(err)
	}
	return validateErrMap, err
}
