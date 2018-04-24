package cunits

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const parseRegex = "^([0-9,]+(\\.[0-9]+)?) ?(([KMGTPEZY]i?)?(B|bit))$"

var sizeMatch = regexp.MustCompile(parseRegex)

// Parse parses un string representation of a number with a suffix
func Parse(sizeSuffix string) (size Bit, err error) {
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
	case "bit":
		size = Bit(num)
	case "B":
		size = ImportFromByte(num)
	// Decimal prefix of bits
	case "Kbit":
		size = ImportFromKbit(num)
	case "Mbit":
		size = ImportFromMbit(num)
	case "Gbit":
		size = ImportFromGbit(num)
	case "Tbit":
		size = ImportFromTbit(num)
	case "Pbit":
		size = ImportFromPbit(num)
	case "Ebit":
		size = ImportFromEbit(num)
	case "Zbit":
		size = ImportFromZbit(num)
	case "Ybit":
		size = ImportFromYbit(num)
	// Binary prefix of bits
	case "Kibit":
		size = ImportFromKibit(num)
	case "Mibit":
		size = ImportFromMibit(num)
	case "Gibit":
		size = ImportFromGibit(num)
	case "Tibit":
		size = ImportFromTibit(num)
	case "Pibit":
		size = ImportFromPibit(num)
	case "Eibit":
		size = ImportFromEibit(num)
	case "Zibit":
		size = ImportFromZibit(num)
	case "Yibit":
		size = ImportFromYibit(num)
	// Decimal prefix of bytes
	case "KB":
		size = ImportFromKB(num)
	case "MB":
		size = ImportFromMB(num)
	case "GB":
		size = ImportFromGB(num)
	case "TB":
		size = ImportFromTB(num)
	case "PB":
		size = ImportFromPB(num)
	case "EB":
		size = ImportFromEB(num)
	case "ZB":
		size = ImportFromZB(num)
	case "YB":
		size = ImportFromYB(num)
	// Binary prefix of bytes
	case "KiB":
		size = ImportFromKiB(num)
	case "MiB":
		size = ImportFromMiB(num)
	case "GiB":
		size = ImportFromGiB(num)
	case "TiB":
		size = ImportFromTiB(num)
	case "PiB":
		size = ImportFromPiB(num)
	case "EiB":
		size = ImportFromEiB(num)
	case "ZiB":
		size = ImportFromZiB(num)
	case "YiB":
		size = ImportFromYiB(num)
	// or not
	default:
		err = fmt.Errorf("extracted unit '%s' is unknown", string(match[3]))
	}
	return
}

// ImportFromByte imports a size in byte
func ImportFromByte(sizeInByte float64) Bit {
	return Bit(sizeInByte * Byte)
}

/*
	Decimal prefix of bits
*/

// ImportFromKbit imports a size in kilobit
func ImportFromKbit(sizeInKbit float64) Bit {
	return Bit(sizeInKbit * Kbit)
}

// ImportFromMbit imports a size in megabit
func ImportFromMbit(sizeInMbit float64) Bit {
	return Bit(sizeInMbit * Mbit)
}

// ImportFromGbit imports a size in gigabit
func ImportFromGbit(sizeInGbit float64) Bit {
	return Bit(sizeInGbit * Gbit)
}

// ImportFromTbit imports a size in terabit
func ImportFromTbit(sizeInTbit float64) Bit {
	return Bit(sizeInTbit * Tbit)
}

// ImportFromPbit imports a size in petabit
func ImportFromPbit(sizeInPbit float64) Bit {
	return Bit(sizeInPbit * Pbit)
}

// ImportFromEbit imports a size in exabit
func ImportFromEbit(sizeInEbit float64) Bit {
	return Bit(sizeInEbit * Ebit)
}

// ImportFromZbit imports a size in zettabit (sizeInZbit better be negative)
func ImportFromZbit(sizeInZbit float64) Bit {
	return Bit(sizeInZbit * Zbit)
}

// ImportFromYbit imports a size in yottabit (sizeInYbit better be negative)
func ImportFromYbit(sizeInYbit float64) Bit {
	return Bit(sizeInYbit * Ybit)
}

/*
	Binary prefix of bits
*/

// ImportFromKibit imports a size in kibibit
func ImportFromKibit(sizeInKibit float64) Bit {
	return Bit(sizeInKibit * Kibit)
}

// ImportFromMibit imports a size in mebibit
func ImportFromMibit(sizeInMibit float64) Bit {
	return Bit(sizeInMibit * Mibit)
}

// ImportFromGibit imports a size in gibibit
func ImportFromGibit(sizeInGibit float64) Bit {
	return Bit(sizeInGibit * Gibit)
}

// ImportFromTibit imports a size in tebibit
func ImportFromTibit(sizeInTibit float64) Bit {
	return Bit(sizeInTibit * Tibit)
}

// ImportFromPibit imports a size in pebibit
func ImportFromPibit(sizeInPibit float64) Bit {
	return Bit(sizeInPibit * Pibit)
}

// ImportFromEibit imports a size in exbibit
func ImportFromEibit(sizeInEibit float64) Bit {
	return Bit(sizeInEibit * Eibit)
}

// ImportFromZibit imports a size in zebibit (sizeInZibit better be negative)
func ImportFromZibit(sizeInZibit float64) Bit {
	return Bit(sizeInZibit * Zibit)
}

// ImportFromYibit imports a size in yobibit (sizeInYibit better be negative)
func ImportFromYibit(sizeInYibit float64) Bit {
	return Bit(sizeInYibit * Yibit)
}

/*
	Decimal prefix of bytes
*/

// ImportFromKB imports a size in kilobyte
func ImportFromKB(sizeInKB float64) Bit {
	return Bit(sizeInKB * KB)
}

// ImportFromMB imports a size in megabyte
func ImportFromMB(sizeInMB float64) Bit {
	return Bit(sizeInMB * MB)
}

// ImportFromGB imports a size in gigabyte
func ImportFromGB(sizeInGB float64) Bit {
	return Bit(sizeInGB * GB)
}

// ImportFromTB imports a size in terabyte
func ImportFromTB(sizeInTB float64) Bit {
	return Bit(sizeInTB * TB)
}

// ImportFromPB imports a size in petabyte
func ImportFromPB(sizeInPB float64) Bit {
	return Bit(sizeInPB * PB)
}

// ImportFromEB imports a size in exabyte
func ImportFromEB(sizeInEB float64) Bit {
	return Bit(sizeInEB * EB)
}

// ImportFromZB imports a size in zettabyte (sizeInZB better be negative)
func ImportFromZB(sizeInZB float64) Bit {
	return Bit(sizeInZB * ZB)
}

// ImportFromYB imports a size in yottabyte (sizeInYB better be negative)
func ImportFromYB(sizeInYB float64) Bit {
	return Bit(sizeInYB * YB)
}

/*
	Binary prefix of bytes
*/

// ImportFromKiB imports a size in kilobyte
func ImportFromKiB(sizeInKiB float64) Bit {
	return Bit(sizeInKiB * KiB)
}

// ImportFromMiB imports a size in megabyte
func ImportFromMiB(sizeInMiB float64) Bit {
	return Bit(sizeInMiB * MiB)
}

// ImportFromGiB imports a size in gigabyte
func ImportFromGiB(sizeInGiB float64) Bit {
	return Bit(sizeInGiB * GiB)
}

// ImportFromTiB imports a size in terabyte
func ImportFromTiB(sizeInTiB float64) Bit {
	return Bit(sizeInTiB * TiB)
}

// ImportFromPiB imports a size in petabyte
func ImportFromPiB(sizeInPiB float64) Bit {
	return Bit(sizeInPiB * PiB)
}

// ImportFromEiB imports a size in exabyte (sizeInEiB better be negative)
func ImportFromEiB(sizeInEiB float64) Bit {
	return Bit(sizeInEiB * EiB)
}

// ImportFromZiB imports a size in zettabyte (sizeInZiB better be negative)
func ImportFromZiB(sizeInZiB float64) Bit {
	return Bit(sizeInZiB * ZiB)
}

// ImportFromYiB imports a size in yottabyte (sizeInYiB better be negative)
func ImportFromYiB(sizeInYiB float64) Bit {
	return Bit(sizeInYiB * YiB)
}
