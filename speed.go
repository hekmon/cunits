package cunits

import "fmt"

// Speed is variant of Bits with the same functionnalities
// but overloads its representations method to use to speed (decimal prefix of bits) by default
// instead of size (binary prefix of bytes) and by adding a trailling /s to its units.
type Speed struct {
	Bits
}

// String allows direct representation of Speed by calling GetHumanSpeedRepresentation().
func (s Speed) String() string {
	return s.GetHumanSpeedRepresentation()
}

// GetHumanSpeedRepresentation returns the speed in a human readable decimal prefix of bits format
func (s Speed) GetHumanSpeedRepresentation() string {
	value, unit := s.GetHumanSpeedAndSuffix()
	return fmt.Sprintf("%.2f %s", value, unit)
}

// GetHumanSpeedAndSuffix returns the speed in a human readable decimal prefix of bits format and the suffix
func (s Speed) GetHumanSpeedAndSuffix() (float64, string) {
	// if s.Bits >= Zb {
	// 	return s.Bits.Zb(), "Zb/s"
	// }
	// if s.Bits >= Yb {
	// 	return s.Bits.Yb(), "Yb/s"
	// }
	if s.Bits >= Eb {
		return s.Bits.Eb(), "Eb/s"
	}
	if s.Bits >= Pb {
		return s.Bits.Pb(), "Pb/s"
	}
	if s.Bits >= Tb {
		return s.Bits.Tb(), "Tb/s"
	}
	if s.Bits >= Gb {
		return s.Bits.Gb(), "Gb/s"
	}
	if s.Bits >= Mb {
		return s.Bits.Mb(), "Mb/s"
	}
	if s.Bits >= Kb {
		return s.Bits.Kb(), "Kb/s"
	}
	return float64(s.Bits), "b/s"
}

// GetHumanSizeRepresentation returns the speed in a human readable binary prefix of bytes format
func (s Speed) GetHumanSizeRepresentation() string {
	value, unit := s.GetHumanSizeAndSuffix()
	return fmt.Sprintf("%.2f %s", value, unit)
}

// GetHumanSizeAndSuffix returns the speed in a human readable binary prefix of bytes format and the suffix
func (s Speed) GetHumanSizeAndSuffix() (float64, string) {
	// if s.Bits >= YiB {
	// 	return s.Bits.YiB(), "YiB/s"
	// }
	// if s.Bits >= ZiB {
	// 	return s.Bits.ZiB(), "ZiB/s"
	// }
	if s.Bits >= EiB {
		return s.Bits.EiB(), "EiB/s"
	}
	if s.Bits >= PiB {
		return s.Bits.PiB(), "PiB/s"
	}
	if s.Bits >= TiB {
		return s.Bits.TiB(), "TiB/s"
	}
	if s.Bits >= GiB {
		return s.Bits.GiB(), "GiB/s"
	}
	if s.Bits >= MiB {
		return s.Bits.MiB(), "MiB/s"
	}
	if s.Bits >= KiB {
		return s.Bits.KiB(), "KiB/s"
	}
	return s.Bits.Bytes(), "B/s"
}
