package cunits

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

/*
	Binary prefix of bytes
*/
