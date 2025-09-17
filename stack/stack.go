package stack

import (
	"fmt"
	"os"
	"strings"

	etherheader "github.com/toastsandwich/custom-tcp-ip-stack/internal/ether_header"
	ipheader "github.com/toastsandwich/custom-tcp-ip-stack/internal/ip_header"
	tcpheader "github.com/toastsandwich/custom-tcp-ip-stack/internal/tcp_header"
	"golang.org/x/sys/unix"
)

const MAXPACKETLEN = 65536

type Stack struct {
	Eth *etherheader.EthernetFrame
	Ip  *ipheader.IPv4Frame
	Tcp *tcpheader.TCPFrame
}

func (s *Stack) String() string {
	b := strings.Builder{}
	b.WriteString("***************" + "\n")
	if s.Eth != nil {
		b.WriteString(s.Eth.String() + "\n")
	}
	if s.Ip != nil {
		b.WriteString(s.Ip.String() + "\n")
	}
	if s.Tcp != nil {
		b.WriteString(s.Tcp.String() + "\n")
	}
	b.WriteString("***************" + "\n")
	return b.String()
}

func Main() {
	fd, err := unix.Socket(unix.AF_PACKET, unix.SOCK_RAW, int(htons(unix.ETH_P_ALL)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer unix.Close(fd)

	// now read for packets
	for {
		stk := &Stack{}
		buf := make([]byte, MAXPACKETLEN)
		n, _, err := unix.Recvfrom(fd, buf, 0)
		if err != nil {
			fmt.Println(err)
		}
		eth, err := etherheader.Parse(buf[:n])
		if err != nil {
			fmt.Println(err)
			continue
		}
		stk.Eth = eth

		if eth.TYPE == unix.ETH_P_IP {
			iph, err := ipheader.Parse(eth.DATA)
			if err != nil {
				fmt.Println(err)
				continue
			}
			stk.Ip = iph

			if iph.Protocol == unix.IPPROTO_TCP {
				tcp, err := tcpheader.Parse(iph.Data)
				if err != nil {
					fmt.Println(err)
					continue
				}
				stk.Tcp = tcp
			}
		}

		fmt.Println(stk)
	}
}

// we are using this to tell kernel that i want data in BigEndian format
func htons(n uint16) uint16 {
	// first shift to left 0x1234 = 0x3400
	// then shift to write 0x0012
	// now combine them 0x3421
	return (n << 8) | (n >> 8)
}
