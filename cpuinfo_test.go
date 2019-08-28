package linuxtool

import "testing"

func TestCPUInfo(t *testing.T) {

	cpuInfo, err := ReadCPUInfo("proc/cpuinfo")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", cpuInfo)

	if len(cpuInfo.Processors) != 8 {
		t.Fatal("wrong processor number : ", len(cpuInfo.Processors))
	}

	if cpuInfo.NumCore() != 8 {
		t.Fatal("wrong core number", cpuInfo.NumCore())
	}

	if cpuInfo.NumPhysicalCPU() != 2 {
		t.Fatal("wrong physical cpu number", cpuInfo.NumPhysicalCPU())
	}

	cpuInfo, err = ReadCPUInfo("proc/cpuinfo_2")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", cpuInfo)

	if len(cpuInfo.Processors) != 4 {
		t.Fatal("wrong processor number : ", len(cpuInfo.Processors))
	}

	if cpuInfo.NumCore() != 4 {
		t.Fatal("wrong core number", cpuInfo.NumCore())
	}

	// not sure at all here
	// does not match with https://github.com/randombit/cpuinfo/blob/master/x86/xeon_l5520
	if cpuInfo.NumPhysicalCPU() != 4 {
		t.Fatal("wrong physical cpu number", cpuInfo.NumPhysicalCPU())
	}

	cpuInfo, err = ReadCPUInfo("proc/cpuinfo_3")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", cpuInfo)

	if len(cpuInfo.Processors) != 4 {
		t.Fatal("wrong processor number : ", len(cpuInfo.Processors))
	}

	if cpuInfo.NumCore() != 2 {
		t.Fatal("wrong core number", cpuInfo.NumCore())
	}

	if cpuInfo.NumPhysicalCPU() != 1 {
		t.Fatal("wrong physical cpu number", cpuInfo.NumPhysicalCPU())
	}
}
