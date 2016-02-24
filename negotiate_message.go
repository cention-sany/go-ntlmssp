package ntlmssp

import (
	"bytes"
	"encoding/binary"
)

type negotiateMessageFields struct {
	messageHeader
	NegotiateFlags negotiateFlags
	_              [5]uint32
	offset1        uint16
	_              uint32
	offset2        uint16
}

//NewNegotiateMessage creates a new NEGOTIATE message with the
//flags that this package supports.
func NewNegotiateMessage() []byte {
	m := negotiateMessageFields{
		messageHeader: newMessageHeader(1),
		offset1:       0x30,
		offset2:       0x30,
	}

	m.NegotiateFlags = negotiateFlagNTLMSSPREQUESTTARGET |
		negotiateFlagNTLMSSPNEGOTIATENTLM |
		negotiateFlagNTLMSSPNEGOTIATEALWAYSSIGN |
		negotiateFlagNTLMSSPNEGOTIATEUNICODE

	b := bytes.Buffer{}
	err := binary.Write(&b, binary.LittleEndian, &m)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}
