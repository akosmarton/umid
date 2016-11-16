# umID
A unique monotonic ID generator based on nanosecond resolution timestamp written in golang.

## Description
The 10-byte umID value consists of:
- a 8-byte value representing the nanoseconds since the Unix epoch,
- a 2-byte counter, starting with zero.

## Examples
### Generate a New umID as []byte
To generate a new umID as []byte, use NewBytes():
```
b := umid.NewAsBytes()
```
In this example, the value of **b** would be:
```
[]byte{0x14, 0x85, 0x70, 0xe5, 0x79, 0xf9, 0x4f, 0x56, 0x0, 0x0}
```
### Generate a New umID as Hexadecimal String
To generate a new umID as Hexadecimal string, use NewString():
```
s := umid.NewAsString()
```
In this example, the value of **s** would be:
```
"148570e579fa687e0000"
```

## Installation
```
go get -u github.com/akosmarton/umid
```
