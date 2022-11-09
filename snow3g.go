package snow3g

type State struct {
	R [3]uint32
	S [16]uint32
}

func New(k, iv []uint32) *State {
	s := new(State)
	s.S[0] = ^k[0]
	s.S[1] = ^k[1]
	s.S[2] = ^k[2]
	s.S[3] = ^k[3]
	s.S[4] = k[0]
	s.S[5] = k[1]
	s.S[6] = k[2]
	s.S[7] = k[3]
	s.S[8] = ^k[0]
	s.S[9] = ^k[1] ^ iv[3]
	s.S[10] = ^k[2] ^ iv[2]
	s.S[11] = ^k[3]
	s.S[12] = k[0] ^ iv[1]
	s.S[13] = k[1]
	s.S[14] = k[2]
	s.S[15] = k[3] ^ iv[0]
	for i := 0; i < 32; i++ {
		F := s.ClockFSM()
		s.InitMode(F)
	}
	return s
}

func (s *State) Read(b []byte) (int, error) {
	s.ClockFSM()
	s.KeystreamMode()
	n := len(b)
	for i := 0; i < n; i += 4 {
		w := s.ClockFSM() ^ s.S[0]
		b[i] = byte(w >> 24)
		if i+1 < n {
			b[i+1] = byte(w >> 16)
		}
		if i+2 < n {
			b[i+2] = byte(w >> 8)
		}
		if i+3 < n {
			b[i+3] = byte(w)
		}
		s.KeystreamMode()
	}
	return n, nil
}

func (s *State) ClockFSM() uint32 {
	F := (s.S[15] + s.R[0]) ^ s.R[1]
	r := s.R[1] + (s.R[2] ^ s.S[5])
	s.R[2] = S2(s.R[1])
	s.R[1] = S1(s.R[0])
	s.R[0] = r
	return F
}

func (s *State) InitMode(F uint32) {
	v := s.S[0] << 8
	v ^= MulAlpha(byte(s.S[0] >> 24))
	v ^= s.S[2]
	v ^= s.S[11] >> 8
	v ^= DivAlpha(byte(s.S[11]))
	v ^= F
	for i := 0; i < 15; i++ {
		s.S[i] = s.S[i+1]
	}
	s.S[15] = v
}

func (s *State) KeystreamMode() {
	v := s.S[0] << 8
	v ^= MulAlpha(byte(s.S[0] >> 24))
	v ^= s.S[2]
	v ^= s.S[11] >> 8
	v ^= DivAlpha(byte(s.S[11]))
	for i := 0; i < 15; i++ {
		s.S[i] = s.S[i+1]
	}
	s.S[15] = v
}

func S1(w uint32) uint32 {
	w0 := byte(w >> 24)
	w1 := byte(w >> 16)
	w2 := byte(w >> 8)
	w3 := byte(w)
	return S1T0[w3] ^ S1T1[w2] ^ S1T2[w1] ^ S1T3[w0]
}

func S2(w uint32) uint32 {
	w0 := byte(w >> 24)
	w1 := byte(w >> 16)
	w2 := byte(w >> 8)
	w3 := byte(w)
	return S2T0[w3] ^ S2T1[w2] ^ S2T2[w1] ^ S2T3[w0]
}

func MulAlpha(c byte) uint32 {
	return MulA[c]
}

func DivAlpha(c byte) uint32 {
	return DivA[c]
}
