package types

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
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
	SloPrecent float64 `json:"slo_precent"`
	Days       int     `json:"slo_days"`
	Name       string  `json:"name"`
}

type ScenarioConf struct {
	Duration int       `yaml:"duration"`
	SLO      []SloConf `yaml:"slo"`
}

type SloConf struct {
	Name           string     `yaml:"name"`
	Slo            float64    `yaml:"slo_precent"`
	PeriodDays     int        `yaml:"days"`
	NormalErrorMax float64    `yaml:"normal_error_max"`
	Events         []SloEvent `yaml:"events"`
}

type SloEvent struct {
	Type     string  `yaml:"type"`
	BurnRate float64 `yaml:"burn"`
	Occurs   int     `yaml:"occurs"`
	Duration int     `yaml:"duration"`
}

func (s *ScenarioConf) GetConf() *ScenarioConf {
	Conf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("opening conf file err   #%v ", err)
	}

	err = yaml.Unmarshal(Conf, s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s
}
