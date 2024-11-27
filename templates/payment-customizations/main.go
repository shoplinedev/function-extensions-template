package main

import (
	"fmt"
	"github.com/mailru/easyjson"
	function "github.com/shoplinedev/function-extensions-template"
	"github.com/shoplinedev/function-extensions-template/templates/payment-customizations/module"
)

func PaymentCustomizationFunction(req *module.PaymentCustomizationFunctionRequest) (result module.PaymentCustomizationFunctionResponse) {
	return module.PaymentCustomizationFunctionResponse{}
}

var _ = fmt.Printf
var _ = function.Log
var _ = easyjson.Marshal
