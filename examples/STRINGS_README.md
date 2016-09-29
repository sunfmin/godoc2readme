use 'godoc cmd/strings' for documentation on the strings command 



Package strings implements simple functions to manipulate UTF-8 encoded strings.

For information about UTF-8 strings in Go, see <a href="https://blog.golang.org/strings">https://blog.golang.org/strings</a>.




* [Compare](#compare)
* [Contains](#contains)
* [Contains Any](#contains-any)
* [Contains Rune](#contains-rune)
* [Count](#count)
* [Equal Fold](#equal-fold)
* [Fields](#fields)
* [Fields Func](#fields-func)
* [Has Prefix](#has-prefix)
* [Has Suffix](#has-suffix)
* [Index](#index)
* [Index Any](#index-any)
* [Index Byte](#index-byte)
* [Index Func](#index-func)
* [Index Rune](#index-rune)
* [Join](#join)
* [Last Index](#last-index)
* [Last Index Any](#last-index-any)
* [Last Index Byte](#last-index-byte)
* [Last Index Func](#last-index-func)
* [Map](#map)
* [Repeat](#repeat)
* [Replace](#replace)
* [Split](#split)
* [Split After](#split-after)
* [Split After N](#split-after-n)
* [Split N](#split-n)
* [Title](#title)
* [To Lower](#to-lower)
* [To Lower Special](#to-lower-special)
* [To Title](#to-title)
* [To Title Special](#to-title-special)
* [To Upper](#to-upper)
* [To Upper Special](#to-upper-special)
* [Trim](#trim)
* [Trim Func](#trim-func)
* [Trim Left](#trim-left)
* [Trim Left Func](#trim-left-func)
* [Trim Prefix](#trim-prefix)
* [Trim Right](#trim-right)
* [Trim Right Func](#trim-right-func)
* [Trim Space](#trim-space)
* [Trim Suffix](#trim-suffix)
* [Type Reader](#type-reader)
  * [New Reader](#reader-new-reader)
  * [Len](#reader-len)
  * [Read](#reader-read)
  * [Read At](#reader-read-at)
  * [Read Byte](#reader-read-byte)
  * [Read Rune](#reader-read-rune)
  * [Reset](#reader-reset)
  * [Seek](#reader-seek)
  * [Size](#reader-size)
  * [Unread Byte](#reader-unread-byte)
  * [Unread Rune](#reader-unread-rune)
  * [Write To](#reader-write-to)
* [Type Replacer](#type-replacer)
  * [New Replacer](#replacer-new-replacer)
  * [Replace](#replacer-replace)
  * [Write String](#replacer-write-string)




## Compare
``` go
func Compare(a, b string) int
```
Compare returns an integer comparing two strings lexicographically.
The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

Compare is included only for symmetry with package bytes.
It is usually clearer and always faster to use the built-in
string comparison operators ==, <, >, and so on.



## Contains
``` go
func Contains(s, substr string) bool
```
Contains reports whether substr is within s.


```go
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	// Output:
	// true
	// false
	// true
	// true
```

## Contains Any
``` go
func ContainsAny(s, chars string) bool
```
ContainsAny reports whether any Unicode code points in chars are within s.


```go
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	// Output:
	// false
	// true
	// false
	// false
```

## Contains Rune
``` go
func ContainsRune(s string, r rune) bool
```
ContainsRune reports whether the Unicode code point r is within s.



## Count
``` go
func Count(s, sep string) int
```
Count counts the number of non-overlapping instances of sep in s.
If sep is an empty string, Count returns 1 + the number of Unicode code points in s.


```go
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
	// Output:
	// 3
	// 5
```

## Equal Fold
``` go
func EqualFold(s, t string) bool
```
EqualFold reports whether s and t, interpreted as UTF-8 strings,
are equal under Unicode case-folding.


```go
	fmt.Println(strings.EqualFold("Go", "go"))
	// Output: true
```

## Fields
``` go
func Fields(s string) []string
```
Fields splits the string s around each instance of one or more consecutive white space
characters, as defined by unicode.IsSpace, returning an array of substrings of s or an
empty list if s contains only white space.


```go
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	// Output: Fields are: ["foo" "bar" "baz"]
```

## Fields Func
``` go
func FieldsFunc(s string, f func(rune) bool) []string
```
FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
and returns an array of slices of s. If all code points in s satisfy f(c) or the
string is empty, an empty slice is returned.
FieldsFunc makes no guarantees about the order in which it calls f(c).
If f does not return consistent results for a given c, FieldsFunc may crash.


```go
	f := func(c rune) bool {
	    return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
	// Output: Fields are: ["foo1" "bar2" "baz3"]
```

## Has Prefix
``` go
func HasPrefix(s, prefix string) bool
```
HasPrefix tests whether the string s begins with prefix.


```go
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	// Output:
	// true
	// false
	// true
```

## Has Suffix
``` go
func HasSuffix(s, suffix string) bool
```
HasSuffix tests whether the string s ends with suffix.


```go
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	// Output:
	// true
	// false
	// false
	// true
```

## Index
``` go
func Index(s, sep string) int
```
Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.


```go
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	// Output:
	// 4
	// -1
```

## Index Any
``` go
func IndexAny(s, chars string) int
```
IndexAny returns the index of the first instance of any Unicode code point
from chars in s, or -1 if no Unicode code point from chars is present in s.


```go
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	// Output:
	// 2
	// -1
```

## Index Byte
``` go
func IndexByte(s string, c byte) int
```
IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.



## Index Func
``` go
func IndexFunc(s string, f func(rune) bool) int
```
IndexFunc returns the index into s of the first Unicode
code point satisfying f(c), or -1 if none do.


```go
	f := func(c rune) bool {
	    return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
	// Output:
	// 7
	// -1
```

## Index Rune
``` go
func IndexRune(s string, r rune) int
```
IndexRune returns the index of the first instance of the Unicode code point
r, or -1 if rune is not present in s.


```go
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
	// Output:
	// 4
	// -1
```

## Join
``` go
func Join(a []string, sep string) string
```
Join concatenates the elements of a to create a single string. The separator string
sep is placed between elements in the resulting string.


```go
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	// Output: foo, bar, baz
```

## Last Index
``` go
func LastIndex(s, sep string) int
```
LastIndex returns the index of the last instance of sep in s, or -1 if sep is not present in s.


```go
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	// Output:
	// 0
	// 3
	// -1
```

## Last Index Any
``` go
func LastIndexAny(s, chars string) int
```
LastIndexAny returns the index of the last instance of any Unicode code
point from chars in s, or -1 if no Unicode code point from chars is
present in s.



## Last Index Byte
``` go
func LastIndexByte(s string, c byte) int
```
LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.



## Last Index Func
``` go
func LastIndexFunc(s string, f func(rune) bool) int
```
LastIndexFunc returns the index into s of the last
Unicode code point satisfying f(c), or -1 if none do.



## Map
``` go
func Map(mapping func(rune) rune, s string) string
```
Map returns a copy of the string s with all its characters modified
according to the mapping function. If mapping returns a negative value, the character is
dropped from the string with no replacement.


```go
	rot13 := func(r rune) rune {
	    switch {
	    case r >= 'A' && r <= 'Z':
	        return 'A' + (r-'A'+13)%26
	    case r >= 'a' && r <= 'z':
	        return 'a' + (r-'a'+13)%26
	    }
	    return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	// Output: 'Gjnf oevyyvt naq gur fyvgul tbcure...
```

## Repeat
``` go
func Repeat(s string, count int) string
```
Repeat returns a new string consisting of count copies of the string s.


```go
	fmt.Println("ba" + strings.Repeat("na", 2))
	// Output: banana
```

## Replace
``` go
func Replace(s, old, new string, n int) string
```
Replace returns a copy of the string s with the first n
non-overlapping instances of old replaced by new.
If old is empty, it matches at the beginning of the string
and after each UTF-8 sequence, yielding up to k+1 replacements
for a k-rune string.
If n < 0, there is no limit on the number of replacements.


```go
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// Output:
	// oinky oinky oink
	// moo moo moo
```

## Split
``` go
func Split(s, sep string) []string
```
Split slices s into all substrings separated by sep and returns a slice of
the substrings between those separators.
If sep is empty, Split splits after each UTF-8 sequence.
It is equivalent to SplitN with a count of -1.


```go
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	// Output:
	// ["a" "b" "c"]
	// ["" "man " "plan " "canal panama"]
	// [" " "x" "y" "z" " "]
	// [""]
```

## Split After
``` go
func SplitAfter(s, sep string) []string
```
SplitAfter slices s into all substrings after each instance of sep and
returns a slice of those substrings.
If sep is empty, SplitAfter splits after each UTF-8 sequence.
It is equivalent to SplitAfterN with a count of -1.


```go
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	// Output: ["a," "b," "c"]
```

## Split After N
``` go
func SplitAfterN(s, sep string, n int) []string
```
SplitAfterN slices s into substrings after each instance of sep and
returns a slice of those substrings.
If sep is empty, SplitAfterN splits after each UTF-8 sequence.
The count determines the number of substrings to return:


```go
n > 0: at most n substrings; the last substring will be the unsplit remainder.
n == 0: the result is nil (zero substrings)
n < 0: all substrings
```


```go
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
	// Output: ["a," "b,c"]
```

## Split N
``` go
func SplitN(s, sep string, n int) []string
```
SplitN slices s into substrings separated by sep and returns a slice of
the substrings between those separators.
If sep is empty, SplitN splits after each UTF-8 sequence.
The count determines the number of substrings to return:


```go
n > 0: at most n substrings; the last substring will be the unsplit remainder.
n == 0: the result is nil (zero substrings)
n < 0: all substrings
```


```go
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	// Output:
	// ["a" "b,c"]
	// [] (nil = true)
```

## Title
``` go
func Title(s string) string
```
Title returns a copy of the string s with all Unicode letters that begin words
mapped to their title case.

BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.


```go
	fmt.Println(strings.Title("her royal highness"))
	// Output: Her Royal Highness
```

## To Lower
``` go
func ToLower(s string) string
```
ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.


```go
	fmt.Println(strings.ToLower("Gopher"))
	// Output: gopher
```

## To Lower Special
``` go
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
```
ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
lower case, giving priority to the special casing rules.



## To Title
``` go
func ToTitle(s string) string
```
ToTitle returns a copy of the string s with all Unicode letters mapped to their title case.


```go
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
	// Output:
	// LOUD NOISES
	// ХЛЕБ
```

## To Title Special
``` go
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
```
ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
title case, giving priority to the special casing rules.



## To Upper
``` go
func ToUpper(s string) string
```
ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.


```go
	fmt.Println(strings.ToUpper("Gopher"))
	// Output: GOPHER
```

## To Upper Special
``` go
func ToUpperSpecial(_case unicode.SpecialCase, s string) string
```
ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
upper case, giving priority to the special casing rules.



## Trim
``` go
func Trim(s string, cutset string) string
```
Trim returns a slice of the string s with all leading and
trailing Unicode code points contained in cutset removed.


```go
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
	// Output: ["Achtung! Achtung"]
```

## Trim Func
``` go
func TrimFunc(s string, f func(rune) bool) string
```
TrimFunc returns a slice of the string s with all leading
and trailing Unicode code points c satisfying f(c) removed.



## Trim Left
``` go
func TrimLeft(s string, cutset string) string
```
TrimLeft returns a slice of the string s with all leading
Unicode code points contained in cutset removed.



## Trim Left Func
``` go
func TrimLeftFunc(s string, f func(rune) bool) string
```
TrimLeftFunc returns a slice of the string s with all leading
Unicode code points c satisfying f(c) removed.



## Trim Prefix
``` go
func TrimPrefix(s, prefix string) string
```
TrimPrefix returns s without the provided leading prefix string.
If s doesn't start with prefix, s is returned unchanged.


```go
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "Goodbye,")
	s = strings.TrimPrefix(s, "Howdy,")
	fmt.Print("Hello" + s)
	// Output: Hello, world!
```

## Trim Right
``` go
func TrimRight(s string, cutset string) string
```
TrimRight returns a slice of the string s, with all trailing
Unicode code points contained in cutset removed.



## Trim Right Func
``` go
func TrimRightFunc(s string, f func(rune) bool) string
```
TrimRightFunc returns a slice of the string s with all trailing
Unicode code points c satisfying f(c) removed.



## Trim Space
``` go
func TrimSpace(s string) string
```
TrimSpace returns a slice of the string s, with all leading
and trailing white space removed, as defined by Unicode.


```go
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
	// Output: a lone gopher
```

## Trim Suffix
``` go
func TrimSuffix(s, suffix string) string
```
TrimSuffix returns s without the provided trailing suffix string.
If s doesn't end with suffix, s is returned unchanged.


```go
	var s = "Hello, goodbye, etc!"
	s = strings.TrimSuffix(s, "goodbye, etc!")
	s = strings.TrimSuffix(s, "planet")
	fmt.Print(s, "world!")
	// Output: Hello, world!
```



## Type: Reader
``` go
type Reader struct {
    // contains filtered or unexported fields
}
```
A Reader implements the io.Reader, io.ReaderAt, io.Seeker, io.WriterTo,
io.ByteScanner, and io.RuneScanner interfaces by reading
from a string.







### Reader: New Reader
``` go
func NewReader(s string) *Reader
```
NewReader returns a new Reader reading from s.
It is similar to bytes.NewBufferString but more efficient and read-only.





### Reader: Len
``` go
func (r *Reader) Len() int
```
Len returns the number of bytes of the unread portion of the
string.




### Reader: Read
``` go
func (r *Reader) Read(b []byte) (n int, err error)
```



### Reader: Read At
``` go
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```



### Reader: Read Byte
``` go
func (r *Reader) ReadByte() (byte, error)
```



### Reader: Read Rune
``` go
func (r *Reader) ReadRune() (ch rune, size int, err error)
```



### Reader: Reset
``` go
func (r *Reader) Reset(s string)
```
Reset resets the Reader to be reading from s.




### Reader: Seek
``` go
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```
Seek implements the io.Seeker interface.




### Reader: Size
``` go
func (r *Reader) Size() int64
```
Size returns the original length of the underlying string.
Size is the number of bytes available for reading via ReadAt.
The returned value is always the same and is not affected by calls
to any other method.




### Reader: Unread Byte
``` go
func (r *Reader) UnreadByte() error
```



### Reader: Unread Rune
``` go
func (r *Reader) UnreadRune() error
```



### Reader: Write To
``` go
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```
WriteTo implements the io.WriterTo interface.




## Type: Replacer
``` go
type Replacer struct {
    // contains filtered or unexported fields
}
```
Replacer replaces a list of strings with replacements.
It is safe for concurrent use by multiple goroutines.







### Replacer: New Replacer
``` go
func NewReplacer(oldnew ...string) *Replacer
```
NewReplacer returns a new Replacer from a list of old, new string pairs.
Replacements are performed in order, without overlapping matches.


```go
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
	// Output: This is &lt;b&gt;HTML&lt;/b&gt;!
```



### Replacer: Replace
``` go
func (r *Replacer) Replace(s string) string
```
Replace returns a copy of s with all replacements performed.




### Replacer: Write String
``` go
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```
WriteString writes s to w with all replacements performed.





