package snow3g

import (
	"bytes"
	"testing"
)

func Test1(t *testing.T) {
	k := []uint32{0x2bd6459f, 0x82c5b300, 0x952c4910, 0x4881ff48}
	iv := []uint32{0xea024714, 0xad5c4d84, 0xdf1f9b25, 0x1c0bf45f}
	z := []byte{0xab, 0xee, 0x97, 0x04, 0x7a, 0xc3, 0x13, 0x73}

	ks := make([]byte, 8)
	s := New(k, iv)
	s.Read(ks)
	if !bytes.Equal(ks, z) {
		t.Errorf("want %x; but got %x\n", z, ks)
	}
}

func Test2(t *testing.T) {
	k := []uint32{0x8ce33e2c, 0xc3c0b5fc, 0x1f3de8a6, 0xdc66b1f3}
	iv := []uint32{0xd3c5d592, 0x327fb11c, 0xde551988, 0xceb2f9b7}
	z := []byte{0xef, 0xf8, 0xa3, 0x42, 0xf7, 0x51, 0x48, 0x0f}

	ks := make([]byte, 8)
	s := New(k, iv)
	s.Read(ks)
	if !bytes.Equal(ks, z) {
		t.Errorf("want %x; but got %x\n", z, ks)
	}
}

func Test3(t *testing.T) {
	k := []uint32{0x4035c668, 0x0af8c6d1, 0xa8ff8667, 0xb1714013}
	iv := []uint32{0x62a54098, 0x1ba6f9b7, 0x4592b0e7, 0x8690f71b}
	z := []byte{0xa8, 0xc8, 0x74, 0xa9, 0x7a, 0xe7, 0xc4, 0xf8}

	ks := make([]byte, 8)
	s := New(k, iv)
	s.Read(ks)
	if !bytes.Equal(ks, z) {
		t.Errorf("want %x; but got %x\n", z, ks)
	}
}

func Test4(t *testing.T) {
	k := []uint32{0x0ded7263, 0x109cf92e, 0x3352255a, 0x140e0f76}
	iv := []uint32{0x6b68079a, 0x41a7c4c9, 0x1befd79f, 0x7fdcc233}
	z := []byte{
		0xd7, 0x12, 0xc0, 0x5c, 0xa9, 0x37, 0xc2, 0xa6,
		0xeb, 0x7e, 0xaa, 0xe3,
	}
	last := []byte{0x9c, 0x0d, 0xb3, 0xaa}

	ks := make([]byte, 10000)
	s := New(k, iv)
	s.Read(ks)
	if !bytes.Equal(ks[:12], z) {
		t.Errorf("want %x; but got %x\n", z, ks[:12])
	}
	if !bytes.Equal(ks[10000-4:10000], last) {
		t.Errorf("want %x; but got %x\n", last, ks[10000-4:10000])
	}
}
