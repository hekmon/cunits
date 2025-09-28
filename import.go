package cunits

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const parseRegex = "^([0-9,]+(\\.[0-9]+)?) ?(([KMGTPEZY]i?)?(B|b))$"

var sizeMatch = regexp.MustCompile(parseRegex)

// Parse parses un string representation of a number with a suffix
func Parse(sizeSuffix string) (size Bits, err error) {
	// Does it match ?
	match := sizeMatch.FindSubmatch([]byte(sizeSuffix))
	if match == nil {
		err = fmt.Errorf("string does not match the parsing regex: %s", parseRegex)
		return
	}
	if len(match) < 4 {
		err = fmt.Errorf("regex matching did not return enough groups")
		return
	}
	// Extract number
	num, err := strconv.ParseFloat(strings.Replace(string(match[1]), ",", "", -1), 64)
	if err != nil {
		err = fmt.Errorf("extracted number '%s' can't be parsed as float64: %v", string(match[1]), err)
		return
	}
	// Findout the unit
	switch string(match[3]) {
	case "b":
		size = Bits(num)
	case "B":
		size = ImportInBytes(num)
	// Decimal prefix of bits
	case "Kb":
		size = ImportInKb(num)
	case "Mb":
		size = ImportInMb(num)
	case "Gb":
		size = ImportInGb(num)
	case "Tb":
		size = ImportInTb(num)
	case "Pb":
		size = ImportInPb(num)
	case "Eb":
		size = ImportInEb(num)
	case "Zb":
		size = ImportInZb(num)
	case "Yb":
		size = ImportInYb(num)
	// Binary prefix of bits
	case "Kib":
		size = ImportInKib(num)
	case "Mib":
		size = ImportInMib(num)
	case "Gib":
		size = ImportInGib(num)
	case "Tib":
		size = ImportInTib(num)
	case "Pib":
		size = ImportInPib(num)
	case "Eib":
		size = ImportInEib(num)
	case "Zib":
		size = ImportInZib(num)
	case "Yib":
		size = ImportInYib(num)
	// Decimal prefix of bytes
	case "KB":
		size = ImportInKB(num)
	case "MB":
		size = ImportInMB(num)
	case "GB":
		size = ImportInGB(num)
	case "TB":
		size = ImportInTB(num)
	case "PB":
		size = ImportInPB(num)
	case "EB":
		size = ImportInEB(num)
	case "ZB":
		size = ImportInZB(num)
	case "YB":
		size = ImportInYB(num)
	// Binary prefix of bytes
	case "KiB":
		size = ImportInKiB(num)
	case "MiB":
		size = ImportInMiB(num)
	case "GiB":
		size = ImportInGiB(num)
	case "TiB":
		size = ImportInTiB(num)
	case "PiB":
		size = ImportInPiB(num)
	case "EiB":
		size = ImportInEiB(num)
	case "ZiB":
		size = ImportInZiB(num)
	case "YiB":
		size = ImportInYiB(num)
	// or not
	default:
		err = fmt.Errorf("extracted unit '%s' is unknown", string(match[3]))
	}
	return
}

// ImportInByte imports a number in byte
func ImportInBytes(sizeInBytes float64) Bits {
	return Bits(sizeInBytes * Byte)
}

/*
	Decimal prefix of bits
*/

// ImportInKb imports a number in kilobit
func ImportInKb(sizeInKb float64) Bits {
	return Bits(sizeInKb * Kb)
}

// ImportInMbit imports a number in megabit
func ImportInMb(sizeInMb float64) Bits {
	return Bits(sizeInMb * Mb)
}

// ImportInGbit imports a number in gigabit
func ImportInGb(sizeInGb float64) Bits {
	return Bits(sizeInGb * Gb)
}

// ImportInTb imports a number in terabit
func ImportInTb(sizeInTb float64) Bits {
	return Bits(sizeInTb * Tb)
}

// ImportInPb imports a number in petabit
func ImportInPb(sizeInPb float64) Bits {
	return Bits(sizeInPb * Pb)
}

// ImportInEb imports a number in exabit
func ImportInEb(sizeInEb float64) Bits {
	return Bits(sizeInEb * Eb)
}

// ImportInZb imports a number in zettabit (sizeInZbit better < 1)
func ImportInZb(sizeInZb float64) Bits {
	return Bits(sizeInZb * Zb)
}

// ImportInYb imports a number in yottabit (sizeInYbit better < 1)
func ImportInYb(sizeInYb float64) Bits {
	return Bits(sizeInYb * Yb)
}

/*
	Binary prefix of bits
*/

// ImportInKib imports a number in kibibit
func ImportInKib(sizeInKib float64) Bits {
	return Bits(sizeInKib * Kib)
}

// ImportInMib imports a number in mebibit
func ImportInMib(sizeInMib float64) Bits {
	return Bits(sizeInMib * Mib)
}

// ImportInGib imports a number in gibibit
func ImportInGib(sizeInGib float64) Bits {
	return Bits(sizeInGib * Gib)
}

// ImportInTib imports a number in tebibit
func ImportInTib(sizeInTib float64) Bits {
	return Bits(sizeInTib * Tib)
}

// ImportInPib imports a number in pebibit
func ImportInPib(sizeInPib float64) Bits {
	return Bits(sizeInPib * Pib)
}

// ImportInEib imports a number in exbibit
func ImportInEib(sizeInEib float64) Bits {
	return Bits(sizeInEib * Eib)
}

// ImportInZib imports a number in zebibit (sizeInZibit better < 1)
func ImportInZib(sizeInZib float64) Bits {
	return Bits(sizeInZib * Zib)
}

// ImportInYib imports a number in yobibit (sizeInYibit better < 1)
func ImportInYib(sizeInYib float64) Bits {
	return Bits(sizeInYib * Yib)
}

/*
	Decimal prefix of bytes
*/

// ImportInKB imports a number in kilobyte
func ImportInKB(sizeInKB float64) Bits {
	return Bits(sizeInKB * KB)
}

// ImportInMB imports a number in megabyte
func ImportInMB(sizeInMB float64) Bits {
	return Bits(sizeInMB * MB)
}

// ImportInGB imports a number in gigabyte
func ImportInGB(sizeInGB float64) Bits {
	return Bits(sizeInGB * GB)
}

// ImportInTB imports a number in terabyte
func ImportInTB(sizeInTB float64) Bits {
	return Bits(sizeInTB * TB)
}

// ImportInPB imports a number in petabyte
func ImportInPB(sizeInPB float64) Bits {
	return Bits(sizeInPB * PB)
}

// ImportInEB imports a number in exabyte
func ImportInEB(sizeInEB float64) Bits {
	return Bits(sizeInEB * EB)
}

// ImportInZB imports a number in zettabyte (sizeInZB better < 1)
func ImportInZB(sizeInZB float64) Bits {
	return Bits(sizeInZB * ZB)
}

// ImportInYB imports a number in yottabyte (sizeInYB better < 1)
func ImportInYB(sizeInYB float64) Bits {
	return Bits(sizeInYB * YB)
}

/*
	Binary prefix of bytes
*/

// ImportInKiB imports a number in kibibyte
func ImportInKiB(sizeInKiB float64) Bits {
	return Bits(sizeInKiB * KiB)
}

// ImportInMiB imports a number in mebibyte
func ImportInMiB(sizeInMiB float64) Bits {
	return Bits(sizeInMiB * MiB)
}

// ImportInGiB imports a number in gibibyte
func ImportInGiB(sizeInGiB float64) Bits {
	return Bits(sizeInGiB * GiB)
}

// ImportInTiB imports a number in tebibyte
func ImportInTiB(sizeInTiB float64) Bits {
	return Bits(sizeInTiB * TiB)
}

// ImportInPiB imports a number in pebibyte
func ImportInPiB(sizeInPiB float64) Bits {
	return Bits(sizeInPiB * PiB)
}

// ImportInEiB imports a number in exbibyte (sizeInEiB better < 1)
func ImportInEiB(sizeInEiB float64) Bits {
	return Bits(sizeInEiB * EiB)
}

// ImportInZiB imports a number in zebibyte (sizeInZiB better < 1)
func ImportInZiB(sizeInZiB float64) Bits {
	return Bits(sizeInZiB * ZiB)
}

// ImportInYiB imports a number in yobibyte (sizeInYiB better < 1)
func ImportInYiB(sizeInYiB float64) Bits {
	return Bits(sizeInYiB * YiB)
}
