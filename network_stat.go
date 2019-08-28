package linuxtool

import (
	"io/ioutil"
	"strings"
)

type NetworkStat struct {
	Iface        string `json:"iface"`
	RxBytes      uint64 `json:"rxbytes"`
	RxPackets    uint64 `json:"rxpackets"`
	RxErrs       uint64 `json:"rxerrs"`
	RxDrop       uint64 `json:"rxdrop"`
	RxFifo       uint64 `json:"rxfifo"`
	RxFrame      uint64 `json:"rxframe"`
	RxCompressed uint64 `json:"rxcompressed"`
	RxMulticast  uint64 `json:"rxmulticast"`
	TxBytes      uint64 `json:"txbytes"`
	TxPackets    uint64 `json:"txpackets"`
	TxErrs       uint64 `json:"txerrs"`
	TxDrop       uint64 `json:"txdrop"`
	TxFifo       uint64 `json:"txfifo"`
	TxColls      uint64 `json:"txcolls"`
	TxCarrier    uint64 `json:"txcarrier"`
	TxCompressed uint64 `json:"txcompressed"`
}

func ReadNetworkStat(path string) ([]NetworkStat, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	// lines[2:] remove /proc/net/dev header
	results := make([]NetworkStat, len(lines[2:])-1)

	for i, line := range lines[2:] {
		// patterns
		// <iface>: 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
		// or
		// <iface>:0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 (without space after colon)
		colon := strings.Index(line, ":")

		if colon > 0 {
			metrics := line[colon+1:]
			fields := strings.Fields(metrics)

			results[i].Iface = strings.Replace(line[0:colon], " ", "", -1)
			results[i].RxBytes = ParseUint64(fields[0])
			results[i].RxPackets = ParseUint64(fields[1])
			results[i].RxErrs = ParseUint64(fields[2])
			results[i].RxDrop = ParseUint64(fields[3])
			results[i].RxFifo = ParseUint64(fields[4])
			results[i].RxFrame = ParseUint64(fields[5])
			results[i].RxCompressed = ParseUint64(fields[6])
			results[i].RxMulticast = ParseUint64(fields[7])
			results[i].TxBytes = ParseUint64(fields[8])
			results[i].TxPackets = ParseUint64(fields[9])
			results[i].TxErrs = ParseUint64(fields[10])
			results[i].TxDrop = ParseUint64(fields[11])
			results[i].TxFifo = ParseUint64(fields[12])
			results[i].TxColls = ParseUint64(fields[13])
			results[i].TxCarrier = ParseUint64(fields[14])
			results[i].TxCompressed = ParseUint64(fields[15])
		}
	}

	return results, nil
}
