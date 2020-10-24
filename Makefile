.PHONY: test
test:
	@go test -coverprofile tmp/cp.out

.PHONY: show-covg
show-covg:
	@go tool cover -html=tmp/cp.out    
