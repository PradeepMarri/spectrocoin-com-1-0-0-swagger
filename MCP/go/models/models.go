package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Orderinformationclass represents the Orderinformationclass schema from the OpenAPI specification
type Orderinformationclass struct {
	Depositaddress string `json:"depositAddress"` // Payment deposit address
	Orderid string `json:"orderId"` // Merchant provided or generated order id
	Orderrequestid int64 `json:"orderRequestId"` // Generated unique order request id
	Receivecurrency string `json:"receiveCurrency"` // Currency to be funded to merchant account
	Redirecturl string `json:"redirectUrl"` // SpectroCoin payment window url where merchant customer may be redirected
	Validuntil int64 `json:"validUntil"` // Timestamp until when order is valid
	Paycurrency string `json:"payCurrency"` // Currency to be paid with
	Payamount float64 `json:"payAmount"` // Amount of payment currency to be paid by merchant customer
	Receiveamount float64 `json:"receiveAmount"` // Amount of receive currency to be funded to merchant account
}

// Orderrequestclassusedtoprovideorderspecificinformation represents the Orderrequestclassusedtoprovideorderspecificinformation schema from the OpenAPI specification
type Orderrequestclassusedtoprovideorderspecificinformation struct {
	Orderid string `json:"orderId,omitempty"` // Custom order ID. Must be unique per API. If not provided it will be generated.
	Payername string `json:"payerName,omitempty"` // Specified payer name.
	Payersurname string `json:"payerSurname,omitempty"` // Specified payer surname.
	Receiveamount float64 `json:"receiveAmount,omitempty"` // Receive amount in receive currency of value that merchant will be funded after merchant customers payment approval. If not provided pay amount will be used to calculate receive amount
	Culture string `json:"culture,omitempty"` // Merchant customer culture payment window to be presented
	Description string `json:"description,omitempty"` // Order description. Will be presented for merchant customer at payment window
	Callbackurl string `json:"callbackUrl,omitempty"` // Url of merchant endpoint callback about order status to be returned
	Failureurl string `json:"failureUrl,omitempty"` // Url of merchant page customer should be redirected after unsuccessful payment
	Sign string `json:"sign"` // Signature required for signing create order request
	Merchantid int64 `json:"merchantId"` // Merchant ID assigned to your account
	Paycurrency string `json:"payCurrency"` // Currency of pay amount
	Payeremail string `json:"payerEmail,omitempty"` // Specified payer email.
	Receivecurrency string `json:"receiveCurrency"` // Currency of receive amount
	Successurl string `json:"successUrl,omitempty"` // Url of merchant page customer should be redirected after successful payment
	Payamount float64 `json:"payAmount,omitempty"` // Pay amount in pay currency of value which should be paid by merchant customer. If not provided receive amount will be used to calculate pay amount
	Apiid int64 `json:"apiId"` // API ID of specific API you have configured on your merchant account
}
