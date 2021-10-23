package harpocrates

import "errors"

var (
	ErrZeroLength error = errors.New("harpocrates: expecting positive input length, receiving zero")
)

func ResizeByteArray(input []byte, toSize int) ([]byte, error) {
	if len(input) == 0 || toSize <= 0 {
		return nil, ErrZeroLength
	}

	var out []byte

	for len(out) < toSize {
		out = append(out, input...)
	}

	if len(out) > toSize {
		out = out[:toSize]
	}

	return out, nil
}
