package cunits

import "fmt"

// Bits represent a size for a computer
type Bits uint64

/*
	Decimal prefix of bits
*/

// Kbit returns the size in kilobit
func (size Bits) Kbit() float64 {
	return float64(size) / Kbit
}

// KbitString returns the size in kilobit with unit suffix
func (size Bits) KbitString() string {
	return fmt.Sprintf("%.2f Kbit", size.Kbit())
}

// Mbit returns the size in megabit
func (size Bits) Mbit() float64 {
	return float64(size) / Mbit
}

// MbitString returns the size in megabit with unit suffix
func (size Bits) MbitString() string {
	return fmt.Sprintf("%.2f Mbit", size.Mbit())
}

// Gbit returns the size in gigabit
func (size Bits) Gbit() float64 {
	return float64(size) / Gbit
}

// GbitString returns the size in gigabit with unit suffix
func (size Bits) GbitString() string {
	return fmt.Sprintf("%.2f Gbit", size.Gbit())
}

// Tbit returns the size in terabit
func (size Bits) Tbit() float64 {
	return float64(size) / Tbit
}

// TbitString returns the size in terabit with unit suffix
func (size Bits) TbitString() string {
	return fmt.Sprintf("%.2f Tbit", size.Tbit())
}

// Pbit returns the size in petabit
func (size Bits) Pbit() float64 {
	return float64(size) / Pbit
}

// PbitString returns the size in petabit with unit suffix
func (size Bits) PbitString() string {
	return fmt.Sprintf("%.2f Pbit", size.Pbit())
}

// Ebit returns the size in exabit
func (size Bits) Ebit() float64 {
	return float64(size) / Ebit
}

// EbitString returns the size in exabit with unit suffix
func (size Bits) EbitString() string {
	return fmt.Sprintf("%.2f Ebit", size.Ebit())
}

// Zbit returns the size in zettabit
func (size Bits) Zbit() float64 {
	return float64(size) / Zbit
}

// ZbitString returns the size in zettabit with unit suffix
func (size Bits) ZbitString() string {
	return fmt.Sprintf("%.2f Zbit", size.Zbit())
}

// Ybit returns the size in yottabit
func (size Bits) Ybit() float64 {
	return float64(size) / Ybit
}

// YbitString returns the size in yottabit with unit suffix
func (size Bits) YbitString() string {
	return fmt.Sprintf("%.2f Ybit", size.Ybit())
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
