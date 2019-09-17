package main

import (
	"fmt"
)

type Digraph []Edge

type Edge struct {
	From, To string
}

func ToDot(g Digraph) (content string) {
	content += "digraph P {\n"

	for _, edge := range g {
		content += fmt.Sprintf("  %q -> %q;\n", edge.From, edge.To)
	}

	content += "}"

	return
}
