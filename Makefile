.PHONY: test
test:
	@go test -race -covermode=atomic -coverprofile cp.out

.PHONY: show-covg
show-covg:
	@go tool cover -html=cp.out    
