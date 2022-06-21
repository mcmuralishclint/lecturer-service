package stripe_handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

// ChargeJSON incoming data for Stripe API
type ChargeJSON struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
}

func ChargeAPI(c *gin.Context) {
	var json ChargeJSON
	c.BindJSON(&json)
	apiKey := os.Getenv("SK_TEST_KEY")
	stripe.Key = apiKey

	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(json.Amount),
		Currency:     stripe.String(string(stripe.CurrencyAUD)),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")}, // this should come from clientside
		ReceiptEmail: stripe.String(json.ReceiptEmail)})

	if err != nil {
		c.String(http.StatusBadRequest, "Request Failed", err)
		return
	}
	c.String(http.StatusCreated, "Succesfully Charged")
}
