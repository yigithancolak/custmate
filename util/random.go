package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/yigithancolak/custmate/graph/model"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

var daysOfWeek = []string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

func RandomIntBetween(min, max int64) int {
	return int(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(8), RandomString(6))
}

func RandomName() string {
	return RandomString(6)
}

func RandomPassword() string {
	return RandomString(8)
}

func RandomPhoneNumber(n int) string {
	var sb strings.Builder

	for i := 0; i < n; i++ {
		numb := numbers[rand.Intn(len(numbers))]
		sb.WriteByte(numb)

	}

	return "+" + sb.String()
}

func RandomDate() string {
	d := RandomIntBetween(1, 30)
	m := RandomIntBetween(1, 12)
	y := RandomIntBetween(2020, 2024)

	if m == 2 {
		if d > 28 {
			d = RandomIntBetween(1, 28)
			//for february
		}
	}

	day := strconv.Itoa(d)
	month := strconv.Itoa(m)
	year := strconv.Itoa(y)

	if d < 10 {
		day = "0" + day
	}

	if m < 10 {
		month = "0" + month
	}

	dateArray := []string{year, month, day}

	return strings.Join(dateArray, "-")
}

func RandomDay() string {
	return daysOfWeek[RandomIntBetween(0, int64(len(daysOfWeek)-1))]
}

func RandomTime() string {
	h := RandomIntBetween(0, 23)
	m := RandomIntBetween(0, 59)

	hour := strconv.Itoa(h)
	if h < 10 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(m)
	if m < 10 {
		minute = "0" + minute
	}

	return hour + ":" + minute
}

func RandomPaymentType() model.PaymentType {
	rand := RandomIntBetween(0, int64(len(model.AllPaymentType)-1))

	return model.AllPaymentType[rand]
}

func RandomCurrency() model.Currency {
	rand := RandomIntBetween(0, int64(len(model.AllCurrency)-1))

	return model.AllCurrency[rand]

}
