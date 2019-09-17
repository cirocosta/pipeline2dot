package main

import (
	"bytes"
	"testing"
)

var pipelineCases = []struct {
	desc, pipeline string
}{
	{
		desc: "simple",
		pipeline: `
resources:
- name: lol
  source:
    aa

jobs:
- name: test
  plan:
  - get: lol
`,
	},
	{
		desc: "with in_parallel",
		pipeline: `
jobs:
  - in_parallel:
      steps:
      - get: concourse
        trigger: true
`,
	},
}

func TestParsing(t *testing.T) {
	for _, tc := range pipelineCases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := Parse(bytes.NewReader([]byte(tc.pipeline)))
			if err != nil {
				t.Errorf("failed to parse pipeline\n%s\n >> %v", tc.pipeline, err)
			}
		})
	}
}
