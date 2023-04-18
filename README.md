# DERO-GetWork-Tester


Simple GetWork Tester.

Used to make sure your mining server is working by checking it is sending jobs and has no error. 
Or to check your nodes max capacity.

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

# Example Usage
```bash
./getwork-tester -count 1000 -daemon-rpc-address localhost:10100 -wallet-address dero1qy07h9mk6xxf2k4x0ymdpezvksjy0talskhpvqmat3xk3d9wczg5jqqvwl0sn
Running (1000) test(s)...

Connections: 1,000 - Jobs: 34,993 - 1,521 per/sec (0 Errors)...^C

```