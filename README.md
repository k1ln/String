[![Go Report Card](https://goreportcard.com/badge/github.com/k1ln/String)](https://goreportcard.com/report/github.com/k1ln/String)
[![Build Status](https://travis-ci.com/k1ln/String.svg?branch=master)](https://travis-ci.com/k1ln/String)
# String

did you ever ask yourself why go has no Object for string? Sou you can do stuff like: 

```
var s String = "Das" 
bl := s.Contains("s")
```

instead of:

```
s := "Das" 
bl := strings.Contains(s,"s")
```

like in javascript for example. And i saw that String handling in go is not very comfortable. You need to do a lot of things yourself what could be normal for a C developer but feels odd if you come from a javascript/php background. There are scenarios where an Object-Oriented approach is favorable in comparison to a functional one. 

It feels more natural "to me" to use the function after the String the function uses. Readibility "for me" is much better like this. I benchmarked all fuctions in comparison from object-oriented to functional approach and didn't encounter any major performance deficits. You can run the benchmarks yourself:

```
go test -bench=.
```


## All functions from the golang "strings" package were added

so this works:
```
s := String("That") 
if (s.Contains("at")) {
  // do some
}
```

In addition I added some comfort functions which you can see below. If you have any more ideas what one would need as functions in string processing which should be added please let me know. I will add them with Joy! Because Go is a wonderful language. And it could even be better if we work on smoothing out the edges so it gets more accessible.



To use the library the right way please import it with a dot in front so it acts like a local library:

```
import . "String"
```

otherwise you have to write "String." in front of all declarations which is not favorable. 

String (Big capital S) is of type string so you can do anything with it you can do with a string. + and [] work as expected. If you want to use it with other functions don't forget to change the variable type before with string() or s.Tostring().

Functions which are not part of the main golang strings package but add to the functionality and comfort. Functions were only tested on Linux:

### **Substr(start, end int) String**
A normal substring function utf-8 compatible. 
```
var s String = "Hello world!"
var ss = s.Substr(3,5);
```
*result: ss="lo wo"*


### **ASCIIsubstr(start, end int) String**
Same as subtring but only for ASCII Strings => is much faster if you only use ASCII.

### **ParseDateLocal(format String,location String) time.Time**
Date parsing in Golang is freaky. I didn't like it. So wrote it it works more like normally.

Like this:
```
s := String("2009-11-17T20:34:58.651387237+01:00",
format:= String("YYYY-MM-DDThh:mm:ss.999999999Z:Z")
timeinBerlin := s.ParseDateLocal(format,"Europe/Berlin")        
```

Parses local date string into time with normally used syntax: 

__*YY*__: last two digits of year.

__*YYYY*__: full Year.


__*M*__: 1 digit month.

__*MM*__: two digit month.

__*MMM*__: three letter month.

__*MMMM*__: full name month 


__*D*__: 1 digit day.

__*DD*__: two digit day.

__*DDD*__: three letter day.

__*DDDD*__: full name day.


__*hh*__: hours 24 

__*hh12*__: hours 12 


__*m*__: single digit minutes

__*mm*__: double digit minutes


__*s*__: single digit seconds

__*ss*__: double digit seconds 


__*9*__ stays for milliseconds


__*Z*__ as golang Z07

__*ZZ*__ as golang Z0700

__*ZZZ*__ as golang Z070000

__*Z:Z*__ as golang Z07:00

__*Z:Z:Z*__ as golang Z07:00:00


if you need a minus before the Z's just add one.


### **Md5() String**
Generates a md5 String out of provided String.
```
var s String = "The fog is getting thicker!And Leon's getting laaarger!"
var b = s.Md5();
```
*result: b="e2c569be17396eca2a2e3c11578123ed"*

### **Sha1() String**
Generates a sha1 String out of provided String.
```
var s String = "This is an example sentence!"
var b = s.Sha1();
```
*result: b="8b7d314a11b489238e9c8f07b117830b0e823a4a"*

### **AesEncrypt(key String) String**
Creates an  Aes Encrypted String with provided key. Uses CTR as encryption algorith. 

####128-bit:####
```
var s String = "Be good to people and people will be good to you"
s = s.AesEncrypt("52cb693d7e8ff8fecb2d9bee9653954b");
s = s.AesDecrypt("52cb693d7e8ff8fecb2d9bee9653954b");
```
*result: s="Be good to people and people will be good to you"*

####192-bit:####
```
var s String = "Be good to people and people will be good to you"
s = s.AesEncrypt("7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8");
s = s.AesDecrypt("7eabf4e67aa790dba95ef3fd99f87613f1c0741e1d915ea8");
```
*result: s="Be good to people and people will be good to you"*

####265-bit:####
```
var s String = "Be good to people and people will be good to you"
s = s.AesEncrypt("57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c");
s = s.AesDecrypt("57bcd105c2230065fcdd8ff312c201cdb896e28fa0967be2e2c43d61e7b7409c");
```
*result: s="Be good to people and people will be good to you"*


### **AesEncryptByte(key String) []byte**
Returns bytes instead of String. Same as AesEncrypt.

### **AesDecrypt(key String) String** 
Decrypts a AES String with provided key. See example above.

### **AesDecryptByte(key String) String**
Please convert the []byte slice to String first. This will work!

### **GenerateAesKeyHex(length int)**
Generate a random hex-key for Aes-decryption out of a String with the possible lengths of 16,24,32 bytes for 128, 192, 256-bit encryption.
```
import (
	."String"
	"fmt"
)

func main () {
	var s String
	s = s.GenerateAesKeyHex(16)
	fmt.Println(s.Tostring())
}
```

### **IsEmail() bool** 
Checks if String is an email.
```
var s String = "kh@kh.com"
bl = s.IsEmail;
```
*result: bl=true*

### **IsUrl() bool**
Checks if String is an URL.
```
var url String = "https://google.de"
bl = url.IsUrl()
```
*result: bl=true*

### **IsWholeNumber() bool**
Checks if String is a whole number.
```
var num String = "12345"
bl = num.IsWholeNumber()
```
*result: bl=true*

### **IsIpV4() bool**
Checks if String is an IPV4 network adress.
```
var ip String = "192.168.0.1"
bl = ip.IsIpV4()
```
*result: bl=true*

### **IsIpV6() bool**
Checks if String is an IPV6 network adress.
```
var ip String = "2a0a:a545:1728:0:1cb9:a3a8:42e1:d9ca"
bl = ip.IsIpV6()
```
*result: bl=true*

### **IsIp() bool** 
Checks if String is ip adress no matter if ipv4 or ipv6.

### **IsHtmlTag() bool** 
Checks if String is an html tag. 
Checks if String is an IPV6 network adress.
```
var tag String = "<html>"
bl = ip.IsHtmlTag()
```
*result: bl=true*

### **IsPhoneNumber() bool**
Checks if String is a phone number. 
```
var phonenumber String = "01776374859663"
bl = phonenumber.IsHtmlTag()
```
*result: bl=true*

### **IsFilePath() bool**
Checks if String is a filepath. 
```
var path String = "/home/user/go"
bl = path.IsFilePath()
```
*result: bl=true*

### **IsUserName (min int, max int) bool**
Checks if String is a valid user name with min and max number of characters.
```
var user String = "k1ln"
bl = user.IsUserName(5,8)
```
*result: bl=false*

### **IsZipCode(country String) bool**
Checks if String is ZipCode of provided country. Country should be provided in the countrycode-format => "DE => Germany, US => USA, FR => France etc..

```
var zip String = "50968"
bl = zip.IsZipCode("DE")
```
*result: bl=true*

### **IsIban(country String) bool**
Checks if String is correct IBAN-Format doesn't check for 2-check digits at the beginning. Perhaps will add this in the future. Use country code for country like in IsZipCode.

```
var iban String = "DE89370400440532013000"
bl = iban.IsIban("DE")
```
*result: bl=true*

### **PwUpperCase (number int) bool**
Checks if String contains number of uppercases.

```
var pw String = "PassWord"
bl = pw.PwUpperCase(2)
```
*result: bl=true*


### **PwSpecialCase(number int) bool**
Checks if String contains number of special cases.
```
var pw String = "!PwSpecialCase?"
bl = pw.PwSpecialCase(2)
```
*result: bl=true*

### **PwDigits** 
Checks if String contains number of digits. 
```
var pw String = "1PwDigits1"
bl = pw.PwSpecialCase(2)
```
*result: bl=true*

### **PwLowerCase(number int) bool**
Checks if String contains number of lowercases.
```
var pw String = "PwLOWERcASE"
bl = pw.PwSpecialCase(2)
```
*result: bl=true*

### **Get() String**
Get contents of a url address.
```
var url String = "http://google.de"
s := url.Get()
```

### **Json() map[String]interface{}**
A very basic JSON parse function of a String. But there is something TODO here so better write your own. Only for very basic usage. 

### **Open() String**
Open File from provided String path and return String of file. 
```
var path String = "/home/user/text.txt"
s := path.Open()
```

### **Exists() bool**
Checks if String of path exists as url or exists as path on the system.
```
var path String = "/home/user/text.txt"
bl := path.Exists()
```
*result: bl=true*

```
var path String = "https://google.de"
bl := path.Exists()
```
*result: bl=true*


### **GetContents() String**
Get String contents of file or url adress provided. 

### **WriteToFile(path String)**
Write String to provided path as file. 
```
var str String = "Put this text into a file"
str.WriteToFile("/home/user/testfile.txt")
```

### **URLEncode() String**
Encode String in URL-Format => Similar to Query Escape.
```
var str String = "uzgdauzgduaszd$&%$&%$&%$"
encoded := str.URLEncode()
```
*result: encoded=uzgdauzgduaszd%24%26%25%24%26%25%24%26%25%24*

### **URLDecode() String**
Decode URL-encoded String.
```
var str String = "uzgdauzgduaszd%24%26%25%24%26%25%24%26%25%24"
encoded := str.URLEncode()
```
*result: encoded=uzgdauzgduaszd$&%$&%$&%$*

### **Post(url String, contenttype String) String**
Basic Post with contenttype. 
```
var str String = "{\"name\":\"K\"}"
var url := "http://httpbin.org/post"
var contenttype := "application/json"
str.Post(url,contenttype)

```


### **Execute() (String,String)**
Executes a command in command line. Returns result in first String and error in second String.
```
var str String = "echo \"Dies +&& Jenes das\" \"bla&&\""
result,err := str.Execute()
```
*result="Dies +&& Jenes das bla&&"*

### **Php() (String,String)**
Run String as php code and get result, error Strings back. You need php-cli installed.
```
var str String = "echo 'hello';"
result,err := str.Php()
```
*result="hello*

### **Python() (String,String)**
Run String in Python if Python installed. Returns result,error. 
```
var str String = "print(\"hello\")"
result,err := str.Python()
```
*result="hello*

### **Node() (String,String)**
Run String in nodejs if nodejs is installed. Returns result,error.
```
var str String = "console.log(\"hello\")"
result,err := str.Node()
```
*result="hello*


### **Perl() (String,String)**
Run String in Perl if installed. Returns result,error.
```
var str String = "print \"hello\";"
result,err := str.Node()
```
*result="hello*


### **PhpFile() (String,String)**
Run php file provided as path in String. Returns result,error.

### **PythonFile() (String,String)** 
Run python file provided as path in String. Returns result,error.

### **NodeFile() (String,String)**
Run nodejs file provided as path in String. Returns result,error.

### **PerlFile() (String,String)**
Run perl file provided as path in String. Returns result,error.

### **Int() int**
Converts String to int.
```
var str String = "12345"
i := str.Int()
```
*result: i=12345

### **Int32() int32**
Converts String to Int32.

### **Int64() int64**
Converts String to Int64.

### **Uint32() uint32**
Converts String to uint32.

### **Uint64() uint64**
Converts String to uint64.

### **Bool() bool** 
Converts String to bool.
```
var str String = "true"
bl := str.Bool()
```
*result: bl=true

### **Float64() float64**
Converts String to float64. 

### **Float32() float32**
Converts String to float32.

### **Uint() uint**
Converts String to uint.

### **StripTags() String**
Strips HTML-Tags from String. 
```
var str String = "<html><body>This is some text in the body</body></html>"
str = str.StripTags()
```
*result: str="This is some text in the body"


### **Find(substring String) int**
Find first appearance of substring in String.
```
var str String = "hello world"
ifind := str.Find("world")
```
*result: ifind=6

### **FindAll(substring String) []int**
Find all appearances of substring in String.
```
var str String = "hello world world wrold world"
ifindarr := str.FindAll("world")
```
*result: ifindarr=[]int{6, 12, 24}

### **Left(length int) String**
Get number of characters from the left of String.
```
var str String = "hello world"
left := str.Left(4)
```
*result: left="hell"

### **Right(length int) String**
Get number of characters from the right of String.
```
var str String = "hello world"
right := str.Right(4)
```
*result: right="orld"

### **Reverse() String**
Reverse the String. Why in the world you want that but you can do it. 
```
var str String = "Kilian"
reverse := str.Reverse()
```
*result: reverse="nailiK"

### **WordCount() map[String]int**
Count all Words in a String and return as ordered map.
```
var str String = "the the new new of da new the du new the of the of du"
arr := str.WordCount()
```
*result: arr=map[String]int{"the": 5, "new": 4, "of": 3, "du": 2, "da": 1}


### **RandomString(length int) String**
Genrate a random String of length length. From seed "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
```
var str String 
randomstr := str.RandomString(12)
```

### **AddLeft(ss String) String**
Add String to Left of String.
```
var str String = "world"
str = str.AddLeft("hello ")
```
*result: str="hello world"

### **AddRight(ss String) String**
Add String right of String.

```
var str String = "hello "
str = str.AddRight("world")
```
*result: str="hello world"

### **AddPos(ss String,pos int) String**
Add ss String to String at position pos.

```
var str String = "Elvis Presley"
str = str.AddPos("the one and only ",6)
```
*result: str="Elvis the one and only Presley"

### **FindInFiles(strpath String) Strings**
Search for String in all files provided by path.
```
var str String = "needle"
str = str.AddPos("/home/user/haystackfiles")
```

