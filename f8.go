package snow3g

import (
	"encoding/binary"
)

func F8(ck []byte, count uint32, bearer, dir uint8, ibs []byte, length uint64) []byte {
	k := make([]uint32, 4)
	for i := 0; i < len(k); i++ {
		k[len(k)-i-1] = binary.BigEndian.Uint32(ck[i*4:])
	}
	iv := make([]uint32, 4)
	iv[3] = count
	iv[2] = uint32(bearer) << 27
	iv[2] |= uint32(dir) << 26
	iv[1] = iv[3]
	iv[0] = iv[2]
	s := New(k, iv)
	obs := make([]byte, len(ibs))
	l := (length + 0x7) >> 3
	n, _ := s.Read(obs[:l])
	shift := length & 0x7
	if shift != 0 {
		obs[l-1] &^= 0xff >> shift
	}
	for i := 0; i < n; i++ {
		obs[i] ^= ibs[i]
	}
	return obs
}
