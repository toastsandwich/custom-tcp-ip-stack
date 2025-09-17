package ipheader

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	MINIPHEADERLEN = 20
	MINIHL         = 5
	MAXOPTIONSLEN  = 40
)

var ErrInvalidBuffer = errors.New("ip: invalid buffer length")

type IPv4Frame struct {
	VIHL           byte // Version + Header Length
	ToS            byte // type of service
	TotalLen       uint16
	Identification uint16

	// FragmentOffset is 16 bits divided into 3 + 13
	// first bit null bit, second bit will be dont fragment bit, third bit more fragment bit
	FragmentOffset uint16

	TimeToLive     byte
	Protocol       byte
	HeaderChecksum uint16
	SourceIP       uint32
	DestinationIP  uint32
	Data           []byte

	Options []byte // this is optional
}

func Parse(buf []byte) (*IPv4Frame, error) {
	if len(buf) < MINIPHEADERLEN {
		return nil, ErrInvalidBuffer
	}

	ipv4 := &IPv4Frame{}

	vihl := buf[0]
	ihl := int(vihl & 0x0F)
	headerlen := ihl * 4

	if len(buf) < headerlen {
		return nil, ErrInvalidBuffer
	}
	ipv4.VIHL = vihl
	ipv4.ToS = buf[1]
	ipv4.TotalLen = binary.BigEndian.Uint16(buf[2:4])
	ipv4.Identification = binary.BigEndian.Uint16(buf[4:6])
	ipv4.FragmentOffset = binary.BigEndian.Uint16(buf[6:8])
	ipv4.TimeToLive = buf[8]
	ipv4.Protocol = buf[9]
	ipv4.HeaderChecksum = binary.BigEndian.Uint16(buf[10:12])
	ipv4.SourceIP = binary.BigEndian.Uint32(buf[12:16])
	ipv4.DestinationIP = binary.BigEndian.Uint32(buf[16:20])
	var optlen int
	if ihl > MINIHL {
		optlen = headerlen - MINIPHEADERLEN
		ipv4.Options = make([]byte, optlen)
		copy(ipv4.Options, buf[20:20+optlen])
	}
	ipv4.Data = make([]byte, len(buf)-headerlen)
	copy(ipv4.Data, buf[headerlen:])
	return ipv4, nil
}

func (i *IPv4Frame) String() string {
	sip := fmt.Sprintf("%d:%d:%d:%d", (i.SourceIP>>24)&0xFF, (i.SourceIP>>16)&0xFF, (i.SourceIP>>8)&0x0FF, (i.SourceIP & 0xFF))
	dip := fmt.Sprintf("%d:%d:%d:%d", (i.DestinationIP>>24)&0xFF, (i.DestinationIP>>16)&0x00F0, (i.DestinationIP>>8)&0x0FF, i.DestinationIP&0xFF)
	return fmt.Sprintf("srcip: %s | dstip: %s | proto: %d", sip, dip, i.Protocol)
}
