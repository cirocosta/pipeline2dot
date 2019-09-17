package main

func ToDigraph(p Pipeline) (g Digraph, err error) {
	for _, job := range p.Jobs {
		deps := unique(getJobDependencies(&job.Plan))

		for _, dep := range deps {
			g = append(g, Edge{
				From: dep,
				To:   job.Name,
			})
		}
	}

	return
}

func unique(slice []string) (u []string) {
	m := map[string]interface{}{}

	for _, item := range slice {
		m[item] = nil
	}

	for k := range m {
		u = append(u, k)
	}

	return
}

func getJobDependencies(s *PlanSequence) (dependencies []string) {
	for _, cfg := range *s {
		dependencies = append(dependencies, getDependencies(&cfg)...)
	}

	return
}

func getDependencies(cfg *PlanConfig) (dependencies []string) {
	if len(cfg.Passed) != 0 {
		return cfg.Passed
	}

	switch {
	case cfg.Abort != nil:
		return getDependencies(cfg.Abort)
	case cfg.Ensure != nil:
		return getDependencies(cfg.Ensure)
	case cfg.Error != nil:
		return getDependencies(cfg.Error)
	case cfg.Failure != nil:
		return getDependencies(cfg.Failure)
	case cfg.Success != nil:
		return getDependencies(cfg.Success)
	case cfg.Try != nil:
		return getDependencies(cfg.Try)

	case cfg.Aggregate != nil:
		return getJobDependencies(cfg.Aggregate)
	case cfg.Do != nil:
		return getJobDependencies(cfg.Do)
	case cfg.InParallel != nil:
		return getJobDependencies(cfg.InParallel)
	}

	return
}
