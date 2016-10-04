package graphviz

import (
	"fmt"
	"io"
)

type Graph struct {
	name  string
	nodes []Node
	edges []Edge
}

type Edge struct {
	from       string
	to         string
	directed   bool
	attributes Attributes
}

func (t Edge) Output(w io.Writer) error {
	fmt.Fprintf(w, "\t%s->%s", t.from, t.to)
	if t.attributes != nil {
		t.attributes.Output(w)
	}
	fmt.Fprintln(w, ";")
	return nil
}

type Attributes map[string]string

func (t Attributes) Output(w io.Writer) error {
	if t != nil {
		fmt.Fprint(w, " [")
		for key, value := range t {
			fmt.Fprintf(w, `%s="%s",`, key, value)
		}
		fmt.Fprint(w, "]")
	}
	return nil
}

type Node struct {
	name       string
	attributes Attributes
}

func (t Node) Output(w io.Writer) error {
	fmt.Fprint(w, "\t")
	fmt.Fprint(w, t.name)
	if t.attributes != nil {
		t.attributes.Output(w)
	}
	fmt.Fprintln(w, ";")
	return nil
}

func NewGraph(name string) *Graph {
	return &Graph{
		name,
		make([]Node, 0, 5),
		make([]Edge, 0, 5),
	}
}

func (t *Graph) AddNode(name string, attributes Attributes) {
	t.nodes = append(t.nodes, Node{
		name:       name,
		attributes: attributes,
	})
}

func (t *Graph) AddEdge(from string, to string, directed bool, attributes Attributes) {
	t.edges = append(t.edges, Edge{
		from:       from,
		to:         to,
		directed:   directed,
		attributes: attributes,
	})
}

func (t Graph) Output(w io.Writer) error {
	fmt.Fprintf(w, "digraph %s {\n", t.name)
	for _, edge := range t.edges {
		edge.Output(w)
	}

	for _, node := range t.nodes {
		node.Output(w)
	}

	fmt.Fprintln(w, "}")

	return nil
}
