package linuxtool

import (
	"io/ioutil"
	"strings"
	"time"
)

type Uptime struct {
	Total float64 `json:"total"`
	Idle  float64 `json:"idle"`
}

func (uptime *Uptime) GetTotalDuration() time.Duration {
	return time.Duration(uptime.Total) * time.Second
}

func (uptime *Uptime) GetIdleDuration() time.Duration {
	return time.Duration(uptime.Idle) * time.Second
}

func (uptime *Uptime) CalculateIdle() float64 {
	// XXX
	// num2/(num1*N)     # N = SMP CPU numbers
	return 0
}

func ReadUptime(path string) (*Uptime, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(string(b))
	uptime := Uptime{}
	if uptime.Total, err = ParseFloat(fields[0]); err != nil {
		return nil, err
	}
	if uptime.Idle, err = ParseFloat(fields[1]); err != nil {
		return nil, err
	}
	return &uptime, nil
}
