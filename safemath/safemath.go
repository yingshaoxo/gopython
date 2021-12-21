package safemath

import "github.com/shopspring/decimal"

func String_to_decimal(numberString string) decimal.Decimal {
	price, err := decimal.NewFromString(numberString)
	if err != nil {
		price, _ = decimal.NewFromString("0.0")
	}
	return price
}

func Decimal_to_string(decimal decimal.Decimal) string {
	return decimal.String()
}
