package main

import (
	"fmt"
	"strings"
)

func OuterPackageGrouper(d map[string][]string, pkg string) map[string][]string {
	useStdlib := false
	nodes := FromPrimitive(d)
	names := nodes.GetAllNodesName()
	newNodes := []Node{}

	for _, name := range names {
		node, _ := nodes.GetNode(name)

		newNode := Node{Name: node.Name}
		for _, d := range node.Dependencies {
			if strings.Contains(d, ".") && !strings.HasPrefix(d, pkg) {
				useStdlib = true
				newNode = newNode.AddDependency("vendor")
			} else {
				newNode = newNode.AddDependency(d)
			}
		}

		newNodes = append(newNodes, newNode)
	}

	if useStdlib {
		newNodes = append(newNodes, Node{"vendor", []string{"stdlib"}})
	}

	return NodeDependencyList(newNodes).ToPrimitive()
}

func OuterPackageAdder(d map[string][]string, pkg string) map[string][]string {
	nodes := FromPrimitive(d)
	names := nodes.GetAllNodesName()
	newNodes := []Node{}
	newFreeNode := Node{Name: "vendor"}

	for _, name := range names {
		node, _ := nodes.GetNode(name)

		newNode := Node{Name: node.Name}
		for _, d := range node.Dependencies {
			if strings.Contains(d, ".") && !strings.HasPrefix(d, pkg) {
				newFreeNode = newFreeNode.AddDependency(d)
			}
			newNode = newNode.AddDependency(d)
		}

		newNodes = append(newNodes, newNode)
	}

	fmt.Println(newFreeNode.Dependencies)

	for _, fn := range newFreeNode.Dependencies {
		newNodes = append(newNodes, Node{fn, []string{}})
	}

	return NodeDependencyList(newNodes).ToPrimitive()
}
