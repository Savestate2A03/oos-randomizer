package rom

// MutableWord returns a special case of MutableRange with a range of a two
// bytes.
func MutableWord(addr Addr, old, new uint16) MutableRange {
	return MutableRange{
		Addr: addr,
		Old:  []byte{byte(old >> 8), byte(old)},
		New:  []byte{byte(new >> 8), byte(new)},
	}
}

var unusedMutables = map[string]Mutable{
	"d0 key chest":   MutableWord(Addr{0x15, 0x53f4}, 0x3003, 0x3003),
	"d0 rupee chest": MutableWord(Addr{0x15, 0x53f8}, 0x2804, 0x2804),

	"d1 key fall":       MutableWord(Addr{0x0b, 0x466f}, 0x3001, 0x3001),
	"d1 map chest":      MutableWord(Addr{0x15, 0x5418}, 0x3302, 0x3302),
	"d1 compass chest":  MutableWord(Addr{0x15, 0x5404}, 0x3202, 0x3202),
	"d1 gasha chest":    MutableWord(Addr{0x15, 0x5400}, 0x3401, 0x3401),
	"d1 bomb chest":     MutableWord(Addr{0x15, 0x5408}, 0x0300, 0x0300),
	"d1 key chest":      MutableWord(Addr{0x15, 0x540c}, 0x3003, 0x3003),
	"d1 boss key chest": MutableWord(Addr{0x15, 0x5410}, 0x3103, 0x3103),
	"d1 ring chest":     MutableWord(Addr{0x15, 0x5414}, 0x2d04, 0x2d04),

	"d2 5-rupee chest":   MutableWord(Addr{0x15, 0x5438}, 0x2801, 0x2801),
	"d2 key fall":        MutableWord(Addr{0x0b, 0x466f}, 0x3001, 0x3001),
	"d2 compass chest":   MutableWord(Addr{0x15, 0x5434}, 0x3202, 0x3202),
	"d2 map chest":       MutableWord(Addr{0x15, 0x5428}, 0x3302, 0x3302),
	"d2 bomb key chest":  MutableWord(Addr{0x15, 0x542c}, 0x3003, 0x3003),
	"d2 blade key chest": MutableWord(Addr{0x15, 0x5430}, 0x3003, 0x3003),
	"d2 10-rupee chest":  MutableWord(Addr{0x15, 0x541c}, 0x2802, 0x2802),
	"d2 boss key chest":  MutableWord(Addr{0x15, 0x5420}, 0x3103, 0x3103),
}
