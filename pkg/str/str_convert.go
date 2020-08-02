package str

import "strconv"

func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func MustToUint8(s string) uint8 {
	n, _ := ToInt(s)

	return uint8(n)
}

func ToInt32(s string) (int32, error) {
	n, err := ToInt(s)

	return int32(n), err
}

func MustToUint32(s string) uint32 {
	n, _ := ToInt32(s)

	return uint32(n)
}
