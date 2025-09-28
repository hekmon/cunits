# ComputerUnits

[![Go Reference](https://pkg.go.dev/badge/github.com/hekmon/cunits/v2.svg)](https://pkg.go.dev/github.com/hekmon/cunits/v2)

ComputerUnits allows to manipulate binary and decimal representations of bits and bytes.

## Installation

```bash
go get -u github.com/hekmon/cunits/v2
```

## Usage

### Constants example

```golang
fmt.Println(cunits.Kb)
fmt.Println(cunits.Kib)
fmt.Println(cunits.KB)
fmt.Println(cunits.KiB)
fmt.Printf("1000 MiB = %f MB\n", float64(1000)*cunits.MiB/cunits.MB)
```

will output:

```text
1000
1024
8000
8192
1000 MiB = 1048.576000 MB
```

### Custom type example

```golang
size := cunits.Bits(58) * cunits.MiB
fmt.Println(size.Mb())
fmt.Println(size.GiBString())
```

will output:

```text
486.539264
0.06 GiB
```

### Parsing example

```golang
size, err := cunits.Parse("7632 MiB")
if err != nil {
    panic(err)
}
fmt.Println(size)
fmt.Println(size.KiB())
fmt.Println(size.KbString())
```

will output:

```text
7.45 GiB
7.815168e+06
64021856.26 Kbit
```
