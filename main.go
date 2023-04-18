package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/deroproject/derohe/rpc"
	"github.com/gorilla/websocket"
)

var daemon_rpc_address string
var connection *websocket.Conn

var counter uint32
var errors uint32
var connections uint32
var data_size uint32

func getwork(daemon_rpc_address string, wallet_address string) {

	for {

		u := url.URL{Scheme: "wss", Host: daemon_rpc_address, Path: "/ws/" + wallet_address}
		// fmt.Println("connecting to ", "url", u.String())

		dialer := websocket.DefaultDialer
		dialer.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {

			atomic.AddUint32(&errors, 1)
			// remove connection

			time.Sleep(100 * time.Millisecond)
			continue
		}

		atomic.AddUint32(&connections, 1)
		var result *rpc.GetBlockTemplate_Result

	wait_for_another_job:

		if err = connection.ReadJSON(&result); err != nil {
			atomic.AddUint32(&errors, 1)
			atomic.AddUint32(&connections, ^uint32(0))
			// remove connection
			connection.Close()
			continue
		}

		atomic.AddUint32(&counter, 1)

		goto wait_for_another_job
	}

}

var startMiner *bool

func main() {

	os.Unsetenv("SOCKS_PROXY")
	os.Unsetenv("SOCKS5_PROXY")
	os.Unsetenv("SOCKS_SERVER")
	os.Unsetenv("SOCKS5_SERVER")
	os.Unsetenv("ALL_PROXY")
	os.Unsetenv("http_proxy")
	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("HTTPS_PROXY")
	os.Unsetenv("https_proxy")

	countFlag := flag.Int("count", 1, "Number of tests to run (10,240 max connections on official)")
	daemonAddress := flag.String("daemon-rpc-address", "localhost:10100", "Daemon address")
	walletAddress := flag.String("wallet-address", "dero1qy.....", "Wallet address")
	flag.Parse()

	// flag usage / print help
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	if len(os.Args[1]) >= 5 {
		daemon_rpc_address = os.Args[1]
	}

	fmt.Printf("Running (%d) test(s)...\n", *countFlag)

	go func() {
		for i := 1; i <= *countFlag; i++ {
			// fmt.Printf("\rRPC Test no: %d...", i)
			go getwork(*daemonAddress, *walletAddress)
			time.Sleep(10 * time.Millisecond)

		}
	}()

	now := time.Now()

	fmt.Printf("\n")
	for {

		i := atomic.LoadUint32(&counter)
		connection_count := atomic.LoadUint32(&connections)

		elapsed := float32(time.Now().Sub(now).Round(time.Second).Seconds())

		speed := float32(0)
		if elapsed >= 1 {
			speed = float32(i) / elapsed
		}
		fmt.Printf("\rConnections: %s - Jobs: %s - %s per/sec (%s Errors)...", NumberToString(connection_count, ','), NumberToString(i, ','), NumberToString(uint32(speed), ','), NumberToString(errors, ','))

		time.Sleep(100 * time.Millisecond)

		// if elapsed >= 60 {
		// 	break
		// }
	}

	fmt.Printf("\nDone! Elapsed %s - Errors: %s\n", time.Duration(time.Now().Sub(now).Round(time.Second)).Round(time.Second), NumberToString(errors, ','))
}

func NumberToString(n uint32, sep rune) string {

	s := strconv.Itoa(int(n))

	startOffset := 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	l := len(s)

	commaIndex := 3 - ((l - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < l; i++ {

		if commaIndex == 3 {
			buff.WriteRune(sep)
			commaIndex = 0
		}
		commaIndex++

		buff.WriteByte(s[i])
	}

	return buff.String()
}
