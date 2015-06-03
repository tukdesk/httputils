package validation

import (
	"fmt"
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
		errors: make([]*ValidationError, 1),
		multi:  false,
	}
}

func NewMulti() *Validation {
	return &Validation{
		errors: make([]*ValidationError, 0),
		multi:  true,
	}
}

func (this *Validation) HasErrors() bool {
	return len(this.errors) > 0
}

func (this *Validation) FirstError() error {
	if len(this.errors) == 0 {
		return nil
	}
	return this.errors[0]
}

func (this *Validation) addError(key, message string) {
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
		message := fmt.Sprintf("%s is required", key)
		this.addError(key, message)
	}
}

func (this *Validation) MinSize(key string, obj interface{}, min int) {
	if !this.should() {
		return
	}

	valid := ValidatorMinSize(obj, min)
	if !valid {
		message := fmt.Sprintf("minimum length of %s is %d", key, min)
		this.addError(key, message)
	}
}

func (this *Validation) MaxSize(key string, obj interface{}, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorMaxSize(obj, max)
	if !valid {
		message := fmt.Sprintf("maximum length of %s is %d", key, max)
		this.addError(key, message)
	}
}

func (this *Validation) Min(key string, obj interface{}, min int) {
	if !this.should() {
		return
	}

	valid := ValidatorMin(obj, min)
	if !valid {
		message := fmt.Sprintf("minimum value of %s is %d", key, min)
		this.addError(key, message)
	}
}

func (this *Validation) Max(key string, obj interface{}, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorMax(obj, max)
	if !valid {
		message := fmt.Sprintf("maximum value of %s is %d", key, max)
		this.addError(key, message)
	}
}

func (this *Validation) Email(key, email string) {
	if !this.should() {
		return
	}

	valid := ValidatorEmail(email)
	if !valid {
		message := fmt.Sprintf("value of %s is not a valid email address", key)
		this.addError(key, message)
	}
}

func (this *Validation) In(key string, obj interface{}, options []interface{}) {
	if !this.should() {
		return
	}

	valid := ValidatorIn(obj, options)
	if !valid {
		message := fmt.Sprintf("value of %s is not in given options", key)
		this.addError(key, message)
	}
}

func (this *Validation) Range(key string, obj interface{}, min, max int) {
	if !this.should() {
		return
	}

	valid := ValidatorRange(obj, min, max)
	if !valid {
		message := fmt.Sprintf("value of %s shoud be in range [%d, %d]", key, min, max)
		this.addError(key, message)
	}
}
