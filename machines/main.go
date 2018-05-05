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
	Name string `yaml:"name"`
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
	for name, node := range m.Nodes {
		transitions := make([]*pb.Transition, 0)
		for name, transition := range node.Transitions {
			t := &pb.Transition{Node: name, Name: transition}
			transitions = append(transitions, t)
		}
		n := &pb.Node{Name: name, Transitions: transitions}
		nodes = append(nodes, n)
	}
	return &pb.DefineRequest{Name: m.Name, Nodes: nodes}
}
