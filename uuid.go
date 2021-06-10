package uuid

import (
    "crypto/rand"
    "fmt"
    "io"
)

//UUID is a 128 bit(16 byte) as defined in RFC 4122.
type UUID [16]byte

// NewV4 generates a random UUID according to RFC 4122.
func NewV4() (UUID, error) {
	var uuid UUID
	buffer := make([]byte, 16)
    n, err := io.ReadFull(rand.Reader, buffer)
    if n != len(uuid) || err != nil {
        return uuid, err // array with nothing useful
    }
	for i := range buffer {
		uuid[i] = buffer[i]
	}
    // variant bits; see section 4.1.1
    uuid[8] = uuid[8]&^0xf0 | 0x80

    // version 4 (pseudo-random); see section 4.1.3
    uuid[6] = uuid[6]&^0xf0 | 0x40
	return uuid, nil
}

// Now you can fmt.Printf("%s", uuidTypeVar) and get your expected result
func (u UUID) String() string {
    return fmt.Sprintf(
		"%x-%x-%x-%x-%x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:],
	)
}
