package main

import "fmt"

type Byte int
type Celcius float32

type BandwidthUsage struct {
	amount []Byte
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Byte
}

func (b *BandwidthUsage) AvgBandUsage() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum
}

func (c *CpuTemp) AvgCpuTemp() int {
	sum := 0
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum
}

func (m *MemoryUsage) AvgMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Byte{5000, 10000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Byte{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		bandwidth,
		temp,
		memory,
	}

	fmt.Println("bandwidth :", dashboard.BandwidthUsage.AvgBandUsage())
	fmt.Println("cpu temp :", dashboard.CpuTemp.AvgCpuTemp())
	fmt.Println("memory usage :", dashboard.MemoryUsage.AvgMemoryUsage())

}
