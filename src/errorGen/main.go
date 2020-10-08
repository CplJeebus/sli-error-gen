package main

import (
	"encoding/json"
	. "errorGen/output"
	. "errorGen/types"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var conf ScenarioConf

	conf.GetConf()

	Errors := make([]ErrorDay, conf.Duration)

	var errorburn float64

	var SillyErrorSet ErrorSet

	SillySlo := Slo{
		Name:       conf.SLO[0].Name,
		Days:       conf.SLO[0].PeriodDays,
		SloPrecent: conf.SLO[0].Slo,
	}

	for i := 0; i < conf.Duration; i++ {
		when := time.Now().AddDate(0, 0, i-conf.Duration)
		randomError := rand.Float64() * conf.SLO[0].NormalErrorMax

		switch {
		case i < 28 && i == 0:
			errorburn = randomError
		case i < 28 && i > 0:
			errorburn = randomError + Errors[i-1].ErrorBurnt
		case i > 28:
			errorburn = randomError + Errors[i-1].ErrorBurnt - Errors[i-conf.SLO[0].PeriodDays].ErrorMins + Burn(conf, i)
			randomError = randomError + Burn(conf, i)
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
	CreatePlot(SillyErrorSet)
}

func Burn(c ScenarioConf, w int) float64 {
	for i := range c.SLO[0].Events {
		switch {
		case c.SLO[0].Events[i].Type == "fast":
			if w == c.SLO[0].Events[i].Occurs {
				return c.SLO[0].Events[i].BurnRate
			}
		case c.SLO[0].Events[i].Type == "slow":
			if w > c.SLO[0].Events[i].Occurs && w < c.SLO[0].Events[i].Occurs+c.SLO[0].Events[i].Duration {
				return c.SLO[0].Events[i].BurnRate
			}
		}
	}

	return 0
}
