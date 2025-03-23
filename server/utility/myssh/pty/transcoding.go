package pty

import "unicode/utf8"

func IsGBK(s []byte) bool {
	if utf8.ValidString(string(s)) {
		return false
	}
	data := s
	length := len(data)
	var i = 0
	for i < length {
		//if data[i] >= 0x81 &&
		//	data[i] <= 0xfe &&
		//	data[i+1] >= 0x40 &&
		//	data[i+1] <= 0xfe &&
		//	data[i+1] != 0xf7 {
		//	i += 2
		//	continue
		//} else {
		//	return false
		//}

		if data[i] >= 129 &&
			data[i] <= 254 &&
			data[i+1] >= 64 &&
			data[i+1] <= 254 &&
			data[i+1] <= 247 {
			i += 2
			return true
		}
	}
	return false
}
