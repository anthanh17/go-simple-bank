package api

import (
	"github.com/anthanh17/simplebank/util"
	"github.com/go-playground/validator/v10"
)

// Define my validator
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}