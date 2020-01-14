package main

import "strings"

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
		newNodes = append(newNodes, Node{"vendor", []string{}})
	}

	return NodeDependencyList(newNodes).ToPrimitive()
}
