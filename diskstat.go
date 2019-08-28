package linuxtool

import (
	"io/ioutil"
	"strings"
	"time"
)

// DiskStat is disk statistics to help measure disk activity.
//
// Note:
// * On a very busy or long-lived system values may wrap.
// * No kernel locks are held while modifying these counters. This implies that
//   minor inaccuracies may occur.
//
// More more info see:
// https://www.kernel.org/doc/Documentation/iostats.txt and
// https://www.kernel.org/doc/Documentation/block/stat.txt
type DiskStat struct {
	Major        int    `json:"major"`         // major device number
	Minor        int    `json:"minor"`         // minor device number
	Name         string `json:"name"`          // device name
	ReadIOs      uint64 `json:"read_ios"`      // number of read I/Os processed
	ReadMerges   uint64 `json:"read_merges"`   // number of read I/Os merged with in-queue I/O
	ReadSectors  uint64 `json:"read_sectors"`  // number of 512 byte sectors read
	ReadTicks    uint64 `json:"read_ticks"`    // total wait time for read requests in milliseconds
	WriteIOs     uint64 `json:"write_ios"`     // number of write I/Os processed
	WriteMerges  uint64 `json:"write_merges"`  // number of write I/Os merged with in-queue I/O
	WriteSectors uint64 `json:"write_sectors"` // number of 512 byte sectors written
	WriteTicks   uint64 `json:"write_ticks"`   // total wait time for write requests in milliseconds
	InFlight     uint64 `json:"in_flight"`     // number of I/Os currently in flight
	IOTicks      uint64 `json:"io_ticks"`      // total time this block device has been active in milliseconds
	TimeInQueue  uint64 `json:"time_in_queue"` // total wait time for all requests in milliseconds
}

// ReadDiskStats reads and parses the file.
//
// Note:
// * Assumes a well formed file and will panic if it isn't.
func ReadDiskStats(path string) ([]DiskStat, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	devices := strings.Split(string(data), "\n")
	results := make([]DiskStat, len(devices)-1)

	for i := range results {
		fields := strings.Fields(devices[i])
		results[i].Major = ParseIntValue(fields[0])
		results[i].Minor = ParseIntValue(fields[1])
		results[i].Name = fields[2]
		results[i].ReadIOs = ParseUint64(fields[3])
		results[i].ReadMerges = ParseUint64(fields[4])
		results[i].ReadSectors = ParseUint64(fields[5])
		results[i].ReadTicks = ParseUint64(fields[6])
		results[i].WriteIOs = ParseUint64(fields[7])
		results[i].WriteMerges = ParseUint64(fields[8])
		results[i].WriteSectors = ParseUint64(fields[9])
		results[i].WriteTicks = ParseUint64(fields[10])
		results[i].InFlight = ParseUint64(fields[11])
		results[i].IOTicks = ParseUint64(fields[12])
		results[i].TimeInQueue = ParseUint64(fields[13])
	}

	return results, nil
}

// GetReadBytes returns the number of bytes read.
func (ds *DiskStat) GetReadBytes() int64 {
	return int64(ds.ReadSectors) * 512
}

// GetReadTicks returns the duration waited for read requests.
func (ds *DiskStat) GetReadTicks() time.Duration {
	return time.Duration(ds.ReadTicks) * time.Millisecond
}

// GetWriteBytes returns the number of bytes written.
func (ds *DiskStat) GetWriteBytes() int64 {
	return int64(ds.WriteSectors) * 512
}

// GetWriteTicks returns the duration waited for write requests.
func (ds *DiskStat) GetWriteTicks() time.Duration {
	return time.Duration(ds.WriteTicks) * time.Millisecond
}

// GetIOTicks returns the duration the disk has been active.
func (ds *DiskStat) GetIOTicks() time.Duration {
	return time.Duration(ds.IOTicks) * time.Millisecond
}

// GetTimeInQueue returns the duration waited for all requests.
func (ds *DiskStat) GetTimeInQueue() time.Duration {
	return time.Duration(ds.TimeInQueue) * time.Millisecond
}
