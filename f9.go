package snow3g

import (
	"encoding/binary"
)

func F9(ik []byte, count, fresh uint32, dir uint8, msg []byte, l uint64) []byte {
	k := make([]uint32, 4)
	for i := 0; i < len(k); i++ {
		k[len(k)-i-1] = binary.BigEndian.Uint32(ik[i*4:])
	}
	iv := make([]uint32, 4)
	iv[3] = count
	iv[2] = fresh
	iv[1] = count
	iv[1] ^= uint32(dir) << 31
	iv[0] = fresh
	iv[0] ^= uint32(dir) << 15
	s := New(k, iv)
	z := make([]byte, 20)
	s.Read(z)
	P := binary.BigEndian.Uint64(z[0:8])
	Q := binary.BigEndian.Uint64(z[8:16])
	OTP := binary.BigEndian.Uint32(z[16:20])
	D := int((l+0x3f)>>6) + 1
	EVAL := uint64(0)
	for i := 0; i < D-2; i++ {
		v := EVAL
		v ^= binary.BigEndian.Uint64(msg[i*8:])
		EVAL = Mul64(v, P, 0x1b)
	}
	rem := make([]byte, 8)
	copy(rem, msg[8*(D-2):])
	v := EVAL
	v ^= binary.BigEndian.Uint64(rem)
	EVAL = Mul64(v, P, 0x1b)
	EVAL ^= l
	EVAL = Mul64(EVAL, Q, 0x1b)
	maci := make([]byte, 4)
	binary.BigEndian.PutUint32(maci, uint32(EVAL>>32)^OTP)
	return maci
}

func Mul64(v, p, c uint64) uint64 {
	r := uint64(0)
	for i := 0; i < 64; i++ {
		if (p>>i)&1 != 0 {
			r ^= Mul64xPow(v, i, c)
		}
	}
	return r
}

func Mul64xPow(v uint64, i int, c uint64) uint64 {
	if i == 0 {
		return v
	}
	return Mul64x(Mul64xPow(v, i-1, c), c)
}

func Mul64x(v, c uint64) uint64 {
	if (v>>63)&1 != 0 {
		return (v << 1) ^ c
	}
	return v << 1
}
