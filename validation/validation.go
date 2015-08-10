package validation

import (
	"fmt"
	"regexp"
)

type ValidationError struct {
	Key     string
	Message string
}

func (this *ValidationError) Error() string {
	return this.Message
}

type Validation struct {
	errors []*ValidationError
	multi  bool
}

func New() *Validation {
	return &Validation{
		errors: make([]*ValidationError, 0, 1),
		multi:  false,
	}
}

func NewMulti() *Validation {
	return &Validation{
		errors: make([]*ValidationError, 0, 20),
		multi:  true,
	}
}

func (this *Validation) HasErrors() bool {
	return len(this.errors) > 0
}

func (this *Validation) Errors() []*ValidationError {
	return this.errors
}

func (this *Validation) AddError(key, message string) {
	this.errors = append(this.errors, &ValidationError{
		Key:     key,
		Message: message,
	})
	return
}

func (this *Validation) should() bool {
	if this.multi {
		return true
	}
	return !this.HasErrors()
}

func (this *Validation) Required(key string, obj interface{}) {
	if !this.should() {
		return
	}

	valid := ValidatorRequired(obj)
	if !valid {
		message := fmt.Sprintf("is required")
		this.AddError(key, message)
	}
}

func (this *Validation) MinSize(key string, obj interface{}, min int) {
	if !this.should() {
		return
	}

	valid := ValidatorMinSize(obj, min)
	if !valid {
		message := fmt.Sprintf("minimum length is %d", min)
		this.AddError(key, message)
	}
}

func (this *Validation) MaxSize(key string, obj interface{}, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorMaxSize(obj, max)
	if !valid {
		message := fmt.Sprintf("maximum length is %d", max)
		this.AddError(key, message)
	}
}

func (this *Validation) Min(key string, obj interface{}, min int) {
	if !this.should() {
		return
	}

	valid := ValidatorMin(obj, min)
	if !valid {
		message := fmt.Sprintf("minimum value is %d", min)
		this.AddError(key, message)
	}
}

func (this *Validation) Max(key string, obj interface{}, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorMax(obj, max)
	if !valid {
		message := fmt.Sprintf("maximum value is %d", max)
		this.AddError(key, message)
	}
}

func (this *Validation) Email(key, email string) {
	if !this.should() {
		return
	}

	valid := ValidatorEmail(email)
	if !valid {
		message := fmt.Sprintf("is not a valid email address")
		this.AddError(key, message)
	}
}

func (this *Validation) In(key string, obj interface{}, options []interface{}) {
	if !this.should() {
		return
	}

	valid := ValidatorIn(obj, options)
	if !valid {
		message := fmt.Sprintf("is not in given options")
		this.AddError(key, message)
	}
}

func (this *Validation) Range(key string, obj interface{}, min, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorRange(obj, min, max)
	if !valid {
		message := fmt.Sprintf("is not in range [%d, %d]", min, max)
		this.AddError(key, message)
	}
}

func (this *Validation) Match(key, val string, re *regexp.Regexp) {
	if !this.should() {
		return
	}

	valid := ValidatorMatch(re, val)
	if !valid {
		message := fmt.Sprintf("is not matched")
		this.AddError(key, message)
	}
}
