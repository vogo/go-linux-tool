package linuxtool

// /proc/meminfo
//
// This file reports statistics about memory usage on the system.
//
// It is used by free(1) to report the amount of free and used memory (both physical and swap)
// on the system as well as the shared memory and buffers used by the kernel.
// Each line of the file consists of a parameter name, followed by a colon,
// the value of the parameter, and an option unit of measurement (e.g., "kB").
// The list below describes the parameter names and the format specifier required to read the field value.
// Except as noted below, all of the fields have been present since at least Linux 2.6.0.
// Some fields are only displayed if the kernel was configured with various options;
// those dependencies are noted in the list.
//
// ref: https://linux.die.net/man/5/proc

import (
	"io/ioutil"
	"reflect"
	"strings"
)

type MemInfo struct {
	MemTotal          uint64 `json:"mem_total"`
	MemFree           uint64 `json:"mem_free"`
	MemAvailable      uint64 `json:"mem_available"`
	Buffers           uint64 `json:"buffers"`
	Cached            uint64 `json:"cached"`
	SwapCached        uint64 `json:"swap_cached"`
	Active            uint64 `json:"active"`
	Inactive          uint64 `json:"inactive"`
	ActiveAnon        uint64 `json:"active_anon" field:"Active(anon)"`
	InactiveAnon      uint64 `json:"inactive_anon" field:"Inactive(anon)"`
	ActiveFile        uint64 `json:"active_file" field:"Active(file)"`
	InactiveFile      uint64 `json:"inactive_file" field:"Inactive(file)"`
	Unevictable       uint64 `json:"unevictable"`
	Mlocked           uint64 `json:"mlocked"`
	SwapTotal         uint64 `json:"swap_total"`
	SwapFree          uint64 `json:"swap_free"`
	Dirty             uint64 `json:"dirty"`
	Writeback         uint64 `json:"write_back"`
	AnonPages         uint64 `json:"anon_pages"`
	Mapped            uint64 `json:"mapped"`
	Shmem             uint64 `json:"shmem"`
	Slab              uint64 `json:"slab"`
	SReclaimable      uint64 `json:"s_reclaimable"`
	SUnreclaim        uint64 `json:"s_unclaim"`
	KernelStack       uint64 `json:"kernel_stack"`
	PageTables        uint64 `json:"page_tables"`
	NFS_Unstable      uint64 `json:"nfs_unstable"`
	Bounce            uint64 `json:"bounce"`
	WritebackTmp      uint64 `json:"writeback_tmp"`
	CommitLimit       uint64 `json:"commit_limit"`
	Committed_AS      uint64 `json:"committed_as"`
	VmallocTotal      uint64 `json:"vmalloc_total"`
	VmallocUsed       uint64 `json:"vmalloc_used"`
	VmallocChunk      uint64 `json:"vmalloc_chunk"`
	HardwareCorrupted uint64 `json:"hardware_corrupted"`
	AnonHugePages     uint64 `json:"anon_huge_pages"`
	HugePages_Total   uint64 `json:"huge_pages_total"`
	HugePages_Free    uint64 `json:"huge_pages_free"`
	HugePages_Rsvd    uint64 `json:"huge_pages_rsvd"`
	HugePages_Surp    uint64 `json:"huge_pages_surp"`
	Hugepagesize      uint64 `json:"hugepagesize"`
	DirectMap4k       uint64 `json:"direct_map_4k"`
	DirectMap2M       uint64 `json:"direct_map_2M"`
	DirectMap1G       uint64 `json:"direct_map_1G"`
}

func ReadMemInfo(path string) (*MemInfo, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	// Maps a meminfo metric to its value (i.e. MemTotal --> 100000)
	statMap := make(map[string]uint64)

	var info = MemInfo{}

	for _, line := range lines {
		fields := strings.SplitN(line, ":", 2)
		if len(fields) < 2 {
			continue
		}
		valFields := strings.Fields(fields[1])
		statMap[fields[0]] = ParseUint64(valFields[0])
	}

	elem := reflect.ValueOf(&info).Elem()
	typeOfElem := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		val, ok := statMap[typeOfElem.Field(i).Name]
		if ok {
			elem.Field(i).SetUint(val)
			continue
		}
		val, ok = statMap[typeOfElem.Field(i).Tag.Get("field")]
		if ok {
			elem.Field(i).SetUint(val)
		}
	}

	return &info, nil
}
