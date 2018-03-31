package cunits

import "fmt"

// Bit represent a size in bit
type Bit uint64

/*
	Decimal prefix of Bit
*/

// Kbit returns the size in kilobit
func (size Bit) Kbit() float64 {
	return float64(size) / Kbit
}

// KBittring returns the size in kilobit with unit suffix
func (size Bit) KBittring() string {
	return fmt.Sprintf("%.2f Kbit", size.Kbit())
}

// Mbit returns the size in megabit
func (size Bit) Mbit() float64 {
	return float64(size) / Mbit
}

// MBittring returns the size in megabit with unit suffix
func (size Bit) MBittring() string {
	return fmt.Sprintf("%.2f Mbit", size.Mbit())
}

// Gbit returns the size in gigabit
func (size Bit) Gbit() float64 {
	return float64(size) / Gbit
}

// GBittring returns the size in gigabit with unit suffix
func (size Bit) GBittring() string {
	return fmt.Sprintf("%.2f Gbit", size.Gbit())
}

// Tbit returns the size in terabit
func (size Bit) Tbit() float64 {
	return float64(size) / Tbit
}

// TBittring returns the size in terabit with unit suffix
func (size Bit) TBittring() string {
	return fmt.Sprintf("%.2f Tbit", size.Tbit())
}

// Pbit returns the size in petabit
func (size Bit) Pbit() float64 {
	return float64(size) / Pbit
}

// PBittring returns the size in petabit with unit suffix
func (size Bit) PBittring() string {
	return fmt.Sprintf("%.2f Pbit", size.Pbit())
}

// Ebit returns the size in exabit
func (size Bit) Ebit() float64 {
	return float64(size) / Ebit
}

// EBittring returns the size in exabit with unit suffix
func (size Bit) EBittring() string {
	return fmt.Sprintf("%.2f Ebit", size.Ebit())
}

// Zbit returns the size in zettabit
func (size Bit) Zbit() float64 {
	return float64(size) / Zbit
}

// ZBittring returns the size in zettabit with unit suffix
func (size Bit) ZBittring() string {
	return fmt.Sprintf("%.2f Zbit", size.Zbit())
}

// Ybit returns the size in yottabit
func (size Bit) Ybit() float64 {
	return float64(size) / Ybit
}

// YBittring returns the size in yottabit with unit suffix
func (size Bit) YBittring() string {
	return fmt.Sprintf("%.2f Ybit", size.Ybit())
}

/*
	Binary prefix of Bit
*/

// Kibit returns the size in kibibit
func (size Bit) Kibit() float64 {
	return float64(size) / Kibit
}

// Mibit returns the size in mebibit
func (size Bit) Mibit() float64 {
	return float64(size) / Mibit
}

// Gibit returns the size in gibibit
func (size Bit) Gibit() float64 {
	return float64(size) / Gibit
}

// Tibit returns the size in tebibit
func (size Bit) Tibit() float64 {
	return float64(size) / Tibit
}

// Pibit returns the size in pebibit
func (size Bit) Pibit() float64 {
	return float64(size) / Pibit
}

// Eibit returns the size in exbibit
func (size Bit) Eibit() float64 {
	return float64(size) / Eibit
}

// Zibit returns the size in zebibit
func (size Bit) Zibit() float64 {
	return float64(size) / Zibit
}

// Yibit returns the size in yobibit
func (size Bit) Yibit() float64 {
	return float64(size) / Yibit
}

/*
	Decimal prefix of bytes
*/

// KB returns the size in kilobyte
func (size Bit) KB() float64 {
	return float64(size) / KB
}

// MB returns the size in megabyte
func (size Bit) MB() float64 {
	return float64(size) / MB
}

// GB returns the size in gigabyte
func (size Bit) GB() float64 {
	return float64(size) / GB
}

// TB returns the size in terabyte
func (size Bit) TB() float64 {
	return float64(size) / TB
}

// PB returns the size in petabyte
func (size Bit) PB() float64 {
	return float64(size) / PB
}

// EB returns the size in exabyte
func (size Bit) EB() float64 {
	return float64(size) / EB
}

// ZB returns the size in zettabyte
func (size Bit) ZB() float64 {
	return float64(size) / ZB
}

// YB returns the size in yottabyte
func (size Bit) YB() float64 {
	return float64(size) / YB
}

/*
	Binary prefix of bytes
*/
