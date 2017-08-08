// go-swagger examples.
//
// The purpose of this application is to provide some
// use cases describing how to generate docs for your API
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import "net/http"

// An OrderBodyParams model.
//
// This is used for operations that want an Order as body of the request
// swagger:parameters createOrder
type OrderBodyParams struct {
	// The order to submit.
	//
	// in: body
	// required: true
	Order *Order `json:"order"`
}

// An OrderResponse response model
//
// This is used for returning a response with a single order as body
//
// swagger:response orderResponse
type OrderResponse struct {
	//Get number of rate limted requests remaining
	//
	RateLimitRemaining string `json:"Rate-Limit-Remaining"`
	//
	// in: body
	Payload *Order `json:"order"`
}

// A ValidationError is a swagger response to represent error
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// CreateOrder swagger:route POST /orders orders createOrder
//
// Handler to create an order.
//
// Responses:
//        200: orderResponse
//        422: validationError
func CreateOrder(w http.ResponseWriter, req *http.Request) {
	// your code here
}

func main() {
}
