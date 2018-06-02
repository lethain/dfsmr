package main

import (
	"testing"

	pb "github.com/lethain/dfsmr/dfsmr"
	"golang.org/x/net/context"
	"github.com/lethain/dfsmr/machines"
)

func TestRegisterMachine(t *testing.T) {
	s := MakeServer()
	ctx := context.Background()
	if len(s.machines) != 0 {
		t.Error("Should be 0 machines, have ", s.machines)
	}
	m := &pb.DefineRequest{}
	s.RegisterMachine(m)
	if len(s.machines) != 1 {
		t.Error("Should be 1 machine, have ", s.machines)
	}
	ms, err := s.Machines(ctx, &pb.MachinesRequest{})
	if err != nil {
		t.Error("Should not error retrieving machines ", err)
	}
	if len(ms.Machines) != 1 {
		t.Error("Should be 1 machine, have ", ms)
	}
}

func TestDefineMachine(t *testing.T) {
	path := "../crawl.fsm.yaml"
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	id := "crawler"
	if m.Id != id {
		t.Error("id should be ", id, "was ", m.Id)
	}
}

func TestMachines(t *testing.T) {
	// build a machine
	path := "../crawl.fsm.yaml"
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	dr := machines.AsDefineRequest(m)

	// setup server
	s := MakeServer()
	ctx := context.Background()

	// no machines registered
	ms, err := s.Machines(ctx, &pb.MachinesRequest{})
	if len(ms.Machines) != 0 {
		t.Error("shouldn't be any registered machines: ", ms.Machines)
	}

	// define machine
	s.Define(ctx, dr)

	// should be registered machine
	ms, err = s.Machines(ctx, &pb.MachinesRequest{})
	if len(ms.Machines) != 1 {
		t.Error("shouldn't be any registered machines: ", ms.Machines)
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
	sr := &pb.TaskMessage{Machine: m.Id}
	_, err = s.Start(ctx, sr)
	if err == nil {
		t.Error("Invalid start request, should have failed ", sr)
	}

	// define a machine
	_, err = s.Define(ctx, dr)
	if err != nil {
		t.Error("failed to define ", err)
	}

	// shouldn't be an instances
	instances, err := s.Instances(ctx, &pb.InstancesRequest{})
	if err != nil {
		t.Error("should be able to retrieve instance ", err)
	}
	if len(instances.Instances) != 0 {
		t.Error("shouldnt be any instances ", instances)
	}

	// start should work for a registered machine
	_, err = s.Start(ctx, sr)
	if err != nil {
		t.Error("start request should have succeeded ", err)
	}

	// instances
	instances, err = s.Instances(ctx, &pb.InstancesRequest{})
	if err != nil {
		t.Error("should be able to retrieve instance ", err)
	}
	if len(instances.Instances) != 1 {
		t.Error("should be one instance ", instances.Instances)
	}


}

func TestReady(t *testing.T) {
	path := "../crawl.fsm.yaml"
	s := MakeServer()
	ctx := context.Background()

	// define crawler
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	dr := machines.AsDefineRequest(m)
	s.Define(ctx, dr)

	// no tasks
	rr := &pb.ReadyRequest{}
	tm, err := s.Ready(ctx, rr)
	if err == nil {
		t.Error("ready should error if no tasks available ", tm)
	}

	// add task
	taskId := "a"
	sr := &pb.TaskMessage{Id: taskId, Machine: m.Id}
	_, err = s.Start(ctx, sr)

	// invalid machine
	rr = &pb.ReadyRequest{Machine: m.Id + "_test"}
	tm, err = s.Ready(ctx, rr)
	if err == nil {
		t.Error("retrieving invalid machine should fail ", tm)
	}

	// invalid node
	rr = &pb.ReadyRequest{Machine: m.Id, Node: "fake"}
	tm, err = s.Ready(ctx, rr)
	if err == nil {
		t.Error("retrieving invalid node should fail ", tm)
	}

	// task exists
	rr = &pb.ReadyRequest{}
	tm, err = s.Ready(ctx, rr)
	if err != nil {
		t.Error("should have retrieved task ", err)
	} else if tm.Id != taskId {
		t.Error("unexpected task id ", tm.Id, " is not ", taskId)
	}

	// all tasks are assigned, so ready should fail
	tm, err = s.Ready(ctx, rr)
	if err == nil {
		t.Error("ready should error if no tasks available ", tm)
	}
}

func TestTransition(t *testing.T) {
	path := "../crawl.fsm.yaml"
	s := MakeServer()
	ctx := context.Background()

	// define crawler
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	dr := machines.AsDefineRequest(m)
	s.Define(ctx, dr)

	// add task
	taskId := "a"
	sr := &pb.TaskMessage{Id: taskId, Machine: m.Id}
	_, err = s.Start(ctx, sr)

	// get request
	rr := &pb.ReadyRequest{}
	tm, err := s.Ready(ctx, rr)

	startState := "crawl"
	if tm.Node != startState {
		t.Error("should be in ", startState, " but is in ", tm.Node)
	}
	// transition crawl-{error}->wait
	// wait-{ok}->crawl
	// transition crawl-{ok}->success


}

func TestRelinquish(t *testing.T) {
	path := "../crawl.fsm.yaml"
	s := MakeServer()
	ctx := context.Background()

	// define crawler
	m, err := machines.FromFile(path)
	if err != nil {
		t.Error("failed to load machine from ", path, err)
	}
	dr := machines.AsDefineRequest(m)
	s.Define(ctx, dr)

	// add task
	taskId := "a"
	sr := &pb.TaskMessage{Id: taskId, Machine: m.Id}
	_, err = s.Start(ctx, sr)

	rr := &pb.ReadyRequest{}
	tm, err := s.Ready(ctx, rr)
	tmId := tm.Id
	if err != nil {
		t.Error("should have retrieved task ", err)
	} else if tm.Id != taskId {
		t.Error("unexpected task id ", tm.Id, " is not ", taskId)
	}

	// all tasks are assigned, so ready should fail
	tm, err = s.Ready(ctx, rr)
	if err == nil {
		t.Error("ready should error if no tasks available ", tm)
	}

	// relinquish task
	rt := &pb.RelinquishRequest{Instance: tmId}
	_, err = s.Relinquish(ctx, rt)

	if err != nil {
		t.Error("couldn't relinquish task: ", err)
	}	

	// relinquished task should be availble again
	tm, err = s.Ready(ctx, rr)
	if err != nil {
		t.Error("should have retrieved task ", err)
	} else if tm.Id != taskId {
		t.Error("unexpected task id ", tm.Id, " is not ", taskId)
	}
}
