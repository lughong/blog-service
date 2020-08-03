package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) ToInt() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustToInt() int {
	v, _ := s.ToInt()

	return v
}

func (s StrTo) ToUInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())

	return uint32(v), err
}

func (s StrTo) MustToUInt32() uint32 {
	v, _ := s.ToUInt32()

	return v
}
