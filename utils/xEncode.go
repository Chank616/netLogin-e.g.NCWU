package utils

func ordat(msg string, i int) uint {
	if len(msg) > i {
		return uint(msg[i])
	}
	return 0
}

func sencode(msg string, key bool) []uint {
	l := len(msg)
	var pwd []uint
	for i := 0; i < l; i += 4 {
		pwd = append(pwd, ordat(msg, i)|ordat(msg, i+1)<<8|ordat(msg, i+2)<<16|ordat(msg, i+3)<<24)
	}
	if key {
		pwd = append(pwd, uint(l))
	}
	return pwd
}

func GetXencode(msg string, key string) []byte {
	pwd := sencode(msg, true)
	pwdk := sencode(key, false)
	n := uint(len(pwd) - 1)
	z := uint(pwd[n])
	var c uint = 0x9E3779B9
	var d uint = 0
	q := int(6 + 52/(n+1))
	for ; q > 0; q-- {
		d = d + c&0xffffffff
		e := d >> 2 & 3
		var p uint = 0
		for ; p < n; p++ {
			y := uint(pwd[p+1])
			m := z>>5 ^ y<<2
			m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
			m = m + (pwdk[(p&3)^e] ^ z)
			pwd[p] = pwd[p] + m&0xffffffff
			z = pwd[p]
		}
		y := pwd[0]
		m := z>>5 ^ y<<2
		m = m + ((y>>3 ^ z<<4) ^ (d ^ y))
		m = m + (pwdk[(p&3)^e] ^ z)
		pwd[n] = pwd[n] + m&0xffffffff
		z = pwd[n]
	}
	var bytes []byte
	for i := 0; i < len(pwd); i++ {
		bytes = append(bytes, uint8(pwd[i]&0xff))
		bytes = append(bytes, uint8(pwd[i]>>8&0xff))
		bytes = append(bytes, uint8(pwd[i]>>16&0xff))
		bytes = append(bytes, uint8(pwd[i]>>24&0xff))
	}
	return bytes
}
