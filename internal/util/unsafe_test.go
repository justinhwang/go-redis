package util

import (
	"testing"

	. "github.com/bsm/ginkgo/v2"
	. "github.com/bsm/gomega"
)

var (
	_tmpBytes  []byte
	_tmpString string
)

func TestBytesToString(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{
			input:  "foo bar baz",
			expect: "foo bar baz",
		},
		{
			input:  "",
			expect: "",
		},
	}

	for _, tt := range tests {
		input := []byte(tt.input)
		It("should equal", func() {
			result := BytesToString(input)
			Expect(result).To(Equal(tt.expect))
		})
		if len(tt.input) == 0 {
			continue
		}
		It("should equal", func() {
			inputModified := []byte(tt.input)
			inputModified[0] = 'x'
			result := BytesToString(inputModified)
			Expect(result).ToNot(Equal(tt.expect))
		})
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		input  string
		expect []byte
	}{
		{
			input:  "foo bar baz",
			expect: []byte("foo bar baz"),
		},
		{
			input:  "",
			expect: nil,
		},
	}

	It("should equal", func() {
		for _, tt := range tests {
			result := StringToBytes(tt.input)
			Expect(result).To(Equal(tt.expect))
		}
	})
}

func BenchmarkStringToBytes(b *testing.B) {
	input := b.Name()

	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_tmpBytes = []byte(input)
		}
	})

	b.Run("unsafe", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_tmpBytes = StringToBytes(input)
		}
	})
}

func BenchmarkBytesToString(b *testing.B) {
	input := []byte(b.Name())

	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_tmpString = string(input)
		}
	})

	b.Run("unsafe", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_tmpString = BytesToString(input)
		}
	})
}
