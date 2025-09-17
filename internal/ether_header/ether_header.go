package etherheader

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

var ErrInvalidBuffer = errors.New("eth: invalid buffer length")

const MINBUFFERLEN = 14

type EthernetFrame struct {
	DMAC [6]byte
	SMAC [6]byte
	TYPE uint16
	DATA []byte
	CRC  uint32
}

func NewEtherFrame(datalen int) *EthernetFrame {
	return &EthernetFrame{
		// PRE:  [7]byte{},
		DMAC: [6]byte{},
		SMAC: [6]byte{},
		DATA: make([]byte, datalen),
	}
}

func Parse(buf []byte) (*EthernetFrame, error) {
	if len(buf) < MINBUFFERLEN {
		return nil, ErrInvalidBuffer
	}

	datalen := len(buf) - MINBUFFERLEN
	EthFrame := NewEtherFrame(datalen)

	copy(EthFrame.DMAC[:], buf[0:6])
	copy(EthFrame.SMAC[:], buf[6:12])
	EthFrame.TYPE = binary.BigEndian.Uint16(buf[12:14])
	copy(EthFrame.DATA[:], buf[14:14+datalen])
	EthFrame.CRC = binary.BigEndian.Uint32(buf[len(buf)-4:])
	return EthFrame, nil
}

func (e *EthernetFrame) String() string {
	return fmt.Sprintf("to: %s from: %s, type: %d data: %s", hex.EncodeToString(e.DMAC[:]), hex.EncodeToString(e.SMAC[:]), e.TYPE, hex.EncodeToString(e.DATA))
}
