// We do safe math here, for example, handle big numbers for financial problems
package safe_math

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

func Add_number(number1 string, number2 string) string {
	number1Decimal := String_to_decimal(number1)
	number2Decimal := String_to_decimal(number2)
	return Decimal_to_string(number1Decimal.Add(number2Decimal))
}

func Subtract_number(number1 string, number2 string) string {
	number1Decimal := String_to_decimal(number1)
	number2Decimal := String_to_decimal(number2)
	return Decimal_to_string(number1Decimal.Sub(number2Decimal))
}

func Multiply_number(number1 string, number2 string) string {
	number1Decimal := String_to_decimal(number1)
	number2Decimal := String_to_decimal(number2)
	return Decimal_to_string(number1Decimal.Mul(number2Decimal))
}

func Divide_number(number1 string, number2 string) string {
	number1Decimal := String_to_decimal(number1)
	number2Decimal := String_to_decimal(number2)
	return Decimal_to_string(number1Decimal.Div(number2Decimal))
}
