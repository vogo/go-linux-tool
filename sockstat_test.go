package linuxtool

import "testing"
import "reflect"

func TestSockStat(t *testing.T) {
	var expected = SockStat{
		SocketsUsed:   231,
		TCPInUse:      27,
		TCPOrphan:     1,
		TCPTimeWait:   23,
		TCPAllocated:  31,
		TCPMemory:     3,
		TCP6InUse:     0,
		UDPInUse:      19,
		UDPMemory:     17,
		UDP6InUse:     0,
		UDPLITEInUse:  0,
		UDPLITE6InUse: 0,
		RAWInUse:      0,
		RAW6InUse:     0,
		FRAGInUse:     0,
		FRAGMemory:    0,
		FRAG6InUse:    0,
		FRAG6Memory:   0,
	}

	sockStat, err := ReadSockStat("proc/sockstat")
	if err != nil {
		t.Fatal("sockstat read fail", err)
	}

	t.Logf("%+v", sockStat)

	if !reflect.DeepEqual(*sockStat, expected) {
		t.Error("not equal to expected")
	}
}
