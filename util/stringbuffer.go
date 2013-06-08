package util

import (
	"bytes"
)

type StringBuffer struct {
	buffer bytes.Buffer
}

func NewStringBuffer(strs ...string) *StringBuffer {
	sb := &StringBuffer{}
	for _, str := range strs {
		sb.Append(str)
	}
	return sb
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
