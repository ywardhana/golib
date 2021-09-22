dep:
	env GO111MODULE=on go mod download

tidy:
	env GO111MODULE=on go mod tidy

vendor:
	env GO111MODULE=on go mod vendor