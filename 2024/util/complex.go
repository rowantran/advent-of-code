package util

func Unpack(c complex64) (int, int) {
	return int(real(c)), int(imag(c))
}

func Pack(a int, b int) complex64 {
	return complex(float32(a), float32(b))
}
