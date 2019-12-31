package main

import (
	"errors"
	"log"
	"strings"
)

type NodeDependencyList []Node

type DependencySet []string

type Node struct {
	Name         string
	Dependencies DependencySet
}

func (n Node) AddDependency(dependency string) Node {
	for _, v := range n.Dependencies {
		if v == dependency {
			return n
		}
	}

	n.Dependencies = append(n.Dependencies, dependency)
	return n
}

func (n Node) AddDependencies(dependencies ...string) Node {
	newN := n
	for _, dependency := range dependencies {
		newN = newN.AddDependency(dependency)
	}
	return newN
}

func FromPrimitive(d map[string][]string) NodeDependencyList {
	lenOfD := len(d)

	nodes := make([]Node, lenOfD, lenOfD)

	i := 0
	for n, ad := range d {
		node := Node{n, ad}
		nodes[i] = node
		i++
	}

	return NodeDependencyList(nodes)
}

func (n NodeDependencyList) ToPrimitive() map[string][]string {
	result := map[string][]string{}
	for _, node := range n {
		result[node.Name] = node.Dependencies
	}
	return result
}

func (n NodeDependencyList) GetAllNodesName() []string {
	lenOfN := len(n)
	result := make([]string, lenOfN, lenOfN)
	for i, node := range n {
		result[i] = node.Name
	}
	return result
}

func (n NodeDependencyList) AddNode(node Node) NodeDependencyList {
	return append(n, node)
}

func (n NodeDependencyList) GetNode(name string) (Node, error) {
	for _, node := range n {
		if node.Name == name {
			return node, nil
		}
	}
	return Node{}, errors.New("Not found")
}

func (n NodeDependencyList) GetNodesStartWith(name string) ([]Node, error) {
	var result []Node
	for _, node := range n {
		if strings.HasPrefix(node.Name, name) {
			result = append(result, node)
		}
	}

	if len(result) == 0 {
		return result, errors.New("Not found")
	}

	return result, nil
}

func (n NodeDependencyList) DeleteNodes(names ...string) (NodeDependencyList, error) {
	newN := n
	var err error

	for _, name := range names {
		newN, err = newN.DeleteNode(name)
		if err != nil {
			return newN, err
		}
	}

	return newN, nil
}

func (n NodeDependencyList) DeleteNode(name string) (NodeDependencyList, error) {
	var result []Node
	deleted := false

	for _, node := range n {
		if node.Name == name {
			deleted = true
			continue
		}
		result = append(result, node)
	}

	if !deleted {
		return result, errors.New("Not found")
	}

	return result, nil
}

func GroupBy(nodeName string, dependencies map[string][]string) map[string][]string {
	n := FromPrimitive(dependencies)

	willBeDeleted, err := n.GetNodesStartWith(nodeName)
	if err != nil {
		log.Println(err)
		return n.ToPrimitive()
	}

	node := Node{Name: nodeName}

	for _, nn := range willBeDeleted {
		node = node.AddDependencies(nn.Dependencies...)
	}

	okToDelete := NodeDependencyList(willBeDeleted).GetAllNodesName()

	n, err = n.DeleteNodes(okToDelete...)
	if err != nil {
		log.Println(err)
	}

	n = n.AddNode(node)

	return n.ToPrimitive()
}
