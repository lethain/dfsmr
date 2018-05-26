package main

import (
	"testing"

	pb "github.com/lethain/dfsmr/dfsmr"
	"github.com/lethain/dfsmr/machines"	
)

func TestRegisterMachine(t *testing.T) {
	s := MakeServer()
	if len(s.machines) != 0 {
		t.Error("Should be 0 machines, have ", s.machines)
	}
	m := &pb.DefineRequest{}
	s.RegisterMachine(m)
	if len(s.machines) != 1 {
		t.Error("Should be 1 machine, have ", s.machines)
	}
	ms := s.Machines()
	if len(ms) != 1 {
		t.Error("Should be 1 machine, have ", ms)
	}
}

func TestDefineMachine(t *testing.T) {
	path := "../crawl.fsm.yaml"
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	name := "crawler"
	if m.Name != name {
		t.Error("name should be ", name, "was ", m.Name)
	}
	
	

	

}
