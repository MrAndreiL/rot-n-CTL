package rotn

// NRotation receives the incoming byte array and returns 
// the respective N-rot encoding for each symbol in the array.
func NRotation(buffer []byte, N int) []byte {
	for i, ch := range buffer {
		if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' {
			if N > 26 {
				N = (N % 26)
			}
			rot_ch := symbolRotation(ch, N)
			buffer[i] = rot_ch
		}
	}
	return buffer
}

// symbolRotation receives a single 'a' -> 'z' or 'A' -> 'Z'
// character and returs the N-rotated form.
func symbolRotation(ch byte, N int) byte {
	if N == 0 || N == 26 { // character remains in place
		return ch
	}
	// if wrapping occurrs 
	if ('A' <= ch && ch <= 'Z' && int(ch) + N > 90) {
		ch = byte(65 + (N - (90 - int(ch))) - 1)
		return ch
	}
	if ('a' <= ch && ch <= 'z' && int(ch) + N > 122) {
		ch = byte(97 + (N - (122 - int(ch))) - 1)
		return ch
	}
	return byte(int(ch) + N)
}
