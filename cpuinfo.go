package linuxtool

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var cpuInfoRegExp = regexp.MustCompile(`([^:]*?)\s*:\s*(.*)$`)

type Processor struct {
	Id         int64    `json:"id"`
	VendorId   string   `json:"vendor_id"`
	Model      int64    `json:"model"`
	ModelName  string   `json:"model_name"`
	Flags      []string `json:"flags"`
	Cores      int64    `json:"cores"`
	MHz        float64  `json:"mhz"`
	CacheSize  int64    `json:"cache_size"` // KB
	PhysicalId int64    `json:"physical_id"`
	CoreId     int64    `json:"core_id"`
}

type CPUInfo struct {
	Processors []Processor `json:"processors"`
}

func (info *CPUInfo) NumCPU() int {
	return len(info.Processors)
}

func (info *CPUInfo) NumCore() int {
	core := make(map[string]bool)

	for _, p := range info.Processors {
		pid := p.PhysicalId
		cid := p.CoreId

		if pid == -1 {
			return info.NumCPU()
		} else {
			// to avoid fmt import
			key := strconv.FormatInt(int64(pid), 10) + ":" + strconv.FormatInt(int64(cid), 10)
			core[key] = true
		}
	}

	return len(core)
}

func (info *CPUInfo) NumPhysicalCPU() int {
	pcpu := make(map[string]bool)

	for _, p := range info.Processors {
		pid := p.PhysicalId

		if pid == -1 {
			return info.NumCPU()
		} else {
			// to avoid fmt import
			key := strconv.FormatInt(int64(pid), 10)
			pcpu[key] = true
		}
	}

	return len(pcpu)
}

func ReadCPUInfo(path string) (*CPUInfo, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := string(b)
	lines := strings.Split(content, "\n")

	var cpuInfo = CPUInfo{}
	var processor = &Processor{CoreId: -1, PhysicalId: -1}

	for i, line := range lines {
		var key string
		var value string

		if len(line) == 0 && i != len(lines)-1 {
			// end of processor
			cpuInfo.Processors = append(cpuInfo.Processors, *processor)
			processor = &Processor{}
			continue
		} else if i == len(lines)-1 {
			continue
		}

		subMatches := cpuInfoRegExp.FindStringSubmatch(line)
		key = subMatches[1]
		value = subMatches[2]

		switch key {
		case "processor":
			processor.Id = ParseInt64(value)
		case "vendor_id":
			processor.VendorId = value
		case "model":
			processor.Model = ParseInt64(value)
		case "model name":
			processor.ModelName = value
		case "flags":
			processor.Flags = strings.Fields(value)
		case "cpu cores":
			processor.Cores = ParseInt64(value)
		case "cpu MHz":
			processor.MHz = ParseFloat64(value)
		case "cache size":
			processor.CacheSize, _ = strconv.ParseInt(value[:strings.IndexAny(value, " \t\n")], 10, 64)
			if strings.HasSuffix(line, "MB") {
				processor.CacheSize *= 1024
			}
		case "physical id":
			processor.PhysicalId = ParseInt64(value)
		case "core id":
			processor.CoreId = ParseInt64(value)
		}
		/*
			processor	: 0
			vendor_id	: GenuineIntel
			cpu family	: 6
			model		: 26
			model name	: Intel(R) Xeon(R) CPU           L5520  @ 2.27GHz
		*/
	}
	return &cpuInfo, nil
}
