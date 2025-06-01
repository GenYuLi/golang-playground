package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"time"
)

func sendCh(testCh chan<- int64) {
	testCh <- 32
}

func dialTLSAsync(ctx context.Context, addr string, host string, timeout time.Duration) (<-chan *tls.Conn, <-chan error) {
	connCh := make(chan *tls.Conn, 1)
	errCh := make(chan error, 1)

	go func() {
		defer close(connCh)
		defer close(errCh)

		dialer := &net.Dialer{
			Timeout: timeout, // TCP SYN/ACK deadline
		}

		tlsCfg := &tls.Config{
			ServerName: host,
		}

		conn, err := tls.DialWithDialer(dialer, "tcp", addr, tlsCfg)
		if err != nil {
			errCh <- err
			return
		}
		if err := conn.HandshakeContext(ctx); err != nil {
			conn_err := conn.Close()
			if conn_err != nil {
				fmt.Println(conn_err)
			}
			errCh <- err
			return
		}

		connCh <- conn
	}()
	return connCh, errCh
}

func getHostName(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		fmt.Println("get error while getting host name:", err)
		panic("damn")
	}
	if net.ParseIP(host) != nil {
		fmt.Println("get error while parsing ip:", err)
		panic("fucked")
	}
	return host
}

func main() {

	testCh := make(chan int64)

	go sendCh(testCh)

	integer := <-testCh

	fmt.Println(integer)

	// ------ define flag -------
	addr := flag.String("addr", "example.com:443", "server address <host:port>")
	timeout := flag.Duration("timeout", 5*time.Second, "overall timeout (e.g. 3s, 500ms)")
	flag.Parse()
	// --------------------------

	host := getHostName(*addr)
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	connCh, errCh := dialTLSAsync(ctx, *addr, host, 3*time.Second)
	var testh int64

	select {

	case conn := <-connCh:
		fmt.Println("TLS connection established:", conn.RemoteAddr())
		testh = 340

		if verify_err := conn.VerifyHostname(host); verify_err != nil {
			panic("Hostname doesn't match with certificate: " + verify_err.Error())
		}

		expiries := conn.ConnectionState().PeerCertificates
		for idx, expiry := range expiries {
			fmt.Printf("%d-th expiry\n", idx)
			fmt.Printf("Issuer Name: %s\n", expiry.Issuer)
			fmt.Printf("Expiry: %s \n", expiry.NotAfter.Format("2006-01-02"))
			fmt.Printf("Common name: %s \n\n", expiry.Issuer.CommonName)
		}
		if conn_err := conn.Close(); conn_err != nil {
			fmt.Println("you fucked up, lol")
		}

	case err := <-errCh:
		fmt.Println("Dial failed:", err)

	case <-ctx.Done():
		fmt.Println("Global timeout reached:", ctx.Err())
	}

	fmt.Println(testh)
}
