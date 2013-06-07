package util

import (
	"bytes"
)

type StringBuffer struct {
	buffer bytes.Buffer
}

func NewStringBuffer() *StringBuffer {
	return &StringBuffer{}
}

func (s *StringBuffer) Append(str string) {
	s.buffer.WriteString(str)
}

func (s *StringBuffer) Buffer() *bytes.Buffer {
	return &s.buffer
}

func (s StringBuffer) String() string {
	return s.buffer.String()
}
