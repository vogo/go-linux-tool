package linuxtool

// # cat /proc/loadavg
//    0.75 0.35 0.25 1/25 1747
//
// The first three fields in this file are load average figures giving the number
// of jobs in the run queue (state R) or waiting for disk I/O (state D)
// averaged over 1, 5, and 15 minutes.
// They are the same as the load average numbers given by uptime(1) and other programs.
//
// The fourth field consists of two numbers separated by a slash (/).
// The first of these is the number of currently executing kernel scheduling entities (processes, threads);
// this will be less than or equal to the number of CPUs.
// The value after the slash is the number of kernel scheduling entities that currently exist on the system.
//
// The fifth field is the PID of the process that was most recently created on the system.
//
// ref: https://linux.die.net/man/5/proc

import (
	"errors"
	"io/ioutil"
	"strings"
)

type LoadAvg struct {
	Last1Min       float64 `json:"last1min"`
	Last5Min       float64 `json:"last5min"`
	Last15Min      float64 `json:"last15min"`
	ProcessRunning uint64  `json:"process_running"`
	ProcessTotal   uint64  `json:"process_total"`
	LastPID        uint64  `json:"last_pid"`
}

func ReadLoadAvg(path string) (*LoadAvg, error) {

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	content := strings.TrimSpace(string(b))
	fields := strings.Fields(content)

	if len(fields) < 5 {
		return nil, errors.New("Cannot parse loadavg: " + content)
	}

	process := strings.Split(fields[3], "/")

	if len(process) != 2 {
		return nil, errors.New("Cannot parse loadavg: " + content)
	}

	loadavg := LoadAvg{}

	if loadavg.Last1Min, err = ParseFloat(fields[0]); err != nil {
		return nil, err
	}

	if loadavg.Last5Min, err = ParseFloat(fields[1]); err != nil {
		return nil, err
	}

	if loadavg.Last15Min, err = ParseFloat(fields[2]); err != nil {
		return nil, err
	}

	if loadavg.ProcessRunning, err = ParseUint(process[0]); err != nil {
		return nil, err
	}

	if loadavg.ProcessTotal, err = ParseUint(process[1]); err != nil {
		return nil, err
	}

	if loadavg.LastPID, err = ParseUint(fields[4]); err != nil {
		return nil, err
	}

	return &loadavg, nil
}
