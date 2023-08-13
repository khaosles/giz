package formater

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/khaosles/giz/convertor"
	"github.com/khaosles/giz/strutil"
	"github.com/khaosles/giz/validator"
	"golang.org/x/exp/constraints"
)

/*
   @File: formater.go
   @Author: khaosles
   @Time: 2023/8/13 10:39
   @Desc:
*/

// Comma add comma to a number value by every 3 numbers from right. ahead by symbol char.
// if value is invalid number string eg "aa", return empty string
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string {
	if validator.IsInt(value) {
		v, err := convertor.ToInt(value)
		if err != nil {
			return ""
		}
		return symbol + commaInt(v)
	}

	if validator.IsFloat(value) {
		v, err := convertor.ToFloat(value)
		if err != nil {
			return ""
		}
		return symbol + commaFloat(v)
	}

	if strutil.IsString(value) {
		v := fmt.Sprintf("%v", value)
		if validator.IsNumberStr(v) {
			return symbol + commaStr(v)
		}
		return ""
	}

	return ""
}

// Pretty data to JSON string.
func Pretty(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// PrettyToWriter pretty encode data to writer.
func PrettyToWriter(v any, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}
