# go-stemming

A simple implementation of the [porter stemming algorithm](https://tartarus.org/martin/PorterStemmer/def.txt.

-----------------------------------------------------------

## Install

```
go get github.com/techjacker/go-stemmer
```

-----------------------------------------------------------

## Example

```
$ go-stemmer

Enter text: Heigh, my hearts! cheerly, cheerly, my hearts! yare, yare! Take in the topsail.

STEM                 COUNT
--------------------------
heigh                    1
heart                    2
cheer                    2
yare                     2
take                     1
topsail                  1
```

-----------------------------------------------------------

## Tests

```
go test
```
