package validation

import (
	"reflect"
	"regexp"
	"unicode/utf8"
)

var reMail = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
var reSlug = regexp.MustCompile("^[a-zA-Z0-9][\\w-]*$")

func ValidatorRequired(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if str, ok := obj.(string); ok {
		return str != ""
	}
	if b, ok := obj.(bool); ok {
		return b
	}
	if i, ok := obj.(int); ok {
		return i != 0
	}
	if i, ok := obj.(int64); ok {
		return i != 0
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() > 0
	}
	return true
}

func ValidatorEmail(email string) bool {
	return reMail.MatchString(email)
}

func ValidatorMinSize(obj interface{}, min int) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) >= min
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() >= min
	}
	return false
}

func ValidatorMaxSize(obj interface{}, max int) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) <= max
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() <= max
	}
	return false
}

func ValidatorMin(obj interface{}, min int) bool {
	if num, ok := obj.(int); ok {
		return num >= min
	}
	if num, ok := obj.(int64); ok {
		return num >= int64(min)
	}
	return false
}

func ValidatorMax(obj interface{}, max int) bool {
	if num, ok := obj.(int); ok {
		return num <= max
	}
	if num, ok := obj.(int64); ok {
		return num <= int64(max)
	}
	return false
}

func ValidatorMatch(re *regexp.Regexp, s string) bool {
	return re.MatchString(s)
}

func ValidatorIn(obj interface{}, options []interface{}) bool {
	for _, one := range options {
		if obj == one {
			return true
		}
	}
	return false
}

func ValidatorNotIn(obj interface{}, options []interface{}) bool {
	for _, one := range options {
		if obj == one {
			return false
		}
	}
	return true
}

func ValidatorRange(obj interface{}, min, max int) bool {
	return ValidatorMin(obj, min) && ValidatorMax(obj, max)
}
