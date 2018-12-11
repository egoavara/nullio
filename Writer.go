package nullio

import "io"

type Writer struct {
}

func (*Writer) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func NewWriter() io.Writer {
	return new(Writer)
}
