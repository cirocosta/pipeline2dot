package main

type Digraph []Edge

type Edge struct {
	From, To string
}

func ToDot(g Digraph) (content string) {
	content += `digraph pipeline-graph {`

	for _, edge := range g {
		content += edge.From + "->" + edge.To + "\n"
	}

	content += "}"

	return
}
