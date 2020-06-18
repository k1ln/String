[![Go Report Card](https://goreportcard.com/badge/github.com/k1ln/String)](https://goreportcard.com/report/github.com/k1ln/String)
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


**All functions from the golang "strings" package were added.** 

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

Functions which are not part of the main golang strings package but add to the functionality and comfort functions were only tested on Linux:

### **Substr(start, end int) String**
A normal substring function utf-8 compatible. 

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

### **AesEncrypt(key String) String**
Creates an  Aes Encrypted String with provided key. Uses CTR as encryption algorithm.

### **AesEncryptByte(key String) []byte**
Returns bytes instead of String. Same as AesEncrypt.

### **AesDecrypt(key String) String** 
Decrypts a AES String with provided key.

### **AesDecryptByte(key String) String**
Please convert the []byte slice to String first. This will work!

### **GenerateAesKeyHex(length int)**
Generate a hex-key out of a String with the possible lengths of 16,24,32 for 128, 192, 256-bit encryption.

### **IsEmail() bool** 
Checks if String is an email.

### **IsUrl() bool**
Checks if String is an URL.

### **IsWholeNumber() bool**
Checks if String is a whole number.

### **IsIpV4() bool**
Checks if String is an IPV4 network adress.

### **IsIpV6() bool**
Checks if String is an IPV6 network adress.

### **IsIp() bool** 
Checks if String is ip adress.

### **IsHtmlTag() bool** 
Checks if String is an html tag. 

### **IsPhoneNumber() bool**
Checks if String is a phone number. 

### **IsFilePath() bool**
Checks if String is a filepath. 

### **IsUserName (min int, max int) bool**
Checks if String is a valid user name with min and max number of characters.

### **IsZipCode(country String) bool**
Checks if String is ZipCode of provided country. Country should be provided in the countrycode-format => "DE => Germany, US => USA, FR => France etc..

### **IsIban(country String) bool**
Checks if String is correct IBAN-Format doesn't check for 2-check digits at the beginning. Perhaps will add this in the future. Use country code for country like in IsZipCode.

### **PwUpperCase (number int) bool**
Checks if String contains number of uppercases.

### **PwSpecialCase(number int) bool**
Checks if String contains number of special cases.

### **PwDigits** 
Checks if String contains number of digits. 

### **PwLowerCase(number int) bool**
Checks if String contains number of lowercases.

### **Get() String**
Get contents of a url adress. 

### **Json() map[String]interface{}**
A very basic JSON parse function of a String. But there is something TODO here so better write your own. Only for very basic usage. 

### **Open() String**
Open File from provided String path and return String on file. 

### **Exists() bool**
Checks if String of path exists as url or exists as path on the system.

### **GetContents() String**
Get String contents of file or url adress provided. 

### **WriteToFile(path String)**
Write String to provided path as file. 

### **URLEncode() String**
Encode String in URL-Format => Similar to Query Escape.

### **URLDecode() String**
Decode URL-encoded String.

### **Post(url String, contenttype String) String**
Basic Post with contenttype. 

### **Execute() (String,String)**
Executes a command in command line. Returns result in first String and error in second String.

### **Php() (String,String)**
Run String as php code and get result, error Strings back. You need php-cli installed.

### **Python() (String,String)**
Run String in Python if Python installed. Returns result,error. 

### **Node() (String,String)**
Run String in nodejs if nodejs is installed. Returns result,error.

### **Perl() (String,String)**
Run String in Perl if installed. Returns result,error.

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

### **Float64() float64**
Converts String to float64. 

### **Float32() float32**
Converts String to float32.

### **Uint() uint**
Converts String to uint.

### **StripTags() String**
Strips HTML-Tags from String. 

### **Find(substring String) int**
Find first appearance of substring in String.

### **FindAll(substring String) []int**
Find all appearances of substring in String.

### **Left(length int) String**
Get number of characters from the left of String.

### **Right(length int) String**
Get number of characters from the right of String.

### **Reverse() String**
Reverse the String. Why in the world you want that but you can do it. 

### **WordCount() map[String]int**
Count all Words in a String and return as ordered map.

### **RandomString(length int) String**
Genrate a random String of length length.

### **AddLeft(ss String) String**
Add String to Left of String.

### **AddRight(ss String) String**
Add String right of String.

### **AddPos(ss String,pos int) String**
Add ss String to String at position po.

### **FindInFiles(strpath String) Strings**
Search for String in all files provided by path.

