package cunits

import "fmt"

// Bits represent an arbitrary number of bits
type Bits uint64

// String allows direct reprensetation of Bits by calling GetHumanSizeRepresentation()
func (size Bits) String() string {
	return size.GetHumanSizeRepresentation()
}

// GetHumanSizeAndSuffix returns the size in a human readable binary prefix of bytes format and the suffix
func (size Bits) GetHumanSizeAndSuffix() (float64, string) {
	if size >= EiB {
		return size.EiB(), "EiB"
	}
	if size >= PiB {
		return size.PiB(), "PiB"
	}
	if size >= TiB {
		return size.TiB(), "TiB"
	}
	if size >= GiB {
		return size.GiB(), "GiB"
	}
	if size >= MiB {
		return size.MiB(), "MiB"
	}
	if size >= KiB {
		return size.KiB(), "KiB"
	}
	return size.Bytes(), "B"
}

// GetHumanSizeRepresentation returns the size in a human readable binary prefix of bytes format
func (size Bits) GetHumanSizeRepresentation() string {
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
	return size.BytesString()
}

// GetHumanSpeedAndSuffix returns the size in a human readable decimal prefix of bits format and the suffix
func (size Bits) GetHumanSpeedAndSuffix() (float64, string) {
	if size >= Eb {
		return size.Eb(), "Eb"
	}
	if size >= Pb {
		return size.Pb(), "Pb"
	}
	if size >= Tb {
		return size.Tb(), "Tb"
	}
	if size >= Gb {
		return size.Gb(), "Gb"
	}
	if size >= Mb {
		return size.Mb(), "Mb"
	}
	if size >= Kb {
		return size.Kb(), "Kb"
	}
	return float64(size), "b"
}

// GetHumanSpeedRepresentation returns the size in a human readable decimal prefix of bits format
func (size Bits) GetHumanSpeedRepresentation() string {
	// if size >= Ybit {
	// 	return size.YbitString()
	// }
	// if size >= Zbit {
	// 	return size.ZbitString()
	// }
	if size >= Eb {
		return size.EbString()
	}
	if size >= Pb {
		return size.PbString()
	}
	if size >= Tb {
		return size.TbString()
	}
	if size >= Gb {
		return size.GbString()
	}
	if size >= Mb {
		return size.MbString()
	}
	if size >= Kb {
		return size.KbString()
	}
	return size.BitsString()
}

/*
	Base forms
*/

// BitString returns the size in bit with unit suffix
func (size Bits) BitsString() string {
	return fmt.Sprintf("%d b", size)
}

// Bytes returns the size in bytes
func (size Bits) Bytes() float64 {
	return float64(size) / 8
}

// ByteString returns the size in byte with unit suffix
func (size Bits) BytesString() string {
	return fmt.Sprintf("%.2f B", size.Bytes())
}

/*
	Decimal prefix of Bits
*/

// Kb returns the size in kilobit
func (size Bits) Kb() float64 {
	return float64(size) / Kb
}

// KbString returns the size in kilobit with unit suffix
func (size Bits) KbString() string {
	return fmt.Sprintf("%.2f Kb", size.Kb())
}

// Mb returns the size in megabit
func (size Bits) Mb() float64 {
	return float64(size) / Mb
}

// MbString returns the size in megabit with unit suffix
func (size Bits) MbString() string {
	return fmt.Sprintf("%.2f Mb", size.Mb())
}

// Gb returns the size in gigabit
func (size Bits) Gb() float64 {
	return float64(size) / Gb
}

// GbString returns the size in gigabit with unit suffix
func (size Bits) GbString() string {
	return fmt.Sprintf("%.2f Gb", size.Gb())
}

// Tb returns the size in terabit
func (size Bits) Tb() float64 {
	return float64(size) / Tb
}

// TbString returns the size in terabit with unit suffix
func (size Bits) TbString() string {
	return fmt.Sprintf("%.2f Tb", size.Tb())
}

// Pb returns the size in petabit
func (size Bits) Pb() float64 {
	return float64(size) / Pb
}

// PbString returns the size in petabit with unit suffix
func (size Bits) PbString() string {
	return fmt.Sprintf("%.2f Pb", size.Pb())
}

// Eb returns the size in exabit
func (size Bits) Eb() float64 {
	return float64(size) / Eb
}

// EbString returns the size in exabit with unit suffix
func (size Bits) EbString() string {
	return fmt.Sprintf("%.2f Eb", size.Eb())
}

// Zb returns the size in zettabit
func (size Bits) Zb() float64 {
	return float64(size) / Zb
}

// ZbString returns the size in zettabit with unit suffix (carefull with sub zeros !)
func (size Bits) ZbString() string {
	return fmt.Sprintf("%f Zb", size.Zb())
}

// Yb returns the size in yottabit
func (size Bits) Yb() float64 {
	return float64(size) / Yb
}

// YbString returns the size in yottabit with unit suffix (carefull with sub zeros !)
func (size Bits) YbString() string {
	return fmt.Sprintf("%f Yb", size.Yb())
}

/*
	Binary prefix of Bits
*/

// Kib returns the size in kibibit
func (size Bits) Kib() float64 {
	return float64(size) / Kib
}

// KibString returns the size in kibibit with unit suffix
func (size Bits) KibString() string {
	return fmt.Sprintf("%.2f Kib", size.Kib())
}

// Mib returns the size in mebibit
func (size Bits) Mib() float64 {
	return float64(size) / Mib
}

// MibString returns the size in mebibit with unit suffix
func (size Bits) MibString() string {
	return fmt.Sprintf("%.2f Mib", size.Mib())
}

// Gib returns the size in gibibit
func (size Bits) Gib() float64 {
	return float64(size) / Gib
}

// GibString returns the size in gibibit with unit suffix
func (size Bits) GibString() string {
	return fmt.Sprintf("%.2f Gib", size.Gib())
}

// Tib returns the size in tebibit
func (size Bits) Tib() float64 {
	return float64(size) / Tib
}

// TibString returns the size in tebibit with unit suffix
func (size Bits) TibString() string {
	return fmt.Sprintf("%.2f Tib", size.Tib())
}

// Pib returns the size in pebibit
func (size Bits) Pib() float64 {
	return float64(size) / Pib
}

// PibString returns the size in pebibit with unit suffix
func (size Bits) PibString() string {
	return fmt.Sprintf("%.2f Pib", size.Pib())
}

// Eib returns the size in exbibit
func (size Bits) Eib() float64 {
	return float64(size) / Eib
}

// EibString returns the size in exbibit with unit suffix
func (size Bits) EibString() string {
	return fmt.Sprintf("%.2f Eib", size.Eib())
}

// Zib returns the size in zebibit
func (size Bits) Zib() float64 {
	return float64(size) / Zib
}

// ZibString returns the size in zebibit with unit suffix (carefull with sub zeros !)
func (size Bits) ZibString() string {
	return fmt.Sprintf("%f Zib", size.Zib())
}

// Yib returns the size in yobibit
func (size Bits) Yib() float64 {
	return float64(size) / Yib
}

// YibString returns the size in yobibit with unit suffix (carefull with sub zeros !)
func (size Bits) YibitString() string {
	return fmt.Sprintf("%f Yib", size.Yib())
}

/*
	Decimal prefix of bytes
*/

// KB returns the size in kilobyte
func (size Bits) KB() float64 {
	return float64(size) / KB
}

// KBString returns the size in kilobyte with unit suffix
func (size Bits) KBString() string {
	return fmt.Sprintf("%.2f KB", size.KB())
}

// MB returns the size in megabyte
func (size Bits) MB() float64 {
	return float64(size) / MB
}

// MBString returns the size in megabyte with unit suffix
func (size Bits) MBString() string {
	return fmt.Sprintf("%.2f MB", size.MB())
}

// GB returns the size in gigabyte
func (size Bits) GB() float64 {
	return float64(size) / GB
}

// GBString returns the size in gigabyte with unit suffix
func (size Bits) GBString() string {
	return fmt.Sprintf("%.2f GB", size.GB())
}

// TB returns the size in terabyte
func (size Bits) TB() float64 {
	return float64(size) / TB
}

// TBString returns the size in terabyte with unit suffix
func (size Bits) TBString() string {
	return fmt.Sprintf("%.2f TB", size.TB())
}

// PB returns the size in petabyte
func (size Bits) PB() float64 {
	return float64(size) / PB
}

// PBString returns the size in petabyte with unit suffix
func (size Bits) PBString() string {
	return fmt.Sprintf("%.2f PB", size.PB())
}

// EB returns the size in exabyte
func (size Bits) EB() float64 {
	return float64(size) / EB
}

// EBString returns the size in exabyte with unit suffix
func (size Bits) EBString() string {
	return fmt.Sprintf("%.2f EB", size.EB())
}

// ZB returns the size in zettabyte
func (size Bits) ZB() float64 {
	return float64(size) / ZB
}

// ZBString returns the size in zettabyte with unit suffix (carefull with sub zeros !)
func (size Bits) ZBString() string {
	return fmt.Sprintf("%f ZB", size.ZB())
}

// YB returns the size in yottabyte
func (size Bits) YB() float64 {
	return float64(size) / YB
}

// YBString returns the size in yottabyte with unit suffix (carefull with sub zeros !)
func (size Bits) YBString() string {
	return fmt.Sprintf("%f YB", size.YB())
}

/*
	Binary prefix of bytes
*/

// KiB returns the size in kibibyte
func (size Bits) KiB() float64 {
	return float64(size) / KiB
}

// KiBString returns the size in kibibyte with unit suffix
func (size Bits) KiBString() string {
	return fmt.Sprintf("%.2f KiB", size.KiB())
}

// MiB returns the size in mebibyte
func (size Bits) MiB() float64 {
	return float64(size) / MiB
}

// MiBString returns the size in mebibyte with unit suffix
func (size Bits) MiBString() string {
	return fmt.Sprintf("%.2f MiB", size.MiB())
}

// GiB returns the size in gibibyte
func (size Bits) GiB() float64 {
	return float64(size) / GiB
}

// GiBString returns the size in gibibyte with unit suffix
func (size Bits) GiBString() string {
	return fmt.Sprintf("%.2f GiB", size.GiB())
}

// TiB returns the size in tebibyte
func (size Bits) TiB() float64 {
	return float64(size) / TiB
}

// TiBString returns the size in tebibyte wit unit suffix
func (size Bits) TiBString() string {
	return fmt.Sprintf("%.2f TiB", size.TiB())
}

// PiB returns the size in pebibyte
func (size Bits) PiB() float64 {
	return float64(size) / PiB
}

// PiBString returns the size in pebibyte with unit suffix
func (size Bits) PiBString() string {
	return fmt.Sprintf("%.2f PiB", size.PiB())
}

// EiB returns the size in exbibyte
func (size Bits) EiB() float64 {
	return float64(size) / EiB
}

// EiBString returns the size in exbibyte with unit suffix (carefull with sub zeros !)
func (size Bits) EiBString() string {
	return fmt.Sprintf("%f EiB", size.EiB())
}

// ZiB returns the size in zebibyte
func (size Bits) ZiB() float64 {
	return float64(size) / ZiB
}

// ZiBString returns the size in zebibyte with unit suffix (carefull with sub zeros !)
func (size Bits) ZiBString() string {
	return fmt.Sprintf("%f ZiB", size.ZiB())
}

// YiB returns the size in yobibyte
func (size Bits) YiB() float64 {
	return float64(size) / YiB
}

// YiBString returns the size in yobibyte with unit suffix (carefull with sub zeros !)
func (size Bits) YiBString() string {
	return fmt.Sprintf("%f YiB", size.YiB())
}
