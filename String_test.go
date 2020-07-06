package String

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
	"unicode"
)

func Test_String_substr(t *testing.T) {
	type args struct {
		start int
		end   int
		str   String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				start: 3,
				end:   5,
				str:   "Hello world!",
			},
			want: "lo wo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.substr(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("String.substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Substr(t *testing.T) {
	type args struct {
		start int
		end   int
		str   String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				start: 3,
				end:   5,
				str:   "Hello world!",
			},
			want: "lo wo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.Substr(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("String.Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_substr_Object(b *testing.B) {
	var s String = "Hello World!"
	for n := 0; n < b.N; n++ {
		s.substr(3, 5)
	}
}

func Test_String_ASCIIsubstr(t *testing.T) {
	type args struct {
		start int
		end   int
		str   String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				start: 3,
				end:   5,
				str:   "Hello world!",
			},
			want: "lo wo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.substr(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("String.ASCIIsubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ASCIIsubstr_Object(b *testing.B) {
	var s String = "Hello World!"
	for n := 0; n < b.N; n++ {
		s.ASCIIsubstr(3, 5)
	}
}

func Test_String_Contains(t *testing.T) {
	type args struct {
		substr String
		str    String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_true",
			args: args{
				substr: "NEEDLEÖ",
				str:    "There is a NEEDLEÖ in the haystack!",
			},
			want: true,
		},
		{
			name: "success_false",
			args: args{
				substr: "PIE",
				str:    "There is a NEEDLEÖ in the haystack!",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.Contains(tt.args.substr); got != tt.want {
				t.Errorf("String.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Contains_Object(b *testing.B) {
	var s String = "There is a NEEDLE in the haystack!"
	for n := 0; n < b.N; n++ {
		s.Contains("NEEDLE")
	}
}

func Benchmark_Contains(b *testing.B) {
	s := "There is a NEEDLE in the haystack!"
	for n := 0; n < b.N; n++ {
		strings.Contains(s, "NEEDLE")
	}
}

func Test_String_Compare(t *testing.T) {
	type args struct {
		str String
		s   String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success_0",
			args: args{
				str: "bbb",
				s:   "bbb",
			},
			want: 0,
		},
		{
			name: "success_-1",
			args: args{
				str: "aaa",
				s:   "bbb",
			},
			want: 1,
		},
		{
			name: "success_1",
			args: args{
				str: "ccc",
				s:   "bbb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Compare(tt.args.str); got != tt.want {
				t.Errorf("String.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompareObject(b *testing.B) {
	var s String = "aaa"
	var ss String = "bbb"
	for n := 0; n < b.N; n++ {
		s.Compare(ss)
	}
}

func BenchmarkCompare(b *testing.B) {
	s := "aaa"
	ss := "bbb"
	for n := 0; n < b.N; n++ {
		strings.Compare(s, ss)
	}
}

func Test_String_ContainsAny(t *testing.T) {
	type args struct {
		chars String
		s     String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_true",
			args: args{
				chars: "N",
				s:     "NEEDLE",
			},
			want: true,
		},
		{
			name: "success_false",
			args: args{
				chars: "P",
				s:     "NEEDLE",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Contains(tt.args.chars); got != tt.want {
				t.Errorf("String.ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsAnyObject(b *testing.B) {
	var s String = "NEEDLE"
	for n := 0; n < b.N; n++ {
		s.ContainsAny("N")
	}
}

func BenchmarkContainsAny(b *testing.B) {
	s := "NEEDLE"
	for n := 0; n < b.N; n++ {
		strings.ContainsAny(s, "N")
	}
}

func Test_String_ContainsRune(t *testing.T) {
	type args struct {
		r   rune
		str String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_true",
			args: args{
				r:   97,
				str: "aaaaaaaa",
			},
			want: true,
		},
		{
			name: "success_false",
			args: args{
				r:   98,
				str: "aaaaaaaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.ContainsRune(tt.args.r); got != tt.want {
				t.Errorf("String.ContainsRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsRuneObject(b *testing.B) {
	var s String = "aaaaaaaa"
	for n := 0; n < b.N; n++ {
		s.ContainsRune(97)
	}
}

func BenchmarkContainsRune(b *testing.B) {
	s := "aaaaaaaa"
	for n := 0; n < b.N; n++ {
		strings.ContainsRune(s, 97)
	}
}

func Test_String_Count(t *testing.T) {
	type args struct {
		substr String
		s      String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				substr: "a",
				s:      "aaaaaaaaaaaaa",
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Count(tt.args.substr); got != tt.want {
				t.Errorf("String.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountObject(b *testing.B) {
	var s String = "aaaaaaaaaaaaa"
	for n := 0; n < b.N; n++ {
		s.Count("a")
	}
}

func BenchmarkCount(b *testing.B) {
	s := "aaaaaaaaaaaaa"
	for n := 0; n < b.N; n++ {
		strings.Count(s, "a")
	}
}

func Test_String_EqualFold(t *testing.T) {
	type args struct {
		substr String
		s      String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_false",
			args: args{
				substr: "bbbbbbbbbbb",
				s:      "aaaaaaaaaaaaa",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				substr: "AAAAAAAAAAAAA",
				s:      "aaaaaaaaaaaaa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.EqualFold(tt.args.substr); got != tt.want {
				t.Errorf("String.EqualFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkEqualFoldObject(b *testing.B) {
	var s String = "aaaaaaaaaaaaa"
	for n := 0; n < b.N; n++ {
		s.EqualFold("AAAAAAAAAAAAA")
	}
}

func BenchmarkEqualFold(b *testing.B) {
	s := "aaaaaaaaaaaaa"
	for n := 0; n < b.N; n++ {
		strings.EqualFold(s, "AAAAAAAAAAAAA")
	}
}

func Test_String_Fields(t *testing.T) {
	type args struct {
		str String
	}
	tests := []struct {
		name string
		args args
		want Strings
	}{
		{
			name: "success",
			args: args{
				str: "this is a string with some whitespaces",
			},
			want: Strings{"this", "is", "a", "string", "with", "some", "whitespaces"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s String = tt.args.str
			if got := s.Fields(); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFieldsObject(b *testing.B) {
	var s String = "this is a string with some whitespaces"
	for n := 0; n < b.N; n++ {
		s.Fields()
	}
}

func BenchmarkFields(b *testing.B) {
	s := "this is a string with some whitespaces"
	for n := 0; n < b.N; n++ {
		strings.Fields(s)
	}
}

//FieldsFunc() is tested with Fields() => Perhaps write an own test sometime

func Test_String_len(t *testing.T) {
	type args struct {
		str String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success_true",
			args: args{
				str: "123456789",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.len(); got != tt.want {
				t.Errorf("String.len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLenObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.len()
	}
}

func BenchmarkLen(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		_ = len(s)
	}
}

func Test_String_HasPrefix(t *testing.T) {
	type args struct {
		s      String
		prefix String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_true",
			args: args{
				s:      "123456789",
				prefix: "123",
			},
			want: true,
		},
		{
			name: "success_false",
			args: args{
				s:      "123456789",
				prefix: "234",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.HasPrefix(tt.args.prefix); got != tt.want {
				t.Errorf("String.HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHasPrefixObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.HasPrefix("123")
	}
}

func BenchmarkHasPrefix(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		strings.HasPrefix(s, "123")
	}
}

func Test_String_length(t *testing.T) {
	type args struct {
		str String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				str: "123456789",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.str.length(); got != tt.want {
				t.Errorf("String.length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLengthObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.length()
	}
}

func BenchmarkLength(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		_ = len(s)
	}
}

func Test_String_HasSuffix(t *testing.T) {
	type args struct {
		s      String
		suffix String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_true",
			args: args{
				s:      "123456789",
				suffix: "789",
			},
			want: true,
		},
		{
			name: "success_false",
			args: args{
				s:      "123456789",
				suffix: "234",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.HasSuffix(tt.args.suffix); got != tt.want {
				t.Errorf("String.HasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHasSuffixObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.HasSuffix("789")
	}
}

func BenchmarkHasSuffix(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		strings.HasSuffix(s, "789")
	}
}

func Test_String_Index(t *testing.T) {
	type args struct {
		s      String
		substr String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:      "123456789",
				substr: "5",
			},
			want: 4,
		},
		{
			name: "fail",
			args: args{
				s:      "123456789",
				substr: "a",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Index(tt.args.substr); got != tt.want {
				t.Errorf("String.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIndexObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.Index("5")
	}
}

func BenchmarkIndex(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		strings.Index(s, "5")
	}
}

func Test_String_IndexAny(t *testing.T) {
	type args struct {
		s      String
		substr String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:      "123456789",
				substr: "58",
			},
			want: 4,
		},
		{
			name: "fail",
			args: args{
				s:      "123456789",
				substr: "ab",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IndexAny(tt.args.substr); got != tt.want {
				t.Errorf("String.IndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIndexAnyObject(b *testing.B) {
	var s String = "123456789"
	for n := 0; n < b.N; n++ {
		s.IndexAny("58")
	}
}

func BenchmarkIndexAny(b *testing.B) {
	s := "123456789"
	for n := 0; n < b.N; n++ {
		strings.IndexAny(s, "58")
	}
}

func Test_String_IndexByte(t *testing.T) {
	type args struct {
		s String
		c byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "12345678a",
				c: 97,
			},
			want: 8,
		},
		{
			name: "fail",
			args: args{
				s: "123456789",
				c: 97,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IndexByte(tt.args.c); got != tt.want {
				t.Errorf("String.IndexByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIndexByteObject(b *testing.B) {
	var s String = "12345678a"
	for n := 0; n < b.N; n++ {
		s.IndexByte(97)
	}
}

func BenchmarkIndexByte(b *testing.B) {
	s := "12345678a"
	for n := 0; n < b.N; n++ {
		strings.IndexByte(s, 97)
	}
}

func Test_String_IndexFunc(t *testing.T) {
	type args struct {
		s String
		f func(rune) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "Hello, 世界",
				f: func(c rune) bool {
					return unicode.Is(unicode.Han, c)
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IndexFunc(tt.args.f); got != tt.want {
				t.Errorf("String.IndexFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIndexFuncObject(b *testing.B) {
	var s String = "Hello, 世界"
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	for n := 0; n < b.N; n++ {
		s.IndexFunc(f)
	}
}

func BenchmarkIndexFunc(b *testing.B) {
	s := "Hello, 世界"
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	for n := 0; n < b.N; n++ {
		strings.IndexFunc(s, f)
	}
}

func Test_String_IndexRune(t *testing.T) {
	type args struct {
		s String
		r rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "chicken",
				r: 'k',
			},
			want: 4,
		},
		{
			name: "success_false",
			args: args{
				s: "chicken",
				r: 'd',
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IndexRune(tt.args.r); got != tt.want {
				t.Errorf("String.IndexRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIndexRuneObject(b *testing.B) {
	var s String = "chicken"
	r := 'k'
	for n := 0; n < b.N; n++ {
		s.IndexRune(r)
	}
}

func BenchmarkIndexRune(b *testing.B) {
	s := "chicken"
	r := 'd'
	for n := 0; n < b.N; n++ {
		strings.IndexRune(s, r)
	}
}

func Test_String_Join(t *testing.T) {
	type args struct {
		s   Strings
		sep String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:   Strings{"foo", "bar", "baz"},
				sep: ", ",
			},
			want: "foo, bar, baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Join(tt.args.sep); got != tt.want {
				t.Errorf("String.IndexJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJoinObject(b *testing.B) {
	var s = Strings{"foo", "bar", "baz"}
	var sep String = ", "
	for n := 0; n < b.N; n++ {
		s.Join(sep)
	}
}

func BenchmarkJoin(b *testing.B) {
	s := []string{"foo", "bar", "baz"}
	sep := ", "
	for n := 0; n < b.N; n++ {
		strings.Join(s, sep)
	}
}

func Test_String_LastIndex(t *testing.T) {
	type args struct {
		s      String
		substr String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:      "go gopher",
				substr: "go",
			},
			want: 3,
		},
		{
			name: "success_false",
			args: args{
				s:      "go gopher",
				substr: "rodent",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.LastIndex(tt.args.substr); got != tt.want {
				t.Errorf("String.LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLastIndexObject(b *testing.B) {
	var s String = "go gopher"
	var substr String = "go"
	for n := 0; n < b.N; n++ {
		s.LastIndex(substr)
	}
}

func BenchmarkLastIndex(b *testing.B) {
	s := "go gopher"
	substr := "go"
	for n := 0; n < b.N; n++ {
		strings.LastIndex(s, substr)
	}
}

func Test_String_LastIndexAny(t *testing.T) {
	type args struct {
		s     String
		chars String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:     "go gopher",
				chars: "go",
			},
			want: 4,
		},
		{
			name: "success_false",
			args: args{
				s:     "go gopher",
				chars: "fail",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.LastIndexAny(tt.args.chars); got != tt.want {
				t.Errorf("String.LastIndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLastIndexAnyObject(b *testing.B) {
	var s String = "go gopher"
	var chars String = "go"
	for n := 0; n < b.N; n++ {
		s.LastIndexAny(chars)
	}
}

func BenchmarkLastIndexAny(b *testing.B) {
	s := "go gopher"
	chars := "go"
	for n := 0; n < b.N; n++ {
		strings.LastIndexAny(s, chars)
	}
}

func Test_String_LastIndexByte(t *testing.T) {
	type args struct {
		s    String
		char byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:    "Hello, world",
				char: 'o',
			},
			want: 8,
		},
		{
			name: "success_false",
			args: args{
				s:    "go gopher",
				char: 'x',
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.LastIndexByte(tt.args.char); got != tt.want {
				t.Errorf("String.LastIndexByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLastIndexByteObject(b *testing.B) {
	var s String = "Hello, world"
	var char byte = 'o'
	for n := 0; n < b.N; n++ {
		s.LastIndexByte(char)
	}
}

func BenchmarkLastIndexByte(b *testing.B) {
	s := "Hello, world"
	var char byte = 'o'
	for n := 0; n < b.N; n++ {
		strings.LastIndexByte(s, char)
	}
}

func Test_String_LastIndexFunc(t *testing.T) {
	type args struct {
		s String
		f func(rune) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "go 123",
				f: unicode.IsNumber,
			},
			want: 5,
		},
		{
			name: "success",
			args: args{
				s: "go adsde",
				f: unicode.IsNumber,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.LastIndexFunc(tt.args.f); got != tt.want {
				t.Errorf("String.LastIndexFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLastIndexFuncObject(b *testing.B) {
	var s String = "go 123"
	f := unicode.IsNumber
	for n := 0; n < b.N; n++ {
		s.LastIndexFunc(f)
	}
}

func BenchmarkLastIndexFunc(b *testing.B) {
	s := "go 123"
	f := unicode.IsNumber
	for n := 0; n < b.N; n++ {
		strings.LastIndexFunc(s, f)
	}
}

func Test_String_Map(t *testing.T) {
	type args struct {
		s       String
		mapping func(r rune) rune
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "'Twas brillig and the slithy gopher...",
				mapping: func(r rune) rune {
					switch {
					case r >= 'A' && r <= 'Z':
						return 'A' + (r-'A'+13)%26
					case r >= 'a' && r <= 'z':
						return 'a' + (r-'a'+13)%26
					}
					return r
				},
			},
			want: "'Gjnf oevyyvt naq gur fyvgul tbcure...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Map(tt.args.mapping); got != tt.want {
				t.Errorf("String.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMapObject(b *testing.B) {
	var s String = "'Twas brillig and the slithy gopher..."
	f := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	for n := 0; n < b.N; n++ {
		s.Map(f)
	}
}

func BenchmarkMap(b *testing.B) {
	s := "'Twas brillig and the slithy gopher..."
	f := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	for n := 0; n < b.N; n++ {
		strings.Map(f, s)
	}
}

func Test_String_Repeat(t *testing.T) {
	type args struct {
		s     String
		count int
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:     "Meh, ",
				count: 5,
			},
			want: "Meh, Meh, Meh, Meh, Meh, ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Repeat(tt.args.count); got != tt.want {
				t.Errorf("String.Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRepeatObject(b *testing.B) {
	var s String = "Meh, "
	count := 5
	for n := 0; n < b.N; n++ {
		s.Repeat(count)
	}
}

func BenchmarkRepeat(b *testing.B) {
	s := "Meh, "
	count := 5
	for n := 0; n < b.N; n++ {
		strings.Repeat(s, count)
	}
}

func Test_String_Replace(t *testing.T) {
	type args struct {
		s   String
		old String
		new String
		n   int
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:   "oink oink oink",
				old: "oink",
				new: "moo",
				n:   2,
			},
			want: "moo moo oink",
		},
		{
			name: "success_all",
			args: args{
				s:   "oink oink oink",
				old: "oink",
				new: "moo",
				n:   -1,
			},
			want: "moo moo moo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Replace(tt.args.old, tt.args.new, tt.args.n); got != tt.want {
				t.Errorf("String.Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceObject(b *testing.B) {
	var s String = "oink oink oink"
	var old String = "oink"
	var new String = "moo"
	nn := 2
	for n := 0; n < b.N; n++ {
		s.Replace(old, new, nn)
	}
}

func BenchmarkReplace(b *testing.B) {
	s := "oink oink oink"
	old := "oink"
	new := "moo"
	nn := 2
	for n := 0; n < b.N; n++ {
		strings.Replace(s, old, new, nn)
	}
}

func Test_String_ReplaceAll(t *testing.T) {
	type args struct {
		s   String
		old String
		new String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:   "oink oink oink",
				old: "oink",
				new: "moo",
			},
			want: "moo moo moo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ReplaceAll(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("String.ReplaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceAllObject(b *testing.B) {
	var s String = "oink oink oink"
	var old String = "oink"
	var new String = "moo"
	for n := 0; n < b.N; n++ {
		s.ReplaceAll(old, new)
	}
}

func BenchmarkReplaceAll(b *testing.B) {
	s := "oink oink oink"
	old := "oink"
	new := "moo"
	for n := 0; n < b.N; n++ {
		strings.ReplaceAll(s, old, new)
	}
}

func Test_String_Split(t *testing.T) {
	type args struct {
		s   String
		sep String
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
		{
			name: "success",
			args: args{
				s:   "moo moo moo",
				sep: " ",
			},
			want: []String{"moo", "moo", "moo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Split(tt.args.sep); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplitObject(b *testing.B) {
	var s String = "moo moo moo"
	var sep String = " "
	for n := 0; n < b.N; n++ {
		s.Split(sep)
	}
}

func BenchmarkSplit(b *testing.B) {
	s := "moo moo moo"
	sep := " "
	for n := 0; n < b.N; n++ {
		strings.Split(s, sep)
	}
}

func Test_String_SplitAfter(t *testing.T) {
	type args struct {
		s   String
		sep String
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
		{
			name: "success",
			args: args{
				s:   "moo moo moo",
				sep: " ",
			},
			want: []String{"moo ", "moo ", "moo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.SplitAfter(tt.args.sep); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.SplitAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplitAfterObject(b *testing.B) {
	var s String = "moo moo moo"
	var sep String = " "
	for n := 0; n < b.N; n++ {
		s.SplitAfter(sep)
	}
}

func BenchmarkSplitAfter(b *testing.B) {
	s := "moo moo moo"
	sep := " "
	for n := 0; n < b.N; n++ {
		strings.SplitAfter(s, sep)
	}
}

func Test_String_SplitAfterN(t *testing.T) {
	type args struct {
		s   String
		sep String
		n   int
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
		{
			name: "success",
			args: args{
				s:   "moo moo moo",
				sep: " ",
				n:   2,
			},
			want: []String{"moo ", "moo moo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.SplitAfterN(tt.args.sep, tt.args.n); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.SplitAfterN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplitAfterNObject(b *testing.B) {
	var s String = "moo moo moo"
	var sep String = " "
	i := 2
	for n := 0; n < b.N; n++ {
		s.SplitAfterN(sep, i)
	}
}

func BenchmarkSplitAfterN(b *testing.B) {
	s := "moo moo moo"
	sep := " "
	i := 2
	for n := 0; n < b.N; n++ {
		strings.SplitAfterN(s, sep, i)
	}
}

func Test_String_SplitN(t *testing.T) {
	type args struct {
		s   String
		sep String
		n   int
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
		{
			name: "success",
			args: args{
				s:   "moo moo moo",
				sep: " ",
				n:   2,
			},
			want: []String{"moo", "moo moo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.SplitN(tt.args.sep, tt.args.n); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSplitNObject(b *testing.B) {
	var s String = "moo moo moo"
	var sep String = " "
	i := 2
	for n := 0; n < b.N; n++ {
		s.SplitN(sep, i)
	}
}

func BenchmarkSplitN(b *testing.B) {
	s := "moo moo moo"
	sep := " "
	i := 2
	for n := 0; n < b.N; n++ {
		strings.SplitN(s, sep, i)
	}
}

func Test_String_Title(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "her royal highness",
			},
			want: "Her Royal Highness",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Title(); got != tt.want {
				t.Errorf("String.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTitleObject(b *testing.B) {
	var s String = "her royal highness"
	for n := 0; n < b.N; n++ {
		s.Title()
	}
}

func BenchmarkTitle(b *testing.B) {
	s := "her royal highness"
	for n := 0; n < b.N; n++ {
		strings.Title(s)
	}
}

func Test_String_ToLower(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "HER ROYAL HIGHNESS",
			},
			want: "her royal highness",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToLower(); got != tt.want {
				t.Errorf("String.ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToLowerObject(b *testing.B) {
	var s String = "HER ROYAL HIGHNESS"
	for n := 0; n < b.N; n++ {
		s.ToLower()
	}
}

func BenchmarkToLower(b *testing.B) {
	s := "HER ROYAL HIGHNESS"
	for n := 0; n < b.N; n++ {
		strings.ToLower(s)
	}
}

func Test_String_ToLowerSpecial(t *testing.T) {
	type args struct {
		s String
		c unicode.SpecialCase
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "Önnek İş",
				c: unicode.TurkishCase,
			},
			want: "önnek iş",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToLowerSpecial(tt.args.c); got != tt.want {
				t.Errorf("String.ToLowerSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToLowerSpecialObject(b *testing.B) {
	var s String = "Önnek İş"
	for n := 0; n < b.N; n++ {
		s.ToLowerSpecial(unicode.TurkishCase)
	}
}

func BenchmarkToLowerSpecial(b *testing.B) {
	s := "Önnek İş"
	for n := 0; n < b.N; n++ {
		strings.ToLowerSpecial(unicode.TurkishCase, s)
	}
}

func Test_String_ToTitle(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "her royal highness",
			},
			want: "HER ROYAL HIGHNESS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToTitle(); got != tt.want {
				t.Errorf("String.ToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToTitleObject(b *testing.B) {
	var s String = "her royal highness"
	for n := 0; n < b.N; n++ {
		s.ToTitle()
	}
}

func BenchmarkToTitle(b *testing.B) {
	s := "her royal highness"
	for n := 0; n < b.N; n++ {
		strings.ToTitle(s)
	}
}

func Test_String_ToUpper(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "her royal highness",
			},
			want: "HER ROYAL HIGHNESS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToUpper(); got != tt.want {
				t.Errorf("String.ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToUpperObject(b *testing.B) {
	var s String = "her royal highness"
	for n := 0; n < b.N; n++ {
		s.ToUpper()
	}
}

func BenchmarkUpper(b *testing.B) {
	s := "her royal highness"
	for n := 0; n < b.N; n++ {
		strings.ToUpper(s)
	}
}

func Test_String_ToTitleSpecial(t *testing.T) {
	type args struct {
		s String
		c unicode.SpecialCase
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "dünyanın ilk borsa yapısı Aizonai kabul edilir",
				c: unicode.TurkishCase,
			},
			want: "DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToTitleSpecial(tt.args.c); got != tt.want {
				t.Errorf("String.ToTitleSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToTitleSpecialObject(b *testing.B) {
	var s String = "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	for n := 0; n < b.N; n++ {
		s.ToTitleSpecial(unicode.TurkishCase)
	}
}

func BenchmarkToTitleSpecial(b *testing.B) {
	s := "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	for n := 0; n < b.N; n++ {
		strings.ToTitleSpecial(unicode.TurkishCase, s)
	}
}

func Test_String_ToUpperSpecial(t *testing.T) {
	type args struct {
		s String
		c unicode.SpecialCase
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "dünyanın ilk borsa yapısı Aizonai kabul edilir",
				c: unicode.TurkishCase,
			},
			want: "DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ToUpperSpecial(tt.args.c); got != tt.want {
				t.Errorf("String.ToUpperSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToUpperSpecialObject(b *testing.B) {
	var s String = "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	for n := 0; n < b.N; n++ {
		s.ToUpperSpecial(unicode.TurkishCase)
	}
}

func BenchmarkToUpperSpecial(b *testing.B) {
	s := "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	for n := 0; n < b.N; n++ {
		strings.ToUpperSpecial(unicode.TurkishCase, s)
	}
}

func Test_String_Trim(t *testing.T) {
	type args struct {
		s      String
		cutset String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:      "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!",
				cutset: "!+",
			},
			want: "Schifffahrtsversicherungsanstaltsgebäudekomplex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Trim(tt.args.cutset); got != tt.want {
				t.Errorf("String.Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimObject(b *testing.B) {
	var s String = "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	var cutset String = "!+"
	for n := 0; n < b.N; n++ {
		s.Trim(cutset)
	}
}

func BenchmarkTrim(b *testing.B) {
	s := "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	cutset := "!+"
	for n := 0; n < b.N; n++ {
		strings.Trim(s, cutset)
	}
}

func Test_String_TrimPrefix(t *testing.T) {
	type args struct {
		s      String
		prefix String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:      "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!",
				prefix: "!!++",
			},
			want: "Schifffahrtsversicherungsanstaltsgebäudekomplex++!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.TrimPrefix(tt.args.prefix); got != tt.want {
				t.Errorf("String.TrimPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimPrefixObject(b *testing.B) {
	var s String = "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	var prefix String = "!!++"
	for n := 0; n < b.N; n++ {
		s.TrimPrefix(prefix)
	}
}

func BenchmarkTrimPrefix(b *testing.B) {
	s := "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	prefix := "!!++"
	for n := 0; n < b.N; n++ {
		strings.TrimPrefix(s, prefix)
	}
}

func Test_String_TrimSpace(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "\t\n Hello, Gophers \n\t\r\n",
			},
			want: "Hello, Gophers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.TrimSpace(); got != tt.want {
				t.Errorf("String.TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimSpaceObject(b *testing.B) {
	var s String = "\t\n Hello, Gophers \n\t\r\n"
	for n := 0; n < b.N; n++ {
		s.TrimSpace()
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	s := "\t\n Hello, Gophers \n\t\r\n"
	for n := 0; n < b.N; n++ {
		strings.TrimSpace(s)
	}
}

func Test_String_TrimSuffix(t *testing.T) {
	type args struct {
		s      String
		suffix String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:      "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!",
				suffix: "++!!",
			},
			want: "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.TrimSuffix(tt.args.suffix); got != tt.want {
				t.Errorf("String.TrimSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTrimSuffixObject(b *testing.B) {
	var s String = "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	var suffix String = "++!!"
	for n := 0; n < b.N; n++ {
		s.TrimSuffix(suffix)
	}
}

func BenchmarkTrimSuffix(b *testing.B) {
	s := "!!++Schifffahrtsversicherungsanstaltsgebäudekomplex++!!"
	suffix := "++!!"
	for n := 0; n < b.N; n++ {
		strings.TrimSuffix(s, suffix)
	}
}

func Test_String_md5(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "The fog is getting thicker!And Leon's getting laaarger!",
			},
			want: "e2c569be17396eca2a2e3c11578123ed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Md5(); got != tt.want {
				t.Errorf("String.md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMd5Object(b *testing.B) {
	var s String = "The fog is getting thicker!And Leon's getting laaarger!"
	for n := 0; n < b.N; n++ {
		s.Md5()
	}
}

func BenchmarkMd5(b *testing.B) {
	s := "The fog is getting thicker!And Leon's getting laaarger!"
	for n := 0; n < b.N; n++ {
		h := md5.New()
		io.WriteString(h, string(s))
		_ = hex.EncodeToString(h.Sum(nil)[:])
	}
}

func Test_String_Sha1(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "This is an example sentence!",
			},
			want: "8b7d314a11b489238e9c8f07b117830b0e823a4a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Sha1(); got != tt.want {
				t.Errorf("String.Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSha1Object(b *testing.B) {
	var s String = "This is an example sentence!"
	for n := 0; n < b.N; n++ {
		s.Sha1()
	}
}

func BenchmarkSha1(b *testing.B) {
	s := "This is an example sentence!"
	for n := 0; n < b.N; n++ {
		h := sha1.New()
		io.WriteString(h, string(s))
		_ = hex.EncodeToString(h.Sum(nil)[:])
	}
}

func Test_String_IsEmail(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "b@1ln.de",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "k.h@bla.de",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "peterpanne@fffmail.xu",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "peterpannefffmail.xu",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsEmail(); got != tt.want {
				t.Errorf("String.Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsEmailObject(b *testing.B) {
	var s String = "b@1ln.de"
	for n := 0; n < b.N; n++ {
		s.IsEmail()
	}
}

func BenchmarkIsEmail(b *testing.B) {
	s := "b@1ln.de"
	var re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	for n := 0; n < b.N; n++ {
		re.MatchString(s)
	}
}

func Test_String_AesEncryptDecrypt(t *testing.T) {
	type args struct {
		s   String
		key String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_16",
			args: args{
				s:   "Be good to people and people will be good to you",
				key: "52cb693d7e8ff8fecb2d9bee9653954b",
			},
			want: "Be good to people and people will be good to you",
		},
		{
			name: "success_24",
			args: args{
				s:   "Be good to people and people will be good to you",
				key: "7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8",
			},
			want: "Be good to people and people will be good to you",
		},
		{
			name: "success_32",
			args: args{
				s:   String("Be good to people and people will be good to you"),
				key: "57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c",
			},
			want: String("Be good to people and people will be good to you"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			that := tt.args.s.AesEncrypt(tt.args.key)
			dec := that.AesDecrypt(tt.args.key)
			if dec != tt.want {
				t.Errorf("String.AesEncryptDecrypt() = (%v), want (%v)", dec, tt.want)
			}
		})
	}
}

func BenchmarkAesEncrypt16(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "52cb693d7e8ff8fecb2d9bee9653954b"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func BenchmarkAesEncrypt24(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func BenchmarkAesEncrypt32(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func Test_String_AesEncryptDecryptByte(t *testing.T) {
	type args struct {
		s   String
		key String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_16",
			args: args{
				s:   "Be good to people and people will be good to you",
				key: "52cb693d7e8ff8fecb2d9bee9653954b",
			},
			want: "Be good to people and people will be good to you",
		},
		{
			name: "success_24",
			args: args{
				s:   "Be good to people and people will be good to you",
				key: "7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8",
			},
			want: "Be good to people and people will be good to you",
		},
		{
			name: "success_32",
			args: args{
				s:   "Be good to people and people will be good to you",
				key: "57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c",
			},
			want: "Be good to people and people will be good to you",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s.AesEncryptByte(tt.args.key)).AesDecryptByte(tt.args.key); got != tt.want {
				t.Errorf("String.AesEncryptDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAesEncrypt16Byte(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "52cb693d7e8ff8fecb2d9bee9653954b"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func BenchmarkAesEncrypt24Byte(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func BenchmarkAesEncrypt32Byte(b *testing.B) {
	var s String = "Be good to people and people will be good to you"
	var key String = "57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c"
	for n := 0; n < b.N; n++ {
		s.AesEncrypt(key).AesDecrypt(key)
	}
}

func Test_String_PwUpperCase(t *testing.T) {
	type args struct {
		s      String
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:      "ASKmsnjd",
				number: 2,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "ASmsnjd",
				number: 2,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "Admsnjd",
				number: 2,
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s:      "admsnjd",
				number: 2,
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s:      "admsnjd",
				number: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.PwUpperCase(tt.args.number); got != tt.want {
				t.Errorf("String.PwUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPwUpperCaseObject(b *testing.B) {
	var s String = "ASKmsnjd"
	for n := 0; n < b.N; n++ {
		s.PwUpperCase(2)
	}
}

func Test_String_PwSpecialCase(t *testing.T) {
	type args struct {
		s      String
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:      "A!Kmsnjd",
				number: 1,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "A\"$msnjd",
				number: 2,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "Admsnjd",
				number: 0,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "$$$hjdgaf",
				number: 4,
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s:      "adm\"\"snjd",
				number: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.PwSpecialCase(tt.args.number); got != tt.want {
				t.Errorf("String.PwSpecialCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPwSpecialCaseObject(b *testing.B) {
	var s String = "ASK$$msnjd"
	for n := 0; n < b.N; n++ {
		s.PwSpecialCase(2)
	}
}

func Test_String_PwDigits(t *testing.T) {
	type args struct {
		s      String
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:      "dkashfhsg",
				number: 0,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "123jdalksjhf",
				number: 2,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "1234dlkajf",
				number: 4,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "12edjwhkhfda",
				number: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.PwDigits(tt.args.number); got != tt.want {
				t.Errorf("String.PwDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPwDigitsObject(b *testing.B) {
	var s String = "ASK$$23"
	for n := 0; n < b.N; n++ {
		s.PwDigits(2)
	}
}

func Test_String_PwLowerCase(t *testing.T) {
	type args struct {
		s      String
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:      "dkashfhsg",
				number: 0,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "123jdalksjhf",
				number: 2,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "1234dlksWER",
				number: 4,
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:      "12edJSKADH",
				number: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.PwLowerCase(tt.args.number); got != tt.want {
				t.Errorf("String.PwDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPwLowerCaseObject(b *testing.B) {
	var s String = "ASK$$asd"
	for n := 0; n < b.N; n++ {
		s.PwLowerCase(2)
	}
}

func Test_String_Get_Json(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_url",
			args: args{
				s: "http://httpbin.org/get",
			},
			want: "http://httpbin.org/get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Get().Json()["url"].(string); got != string(tt.want) {
				t.Errorf("String.Get_Json() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Open(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_url",
			args: args{
				s: "testtext",
			},
			want: "This is a Test Textfile!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Open(); got != tt.want {
				t.Errorf("String.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOpen(b *testing.B) {
	var s String = "testtext"
	for n := 0; n < b.N; n++ {
		s.Open()
	}
}

func Test_Get_Contents(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_url",
			args: args{
				s: "testtext",
			},
			want: "This is a Test Textfile!",
		},
		{
			name: "success_url",
			args: args{
				s: "http://httpbin.org/status/200",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.GetContents(); got != tt.want {
				t.Errorf("String.GetContents() = %v, want %v %v", got, tt.want, tt.args.s.substr(0, 7))
			}
		})
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Test_WriteToFile(t *testing.T) {
	type args struct {
		s    String
		path String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_url",
			args: args{
				s:    "testtext",
				path: "testWriteToFile",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = os.Remove(string(tt.args.path))
			tt.args.s.WriteToFile(tt.args.path)
			var got bool
			if fileExists(string(tt.args.path)) {
				got = true
			} else {
				got = false
			}
			if got != tt.want {
				t.Errorf("String.WriteToFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkWriteToFile(b *testing.B) {
	var s String = "testtext"
	_ = os.Remove("benchmarkWriteToFile")
	for n := 0; n < b.N; n++ {
		s.WriteToFile("benchmarkWriteToFile")
	}
}

func Test_Exists(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success_url",
			args: args{
				s: "http://google.de",
			},
			want: true,
		},
		{
			name: "fail_url",
			args: args{
				s: "http://akfsljdhlkshdfljshdlkfhklwhjre.de",
			},
			want: false,
		},
		{
			name: "success_file",
			args: args{
				s: "/bin/ls",
			},
			want: true,
		},
		{
			name: "fail_file",
			args: args{
				s: "/bin/sahdfuhoiuwhiuhf",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Exists(); got != tt.want {
				t.Errorf("String.Exists() = %v, want %v %v", got, tt.want, tt.args.s.substr(0, 7))
			}
		})
	}
}

func BenchmarkExists(b *testing.B) {
	var s String = "/bin/ls"
	for n := 0; n < b.N; n++ {
		s.Exists()
	}
}

func Test_URLEncode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_url",
			args: args{
				s: "uzgdauzgduaszd$&%$&%$&%$",
			},
			want: "uzgdauzgduaszd%24%26%25%24%26%25%24%26%25%24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.URLEncode(); got != tt.want {
				t.Errorf("String.URLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkURLEncode(b *testing.B) {
	var s String = "uzgdauzgduaszd$&%$&%$&%$"
	for n := 0; n < b.N; n++ {
		s.URLEncode()
	}
}

func Test_URLDecode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success_url",
			args: args{
				s: "uzgdauzgduaszd%24%26%25%24%26%25%24%26%25%24",
			},
			want: "uzgdauzgduaszd$&%$&%$&%$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.URLDecode(); got != tt.want {
				t.Errorf("String.URLDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkURLDecode(b *testing.B) {
	var s String = "uzgdauzgduaszd%24%26%25%24%26%25%24%26%25%24"
	for n := 0; n < b.N; n++ {
		s.URLDecode()
	}
}

func Test_B64Encode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "3869876asduhfjahsdu92)(/)(/",
			},
			want: "Mzg2OTg3NmFzZHVoZmphaHNkdTkyKSgvKSgv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.B64Encode(); got != tt.want {
				t.Errorf("String.B64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkB64Encode(b *testing.B) {
	var s String = "3869876asduhfjahsdu92)(/)(/"
	for n := 0; n < b.N; n++ {
		s.B64Encode()
	}
}

func Test_B64Decode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "Mzg2OTg3NmFzZHVoZmphaHNkdTkyKSgvKSgv",
			},
			want: "3869876asduhfjahsdu92)(/)(/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.B64Decode(); got != tt.want {
				t.Errorf("String.B64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkB64Decode(b *testing.B) {
	var s String = "Mzg2OTg3NmFzZHVoZmphaHNkdTkyKSgvKSgv"
	for n := 0; n < b.N; n++ {
		s.B64Decode()
	}
}

func Test_B64URLEncode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "hello world12345!?$*&()'-@~",
			},
			want: "aGVsbG8gd29ybGQxMjM0NSE_JComKCknLUB-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.B64URLEncode(); got != tt.want {
				t.Errorf("String.B64URLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkB64URLEncode(b *testing.B) {
	var s String = "hello world12345!?$*&()'-@~"
	for n := 0; n < b.N; n++ {
		s.B64URLEncode()
	}
}

func Test_B64URLDecode(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "aGVsbG8gd29ybGQxMjM0NSE_JComKCknLUB-",
			},
			want: "hello world12345!?$*&()'-@~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.B64URLDecode(); got != tt.want {
				t.Errorf("String.B64URLDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkB64URLDecode(b *testing.B) {
	var s String = "aGVsbG8gd29ybGQxMjM0NSE_JComKCknLUB-"
	for n := 0; n < b.N; n++ {
		s.B64URLDecode()
	}
}

func Test_String_Post_Json(t *testing.T) {
	type args struct {
		s           String
		url         String
		contenttype String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:           "{\"name\":\"K\"}",
				url:         "http://httpbin.org/post",
				contenttype: "application/json",
			},
			want: "K",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Post(tt.args.url, tt.args.contenttype).Json()["json"].(map[string]interface{})["name"].(string); got != string(tt.want) {
				t.Errorf("String.Post_Json() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Int(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "42",
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Int(); got != tt.want {
				t.Errorf("String.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Execute(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success+0",
			args: args{
				s: "echo \"Dies +&& Jenes das\" \"bla&&\"",
			},
			want: "Dies +&& Jenes das bla&&",
		},
		{
			name: "success+1",
			args: args{
				s: "echo Dies + Jenes das bla",
			},
			want: "Dies + Jenes das bla",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.args.s.Execute(); got != tt.want {
				t.Errorf("String.Execute() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_String_Php(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "echo 'hello';",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.args.s.Php(); got != tt.want {
				t.Errorf("String.Php() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Python(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "print(\"hello\")",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.args.s.Python(); got != tt.want {
				t.Errorf("String.Python() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_Node(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "console.log(\"hello\")",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.args.s.Node(); got != tt.want {
				t.Errorf("String.Node() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_Perl(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "print \"hello\";",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.args.s.Perl(); got != tt.want {
				t.Errorf("String.Perl() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_PhpFile(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "test_php_file.php",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := os.Getwd()
			ss := String(path)
			if got, _ := (ss + "/" + tt.args.s).PhpFile(); got != tt.want {
				t.Errorf("String.PhpFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_PythonFile(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "test_python_file.py",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := os.Getwd()
			ss := String(path)
			if got, _ := (ss + "/" + tt.args.s).PythonFile(); got != tt.want {
				t.Errorf("String.PythonFile() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_NodeFile(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "test_node_file.js",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := os.Getwd()
			ss := String(path)
			if got, _ := (ss + "/" + tt.args.s).NodeFile(); got != tt.want {
				t.Errorf("String.NodeFile() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_PerlFile(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "test_perl_file.pl",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := os.Getwd()
			ss := String(path)
			if got, _ := (ss + "/" + tt.args.s).PerlFile(); got != tt.want {
				t.Errorf("String.PerlFile() = %v, want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}

func Test_String_ParseDate(t *testing.T) {
	type args struct {
		s      String
		format String
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success",
			args: args{
				s:      "2009-11-17T20:34:58.651387237Z",
				format: "YYYY-MM-DDThh:mm:ss.999999999Z:Z",
			},
			want: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ParseDate(tt.args.format); got != tt.want {
				t.Errorf("String.ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func returnlocal(s string) *time.Location {
	loc, _ := time.LoadLocation(s)
	return loc
}

func Test_String_ParseDateLocal(t *testing.T) {
	type args struct {
		s      String
		format String
		local  String
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "success",
			args: args{
				s:      "2009-11-17T20:34:58.651387237+01:00",
				format: "YYYY-MM-DDThh:mm:ss.999999999Z:Z",
				local:  "Europe/Berlin",
			},
			want: time.Date(2009, 11, 17, 20, 34, 58, 651387237, returnlocal("Europe/Berlin")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ParseDateLocal(tt.args.format, "Europe/Berlin"); got.Sub(tt.want) != 0 {
				t.Errorf("String.ParseDateLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Bool(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "true",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Bool(); got != tt.want {
				t.Errorf("String.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Float64(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "success",
			args: args{
				s: "3.1415926535",
			},
			want: 3.1415926535,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Float64(); got != tt.want {
				t.Errorf("String.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Float32(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "success",
			args: args{
				s: "3.1415926535",
			},
			want: 3.1415927410125732,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Float32(); got != tt.want {
				t.Errorf("String.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Uint(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "success",
			args: args{
				s: "3",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Uint(); got != tt.want {
				t.Errorf("String.Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Uint32(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "success",
			args: args{
				s: "3",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Uint32(); got != tt.want {
				t.Errorf("String.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Uint64(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "success",
			args: args{
				s: "3",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Uint64(); got != tt.want {
				t.Errorf("String.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Int64(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "success",
			args: args{
				s: "3",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Int64(); got != tt.want {
				t.Errorf("String.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Int32(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "success",
			args: args{
				s: "3",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Int32(); got != tt.want {
				t.Errorf("String.Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_StripTags(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "<html><style>lestyle</style><script>this is some javascript</script><body><div>This is some text.</div><div>This is some other text.</div></body></html>",
			},
			want: "This is some text.This is some other text.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.StripTags(); got != tt.want {
				t.Errorf("String.StripTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStripTags(b *testing.B) {
	var s String = "<html><body><div>This is some text.</div><div>This is some other text.</div></body></html>"
	for n := 0; n < b.N; n++ {
		s.StripTags()
	}
}

func Test_String_Find(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "hello world",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Find("world"); got != tt.want {
				t.Errorf("String.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_String_FindAll(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				s: "hello world world wrold world",
			},
			want: []int{6, 12, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.FindAll("world"); testEq(got, tt.want) != true {
				t.Errorf("String.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsUrl(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "http://1ln.de",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "http://www.1ln.de",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "hello world world wrold world",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "das.de",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "https://google.de",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsUrl(); got != tt.want {
				t.Errorf("String.IsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_String_IsComplexPw(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "http://1ln.de",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "x!1Mkdjei93",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "1234",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "?783Mmhdj3!",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got:= tt.args.s.IsComplexPw(); got != tt.want {
				t.Errorf("String.IsComplexPw() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func Test_String_IsWholeNumber(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "1234543",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "123er323",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "1234.324",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsWholeNumber(); got != tt.want {
				t.Errorf("String.IsWholeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsIpV4(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "192.168.0.1",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "255.255.255.0",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "1234.324.1223.222",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "12.4.1.2",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "123.284.133.222",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsIpV4(); got != tt.want {
				t.Errorf("String.IsIpV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsIpV6(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:1cb9:a3a8:42e1:d9ca",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "fe80::8241:201f:2083:d021 ",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "1234.324.1223.222",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:c7f0:5df4:4db7:4e8c",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:c7f0:5df4:;;4db7:4e8c:dae3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsIpV6(); got != tt.want {
				t.Errorf("String.IsIpV6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsIp(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:1cb9:a3a8:42e1:d9ca",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "fe80::8241:201f:2083:d021 ",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "192.224.122.222",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:c7f0:5df4:4db7:4e8c",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:c7f0:5df4:4db7:4e8c:dae3",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "adsasfadfawsadcsdcsaddewaed",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsIp(); got != tt.want {
				t.Errorf("String.IsIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsHtmlTag(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "<html></html>",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "<div>blabla</div>",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "<1234.324.1223.222>",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "2a0a:a545:1728:0:c7f0:5df4:4db7:4e8c:dae3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsHtmlTag(); got != tt.want {
				t.Errorf("String.IsHtmlTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsPhoneNumber(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success1",
			args: args{
				s: "+49 (221) 8671714",
			},
			want: true,
		},
		{
			name: "success2",
			args: args{
				s: "+1 (221) 475647387",
			},
			want: true,
		},
		{
			name: "success3",
			args: args{
				s: "+49 (221) 34553445",
			},
			want: true,
		},
		{
			name: "success4",
			args: args{
				s: "8439763987gfhd",
			},
			want: false,
		},
		{
			name: "success5",
			args: args{
				s: "++843976398713213",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsPhoneNumber(); got != tt.want {
				t.Errorf("String.IsPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsFilePath(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "/usr/share/nginx",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "/var/www",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "+49 (221) 34553445",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsFilePath(); got != tt.want {
				t.Errorf("String.IsFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsUserName(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "fdkfhrqw43243",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s: "/usr/share/nginx",
			},
			want: false,
		},
		{
			name: "success",
			args: args{
				s: "+49 (221) 34553445",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsUserName(8, 20); got != tt.want {
				t.Errorf("String.IsUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsZipCode(t *testing.T) {
	type args struct {
		s       String
		country String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:       "58765",
				country: "DE",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:       "97439",
				country: "FR",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:       "587",
				country: "DE",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsZipCode(tt.args.country); got != tt.want {
				t.Errorf("String.IsZipCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_IsIban(t *testing.T) {
	type args struct {
		s       String
		country String
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s:       "DE89370400440532013000",
				country: "DE",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:       "FR1420041010050500013M02606",
				country: "FR",
			},
			want: true,
		},
		{
			name: "success",
			args: args{
				s:       "FR1420041010050500013M02606",
				country: "DE",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.IsIban(tt.args.country); got != tt.want {
				t.Errorf("String.IsIban() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Left(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "hello world",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Left(5); got != tt.want {
				t.Errorf("String.Left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Right(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "hello world",
			},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Right(5); got != tt.want {
				t.Errorf("String.Right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_Reverse(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s: "Kilian",
			},
			want: "nailiK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Reverse(); got != tt.want {
				t.Errorf("String.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_WordCount(t *testing.T) {
	type args struct {
		s String
	}
	tests := []struct {
		name string
		args args
		want map[String]int
	}{
		{
			name: "success",
			args: args{
				s: "the the new new of da new the du new the of the of du",
			},
			want: map[String]int{"the": 5, "new": 4, "of": 3, "du": 2, "da": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.WordCount(); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.WordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_RandomString(t *testing.T) {
	type args struct {
		s      String
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s:      "",
				length: 12,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.RandomString(tt.args.length); len(got) != tt.want {
				t.Errorf("String.RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_AddLeft(t *testing.T) {
	type args struct {
		s  String
		ss String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:  "Hertel",
				ss: "Kilian ",
			},
			want: "Kilian Hertel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.AddLeft(tt.args.ss); got != tt.want {
				t.Errorf("String.AddLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_AddRight(t *testing.T) {
	type args struct {
		s  String
		ss String
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:  "Kilian",
				ss: " Hertel",
			},
			want: "Kilian Hertel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.AddRight(tt.args.ss); got != tt.want {
				t.Errorf("String.AddRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_String_AddPos(t *testing.T) {
	type args struct {
		s   String
		ss  String
		pos int
	}
	tests := []struct {
		name string
		args args
		want String
	}{
		{
			name: "success",
			args: args{
				s:   "Kilian Hertel",
				ss:  "the one and only ",
				pos: 7,
			},
			want: "Kilian the one and only Hertel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.AddPos(tt.args.ss, tt.args.pos); got != tt.want {
				t.Errorf("String.AddPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
/*
func Test_String_FindInFiles(t *testing.T) {
	type args struct {
		s    String
		path String
	}
	tests := []struct {
		name string
		args args
		want Strings
	}{
		{
			name: "success",
			args: args{
				s:    "needle",
				path: "/FindInFilesTestFolder",
			},
			want: Strings{
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:28 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:28 - Pos:79",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:32 - Pos:18",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:58 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:73 - Pos:90",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:91 - Pos:108",
				"/home/k/go/src/String/FindInFilesTestFolder/1 - Line:811729 - Pos:99",
				"/home/k/go/src/String/FindInFilesTestFolder/2 - Line:7 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/2 - Line:32 - Pos:18",
				"/home/k/go/src/String/FindInFilesTestFolder/2 - Line:58 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/2 - Line:73 - Pos:90",
				"/home/k/go/src/String/FindInFilesTestFolder/2 - Line:91 - Pos:108",
				"/home/k/go/src/String/FindInFilesTestFolder/3 - Line:1 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/3 - Line:32 - Pos:18",
				"/home/k/go/src/String/FindInFilesTestFolder/3 - Line:58 - Pos:63",
				"/home/k/go/src/String/FindInFilesTestFolder/3 - Line:73 - Pos:90",
				"/home/k/go/src/String/FindInFilesTestFolder/3 - Line:91 - Pos:108",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd, _ := os.Getwd()
			var path String = String(wd) + tt.args.path
			if got := tt.args.s.FindInFiles(path); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.FindInFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func Test_Strings_substr(t *testing.T) {
	type args struct {
		s     Strings
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want Strings
	}{
		{
			name: "success",
			args: args{
				s: Strings{
					"Hello world!",
					"Hello world!",
					"Hello world!",
				},
				start: 3,
				end:   5,
			},
			want: Strings{
				"lo wo",
				"lo wo",
				"lo wo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.substr(tt.args.start, tt.args.end); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Strings_ASCIIsubstr(t *testing.T) {
	type args struct {
		s     Strings
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want Strings
	}{
		{
			name: "success",
			args: args{
				s: Strings{
					"Hello world!",
					"Hello world!",
					"Hello world!",
				},
				start: 3,
				end:   5,
			},
			want: Strings{
				"lo wo",
				"lo wo",
				"lo wo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.ASCIIsubstr(tt.args.start, tt.args.end); reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("String.ASCIIsubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Strings_Compare(t *testing.T) {
	type args struct {
		s  Strings
		ss Strings
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: Strings{
					"Hello world!",
					"Hello world!",
					"Hello world!",
				},
				ss: Strings{
					"Hello world!",
					"Hello world!",
					"Hello world!",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Compare(tt.args.ss); got != tt.want {
				t.Errorf("String.ASCIIsubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
