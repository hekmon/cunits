package cunits

import "fmt"

// Bit represent a size in bit
type Bit uint64

func (size Bit) String() string {
	// if size >= YiB {
	// 	return size.YiBString()
	// }
	// if size >= ZiB {
	// 	return size.ZiBString()
	// }
	if size >= EiB {
		return size.EiBString()
	}
	if size >= PiB {
		return size.PiBString()
	}
	if size >= TiB {
		return size.TiBString()
	}
	if size >= GiB {
		return size.GiBString()
	}
	if size >= MiB {
		return size.MiBString()
	}
	if size >= KiB {
		return size.KiBString()
	}
	return size.ByteString()
}

/*
	In bytes
*/

// Byte returns the size in byte
func (size Bit) Byte() float64 {
	return float64(size) / 8
}

// ByteString returns the size in byte with unit suffix
func (size Bit) ByteString() string {
	return fmt.Sprintf("%.2f byte", size.Byte())
}

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

// ZBittring returns the size in zettabit with unit suffix (carefull with sub zeros !)
func (size Bit) ZBittring() string {
	return fmt.Sprintf("%f Zbit", size.Zbit())
}

// Ybit returns the size in yottabit
func (size Bit) Ybit() float64 {
	return float64(size) / Ybit
}

// YBittring returns the size in yottabit with unit suffix (carefull with sub zeros !)
func (size Bit) YBittring() string {
	return fmt.Sprintf("%f Ybit", size.Ybit())
}

/*
	Binary prefix of Bit
*/

// Kibit returns the size in kibibit
func (size Bit) Kibit() float64 {
	return float64(size) / Kibit
}

// KibitString returns the size in kibibit with unit suffix
func (size Bit) KibitString() string {
	return fmt.Sprintf("%.2f Kibit", size.Kibit())
}

// Mibit returns the size in mebibit
func (size Bit) Mibit() float64 {
	return float64(size) / Mibit
}

// MibitString returns the size in mebibit with unit suffix
func (size Bit) MibitString() string {
	return fmt.Sprintf("%.2f Mibit", size.Mibit())
}

// Gibit returns the size in gibibit
func (size Bit) Gibit() float64 {
	return float64(size) / Gibit
}

// GibitString returns the size in gibibit with unit suffix
func (size Bit) GibitString() string {
	return fmt.Sprintf("%.2f Gibit", size.Gibit())
}

// Tibit returns the size in tebibit
func (size Bit) Tibit() float64 {
	return float64(size) / Tibit
}

// TibitString returns the size in tebibit with unit suffix
func (size Bit) TibitString() string {
	return fmt.Sprintf("%.2f Tibit", size.Tibit())
}

// Pibit returns the size in pebibit
func (size Bit) Pibit() float64 {
	return float64(size) / Pibit
}

// PibitString returns the size in pebibit with unit suffix
func (size Bit) PibitString() string {
	return fmt.Sprintf("%.2f Pibit", size.Pibit())
}

// Eibit returns the size in exbibit
func (size Bit) Eibit() float64 {
	return float64(size) / Eibit
}

// EibitString returns the size in exbibit with unit suffix
func (size Bit) EibitString() string {
	return fmt.Sprintf("%.2f Eibit", size.Eibit())
}

// Zibit returns the size in zebibit
func (size Bit) Zibit() float64 {
	return float64(size) / Zibit
}

// ZibitString returns the size in zebibit with unit suffix (carefull with sub zeros !)
func (size Bit) ZibitString() string {
	return fmt.Sprintf("%f Zibit", size.Zibit())
}

// Yibit returns the size in yobibit
func (size Bit) Yibit() float64 {
	return float64(size) / Yibit
}

// YibitString returns the size in yobibit with unit suffix (carefull with sub zeros !)
func (size Bit) YibitString() string {
	return fmt.Sprintf("%f Yibit", size.Yibit())
}

/*
	Decimal prefix of bytes
*/

// KB returns the size in kilobyte
func (size Bit) KB() float64 {
	return float64(size) / KB
}

// KBString returns the size in kilobyte with unit suffix
func (size Bit) KBString() string {
	return fmt.Sprintf("%.2f KB", size.KB())
}

// MB returns the size in megabyte
func (size Bit) MB() float64 {
	return float64(size) / MB
}

// MBString returns the size in megabyte with unit suffix
func (size Bit) MBString() string {
	return fmt.Sprintf("%.2f MB", size.MB())
}

// GB returns the size in gigabyte
func (size Bit) GB() float64 {
	return float64(size) / GB
}

// GBString returns the size in gigabyte with unit suffix
func (size Bit) GBString() string {
	return fmt.Sprintf("%.2f GB", size.GB())
}

// TB returns the size in terabyte
func (size Bit) TB() float64 {
	return float64(size) / TB
}

// TBString returns the size in terabyte with unit suffix
func (size Bit) TBString() string {
	return fmt.Sprintf("%.2f TB", size.TB())
}

// PB returns the size in petabyte
func (size Bit) PB() float64 {
	return float64(size) / PB
}

// PBString returns the size in petabyte with unit suffix
func (size Bit) PBString() string {
	return fmt.Sprintf("%.2f PB", size.PB())
}

// EB returns the size in exabyte
func (size Bit) EB() float64 {
	return float64(size) / EB
}

// EBString returns the size in exabyte with unit suffix
func (size Bit) EBString() string {
	return fmt.Sprintf("%.2f EB", size.EB())
}

// ZB returns the size in zettabyte
func (size Bit) ZB() float64 {
	return float64(size) / ZB
}

// ZBString returns the size in zettabyte with unit suffix (carefull with sub zeros !)
func (size Bit) ZBString() string {
	return fmt.Sprintf("%f ZB", size.ZB())
}

// YB returns the size in yottabyte
func (size Bit) YB() float64 {
	return float64(size) / YB
}

// YBString returns the size in yottabyte with unit suffix (carefull with sub zeros !)
func (size Bit) YBString() string {
	return fmt.Sprintf("%f YB", size.YB())
}

/*
	Binary prefix of bytes
*/

// KiB returns the size in kibibyte
func (size Bit) KiB() float64 {
	return float64(size) / KiB
}

// KiBString returns the size in kibibyte with unit suffix
func (size Bit) KiBString() string {
	return fmt.Sprintf("%.2f KiB", size.KiB())
}

// MiB returns the size in mebibyte
func (size Bit) MiB() float64 {
	return float64(size) / MiB
}

// MiBString returns the size in mebibyte with unit suffix
func (size Bit) MiBString() string {
	return fmt.Sprintf("%.2f MiB", size.MiB())
}

// GiB returns the size in gibibyte
func (size Bit) GiB() float64 {
	return float64(size) / GiB
}

// GiBString returns the size in gibibyte with unit suffix
func (size Bit) GiBString() string {
	return fmt.Sprintf("%.2f GiB", size.GiB())
}

// TiB returns the size in tebibyte
func (size Bit) TiB() float64 {
	return float64(size) / TiB
}

// TiBString returns the size in tebibyte wit unit suffix
func (size Bit) TiBString() string {
	return fmt.Sprintf("%.2f TiB", size.TiB())
}

// PiB returns the size in pebibyte
func (size Bit) PiB() float64 {
	return float64(size) / PiB
}

// PiBString returns the size in pebibyte with unit suffix
func (size Bit) PiBString() string {
	return fmt.Sprintf("%.2f PiB", size.PiB())
}

// EiB returns the size in exbibyte
func (size Bit) EiB() float64 {
	return float64(size) / EiB
}

// EiBString returns the size in exbibyte with unit suffix (carefull with sub zeros !)
func (size Bit) EiBString() string {
	return fmt.Sprintf("%f EiB", size.EiB())
}

// ZiB returns the size in zebibyte
func (size Bit) ZiB() float64 {
	return float64(size) / ZiB
}

// ZiBString returns the size in zebibyte with unit suffix (carefull with sub zeros !)
func (size Bit) ZiBString() string {
	return fmt.Sprintf("%f ZiB", size.ZiB())
}

// YiB returns the size in yobibyte
func (size Bit) YiB() float64 {
	return float64(size) / YiB
}

// YiBString returns the size in yobibyte with unit suffix (carefull with sub zeros !)
func (size Bit) YiBString() string {
	return fmt.Sprintf("%f YiB", size.YiB())
}
