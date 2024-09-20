package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	lowlevelfunctions "github.com/NikoMalik/low-level-functions"
)

var ErrInvalidUUID = errors.New("invalid UUID")

type UUID [16]byte

// String returns the string representation of this UUID.
func (u *UUID) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:])
}

// Bytes returns the bytes representation of this UUID.
func (u *UUID) Bytes() []byte {
	return u[:]
}

// Equals returns true if this UUID equals another UUID by value.
// Equals returns true if this UUID equals another UUID by value.
func (u *UUID) Equals(another *UUID) bool {
	if u == nil || another == nil {
		return u == another
	}
	return lowlevelfunctions.Equal(u[:], another[:])
}

// New creates a UUID with random value.
func New() *UUID {
	var uuid UUID
	rand.Read(uuid.Bytes())
	uuid[6] = (uuid[6] & 0x0F) | (4 << 4)
	uuid[8] = (uuid[8] & 0x3F) | 0x80
	return &uuid
}

// ParseBytes converts a UUID in byte form to object.
func ParseBytes(b []byte) (UUID, error) {
	var uuid UUID
	if len(b) != 16 {
		return uuid, ErrInvalidUUID
	}
	copy(uuid[:], b)
	return uuid, nil
}

// ParseString converts a UUID in string form to object.
func ParseString(str string) (UUID, error) {
	var uuid UUID
	if len(str) != 36 {
		return uuid, ErrInvalidUUID
	}

	// Check for hyphens and process the UUID string
	if str[8] != '-' || str[13] != '-' || str[18] != '-' || str[23] != '-' {
		return uuid, ErrInvalidUUID
	}

	// Remove hyphens
	str = str[:8] + str[9:13] + str[14:18] + str[19:23] + str[24:]

	_, err := hex.Decode(uuid[:], lowlevelfunctions.StringToBytes(str))
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
