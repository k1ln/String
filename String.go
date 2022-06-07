package String

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	mrand "math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type String string
type Strings []String

var Timeout = 5
var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func (s String) substr(start, end int) String {
	counter, startIdx := 0, 0
	end = end + start
	for i := range s {
		if counter == start {
			startIdx = i
		}
		if counter == end {
			return s[startIdx:i]
		}
		counter++
	}
	return s[startIdx:]
}

func (s String) Substr(start, end int) String {
	counter, startIdx := 0, 0
	end = end + start
	for i := range s {
		if counter == start {
			startIdx = i
		}
		if counter == end {
			return s[startIdx:i]
		}
		counter++
	}
	return s[startIdx:]
}

func (s String) Tostring() string {
	return string(s)
}

func (s String) ASCIIsubstr(start, end int) String {
	lens := start + end
	if lens > len(s)-start {
		lens = len(s) - start + 1
	}
	return s[start:lens]
}

func (s String) Compare(b String) int {
	if s == b {
		return 0
	}
	if s < b {
		return -1
	}
	return +1
}

func (s String) Contains(substr String) bool {
	return strings.Index(string(s), string(substr)) >= 0
}

func (s String) ContainsAny(chars String) int {
	return strings.IndexAny(string(s), string(chars))
}

func (s String) ContainsRune(r rune) bool {
	return strings.IndexRune(string(s), r) >= 0
}

func (s String) Count(substr String) int {
	return strings.Count(string(s), string(substr))
}

func (s String) EqualFold(t String) bool {
	return strings.EqualFold(string(s), string(t))
}

func (s String) Fields() Strings {
	// First count the fields.
	// This is an exact count if s is ASCII, otherwise it is an approximation.
	n := 0
	wasSpace := 1
	// setBits is used to track which bits are set in the bytes of s.
	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
		r := s[i]
		setBits |= r
		isSpace := int(asciiSpace[r])
		n += wasSpace & ^isSpace
		wasSpace = isSpace
	}

	if setBits >= utf8.RuneSelf {
		// Some runes in the input string are not ASCII.
		return s.FieldsFunc(unicode.IsSpace)
	}
	// ASCII fast path
	a := make([]String, n)
	na := 0
	fieldStart := 0
	i := 0
	// Skip spaces in the front of the input.
	for i < len(s) && asciiSpace[s[i]] != 0 {
		i++
	}
	fieldStart = i
	for i < len(s) {
		if asciiSpace[s[i]] == 0 {
			i++
			continue
		}
		a[na] = s[fieldStart:i]
		na++
		i++
		// Skip spaces in between fields.
		for i < len(s) && asciiSpace[s[i]] != 0 {
			i++
		}
		fieldStart = i
	}
	if fieldStart < len(s) { // Last field might end at EOF.
		a[na] = s[fieldStart:]
	}
	return a
}

func (s String) FieldsFunc(f func(rune) bool) []String {
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	// Find the field start and end indices.
	wasField := false
	fromIndex := 0
	for i, rune := range s {
		if f(rune) {
			if wasField {
				spans = append(spans, span{start: fromIndex, end: i})
				wasField = false
			}
		} else {
			if !wasField {
				fromIndex = i
				wasField = true
			}
		}
	}

	// Last field might end at EOF.
	if wasField {
		spans = append(spans, span{fromIndex, len(s)})
	}

	// Create strings from recorded field indices.
	a := make([]String, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}

func (s String) HasPrefix(prefix String) bool {
	l := len(prefix)
	return len(s) >= l && s[0:l] == prefix
}

func (s String) HasSuffix(suffix String) bool {
	l := len(suffix)
	ls := len(s)
	return ls >= l && s[ls-l:] == suffix
}

func (s String) len() int {
	return len(s)
}

func (s String) length() int {
	return len(s)
}

func (s String) Index(substr String) int {
	return strings.Index(string(s), string(substr))
}

func (s String) IndexAny(chars String) int {
	return strings.IndexAny(string(s), string(chars))
}

func (s String) IndexByte(c byte) int {
	return strings.IndexByte(string(s), c)
}

func (s String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

func (s String) IndexRune(r rune) int {
	switch {
	case 0 <= r && r < utf8.RuneSelf:
		return s.IndexByte(byte(r))
	case r == utf8.RuneError:
		for i, r := range s {
			if r == utf8.RuneError {
				return i
			}
		}
		return -1
	case !utf8.ValidRune(r):
		return -1
	default:
		return s.Index(String(r))
	}
}

func (s Strings) Join(sep String) String {
	ls := len(s)
	ssep := string(sep)
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0]
	}
	n := len(sep) * (ls - 1)
	for i := 0; i < ls; i++ {
		n += len(s[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(s[0]))
	for _, ss := range s[1:] {
		b.WriteString(ssep)
		b.WriteString(string(ss))
	}
	return String(b.String())
}

func (s String) LastIndex(substr String) int {
	return strings.LastIndex(string(s), string(substr))
}

func (s String) LastIndexAny(chars String) int {
	return strings.LastIndexAny(string(s), string(chars))
}

func (s String) LastIndexByte(c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func (s String) LastIndexFunc(f func(rune) bool) int {
	return s.lastIndexFunc(f, true)
}

func (s String) lastIndexFunc(f func(rune) bool, truth bool) int {
	for i := len(s); i > 0; {
		r, size := utf8.DecodeLastRuneInString(string(s)[0:i])
		i -= size
		if f(r) == truth {
			return i
		}
	}
	return -1
}

func (s String) Map(mapping func(rune) rune) String {
	return String(strings.Map(mapping, string(s)))
}

func (s String) Repeat(count int) String {
	ss := string(s)
	ls := len(s)
	if count == 0 {
		return ""
	}

	// Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237
	if count < 0 {
		panic("strings: negative Repeat count")
	} else if ls*count/count != ls {
		panic("strings: Repeat count causes overflow")
	}

	n := ls * count
	var b strings.Builder
	b.Grow(n)
	b.WriteString(ss)
	for b.Len() < n {
		if b.Len() <= n/2 {
			b.WriteString(b.String())
		} else {
			b.WriteString(b.String()[:n-b.Len()])
			break
		}
	}
	return String(b.String())
}

func (s String) Replace(old, new String, n int) String {
	return String(strings.Replace(string(s), string(old), string(new), n))
}

func (s String) ReplaceAll(old, new String) String {
	return s.Replace(old, new, -1)
}

func (s String) Split(sep String) []String {
	return s.genSplit(sep, 0, -1)
}

func (s String) genSplit(sep String, sepSave, n int) []String {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return s.explode(n)
	}
	if n < 0 {
		n = s.Count(sep) + 1
	}

	a := make([]String, n)
	n--
	i := 0
	for i < n {
		m := s.Index(sep)
		if m < 0 {
			break
		}
		a[i] = s[:m+sepSave]
		s = s[m+len(sep):]
		i++
	}
	a[i] = s
	return a[:i+1]
}

func (s String) explode(n int) []String {
	l := utf8.RuneCountInString(string(s))
	if n < 0 || n > l {
		n = l
	}
	a := make([]String, n)
	for i := 0; i < n-1; i++ {
		ch, size := utf8.DecodeRuneInString(string(s))
		a[i] = s[:size]
		s = s[size:]
		if ch == utf8.RuneError {
			a[i] = String(utf8.RuneError)
		}
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}

func (s String) SplitAfter(sep String) []String {
	return s.genSplit(sep, len(sep), -1)
}

func (s String) SplitAfterN(sep String, n int) []String {
	return s.genSplit(sep, len(sep), n)
}

func (s String) SplitN(sep String, n int) []String {
	return s.genSplit(sep, 0, n)
}

func (s String) Title() String {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return s.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		})
}

func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

func (s String) ToLower() String {
	isASCII, hasUpper := true, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			break
		}
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}

	if isASCII { // optimize for ASCII-only strings.
		if !hasUpper {
			return s
		}
		var b strings.Builder
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'A' <= c && c <= 'Z' {
				c += 'a' - 'A'
			}
			b.WriteByte(c)
		}
		return String(b.String())
	}
	return s.Map(unicode.ToLower)
}

func (s String) ToLowerSpecial(c unicode.SpecialCase) String {
	return s.Map(c.ToLower)
}

func (s String) ToTitle() String { return s.Map(unicode.ToTitle) }

func (s String) ToUpper() String {
	isASCII, hasLower := true, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			break
		}
		hasLower = hasLower || ('a' <= c && c <= 'z')
	}

	if isASCII { // optimize for ASCII-only strings.
		if !hasLower {
			return s
		}
		var b strings.Builder
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
			b.WriteByte(c)
		}
		return String(b.String())
	}
	return s.Map(unicode.ToUpper)
}

func (s String) ToTitleSpecial(c unicode.SpecialCase) String {
	return s.Map(c.ToTitle)
}

func (s String) ToUpperSpecial(c unicode.SpecialCase) String {
	return s.Map(c.ToUpper)
}

// I dont know how to test this TODO
func (s String) ToValidUTF8(replacement String) String {
	var b strings.Builder
	str := string(s)

	for i, c := range str {
		if c != utf8.RuneError {
			continue
		}

		_, wid := utf8.DecodeRuneInString(str[i:])
		if wid == 1 {
			b.Grow(len(str) + len(replacement))
			b.WriteString(str[:i])
			str = str[i:]
			break
		}
	}

	// Fast path for unchanged input
	if b.Cap() == 0 { // didn't call b.Grow above
		return String(str)
	}

	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(str); {
		c := str[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b.WriteByte(c)
			continue
		}
		_, wid := utf8.DecodeRuneInString(str[i:])
		if wid == 1 {
			i++
			if !invalid {
				invalid = true
				b.WriteString(string(replacement))
			}
			continue
		}
		invalid = false
		b.WriteString(str[i : i+wid])
		i += wid
	}

	return String(b.String())
}

func (s String) Trim(cutset String) String {
	if s == "" || cutset == "" {
		return s
	}
	return s.TrimFunc(cutset.makeCutsetFunc())
}

func (s String) TrimFunc(f func(rune) bool) String {
	return s.TrimLeftFunc(f).TrimRightFunc(f)
}

func (s String) TrimRightFunc(f func(rune) bool) String {
	i := s.lastIndexFunc(f, false)
	if i >= 0 && s[i] >= utf8.RuneSelf {
		_, wid := utf8.DecodeRuneInString(string(s)[i:])
		i += wid
	} else {
		i++
	}
	return s[0:i]
}

func (s String) TrimLeftFunc(f func(rune) bool) String {
	i := s.indexFunc(f, false)
	if i == -1 {
		return ""
	}
	return s[i:]
}

func (s String) indexFunc(f func(rune) bool, truth bool) int {
	for i, r := range s {
		if f(r) == truth {
			return i
		}
	}
	return -1
}

func (cutset String) makeCutsetFunc() func(rune) bool {
	if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
		return func(r rune) bool {
			return r == rune(cutset[0])
		}
	}
	if as, isASCII := cutset.makeASCIISet(); isASCII {
		return func(r rune) bool {
			return r < utf8.RuneSelf && as.contains(byte(r))
		}
	}
	return func(r rune) bool { return cutset.IndexRune(r) >= 0 }
}

type asciiSet [8]uint32

func (chars String) makeASCIISet() (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c>>5] |= 1 << uint(c&31)
	}
	return as, true
}

func (as *asciiSet) contains(c byte) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}

func (s String) TrimLeft(cutset String) String {
	if s == "" || cutset == "" {
		return s
	}
	return s.TrimLeftFunc(cutset.makeCutsetFunc())
}

func (s String) TrimPrefix(prefix String) String {
	if s.HasPrefix(prefix) {
		return s[len(prefix):]
	}
	return s
}

func (s String) TrimRight(cutset String) String {
	if s == "" || cutset == "" {
		return s
	}
	return s.TrimRightFunc(cutset.makeCutsetFunc())
}

func (s String) TrimSpace() String {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return s[start:].TrimFunc(unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// Now look for the first ASCII non-space byte from the end
	stop := len(s)
	for ; stop > start; stop-- {
		c := s[stop-1]
		if c >= utf8.RuneSelf {
			return s[start:stop].TrimFunc(unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}

	// At this point s[start:stop] starts and ends with an ASCII
	// non-space bytes, so we're done. Non-ASCII cases have already
	// been handled above.
	return s[start:stop]
}

func (s String) TrimSuffix(suffix String) String {
	if s.HasSuffix(suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

//own functions

func (s String) Md5() String {
	h := md5.New()
	io.WriteString(h, string(s))
	return String(hex.EncodeToString(h.Sum(nil)[:]))
}

func (s String) Sha1() String {
	h := sha1.New()
	io.WriteString(h, string(s))
	return String(hex.EncodeToString(h.Sum(nil)[:]))
}

//uses CTR
func (s String) AesEncrypt(key String) String {
	hexkey, _ := hex.DecodeString(string(key))
	str := []byte(s)
	aescipher, err := aes.NewCipher(hexkey)
	if err != nil {
		panic(err)
	}
	blocksize := aes.BlockSize

	var cypher []byte = make([]byte, blocksize+len(s))
	IV := cypher[:blocksize]
	if _, err := io.ReadFull(rand.Reader, IV); err != nil {
		panic(err)
	}
	//encrypt
	ctr := cipher.NewCTR(aescipher, IV)
	ctr.XORKeyStream(cypher[blocksize:], str)
	//IV is added to the beginning of the cypher
	return String(b64.StdEncoding.EncodeToString(append(IV, cypher[blocksize:]...)))
}

//returns byte if you want to return binary
func (s String) AesEncryptByte(key String) []byte {
	hexkey, _ := hex.DecodeString(string(key))
	str := []byte(s)
	aescipher, err := aes.NewCipher(hexkey)
	if err != nil {
		panic(err)
	}
	blocksize := aes.BlockSize

	var cypher []byte = make([]byte, blocksize+len(s))
	IV := cypher[:blocksize]
	if _, err := io.ReadFull(rand.Reader, IV); err != nil {
		panic(err)
	}
	//encrypt
	ctr := cipher.NewCTR(aescipher, IV)
	ctr.XORKeyStream(cypher[blocksize:], str)
	//IV is added to the beginning of the cypher
	return append(IV, cypher[blocksize:]...)
}

//needs base64encoded string
func (s String) AesDecrypt(key String) String {
	stra, _ := b64.StdEncoding.DecodeString(string(s))
	str := []byte(stra)
	hexkey, _ := hex.DecodeString(string(key))
	aescipher, err := aes.NewCipher(hexkey)
	if err != nil {
		panic(err)
	}
	//decrypt
	//IV removed from beginning of the cypher
	IV := str[:aes.BlockSize]
	text := make([]byte, len(str[aes.BlockSize:]))
	ctr := cipher.NewCTR(aescipher, IV)
	ctr.XORKeyStream(text, str[aes.BlockSize:])
	return String(text)
}

//Convert byte to String object to use it. It will work
func (s String) AesDecryptByte(key String) String {
	str := []byte(s)
	hexkey, _ := hex.DecodeString(string(key))
	aescipher, err := aes.NewCipher(hexkey)
	if err != nil {
		panic(err)
	}
	//decrypt
	//IV removed from beginning of the cypher
	IV := str[:aes.BlockSize]
	text := make([]byte, len(str[aes.BlockSize:]))
	ctr := cipher.NewCTR(aescipher, IV)
	ctr.XORKeyStream(text, str[aes.BlockSize:])
	return String(text)
}

// generate 16 24 or 32 byte key for 128 192 or 256-bit Encryption
func (s String) GenerateAesKeyHex(length int) String {
	if length != 16 && length != 24 && length != 32 {
		panic("Please use 16,24 or 32 as Key length")
	}
	var key []byte = make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	//fmt.Println(key)
	return String(hex.EncodeToString(key))
}

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var reUserNamestr String = "^[a-z0-9_-]{{min},{max}}$"
var reUrl = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#()?&//=]*)")

//var reComplexPw = regexp.MustCompile("(?=(.*[0-9]))(?=.*[\\!@#$%^&*()\\[\\]{}\\-_+=~`|:;\"'<>,./?])(?=.*[a-z])(?=(.*[A-Z]))(?=(.*)).{8,}")
var reWholeNumber = regexp.MustCompile("^\\d+$")
var reIsIpV4 = regexp.MustCompile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")
var reIsIpV6 = regexp.MustCompile("(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))")
var reIsIp = regexp.MustCompile("((^\\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\\s*$)|(^\\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:)))(%.+)?\\s*$))")
var reIsHtmlTag = regexp.MustCompile("<\\/?[\\w\\s]*>|<.+[\\W]>")
var reIsPhoneNumber = regexp.MustCompile("^(?:(?:\\(?(?:00|\\+)([1-4]\\d\\d|[1-9]\\d?)\\)?)?[\\-\\.\\ \\\\\\/]?)?((?:\\(?\\d{1,}\\)?[\\-\\.\\ \\\\\\/]?){0,})(?:[\\-\\.\\ \\\\\\/]?(?:#|ext\\.?|extension|x)[\\-\\.\\ \\\\\\/]?(\\d+))?$")
var reIsFilePath = regexp.MustCompile("^(.+)/([^/]+)$")

var reZipMap = map[String]String{
	"GB": "GIR[ ]?0AA|((AB|AL|B|BA|BB|BD|BH|BL|BN|BR|BS|BT|CA|CB|CF|CH|CM|CO|CR|CT|CV|CW|DA|DD|DE|DG|DH|DL|DN|DT|DY|E|EC|EH|EN|EX|FK|FY|G|GL|GY|GU|HA|HD|HG|HP|HR|HS|HU|HX|IG|IM|IP|IV|JE|KA|KT|KW|KY|L|LA|LD|LE|LL|LN|LS|LU|M|ME|MK|ML|N|NE|NG|NN|NP|NR|NW|OL|OX|PA|PE|PH|PL|PO|PR|RG|RH|RM|S|SA|SE|SG|SK|SL|SM|SN|SO|SP|SR|SS|ST|SW|SY|TA|TD|TF|TN|TQ|TR|TS|TW|UB|W|WA|WC|WD|WF|WN|WR|WS|WV|YO|ZE)(\\d[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}))|BFPO[ ]?\\d{1,4}",
	"JE": "JE\\d[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}",
	"GG": "GY\\d[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}",
	"IM": "IM\\d[\\dA-Z]?[ ]?\\d[ABD-HJLN-UW-Z]{2}",
	"US": "\\d{5}([ \\-]\\d{4})?",
	"CA": "[ABCEGHJKLMNPRSTVXY]\\d[ABCEGHJ-NPRSTV-Z][ ]?\\d[ABCEGHJ-NPRSTV-Z]\\d",
	"DE": "\\d{5}|\\d{4}",
	"JP": "\\d{3}-\\d{4}",
	"FR": "\\d{2}[ ]?\\d{3}",
	"AU": "\\d{4}",
	"IT": "\\d{5}",
	"CH": "\\d{4}",
	"AT": "\\d{4}",
	"ES": "\\d{5}",
	"NL": "\\d{4}[ ]?[A-Z]{2}",
	"BE": "\\d{4}",
	"DK": "\\d{4}",
	"SE": "\\d{3}[ ]?\\d{2}",
	"NO": "\\d{4}",
	"BR": "\\d{5}[\\-]?\\d{3}",
	"PT": "\\d{4}([\\-]\\d{3})?",
	"FI": "\\d{5}",
	"AX": "22\\d{3}",
	"KR": "\\d{3}[\\-]\\d{3}",
	"CN": "\\d{6}",
	"TW": "\\d{3}(\\d{2})?",
	"SG": "\\d{6}",
	"DZ": "\\d{5}",
	"AD": "AD\\d{3}",
	"AR": "([A-HJ-NP-Z])?\\d{4}([A-Z]{3})?",
	"AM": "(37)?\\d{4}",
	"AZ": "\\d{4}",
	"BH": "((1[0-2]|[2-9])\\d{2})?",
	"BD": "\\d{4}",
	"BB": "(BB\\d{5})?",
	"BY": "\\d{6}",
	"BM": "[A-Z]{2}[ ]?[A-Z0-9]{2}",
	"BA": "\\d{5}",
	"IO": "BBND 1ZZ",
	"BN": "[A-Z]{2}[ ]?\\d{4}",
	"BG": "\\d{4}",
	"KH": "\\d{5}",
	"CV": "\\d{4}",
	"CL": "\\d{7}",
	"CR": "\\d{4,5}|\\d{3}-\\d{4}",
	"HR": "\\d{5}",
	"CY": "\\d{4}",
	"CZ": "\\d{3}[ ]?\\d{2}",
	"DO": "\\d{5}",
	"EC": "([A-Z]\\d{4}[A-Z]|(?:[A-Z]{2})?\\d{6})?",
	"EG": "\\d{5}",
	"EE": "\\d{5}",
	"FO": "\\d{3}",
	"GE": "\\d{4}",
	"GR": "\\d{3}[ ]?\\d{2}",
	"GL": "39\\d{2}",
	"GT": "\\d{5}",
	"HT": "\\d{4}",
	"HN": "(?:\\d{5})?",
	"HU": "\\d{4}",
	"IS": "\\d{3}",
	"IN": "\\d{6}",
	"ID": "\\d{5}",
	"IL": "\\d{5}",
	"JO": "\\d{5}",
	"KZ": "\\d{6}",
	"KE": "\\d{5}",
	"KW": "\\d{5}",
	"LA": "\\d{5}",
	"LV": "\\d{4}",
	"LB": "(\\d{4}([ ]?\\d{4})?)?",
	"LI": "(948[5-9])|(949[0-7])",
	"LT": "\\d{5}",
	"LU": "\\d{4}",
	"MK": "\\d{4}",
	"MY": "\\d{5}",
	"MV": "\\d{5}",
	"MT": "[A-Z]{3}[ ]?\\d{2,4}",
	"MU": "(\\d{3}[A-Z]{2}\\d{3})?",
	"MX": "\\d{5}",
	"MD": "\\d{4}",
	"MC": "980\\d{2}",
	"MA": "\\d{5}",
	"NP": "\\d{5}",
	"NZ": "\\d{4}",
	"NI": "((\\d{4}-)?\\d{3}-\\d{3}(-\\d{1})?)?",
	"NG": "(\\d{6})?",
	"OM": "(PC )?\\d{3}",
	"PK": "\\d{5}",
	"PY": "\\d{4}",
	"PH": "\\d{4}",
	"PL": "\\d{2}-\\d{3}",
	"PR": "00[679]\\d{2}([ \\-]\\d{4})?",
	"RO": "\\d{6}",
	"RU": "\\d{6}",
	"SM": "4789\\d",
	"SA": "\\d{5}",
	"SN": "\\d{5}",
	"SK": "\\d{3}[ ]?\\d{2}",
	"SI": "\\d{4}",
	"ZA": "\\d{4}",
	"LK": "\\d{5}",
	"TJ": "\\d{6}",
	"TH": "\\d{5}",
	"TN": "\\d{4}",
	"TR": "\\d{5}",
	"TM": "\\d{6}",
	"UA": "\\d{5}",
	"UY": "\\d{5}",
	"UZ": "\\d{6}",
	"VA": "00120",
	"VE": "\\d{4}",
	"ZM": "\\d{5}",
	"AS": "96799",
	"CC": "6799",
	"CK": "\\d{4}",
	"RS": "\\d{6}",
	"ME": "8\\d{4}",
	"CS": "\\d{5}",
	"YU": "\\d{5}",
	"CX": "6798",
	"ET": "\\d{4}",
	"FK": "FIQQ 1ZZ",
	"NF": "2899",
	"FM": "(9694[1-4])([ \\-]\\d{4})?",
	"GF": "9[78]3\\d{2}",
	"GN": "\\d{3}",
	"GP": "9[78][01]\\d{2}",
	"GS": "SIQQ 1ZZ",
	"GU": "969[123]\\d([ \\-]\\d{4})?",
	"GW": "\\d{4}",
	"HM": "\\d{4}",
	"IQ": "\\d{5}",
	"KG": "\\d{6}",
	"LR": "\\d{4}",
	"LS": "\\d{3}",
	"MG": "\\d{3}",
	"MH": "969[67]\\d([ \\-]\\d{4})?",
	"MN": "\\d{6}",
	"MP": "9695[012]([ \\-]\\d{4})?",
	"MQ": "9[78]2\\d{2}",
	"NC": "988\\d{2}",
	"NE": "\\d{4}",
	"VI": "008(([0-4]\\d)|(5[01]))([ \\-]\\d{4})?",
	"PF": "987\\d{2}",
	"PG": "\\d{3}",
	"PM": "9[78]5\\d{2}",
	"PN": "PCRN 1ZZ",
	"PW": "96940",
	"RE": "9[78]4\\d{2}",
	"SH": "(ASCN|STHL) 1ZZ",
	"SJ": "\\d{4}",
	"SO": "\\d{5}",
	"SZ": "[HLMS]\\d{3}",
	"TC": "TKCA 1ZZ",
	"WF": "986\\d{2}",
	"XK": "\\d{5}",
	"YT": "976\\d{2}",
}

var reIbanMap = map[String]String{
	"AL": "^AL\\d{10}[0-9A-Z]{16}$",
	"AD": "^AD\\d{10}[0-9A-Z]{12}$",
	"AT": "^AT\\d{18}$",
	"BH": "^BH\\d{2}[A-Z]{4}[0-9A-Z]{14}$",
	"BE": "^BE\\d{14}$",
	"BA": "^BA\\d{18}$",
	"BG": "^BG\\d{2}[A-Z]{4}\\d{6}[0-9A-Z]{8}$",
	"HR": "^HR\\d{19}$",
	"CY": "^CY\\d{10}[0-9A-Z]{16}$",
	"CZ": "^CZ\\d{22}$",
	"DK": "^DK\\d{16}$|^FO\\d{16}$|^GL\\d{16}$",
	"DO": "^DO\\d{2}[0-9A-Z]{4}\\d{20}$",
	"EE": "^EE\\d{18}$",
	"FI": "^FI\\d{16}$",
	"FR": "^FR\\d{12}[0-9A-Z]{11}\\d{2}$",
	"GE": "^GE\\d{2}[A-Z]{2}\\d{16}$",
	"DE": "^DE\\d{20}$",
	"GI": "^GI\\d{2}[A-Z]{4}[0-9A-Z]{15}$",
	"GR": "^GR\\d{9}[0-9A-Z]{16}$",
	"HU": "^HU\\d{26}$",
	"IS": "^IS\\d{24}$",
	"IE": "^IE\\d{2}[A-Z]{4}\\d{14}$",
	"IL": "^IL\\d{21}$",
	"IT": "^IT\\d{2}[A-Z]\\d{10}[0-9A-Z]{12}$",
	"KZ": "^[A-Z]{2}\\d{5}[0-9A-Z]{13}$",
	"KW": "^KW\\d{2}[A-Z]{4}22!$",
	"LV": "^LV\\d{2}[A-Z]{4}[0-9A-Z]{13}$",
	"LB": "^LB\\d{6}[0-9A-Z]{20}$",
	"LI": "^LI\\d{7}[0-9A-Z]{12}$",
	"LT": "^LT\\d{18}$",
	"LU": "^LU\\d{5}[0-9A-Z]{13}$",
	"MK": "^MK\\d{5}[0-9A-Z]{10}\\d{2}$",
	"MT": "^MT\\d{2}[A-Z]{4}\\d{5}[0-9A-Z]{18}$",
	"MR": "^MR13\\d{23}$",
	"MU": "^MU\\d{2}[A-Z]{4}\\d{19}[A-Z]{3}$",
	"MC": "^MC\\d{12}[0-9A-Z]{11}\\d{2}$",
	"ME": "^ME\\d{20}$",
	"NL": "^NL\\d{2}[A-Z]{4}\\d{10}$",
	"NO": "^NO\\d{13}$",
	"PL": "^PL\\d{10}[0-9A-Z]{,16}n$",
	"PT": "^PT\\d{23}$",
	"RO": "^RO\\d{2}[A-Z]{4}[0-9A-Z]{16}$",
	"SM": "^SM\\d{2}[A-Z]\\d{10}[0-9A-Z]{12}$",
	"SA": "^SA\\d{4}[0-9A-Z]{18}$",
	"RS": "^RS\\d{20}$",
	"SK": "^SK\\d{22}$",
	"SI": "^SI\\d{17}$",
	"ES": "^ES\\d{22}$",
	"SE": "^SE\\d{22}$",
	"CH": "^CH\\d{7}[0-9A-Z]{12}$",
	"TN": "^TN59\\d{20}$",
	"TR": "^TR\\d{7}[0-9A-Z]{17}$",
	"AE": "^AE\\d{21}$",
	"GB": "^GB\\d{2}[A-Z]{4}\\d{14}$",
}

func (s String) string() string {
	return string(s)
}

func (s String) IsEmail() bool {
	return reEmail.MatchString(s.string())
}

func (s String) IsUrl() bool {
	return reUrl.MatchString(s.string())
}

/*
func (s String) IsComplexPw() bool{
        return reComplexPw.MatchString(s.string())
}
*/

func (s String) IsWholeNumber() bool {
	return reWholeNumber.MatchString(s.string())
}

func (s String) IsIpV4() bool {
	return reIsIpV4.MatchString(s.string())
}

func (s String) IsIpV6() bool {
	return reIsIpV6.MatchString(s.string())
}

func (s String) IsIp() bool {
	return reIsIp.MatchString(s.string())
}

func (s String) IsHtmlTag() bool {
	return reIsHtmlTag.MatchString(s.string())
}

func (s String) IsPhoneNumber() bool {
	return reIsPhoneNumber.MatchString(s.string())
}

func (s String) IsFilePath() bool {
	return reIsFilePath.MatchString(s.string())
}

func (s String) IsUserName(min int, max int) bool {
	reUserName := reUserNamestr.ReplaceAll("{min}", String(strconv.Itoa(min))).ReplaceAll("{max}", String(strconv.Itoa(max)))
	rexUserName := regexp.MustCompile(reUserName.string())
	return rexUserName.MatchString(string(s))
}

func (s String) IsZipCode(country String) bool {
	var reIsZip = regexp.MustCompile(reZipMap[country].string())
	return reIsZip.MatchString(s.string())
}

func (s String) IsIban(country String) bool {
	var reIsIban = regexp.MustCompile(reIbanMap[country].string())
	return reIsIban.MatchString(s.string())
}

var reIsZipFast = regexp.MustCompile("\\d{5}|\\d{4}")

func (s String) PrecompileIsZipCodeFast(country String) {
	reIsZipFast = regexp.MustCompile(reZipMap[country].string())
}

func (s String) IsZipCodeFast(country String) bool {
	return reIsZipFast.MatchString(s.string())
}

var reIsIbanFast = regexp.MustCompile("^(.+)/([^/]+)$")

func (s String) PrecompileIsIbanFast(country String) {
	reIsIbanFast = regexp.MustCompile(reIbanMap[country].string())
}

func (s String) IsIbanFast(country String) bool {
	var reIsIban = regexp.MustCompile(reIbanMap[country].string())
	return reIsIban.MatchString(s.string())
}

func (s String) PwUpperCase(number int) bool {
	if number == 0 {
		return true
	} else {
		regxstring := "^"
		i := 0
		for i < number {
			regxstring += ".*[A-Z]"
			i++
		}
		var re = regexp.MustCompile(regxstring)
		return re.MatchString(string(s))
	}
}

func (s String) PwSpecialCase(number int) bool {
	if number == 0 {
		return true
	} else {
		regxstring := "^"
		i := 0
		for i < number {
			regxstring += ".*[!@#$%^&*(),.?\":{}|<>]"
			i++
		}
		var re = regexp.MustCompile(regxstring)
		return re.MatchString(string(s))
	}
}

func (s String) PwDigits(number int) bool {
	if number == 0 {
		return true
	} else {
		regxstring := "^"
		i := 0
		for i < number {
			regxstring += ".*[0-9]"
			i++
		}
		var re = regexp.MustCompile(regxstring)
		return re.MatchString(string(s))
	}
}

func (s String) PwLowerCase(number int) bool {
	if number == 0 {
		return true
	} else {
		regxstring := "^"
		i := 0
		for i < number {
			regxstring += ".*[0-9]"
			i++
		}
		var re = regexp.MustCompile(regxstring)
		return re.MatchString(string(s))
	}
}

func (s String) Get() String {
	client := http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
	}
	resp, err := client.Get(string(s))
	if err != nil {
		return String(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return String(err)
	}
	return String(body)
}

//TODO I need something better here
func (s String) Json() map[String]interface{} {
	var result map[String]interface{}
	//for each in continue interface is varienat though
	json.Unmarshal([]byte(s), &result)
	return result
}

func (s String) Open() String {
	return String(s.OpenByte())
}

func (s String) OpenByte() []byte {
	b, err := ioutil.ReadFile(string(s)) // just pass the file name
	if err != nil {
		panic(err)
	}
	return b
}

func (s String) Exists() bool {
	if s.substr(0, 7) == "http://" || s.substr(0, 8) == "https://" {
		resp, err := http.Get(string(s))
		if err != nil {
			return false
		}
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			return true
		} else {
			return false
		}
	} else {
		info, err := os.Stat(string(s))
		if os.IsNotExist(err) {
			return false
		}
		return !info.IsDir()
	}
}

func (s String) GetContents() String {
	if s.substr(0, 7) == "http://" || s.substr(0, 8) == "https://" {
		return s.Get()
	} else {
		return s.Open()
	}
}

func (s String) WriteToFile(path String) {
	f, err := os.OpenFile(string(path),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(s)); err != nil {
		panic(err)
	}
}

func (s String) URLEncode() String {
	return String(url.QueryEscape(string(s)))
}

func (s String) URLDecode() String {
	dec, _ := url.QueryUnescape(string(s))
	return String(dec)
}

func (s String) B64Encode() String {
	return String(b64.StdEncoding.EncodeToString([]byte(s)))
}

func (s String) B64Decode() String {
	dec, _ := b64.StdEncoding.DecodeString(string(s))
	return String(dec)
}

func (s String) B64URLEncode() String {
	return String(b64.URLEncoding.EncodeToString([]byte(s)))
}

func (s String) B64URLDecode() String {
	dec, _ := b64.URLEncoding.DecodeString(string(s))
	return String(dec)
}

func (s String) Post(url String, contenttype String) String {
	req, err := http.NewRequest("POST", string(url), bytes.NewBuffer([]byte(s)))
	req.Header.Set("Content-Type", string(contenttype))

	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return String(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return String(body)
}

func (s String) CreateCommandFields() []String {
	i := 0
	inquotes := false
	insinglequotes := false
	oldi := 0
	var commands []String
	for i < len(s) {
		if i == len(s)-1 {
			if s.substr(len(s)-1, 1) == "'" || s.substr(len(s)-1, 1) == "\"" {
				commands = append(commands, s.substr(oldi, i-oldi))
			} else {
				commands = append(commands, s.substr(oldi, i-oldi+1))
			}

		} else if s.substr(i, 1) == " " &&
			inquotes == false &&
			insinglequotes == false &&
			s.substr(i-1, 1) != "\"" &&
			s.substr(i-1, 1) != "'" {
			commands = append(commands, s.substr(oldi, i-oldi))
			//fmt.Println("inquotes false space:" + s.substr(oldi,i-oldi))
			oldi = i + 1
		} else if inquotes == true &&
			s.substr(i, 1) == "\"" {
			//fmt.Println("inquotes quote:" + s.substr(oldi,i-oldi+1))
			inquotes = false
			commands = append(commands, s.substr(oldi, i-oldi))
			oldi = i + 1
		} else if insinglequotes == true &&
			s.substr(i, 1) == "'" {
			//fmt.Println("inquotes quote:" + s.substr(oldi,i-oldi+1))
			insinglequotes = false
			commands = append(commands, s.substr(oldi, i-oldi))
			oldi = i + 1
		} else if insinglequotes == false && inquotes == false && s.substr(i, 1) == "\"" {
			//fmt.Println("inquotes false quote:" + s.substr(oldi,i-oldi))
			inquotes = true
			i++
			oldi = i
		} else if inquotes == false && s.substr(i, 1) == "'" {
			//fmt.Println("inquotes false quote:" + s.substr(oldi,i-oldi))
			insinglequotes = true
			i++
			oldi = i
		}
		i++
	}
	return commands
}

func (s String) Execute() (String, String) {
	commands := s.CreateCommandFields()
	if len(commands) > 0 {
		args := []string{}
		icommands := 1
		for icommands < len(commands) {
			//fmt.Println(string(commands[icommands]))
			args = append(args, string(commands[icommands]))
			icommands++
		}
		cmd := exec.Command(string(commands[0]), args...)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			return String(err)
		}
		return String(stdout.Bytes()[:len(stdout.Bytes())]).TrimRight("\n"), String(stderr.Bytes()[:len(stdout.Bytes())]).TrimRight("\n")
	} else {
		return "", "no commands selected"
	}
	// /bin /usr/bin if commands in there than run
}

func (s String) Php() (String, String) {
	var ss String = "php -r \"" + s + "\""
	return ss.Execute()
}

func (s String) Python() (String, String) {
	var ss String = "python -c '" + s + "'"
	return ss.Execute()
}

func (s String) Node() (String, String) {
	var ss String = "node -e '" + s + "'"
	return ss.Execute()
}

func (s String) Perl() (String, String) {
	var ss String = "perl -e '" + s + "'"
	return ss.Execute()
}

func (s String) PhpFile() (String, String) {
	var ss String = "php -f " + s
	return ss.Execute()
}

func (s String) PythonFile() (String, String) {
	var ss String = "python " + s
	return ss.Execute()
}

func (s String) NodeFile() (String, String) {
	var ss String = "node " + s
	return ss.Execute()
}

func (s String) PerlFile() (String, String) {
	var ss String = "perl " + s
	return ss.Execute()
}

func (s String) ParseDateLocal(format String, location String) time.Time {
	loc, err := time.LoadLocation(string(location))
	if err != nil {
		panic(err)
	}
	if format.Contains("YY") {
		if format.Contains("YYYY") {
			format = format.ReplaceAll("YYYY", "2006")
		} else {
			format = format.ReplaceAll("YY", "06")
		}
	}

	if format.Contains("M") {
		if format.Contains("MM") {
			if format.Contains("MMM") {
				if format.Contains("MMMM") {
					format = format.ReplaceAll("MMMM", "January")
					if format.Contains("MMM") {
						format = format.ReplaceAll("MMM", "Jan")
					}
					if format.Contains("MM") {
						format = format.ReplaceAll("MM", "01")
					}
					if format.Contains("M") {
						format = format.ReplaceAll("M", "1")
					}
				} else {
					format = format.ReplaceAll("MMM", "Jan")
					if format.Contains("MM") {
						format = format.ReplaceAll("MM", "01")
					}
					if format.Contains("M") {
						format = format.ReplaceAll("M", "1")
					}
				}
			} else {
				format = format.ReplaceAll("MM", "01")
				if format.Contains("M") {
					format = format.ReplaceAll("M", "1")
				}
			}
		} else {
			format = format.ReplaceAll("M", "1")
		}
	}

	if format.Contains("D") {
		if format.Contains("DD") {
			if format.Contains("DDD") {
				if format.Contains("DDDD") {
					format = format.ReplaceAll("DDDD", "Monday")
					if format.Contains("DDD") {
						format = format.ReplaceAll("DDD", "Mon")
					}
					if format.Contains("DD") {
						format = format.ReplaceAll("DD", "02")
					}
					if format.Contains("D") {
						format = format.ReplaceAll("D", "2")
					}
				} else {
					format = format.ReplaceAll("DDD", "Mon")
					if format.Contains("DD") {
						format = format.ReplaceAll("DD", "02")
					}
					if format.Contains("D") {
						format = format.ReplaceAll("D", "2")
					}
				}
			} else {
				format = format.ReplaceAll("DD", "02")
				if format.Contains("D") {
					format = format.ReplaceAll("D", "2")
				}
			}
		} else {
			format = format.ReplaceAll("D", "2")
		}
	}

	if format.Contains("hh") {
		if format.Contains("hh12") {
			format = format.ReplaceAll("hh12", "03")
		} else {
			format = format.ReplaceAll("hh", "15")
		}
	}
	if format.Contains("h12") {
		format = format.ReplaceAll("h12", "3")
	}

	if format.Contains("m") {
		if format.Contains("mm") {
			format = format.ReplaceAll("mm", "04")
		} else {
			format = format.ReplaceAll("m", "4")
		}
	}

	if format.Contains("s") {
		if format.Contains("ss") {
			format = format.ReplaceAll("ss", "05")
		} else {
			format = format.ReplaceAll("s", "5")
		}
	}

	if format.Contains("Z") {
		if format.Contains("-ZZ") {
			if format.Contains("-ZZZ") {
				format = format.ReplaceAll("-ZZZ", "-070000")
			} else {
				format = format.ReplaceAll("-ZZ", "-0700")
			}
		} else if format.Contains("-Z:Z") {
			if format.Contains("-Z:Z:Z") {
				format = format.ReplaceAll("-Z:Z:Z", "-07:00:00")
			} else {
				format = format.ReplaceAll("-Z:Z", "-07:00")
			}
		} else if format.Contains("ZZ") {
			if format.Contains("ZZZ") {
				format = format.ReplaceAll("ZZZ", "Z070000")
			} else {
				format = format.ReplaceAll("ZZ", "Z0700")
			}
		} else if format.Contains("Z:Z") {
			if format.Contains("Z:Z:Z") {
				format = format.ReplaceAll("Z:Z:Z", "Z07:00:00")
			} else {
				format = format.ReplaceAll("Z:Z", "Z07:00")
			}
		} else if format.Contains("-Z") {
			format = format.ReplaceAll("-Z", "-07")
		} else if format.Contains("Z") {
			format = format.ReplaceAll("Z", "Z07")
		}

	}
	//fmt.Println(format)
	timee, err := time.ParseInLocation(string(format), string(s), loc)

	if err != nil {
		panic(err)
	}

	return timee
}

func (s String) ParseDate(format String) time.Time {
	if format.Contains("YY") {
		if format.Contains("YYYY") {
			format = format.ReplaceAll("YYYY", "2006")
		} else {
			format = format.ReplaceAll("YY", "06")
		}
	}

	if format.Contains("M") {
		if format.Contains("MM") {
			if format.Contains("MMM") {
				if format.Contains("MMMM") {
					format = format.ReplaceAll("MMMM", "January")
					if format.Contains("MMM") {
						format = format.ReplaceAll("MMM", "Jan")
					}
					if format.Contains("MM") {
						format = format.ReplaceAll("MM", "01")
					}
					if format.Contains("M") {
						format = format.ReplaceAll("M", "1")
					}
				} else {
					format = format.ReplaceAll("MMM", "Jan")
					if format.Contains("MM") {
						format = format.ReplaceAll("MM", "01")
					}
					if format.Contains("M") {
						format = format.ReplaceAll("M", "1")
					}
				}
			} else {
				format = format.ReplaceAll("MM", "01")
				if format.Contains("M") {
					format = format.ReplaceAll("M", "1")
				}
			}
		} else {
			format = format.ReplaceAll("M", "1")
		}
	}

	if format.Contains("D") {
		if format.Contains("DD") {
			if format.Contains("DDD") {
				if format.Contains("DDDD") {
					format = format.ReplaceAll("DDDD", "Monday")
					if format.Contains("DDD") {
						format = format.ReplaceAll("DDD", "Mon")
					}
					if format.Contains("DD") {
						format = format.ReplaceAll("DD", "02")
					}
					if format.Contains("D") {
						format = format.ReplaceAll("D", "2")
					}
				} else {
					format = format.ReplaceAll("DDD", "Mon")
					if format.Contains("DD") {
						format = format.ReplaceAll("DD", "02")
					}
					if format.Contains("D") {
						format = format.ReplaceAll("D", "2")
					}
				}
			} else {
				format = format.ReplaceAll("DD", "02")
				if format.Contains("D") {
					format = format.ReplaceAll("D", "2")
				}
			}
		} else {
			format = format.ReplaceAll("D", "2")
		}
	}

	if format.Contains("hh") {
		if format.Contains("hh12") {
			format = format.ReplaceAll("hh12", "03")
		} else {
			format = format.ReplaceAll("hh", "15")
		}
	}
	if format.Contains("h12") {
		format = format.ReplaceAll("h12", "3")
	}

	if format.Contains("m") {
		if format.Contains("mm") {
			format = format.ReplaceAll("mm", "04")
		} else {
			format = format.ReplaceAll("m", "4")
		}
	}

	if format.Contains("s") {
		if format.Contains("ss") {
			format = format.ReplaceAll("ss", "05")
		} else {
			format = format.ReplaceAll("s", "5")
		}
	}

	if format.Contains("Z") {
		if format.Contains("-ZZ") {
			if format.Contains("-ZZZ") {
				format = format.ReplaceAll("-ZZZ", "-070000")
			} else {
				format = format.ReplaceAll("-ZZ", "-0700")
			}
		} else if format.Contains("-Z:Z") {
			if format.Contains("-Z:Z:Z") {
				format = format.ReplaceAll("-Z:Z:Z", "-07:00:00")
			} else {
				format = format.ReplaceAll("-Z:Z", "-07:00")
			}
		} else if format.Contains("ZZ") {
			if format.Contains("ZZZ") {
				format = format.ReplaceAll("ZZZ", "Z070000")
			} else {
				format = format.ReplaceAll("ZZ", "Z0700")
			}
		} else if format.Contains("Z:Z") {
			if format.Contains("Z:Z:Z") {
				format = format.ReplaceAll("Z:Z:Z", "Z07:00:00")
			} else {
				format = format.ReplaceAll("Z:Z", "Z07:00")
			}
		} else if format.Contains("-Z") {
			format = format.ReplaceAll("-Z", "-07")
		} else if format.Contains("Z") {
			format = format.ReplaceAll("Z", "Z07")
		}

	}
	//fmt.Println(format)
	timee, err := time.Parse(string(format), string(s))

	if err != nil {
		panic(err)
	}

	return timee
}

func (s String) Int() int {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return i
}

func (s String) Int32() int32 {
	i, err := strconv.ParseInt(string(s), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(i)
}

func (s String) Int64() int64 {
	i, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (s String) Uint32() uint32 {
	i, err := strconv.ParseUint(string(s), 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(i)
}

func (s String) Uint64() uint64 {
	i, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (s String) Bool() bool {
	b, err := strconv.ParseBool(string(s))
	if err != nil {
		panic(err)
	}
	return b
}

func (s String) Float64() float64 {
	f, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func (s String) Float32() float32 {
	f, err := strconv.ParseFloat(string(s), 32)
	if err != nil {
		panic(err)
	}
	return float32(f)
}

func (s String) Uint() uint {
	u, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(u)
}

func (s String) StripTags() String {
	var strippedstring String = ""
	var i = 0
	var iold = 0
	var ommittext = false
	for _, char := range s {
		if char == '<' {
			if ommittext == false {
				strippedstring += s.substr(iold, i-iold)
			}
			if s.substr(i, 7) == "<script" || s.substr(i, 6) == "<style" {
				ommittext = true
			}
			if ommittext == true {
				iold = i + 1
			}
		}
		if char == '>' {
			//fmt.Println(s.substr(i-8,8))
			//fmt.Println(s.substr(i-7,7))
			if s.substr(i-7, 8) == "/script>" || s.substr(i-6, 7) == "/style>" {
				ommittext = false
			}
			iold = i + 1
		}
		i++
	}
	return strippedstring
}

func (s String) Find(substring String) int {
	return s.Index(substring)
}

func (s String) FindAll(substring String) []int {
	ilength := len(s)
	i := 0
	var intarr []int
	for i < ilength {
		is := s.substr(i, ilength-i).Index(substring)
		//fmt.Println(i)
		//fmt.Println("is")
		//fmt.Println(is)
		if is == -1 {
			i = ilength
		} else {
			i = i + is
			intarr = append(intarr, i)
		}
		i++
	}
	return intarr
}

func (s String) Left(length int) String {
	return s.substr(0, length)
}

func (s String) Right(length int) String {
	return s.substr(len(s)-length, length)
}

func (s String) Reverse() String {
	i := len(s) - 1
	var ss String = ""
	for i >= 0 {
		ss += s.substr(i, 1)
		i--
	}
	return ss
}

type Strint struct {
	s String
	v int
}

func cout(a ...interface{}) {
	fmt.Println(a...)
}

func (s String) WordCount() map[String]int {
	m := make(map[String]int)
	arr := s.Split(" ")
	i := 0
	arrlen := len(arr)
	msort := []Strint{}
	isort := 0
	for i < arrlen {
		arri := arr[i]
		isort = 0
		lenmsort := len(msort)
		if lenmsort == 0 {
			mmsort := Strint{arri, 1}
			msort = append(msort, mmsort)
			//fmt.Println(msort)
		} else {
			isort = 0
			found := false
			for isort < lenmsort {
				if arri == msort[isort].s {
					found = true
					msort[isort].v++
				}
				isort++
			}
			if found == false {
				mmsort := Strint{arri, 1}
				msort = append(msort, mmsort)
				//fmt.Println(msort)
			}
		}
		i++
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(msort, func(i, j int) bool {
		return msort[i].v > msort[j].v
	})
	/*i=0
	  for i < len(msort) {
	          cout(msort[i].s)
	          cout(msort[i].v)
	          i++
	  }*/
	m = make(map[String]int)
	//cout(len(msort))
	i = 0
	for i < len(msort) {
		m[msort[i].s] = msort[i].v
		i++
	}
	//cout(m)
	return m
}

const RandomStringCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var RandomStringseededRand *mrand.Rand = mrand.New(
	mrand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[RandomStringseededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s String) RandomString(length int) String {
	return String(StringWithCharset(length, RandomStringCharset))
}

func (s String) pushTo(ss Strings) bool {
	ss[len(ss)] = s
	return true
}

func (s String) AddLeft(ss String) String {
	return ss + s
}

func (s String) AddRight(ss String) String {
	return s + ss
}

func (s String) AddPos(ss String, pos int) String {
	return s.substr(0, pos) + ss + s.substr(pos, s.len()-pos)
}

func (s String) FindInFiles(strpath String) Strings {
	var ss Strings
	//cout(strpath)
	e := filepath.Walk("/home/k/go/src/String/FindInFilesTestFolder", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		//cout(path)
		return nil
	})
	e = filepath.Walk(string(strpath), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		//cout("Drin")
		//cout(path)
		file, err := os.Open(path)
		if err != nil {
			//cout(err)
			return nil
		} else {
			//cout("tiefer")
			defer file.Close()
			//cout("tiefer2")
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			// This is our buffer now
			//cout("tiefer3")
			i := 1
			for scanner.Scan() {
				var scannertext String = String(scanner.Text())
				scannerarr := scannertext.FindAll(s)
				ii := 0
				scannerarrlen := len(scannerarr)
				//cout(scannerarrlen)
				for ii < scannerarrlen {
					ss = append(ss, String(path+" - Line:"+strconv.Itoa(i)+" - Pos:"+strconv.Itoa(scannerarr[ii])))
					ii++
				}
				i++
			}
		}
		return nil
	})
	if e != nil {
		cout("Error")
		cout(e)
	}
	return ss
}

/*
// Starting Strings section
func (s Strings) len() int {
        return len(s)
}

func (arrs Strings) Filter(test func(String) bool) Strings {
        var ret Strings
        for _, s := range arrs {
                if test(s) {
                        ret = append(ret, s)
                }
        }
        return ret
}

func (s Strings) substr(start, end int) Strings {
        for i := range s {
                s[i] = s[i].substr(start, end)
        }
        return s
}

func (s Strings) ASCIIsubstr(start, end int) Strings {
        for i := range s {
                s[i] = s[i].ASCIIsubstr(start, end)
        }
        return s
}

func (s Strings) Compare(ss Strings) bool {
        b := reflect.DeepEqual(s, ss)
        return b
}

func (s Strings) Contains(ss String) []bool {
        var sss []bool
        for _, val := range s {
                ii := val.Contains(ss)
                sss = append(sss, ii)
        }
        return sss
}

func (s Strings) ContainsFilter(ss String) Strings {
        var sss Strings
        for _, val := range s {
                if val.Contains(ss) {
                        sss = append(sss, val)
                }
        }
        return sss
}

func (s Strings) ContainsAny(chars String) []int {
        var sss []int
        for _, val := range s {
                ii := val.ContainsAny(chars)
                sss = append(sss, ii)
        }
        return sss
}

func (s Strings) ContainsRune(r rune) []bool {
        var sss []bool
        for _, val := range s {
                ii := val.ContainsRune(r)
                sss = append(sss, ii)
        }
        return sss
}

func (s Strings) ContainsRuneFilter(r rune) Strings {
        var sss Strings
        for _, val := range s {
                if val.ContainsRune(r) {
                        sss = append(sss, val)
                }
        }
        return sss
}

func (s Strings) Count(substr String) []int {
        var sss []int
        for _, val := range s {
                ii := val.Count(substr)
                sss = append(sss, ii)
        }
        return sss
}

/*func (s Strings) Fields() map[String]Strings {
        sss := make(map[String]Strings)
        for _, val := range s {
                fields := val.Fields()
                sss = append(sss,fields)
        }
        return sss
}*/

/*func (s Strings) FieldsFunc(f func(rune) bool) map[String]Strings {
        sss := make(map[String]Strings)
        for _, val := range s {
                fields := val.FieldsFunc(f)
                sss = append(sss,fields)
        }
        return sss
}*/
/*
func (s Strings) HasPrefix(ss String) []bool {
        var sss []bool
        for _, val := range s {
                ii := val.HasPrefix(ss)
                sss = append(sss, ii)
        }
        return sss
}

func (s Strings) HasPrefixFilter(ss String) Strings {
        var sss Strings
        for _, val := range s {
                if val.HasPrefix(ss) {
                        sss = append(sss, val)
                }
        }
        return sss
}

func (s Strings) HasSuffix(ss String) []bool {
        var sss []bool
        for _, val := range s {
                ii := val.HasSuffix(ss)
                sss = append(sss, ii)
        }
        return sss
}

func (s Strings) HasSuffixFilter(ss String) Strings {
        var sss Strings
        for _, val := range s {
                if val.HasSuffix(ss) {
                        sss = append(sss, val)
                }
        }
        return sss
}

func (s Strings) Index(ss String) []bool {
        var sss []bool
        for _, val := range s {
                ii := val.HasSuffix(ss)
                sss = append(sss, ii)
        }
        return sss
}
