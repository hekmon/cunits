package cunits

// Bits represent a size for a computer
type Bits uint64

/*
	Decimal prefix of bits
*/

// Kbit returns the size in kilobit
func (size Bits) Kbit() float64 {
	return float64(size) / Kbit
}

// Mbit returns the size in megabit
func (size Bits) Mbit() float64 {
	return float64(size) / Mbit
}

// Gbit returns the size in gigabit
func (size Bits) Gbit() float64 {
	return float64(size) / Gbit
}

// Tbit returns the size in terabit
func (size Bits) Tbit() float64 {
	return float64(size) / Tbit
}

// Pbit returns the size in petabit
func (size Bits) Pbit() float64 {
	return float64(size) / Pbit
}

// Ebit returns the size in exabit
func (size Bits) Ebit() float64 {
	return float64(size) / Ebit
}

// Zbit returns the size in zettabit
func (size Bits) Zbit() float64 {
	return float64(size) / Zbit
}

// Ybit returns the size in yottabit
func (size Bits) Ybit() float64 {
	return float64(size) / Ybit
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
