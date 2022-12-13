test:  		## Run unit test with coverage
	go clean -testcache
	go test -cover -coverprofile=c.out ./atm ./bank ./bank/account ./cashbin ./db

testcover:	## Open coverage file as html
	go tool cover -html=c.out

build: 		## Build executable
	set	GOOS=$(os)& set	GOARCH=$(arch)& go build -o ./$(name) main.go
