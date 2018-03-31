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

// Kibit returns the size in kibibit
func (size Bits) Kibit() float64 {
	return float64(size) / Kibit
}

// Mibit returns the size in mebibit
func (size Bits) Mibit() float64 {
	return float64(size) / Mibit
}

// Gibit returns the size in gibibit
func (size Bits) Gibit() float64 {
	return float64(size) / Gibit
}

// Tibit returns the size in tebibit
func (size Bits) Tibit() float64 {
	return float64(size) / Tibit
}

// Pibit returns the size in pebibit
func (size Bits) Pibit() float64 {
	return float64(size) / Pibit
}

// Eibit returns the size in exbibit
func (size Bits) Eibit() float64 {
	return float64(size) / Eibit
}

// Zibit returns the size in zebibit
func (size Bits) Zibit() float64 {
	return float64(size) / Zibit
}

// Yibit returns the size in yobibit
func (size Bits) Yibit() float64 {
	return float64(size) / Yibit
}

/*
	Decimal prefix of bytes
*/

/*
	Binary prefix of bytes
*/
