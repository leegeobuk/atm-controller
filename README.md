# atm-controller
atm-controller is a simple ATM Controller that hugely simplified real-world atm-controller.  

# Prerequisite  
- git
- Go 1.18+
- Makefile

# Clone project  
## HTTPs  
```bash
git clone https://github.com/leegeobuk/atm-controller.git
```
## SSH
```bash
git clone git@github.com:leegeobuk/atm-controller.git
```

# Test
## Unit test
```bash
make test
```

## coverage.html
```bash
make testcover
```

# Build
```bash
make build os=<os> arch=<arch> name=<name>
```
## Linux
```bash
make build os=linux arch=amd64 name=atm-linux-amd64
```

## MacOS
```bash
make build os=darwin arch=amd64 name=atm-darwin-amd64
```

## Windows
```bash
make build os=windows arch=amd64 name=atm-windows-amd64.exe
```

After the ```make build``` command, then executable binary will be created with name you entered in ```make build```  
Full list of supported OS and architecture: https://github.com/golang/go/blob/master/src/go/build/syslist.go

# Run
## Executable
if name = atm-linux-amd64 then run  
```bash
./atm-linux-amd64
```

## Go run command
```bash
go run main.go
```