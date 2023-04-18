# DERO-GetWork-Tester


Simple GetWork Tester.

Used to make sure your mining server is working by checking it is sending jobs and has no error. 
Or to check your node's it's max capacity.

# Getting Started

- Download and compile.

```bash
git clone https://github.com/Hansen333/DERO-GetWork-Tester.git
cd DERO-GetWork-Tester
go build
./getwork-tester
```

# Usage

```bash
Usage of ./getwork-tester:
  -count int
    	Number of tests to run (10,240 max connections on official) (default 1)
  -daemon-rpc-address string
    	Daemon address (default "localhost:10100")
  -wallet-address string
    	Wallet address (default "dero1qy.....")
```