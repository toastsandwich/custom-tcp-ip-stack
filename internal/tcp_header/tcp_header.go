package tcpheader

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	MIN_TCP_HEADER_LEN = 20
)

var ErrInvalidBuffer = errors.New("tcp: invalid buffer")

type TCPFrame struct {
	SrcPort uint16 // 2
	DstPort uint16 // 2
	Seq     uint32 // 4
	Ack     uint32 // 4

	// HLEN, Reserved, URG, ACK, PSH, RST, SYN, FIN
	HLENRUAPRSF uint16 // 2

	WinSize        uint16 // 2
	Checksum       uint16 // 2
	UrgentPointer  uint16 // 2
	OptionsPadding []byte
	Data           []byte
}

func Parse(buf []byte) (*TCPFrame, error) {
	if len(buf) < MIN_TCP_HEADER_LEN {
		return nil, ErrInvalidBuffer
	}

	tcp := &TCPFrame{}

	tcp.SrcPort = binary.BigEndian.Uint16(buf[0:2])
	tcp.DstPort = binary.BigEndian.Uint16(buf[2:4])
	tcp.Seq = binary.BigEndian.Uint32(buf[4:8])
	tcp.Ack = binary.BigEndian.Uint32(buf[8:12])
	tcp.HLENRUAPRSF = binary.BigEndian.Uint16(buf[12:14])

	tcp.WinSize = binary.BigEndian.Uint16(buf[14:16])
	tcp.Checksum = binary.BigEndian.Uint16(buf[16:18])
	tcp.UrgentPointer = binary.BigEndian.Uint16(buf[18:20])

	// why 12 ? beacuse we need the first 4 bits, so move to right by size - 4 and then multiply by 4
	hlen := int((tcp.HLENRUAPRSF >> 12) * 4)

	optlen := hlen - MIN_TCP_HEADER_LEN
	if optlen > 0 {
		tcp.OptionsPadding = make([]byte, optlen)
		copy(tcp.OptionsPadding[:], buf[20:20+optlen])
	}

	if len(buf) > hlen {
		tcp.Data = make([]byte, len(buf)-hlen)
		copy(tcp.Data[:], buf[hlen:])
	}

	return tcp, nil
}

func (t *TCPFrame) String() string {
	return fmt.Sprintf("src: %d dst: %d\ndata: %s", t.SrcPort, t.DstPort, string(t.Data))
}
