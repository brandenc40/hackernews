.PHONY: test
test:
	@go test -race -covermode=atomic -coverprofile cp.out

.PHONY: show-covg
show-covg: test
	@go tool cover -html=cp.out    

.PHONY: escape-analysis
escape-analysis:
	@go build -gcflags="-m -l"

.PHONY: maligned
maligned:
	@golangci-lint run --disable-all -E maligned item.go
