package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type ErrorSet struct {
	ErrorDays []ErrorDay `json:"error_days"`
	SloDef    Slo        `json:"slo"`
}

type ErrorDay struct {
	Date       time.Time `json:"date"`
	ErrorMins  float64   `json:"error_mins"`
	ErrorBurnt float64   `json:"error_burnt"`
}

type Slo struct {
	SloPrecent float32 `json:"slo_precent"`
	Days       int     `json:"slo_days"`
	Name       string  `json:"name"`
}

func main() {
	Errors := make([]ErrorDay, 180)

	var errorburn float64

	var SillyErrorSet ErrorSet

	SillySlo := Slo{
		Name:       "Dummy Data",
		Days:       28,
		SloPrecent: 99.9,
	}

	for i := 0; i < 180; i++ {
		when := time.Now().AddDate(0, 0, i-180)
		randomError := rand.Float64() * 10

		switch {
		case i < 28 && i == 0:
			errorburn = randomError
		case i < 28 && i > 0:
			errorburn = randomError + Errors[i-1].ErrorBurnt
		case i > 28:
			errorburn = randomError + Errors[i-1].ErrorBurnt - Errors[i-28].ErrorMins
		}

		Errors[i] = ErrorDay{
			Date:       when,
			ErrorMins:  randomError,
			ErrorBurnt: errorburn,
		}
	}

	SillyErrorSet = ErrorSet{
		ErrorDays: Errors,
		SloDef:    SillySlo,
	}
	out, _ := json.Marshal(SillyErrorSet)
	//fmt.Print(len(out))
	fmt.Println(string(out))
}
