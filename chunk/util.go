package chunk

func allocBytes(in []byte) []byte {
	buf := make([]byte, len(in))
	copy(buf, in)
	return buf
}
