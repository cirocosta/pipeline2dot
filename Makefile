SAMPLES = 


install:
	go install -v .


test:
	go test -v ./...

samples:
	find ./samples -name "*.yml" | \
		xargs -I[] /bin/bash -c \
			'pipeline2dot -i [] | dot -Tpng > [].png'
.PHONY: samples
