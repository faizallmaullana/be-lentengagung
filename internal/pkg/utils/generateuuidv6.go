package utils

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateUUIDV6() uuid.UUID {
	b := make([]byte, 16)

	// timestamp in microseconds
	ts := uint64(time.Now().UnixNano() / 1e3)

	// put top 48 bits of timestamp into b[0..5] big-endian
	b[0] = byte(ts >> 40)
	b[1] = byte(ts >> 32)
	b[2] = byte(ts >> 24)
	b[3] = byte(ts >> 16)
	b[4] = byte(ts >> 8)
	b[5] = byte(ts)

	// fill remaining bytes with randomness
	if _, err := rand.Read(b[6:]); err != nil {
		// fallback to zeroed random area on failure
		for i := 6; i < 16; i++ {
			b[i] = 0
		}
	}

	// set version to 6 (top 4 bits of b[6])
	b[6] = (b[6] & 0x0f) | (6 << 4)

	// set RFC4122 variant (top two bits of b[8] to 10)
	b[8] = (b[8] & 0x3f) | 0x80

	u, _ := uuid.FromBytes(b)
	return u
}
