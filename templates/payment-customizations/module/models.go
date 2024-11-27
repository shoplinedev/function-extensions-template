package module

type PaymentCustomizationFunctionRequest struct {
	Checkout                *Checkout               `json:"checkout"`
	Localization            *Localization           `json:"localization"`
	PaymentCustomizations   []*PaymentCustomization `json:"paymentCustomizations"`
	PresentmentCurrencyRate *ExchangeRate           `json:"presentmentCurrencyRate"`
}

type Operation struct {
	Hide   *HideOperation   `json:"hide"`
	Move   *MoveOperation   `json:"move"`
	Rename *RenameOperation `json:"rename"`
}

//easyjson:json
type PaymentCustomizationFunctionResponses []PaymentCustomizationFunctionResponse

type PaymentCustomizationFunctionResponse struct {
	Operations []Operation `json:"operations"`
}

type HideOperation struct {
	PaymentMethodId *string `json:"paymentMethodId"`
}

type MoveOperation struct {
	Index           *int64  `json:"index"`
	PaymentMethodId *string `json:"paymentMethodId"`
}

type RenameOperation struct {
	Name            *string `json:"name"`
	PaymentMethodId *string `json:"paymentMethodId"`
}

type Enum__TypeKind string

const Enum__TypeKindSCALAR Enum__TypeKind = "SCALAR"
const Enum__TypeKindOBJECT Enum__TypeKind = "OBJECT"
const Enum__TypeKindINTERFACE Enum__TypeKind = "INTERFACE"
const Enum__TypeKindUNION Enum__TypeKind = "UNION"
const Enum__TypeKindENUM Enum__TypeKind = "ENUM"
const Enum__TypeKindINPUT_OBJECT Enum__TypeKind = "INPUT_OBJECT"
const Enum__TypeKindLIST Enum__TypeKind = "LIST"
const Enum__TypeKindNON_NULL Enum__TypeKind = "NON_NULL"

type Enum__DirectiveLocation string

const Enum__DirectiveLocationQUERY Enum__DirectiveLocation = "QUERY"
const Enum__DirectiveLocationMUTATION Enum__DirectiveLocation = "MUTATION"
const Enum__DirectiveLocationSUBSCRIPTION Enum__DirectiveLocation = "SUBSCRIPTION"
const Enum__DirectiveLocationFIELD Enum__DirectiveLocation = "FIELD"
const Enum__DirectiveLocationFRAGMENT_DEFINITION Enum__DirectiveLocation = "FRAGMENT_DEFINITION"
const Enum__DirectiveLocationFRAGMENT_SPREAD Enum__DirectiveLocation = "FRAGMENT_SPREAD"
const Enum__DirectiveLocationINLINE_FRAGMENT Enum__DirectiveLocation = "INLINE_FRAGMENT"
const Enum__DirectiveLocationSCHEMA Enum__DirectiveLocation = "SCHEMA"
const Enum__DirectiveLocationSCALAR Enum__DirectiveLocation = "SCALAR"
const Enum__DirectiveLocationOBJECT Enum__DirectiveLocation = "OBJECT"
const Enum__DirectiveLocationFIELD_DEFINITION Enum__DirectiveLocation = "FIELD_DEFINITION"
const Enum__DirectiveLocationARGUMENT_DEFINITION Enum__DirectiveLocation = "ARGUMENT_DEFINITION"
const Enum__DirectiveLocationINTERFACE Enum__DirectiveLocation = "INTERFACE"
const Enum__DirectiveLocationUNION Enum__DirectiveLocation = "UNION"
const Enum__DirectiveLocationENUM Enum__DirectiveLocation = "ENUM"
const Enum__DirectiveLocationENUM_VALUE Enum__DirectiveLocation = "ENUM_VALUE"
const Enum__DirectiveLocationINPUT_OBJECT Enum__DirectiveLocation = "INPUT_OBJECT"
const Enum__DirectiveLocationINPUT_FIELD_DEFINITION Enum__DirectiveLocation = "INPUT_FIELD_DEFINITION"

type __Schema struct {
	Types            []__Type      `json:"types"`
	QueryType        __Type        `json:"queryType"`
	MutationType     *__Type       `json:"mutationType"`
	SubscriptionType *__Type       `json:"subscriptionType"`
	Directives       []__Directive `json:"directives"`
}

type __Type struct {
	Kind          Enum__TypeKind `json:"kind"`
	Name          *string        `json:"name"`
	Description   *string        `json:"description"`
	Fields        []__Field      `json:"fields"`
	Interfaces    []__Type       `json:"interfaces"`
	PossibleTypes []__Type       `json:"possibleTypes"`
	EnumValues    []__EnumValue  `json:"enumValues"`
	InputFields   []__InputValue `json:"inputFields"`
	OfType        *__Type        `json:"ofType"`
}

type __Field struct {
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	Args              []__InputValue `json:"args"`
	Type              __Type         `json:"type"`
	IsDeprecated      bool           `json:"isDeprecated"`
	DeprecationReason *string        `json:"deprecationReason"`
}

type __InputValue struct {
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	Type         __Type  `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

type __EnumValue struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason *string `json:"deprecationReason"`
}

type __Directive struct {
	Name        string                    `json:"name"`
	Description *string                   `json:"description"`
	Locations   []Enum__DirectiveLocation `json:"locations"`
	Args        []__InputValue            `json:"args"`
}

type BuyerIdentity struct {
	Customer *Customer `json:"customer"`
	Email    *string   `json:"email"`
	Phone    *string   `json:"phone"`
}

type Checkout struct {
	BuyerIdentity  *BuyerIdentity           `json:"buyerIdentity"`
	Cost           CheckoutCost             `json:"cost"`
	DeliveryGroups []*CheckoutDeliveryGroup `json:"deliveryGroups"`
	Lines          []*CheckoutLine          `json:"lines"`
}

type CheckoutCost struct {
	SubtotalAmount Money  `json:"subtotalAmount"`
	TotalAmount    Money  `json:"totalAmount"`
	TotalTaxAmount *Money `json:"totalTaxAmount"`
}

type CheckoutDeliveryGroup struct {
	CheckoutLines          []*CheckoutLine           `json:"checkoutLines"`
	DeliveryAddress        *MailingAddress           `json:"deliveryAddress"`
	DeliveryOptions        []*CheckoutDeliveryOption `json:"deliveryOptions"`
	Id                     []*string                 `json:"id"`
	SelectedDeliveryOption *CheckoutDeliveryOption   `json:"selectedDeliveryOption"`
}

type CheckoutDeliveryOption struct {
	Cost               Money   `json:"cost"`
	DeliveryMethodType *string `json:"deliveryMethodType"`
	Description        string  `json:"description"`
	Handle             string  `json:"handle"`
	Title              *string `json:"title"`
}

type CheckoutLine struct {
	Cost        CheckoutLineCost `json:"cost"`
	Id          string           `json:"id"`
	Merchandise ProductVariant   `json:"merchandise"`
	Quantity    int64            `json:"quantity"`
}

type CheckoutLineCost struct {
	SubtotalAmount Money `json:"subtotalAmount"`
	TotalAmount    Money `json:"totalAmount"`
}

type Customer struct {
	DisplayName string  `json:"displayName"`
	Email       *string `json:"email"`
	Id          string  `json:"id"`
}

type ExchangeRate struct {
	Rate string `json:"rate"`
}

type Input struct {
	Checkout                Checkout                             `json:"checkout"`
	Localization            Localization                         `json:"localization"`
	PaymentCustomizations   []*PaymentCustomization              `json:"paymentCustomizations"`
	PaymentMethods          []*PaymentCustomizationPaymentMethod `json:"paymentMethods"`
	PresentmentCurrencyRate ExchangeRate                         `json:"presentmentCurrencyRate"`
	__schema                __Schema                             `json:"__schema"`
	__type                  *__Type                              `json:"__type"`
}

type Localization struct {
	Country  string `json:"country"`
	Language string `json:"language"`
}

type MailingAddress struct {
	Address1     *string `json:"address1"`
	Address2     *string `json:"address2"`
	City         *string `json:"city"`
	Company      *string `json:"company"`
	CountryCode  *string `json:"countryCode"`
	FirstName    *string `json:"firstName"`
	LastName     *string `json:"lastName"`
	Name         *string `json:"name"`
	Phone        *string `json:"phone"`
	ProvinceCode *string `json:"provinceCode"`
	Zip          *string `json:"zip"`
}

type Metafield struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Money struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

type PaymentCustomization struct {
	Metafield *Metafield `json:"metafield"`
}

type PaymentCustomizationPaymentMethod struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Handle     string `json:"handle"`
	Id         string `json:"id"`
	IsGiftCard bool   `json:"isGiftCard"`
}

type ProductVariant struct {
	Id               string   `json:"id"`
	Product          Product  `json:"product"`
	RequiresShipping bool     `json:"requiresShipping"`
	Sku              *string  `json:"sku"`
	Weight           *float64 `json:"weight"`
	WeightUnit       *string  `json:"weightUnit"`
}
