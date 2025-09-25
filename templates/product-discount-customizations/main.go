package main

import (
	"fmt"
	"github.com/mailru/easyjson"
	function "github.com/shoplinedev/function-extensions-template"
	"gopkg.inshopline.com/gsoul/function-proxy/template/product-discount-customizations/module"
)

func ProductDiscountCustomizationFunction(req *module.ProductDiscountCustomizationFunctionRequest) (result module.ProductDiscountCustomizationFunctionResponse) {
	return module.ProductDiscountCustomizationFunctionResponse{}
}

var _ = fmt.Printf
var _ = function.Log
var _ = easyjson.Marshal
