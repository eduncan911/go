package uuid

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

var (
	// EmptyID is a placeholder for an empty ID
	EmptyID ID
	seq     uint32
	node    = nodeID()
)

// ID is a 16 byte array that represents the UUID
type ID [16]byte

func nodeID() uint32 {
	n := uuid.NodeID()
	return binary.BigEndian.Uint32(n)
}

// NewID returns a new UUID
func NewID() ID {
	var uuid [16]byte
	var now = time.Now().UTC()

	nano := now.UnixNano()
	incr := atomic.AddUint32(&seq, 1)

	binary.BigEndian.PutUint64(uuid[0:], uint64(nano))
	binary.BigEndian.PutUint32(uuid[8:], incr)
	binary.BigEndian.PutUint32(uuid[12:], node)

	return uuid
}

// Bytes return a bytes array
func (id ID) Bytes() []byte {
	return id[:]
}

// Time returns the time at which the ID was created
func (id ID) Time() time.Time {
	bytes := id[:]
	nsec := binary.BigEndian.Uint64(bytes)
	return time.Unix(0, int64(nsec)).UTC()
}

// Equals compares the specified ID to the underlying ID
func (id ID) Equals(other ID) bool {
	return bytes.Equal(id.Bytes(), other.Bytes())
}

// IsEmpty returns a boolean indicating the ID is empty
func (id ID) IsEmpty() bool {
	for _, x := range id {
		if x != 0 {
			return false
		}
	}
	return true
}

// String returns a string-readable representation of the ID
func (id ID) String() string {
	return hex.EncodeToString(id[:])
}

// ParseID parses the specified string version of ID and returns an ID or an Error
func ParseID(value string) (id ID, err error) {
	if len(value) == 0 {
		err = fmt.Errorf("Invalid id: value is empty")
		return
	}

	var b []byte
	orgValue := value

	if len(value) != 32 {
		value = strings.Map(func(r rune) rune {
			if r == '-' || r == '{' || r == '}' {
				return -1
			}
			return r
		}, value)
	}

	if b, err = hex.DecodeString(value); err != nil {
		err = fmt.Errorf("invalid id %v: %v", orgValue, err.Error())
		return
	}

	if len(b) != 16 {
		err = fmt.Errorf("invalid id %v: did not convert to a 16 byte array", orgValue)
		return
	}

	for index, value := range b {
		id[index] = value
	}

	return
}

// MarshalJSON handles marshaling of this type
func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsEmpty() {
		return []byte("\"\""), nil
	}

	jsonString := `"` + hex.EncodeToString(id[:]) + `"`
	return []byte(jsonString), nil
}

// UnmarshalJSON handles unmarshaling this type
func (id *ID) UnmarshalJSON(data []byte) error {
	jsonString := string(data)
	valueString := strings.Trim(jsonString, "\"")

	value, err := ParseID(valueString)
	if err != nil {
		return err
	}

	*id = value
	return nil
}
