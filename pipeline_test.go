package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestConversion(t *testing.T) {
	var cases = []struct {
		desc     string
		content  string
		expected Digraph
	}{
		{
			desc: "no deps",
			content: `
jobs:
- name: test
  plan:
  - get: resource
`,
		},

		{
			desc: "single job, with deps",
			content: `
jobs:
- name: test
  plan:
  - get: resource
    passed: [a,b,c]
`,
			expected: Digraph{
				{From: "a", To: "test"},
				{From: "b", To: "test"},
				{From: "c", To: "test"},
			},
		},

		{
			desc: "multiple jobs, with deps",
			content: `
jobs:
- name: test
  plan:
  - get: resource
    passed: [a]

- name: test2
  plan:
  - get: resource
    passed: [test]
`,
			expected: Digraph{
				{From: "a", To: "test"},
				{From: "test", To: "test2"},
			},
		},

		{
			desc: "with seqs, with deps",
			content: `
jobs:
- name: test
  plan:
  - in_parallel:
    - get: resource
      passed: [a]
  - get: resource
    passed: [b]
`,
			expected: Digraph{
				{From: "a", To: "test"},
				{From: "b", To: "test"},
			},
		},

		{
			desc: "with duplicate deps",
			content: `
jobs:
- name: test
  plan:
  - in_parallel:
    - get: resource
      passed: [a,a,a,a]
`,
			expected: Digraph{
				{From: "a", To: "test"},
			},
		},
	}

	var (
		err      error
		pipeline Pipeline
		actual   Digraph
	)

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			pipeline, err = Parse(bytes.NewReader([]byte(tc.content)))
			if err != nil {
				t.Errorf("should've not failed parsing")
				return
			}

			actual, err = ToDigraph(pipeline)
			if err != nil {
				t.Errorf("should've not failed converting to digraph")
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("%+v != %+v", actual, tc.expected)
				return
			}
		})
	}
}
