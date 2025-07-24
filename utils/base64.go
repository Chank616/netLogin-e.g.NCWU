package utils

// 更改了码表的base64
const (
	base64Table = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"
)

func Base64Encode(data []byte) string {
	result := make([]byte, 0, len(data)*4/3+3)

	for i := 0; i < len(data); i += 3 {
		chunk := data[i:]
		if len(chunk) > 3 {
			chunk = chunk[:3]
		}

		b1, b2, b3 := chunk[0], byte(0), byte(0)
		if len(chunk) > 1 {
			b2 = chunk[1]
		}
		if len(chunk) > 2 {
			b3 = chunk[2]
		}

		result = append(result, base64Table[b1>>2])
		result = append(result, base64Table[((b1&0x03)<<4)|(b2>>4)])
		if len(chunk) > 1 {
			result = append(result, base64Table[((b2&0x0f)<<2)|(b3>>6)])
		} else {
			result = append(result, '=')
		}
		if len(chunk) > 2 {
			result = append(result, base64Table[b3&0x3f])
		} else {
			result = append(result, '=')
		}
	}

	return string(result)
}
