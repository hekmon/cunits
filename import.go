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

// ImportFromYbit imports a size in yottabit (sizeInZbit better be negative)
func ImportFromYbit(sizeInYbit float64) Bit {
	return Bit(sizeInYbit * Ybit)
}

/*
	Binary prefix of bits
*/

/*
	Decimal prefix of bytes
*/

/*
	Binary prefix of bytes
*/
