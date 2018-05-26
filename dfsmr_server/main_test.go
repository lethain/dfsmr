package main

import (
	"testing"

	pb "github.com/lethain/dfsmr/dfsmr"
	"golang.org/x/net/context"	
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

func TestStart(t *testing.T) {
	path := "../crawl.fsm.yaml"

	// create server
	s := MakeServer()
	ctx := context.Background()	

	// load a machine
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	dr := machines.AsDefineRequest(m)

	// there should be no machines
	//ms := s.Machines(ctx)

	// start for non-existant machine should fail
	sr := &pb.StartRequest{Name: m.Name}	
	_, err = s.Start(ctx, sr)
	if err == nil {
		t.Error("Invalid start request, should have failed ", sr)
	}

	// define a machine
	_, err = s.Define(ctx, dr)
	if err != nil {
		t.Error("failed to define ", err)
	}

	// start should work for a registered machine
	_, err = s.Start(ctx, sr)
	if err != nil {
		t.Error("start request should have succeeded ", err)
	}
	

}
