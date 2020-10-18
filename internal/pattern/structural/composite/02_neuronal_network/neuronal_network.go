package main

import "fmt"

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	name string
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron)  {
	n.Out = append(n.Out, other)
	fmt.Printf("Connecting %s with %s eachother\n", n.name, other.name)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)

	for i:= range n.Neurons {
		result = append(result, &n.Neurons[i])
	}

	return result
}

func NewNeuronLayer(name string, count int) *NeuronLayer {
	neurons := make([]Neuron, count)

	for i:=0; i<count; i++ {
		neurons[i].name = fmt.Sprintf("%s.neuron%v", name, count)
	}

	return &NeuronLayer{
		Neurons: neurons,
	}
}

func main() {
	neuron1, neuron2 := &Neuron{name: "neuron1"}, &Neuron{name: "neuron2"}
	layer1, layer2 := NewNeuronLayer("layer1",3), NewNeuronLayer("layer2", 4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}
