test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out ./...

testcover:	## Open coverage file as html
	go tool cover -html=c.out
