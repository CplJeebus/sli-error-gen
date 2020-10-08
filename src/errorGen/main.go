package main

import (
	"encoding/json"
	. "errorGen/types"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var conf ScenarioConf
	conf.GetConf()
	fmt.Println(conf)

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
