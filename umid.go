package umid

import (
	"encoding/binary"
	"encoding/hex"
	"sync"
	"time"
)

type umID struct {
	timestamp uint64
	counter   uint16
}

var internal struct {
	sync.Mutex
	lastUmID umID
}

// NewAsBytes returns a unique monotonic ID (umID) as []byte
func NewAsBytes() []byte {
	m := umID{}
	m.timestamp = uint64(time.Now().UnixNano())

	internal.Lock()

	if m.timestamp == internal.lastUmID.timestamp {
		m.counter = internal.lastUmID.counter + 1
	}

	internal.lastUmID = m

	internal.Unlock()

	tb := make([]byte, 8)
	binary.BigEndian.PutUint64(tb, m.timestamp)
	cb := make([]byte, 2)
	binary.BigEndian.PutUint16(cb, m.counter)

	return append(tb, cb...)
}

// NewAsString returns a unique monotonic ID (umID) as hexadecimal string
func NewAsString() string {
	return hex.EncodeToString(NewAsBytes())
}
