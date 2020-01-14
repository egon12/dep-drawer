package main

import "strings"

func GroupStdlibDependency(d map[string][]string) map[string][]string {
	useStdlib := false
	nodes := FromPrimitive(d)
	names := nodes.GetAllNodesName()
	newNodes := []Node{}

	for _, name := range names {
		node, _ := nodes.GetNode(name)

		newNode := Node{Name: node.Name}
		for _, d := range node.Dependencies {
			if strings.Contains(d, ".") || d == "vendor" {
				newNode = newNode.AddDependency(d)
			} else {
				useStdlib = true
				newNode = newNode.AddDependency("stdlib")
			}
		}

		newNodes = append(newNodes, newNode)
	}

	if useStdlib {
		newNodes = append(newNodes, Node{"stdlib", []string{}})
	}

	return NodeDependencyList(newNodes).ToPrimitive()

}
