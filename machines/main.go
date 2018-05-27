package machines

import (
	"io/ioutil"
	
	"gopkg.in/yaml.v2"

	pb "github.com/lethain/dfsmr/dfsmr"	
)


type Node struct {
	Input string `yaml:"input"`
	Transitions map[string]string `yaml:"transitions"`
	Final bool `yaml:"final"`
}

type Machine struct {
	Id string `yaml:"id"`
	Nodes map[string]Node `yaml:"nodes"`
}

func FromYAML(data []byte) (*Machine, error) {
	m := &Machine{}	
	err := yaml.Unmarshal(data, m)
	return m, err
}

func FromFile(filepath string) (*Machine, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return FromYAML(data)
}

func AsDefineRequest(m *Machine) *pb.DefineRequest {
	nodes := make([]*pb.Node, 0)
	for id, node := range m.Nodes {
		transitions := make([]*pb.Transition, 0)
		for id, transition := range node.Transitions {
			t := &pb.Transition{Id: id, Node: transition}
			transitions = append(transitions, t)
		}
		n := &pb.Node{Id: id, Transitions: transitions}
		nodes = append(nodes, n)
	}
	return &pb.DefineRequest{Id: m.Id, Nodes: nodes}
}
