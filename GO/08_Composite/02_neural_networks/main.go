package main

// interface to implement Composite functions to treat Neuron and NeuronLayer as scalar object
type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

//factory func
func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{
		Neurons: make([]Neuron, count),
	}
}

func (nl *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)

	for i := range nl.Neurons {
		result = append(result, &nl.Neurons[i])
	}
	return result
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	// This connect func needs to be a single method, not 4 different methods with different signatures
	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

}
