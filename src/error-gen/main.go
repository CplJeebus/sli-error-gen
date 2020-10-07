package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ErrorSet struct {
	ErrorDays []ErrorDay `json:"error_days"`
	SloDef    Slo        `json:"slo"`
}

type ErrorDay struct {
	Date      time.Time `json:"date"`
	ErrorMins float64   `json:"ErrorMins"`
}

type Slo struct {
	SloPrecent float32 `json:"slo_precent"`
	Days       int     `json:"slo_days"`
	Name       string  `json:"name"`
}

func main() {
	Errors := make([]ErrorDay, 200)

	var error float64

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
			error = randomError
		case i < 28 && i > 0:
			error = randomError + Errors[i-1].ErrorMins
		case i > 28:
			error = randomError + Errors[i-1].ErrorMins - Errors[i].ErrorMins
		}

		Errors[i] = ErrorDay{
			Date:      when,
			ErrorMins: error,
		}
	}

	SillyErrorSet = ErrorSet{
		ErrorDays: Errors,
		SloDef:    SillySlo,
	}
	fmt.Print(SillyErrorSet)
}
