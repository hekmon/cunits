package cunits

// Decimal prefix of bits
const (
	// Kbit represents a kilobit
	Kbit = 1000
	// Mbit represents a megabit
	Mbit = 1000000
	// Gbit represents a gigabit
	Gbit = 1000000000
	// Tbit represents a terabit
	Tbit = 1000000000000
	// Pbit represents a petabit
	Pbit = 1000000000000000
	// Ebit represents an exabit
	Ebit = 1000000000000000000
	// Zbit represents a zettabit (overflows int64)
	Zbit = 1000000000000000000000
	// Ybit represents a yottabit (overflows int64)
	Ybit = 1000000000000000000000000
)

// Binary prefix of bits
const (
	// Kibit represents a kibibit
	Kibit = 1 << 10
	// Mibit represents a mebibit
	Mibit = 1 << 20
	// Gibit represents a gibibit
	Gibit = 1 << 30
	// Tibit represents a tebibit
	Tibit = 1 << 40
	// Pibit represents a pebibit
	Pibit = 1 << 50
	// Eibit represents an exbibit
	Eibit = 1 << 60
	// Zibit represents a zebibit (overflows int64)
	Zibit = 1 << 70
	// Yibit represents a yobibit (overflows int64)
	Yibit = 1 << 80
)

// Decimal prefix of bytes
const (
	// KB represents a kilobyte
	KB = Kbit * 8
	// MB represents a megabyte
	MB = Mbit * 8
	// GB represents a gigabyte
	GB = Gbit * 8
	// TB represents a terabyte
	TB = Tbit * 8
	// PB represents a petabyte
	PB = Pbit * 8
	// EB represents an exabyte
	EB = Ebit * 8
	// ZB represents a zettabyte (overflows int64)
	ZB = Zbit * 8
	// YB represents a yottabyte (overflows int64)
	YB = Ybit * 8
)

// Binary prefix of bytes
const (
	// KiB represents a kibibyte
	KiB = Kibit * 8
	// MiB represents a mebibyte
	MiB = Mibit * 8
	// GiB represents a gibibyte
	GiB = Gibit * 8
	// TiB represents a tebibyte
	TiB = Tibit * 8
	// PiB represents a pebibyte
	PiB = Pibit * 8
	// EiB represents an exbibyte (overflows int64)
	EiB = Eibit * 8
	// ZiB represents a zebibyte (overflows int64)
	ZiB = Zbit * 8
	// YiB represents a yobibyte (overflows int64)
	YiB = Ybit * 8
)
