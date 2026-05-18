package main

import (
	"fmt"
	"github.com/mailru/easyjson"
	function "github.com/shoplinedev/function-extensions-template"
	"github.com/shoplinedev/function-extensions-template/templates/cart-flex/module"
)

func CartFlexFunction(req *module.CartFlexFunctionRequest) (result module.CartFlexFunctionResponse) {
	return module.CartFlexFunctionResponse{}
}

var _ = fmt.Printf
var _ = function.Log
var _ = easyjson.Marshal
