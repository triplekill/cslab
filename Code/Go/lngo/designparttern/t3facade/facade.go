package t3facade

import (
	"fmt"
)

type CPU interface {
	Process()
}

type CpuImp struct{}

func (c *CpuImp) Process() {
	fmt.Println("CpuImp")
}

type Mem interface {
	MixinMem()
}

type MemImp struct{}

func (m *MemImp) MixinMem() {
	fmt.Println("MemImp")
}

type Other interface {
	MixinO()
}
type OtherImp struct{}

func (o *OtherImp) MixinO() {
	fmt.Println("OtherImp")
}

type ComputerFacade struct{}

func (cf *ComputerFacade) SimpleComputer() {
	var cpu CPU = &CpuImp{}
	var mem Mem = &MemImp{}
	var others Other = &OtherImp{}
	cpu.Process()
	mem.MixinMem()
	others.MixinO()
}

func Client() {
	computer := ComputerFacade{}
	computer.SimpleComputer()
}
