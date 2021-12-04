package proto

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

// .
func ValidateNavigation(
	field string,
	getValueFunc func(interface{}) interface{},
	rules ...validation.Rule,
) *validateNavigationRule {
	return &validateNavigationRule{
		field:        field,
		getValueFunc: getValueFunc,
		rules:        rules,
	}
}

type validateNavigationRule struct {
	field        string
	getValueFunc func(interface{}) interface{}
	rules        []validation.Rule
}

// .
func (v *validateNavigationRule) Validate(value interface{}) error {
	navValue := v.getValueFunc(value)
	for _, rule := range v.rules {
		if err := rule.Validate(navValue); err != nil {
			return fmt.Errorf("%s: %w", v.field, err)
		}
	}
	return nil
}
