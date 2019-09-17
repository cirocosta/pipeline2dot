package main

import (
	"testing"
)

var graphPrintingCases = []struct {
	desc  string
	graph Digraph
	out   string
}{
	{
		desc: "empty",
		out:  "digraph P {\n}",
	},
	{
		desc: "single edge",
		graph: Digraph{
			{From: "a", To: "b"},
		},
		out: `digraph P {
  "a" -> "b";
}`,
	},
	{
		desc: "multiple edge",
		graph: Digraph{
			{From: "a", To: "b"},
			{From: "b", To: "c"},
		},
		out: `digraph P {
  "a" -> "b";
  "b" -> "c";
}`,
	},
}

func TestGraph(t *testing.T) {
	for _, tc := range graphPrintingCases {
		t.Run(tc.desc, func(t *testing.T) {
			out := ToDot(tc.graph)

			if out == tc.out {
				return
			}

			t.Errorf("%s != %s", out, tc.out)
		})
	}
}
