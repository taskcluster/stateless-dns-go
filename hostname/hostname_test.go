package hostname_test

import (
	"fmt"
	"net"
	"time"

	"github.com/taskcluster/stateless-dns-go/hostname"
)

func ExampleNew() {
	ip := net.IPv4(byte(203), byte(43), byte(55), byte(2))
	subdomain := "foo.com"
	expires := time.Now().Add(15 * time.Minute)
	secret := "turnip4tea"
	fmt.Println(hostname.New(ip, subdomain, expires, secret))
}
