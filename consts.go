package cunits

// Byte represent the number of bits a byte contains
const Byte = 8

// Decimal prefix of bits
const (
	// Kb represents a kilobit
	Kb = 1000
	// Mb represents a megabit
	Mb = 1_000_000
	// Gb represents a gigabit
	Gb = 1_000_000_000
	// Tb represents a terabit
	Tb = 1_000_000_000_000
	// Pb represents a petabit
	Pb = 1_000_000_000_000_000
	// Eb represents an exabit
	Eb = 1_000_000_000_000_000_000
	// Zb represents a zettabit (overflows int64)
	Zb = 1_000_000_000_000_000_000_000
	// Yb represents a yottabit (overflows int64)
	Yb = 1_000_000_000_000_000_000_000_000
)

// Binary prefix of bits
const (
	// Kib represents a kibibit
	Kib = 1 << 10
	// Mib represents a mebibit
	Mib = 1 << 20
	// Gib represents a gibibit
	Gib = 1 << 30
	// Tib represents a tebibit
	Tib = 1 << 40
	// Pib represents a pebibit
	Pib = 1 << 50
	// Eib represents an exbibit
	Eib = 1 << 60
	// Zib represents a zebibit (overflows int64)
	Zib = 1 << 70
	// Yib represents a yobibit (overflows int64)
	Yib = 1 << 80
)

// Decimal prefix of bytes
const (
	// KB represents a kilobyte
	KB = Kb * Byte
	// MB represents a megabyte
	MB = Mb * Byte
	// GB represents a gigabyte
	GB = Gb * Byte
	// TB represents a terabyte
	TB = Tb * Byte
	// PB represents a petabyte
	PB = Pb * Byte
	// EB represents an exabyte
	EB = Eb * Byte
	// ZB represents a zettabyte (overflows int64)
	ZB = Zb * Byte
	// YB represents a yottabyte (overflows int64)
	YB = Yb * Byte
)

// Binary prefix of bytes
const (
	// KiB represents a kibibyte
	KiB = Kib * Byte
	// MiB represents a mebibyte
	MiB = Mib * Byte
	// GiB represents a gibibyte
	GiB = Gib * Byte
	// TiB represents a tebibyte
	TiB = Tib * Byte
	// PiB represents a pebibyte
	PiB = Pib * Byte
	// EiB represents an exbibyte (overflows int64)
	EiB = Eib * Byte
	// ZiB represents a zebibyte (overflows int64)
	ZiB = Zb * Byte
	// YiB represents a yobibyte (overflows int64)
	YiB = Yb * Byte
)
