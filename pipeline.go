package main

import (
	"gopkg.in/yaml.v3"
	"io"

	"github.com/pkg/errors"
)

type Pipeline struct {
	Jobs []Job `yaml:"jobs"`
}

type Job struct {
	Name string `yaml:"name"`
	Plan PlanSequence
}

type PlanSequence []PlanConfig

type PlanConfig struct {
	RawName string `yaml:"name,omitempty"`

	Get    string   `yaml:"get,omitempty"`
	Passed []string `yaml:"passed,omitempty"`

	Aggregate  *PlanSequence `yaml:"aggregate,omitempty"`
	Do         *PlanSequence `yaml:"do,omitempty"`
	InParallel *PlanSequence `yaml:"in_parallel,omitempty"`

	Abort   *PlanConfig `yaml:"on_abort,omitempty"`
	Ensure  *PlanConfig `yaml:"ensure,omitempty"`
	Error   *PlanConfig `yaml:"on_error,omitempty"`
	Failure *PlanConfig `yaml:"on_failure,omitempty"`
	Success *PlanConfig `yaml:"on_success,omitempty"`
	Try     *PlanConfig `yaml:"try,omitempty"`
}

func Parse(r io.Reader) (p Pipeline, err error) {
	d := yaml.NewDecoder(r)

	err = d.Decode(&p)
	if err != nil {
		errors.Wrapf(err,
			"failed to unmarshal json into pipeline")
		return
	}

	return
}
