test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out ./atm ./bank/... ./cashbin ./db

testcover:	## Open coverage file as html
	go tool cover -html=c.out
