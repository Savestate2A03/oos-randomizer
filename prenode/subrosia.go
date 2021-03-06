package prenode

// subrosia has several large areas which are guaranteed to be traverseable as
// long as you can get there in the first place:
//
// 1. "temple": rosa portal, dance hall, temple, smithy
// 2. "beach": swamp portal, market, beach
// 3. "hide and seek": H&S, mountain portal, spring tower
// 4. "pirate house": village portal, pirates, desert portal
// 5. "furnace": lake portal, furnace, bomb flower
// 6. "bridge": bridge area (large but not visited in any%)
//
// the other locations are isolated and only traverseable with some combination
// of jumping and boulder removal.

// semi-unrelated note:
//
// if you run into the screen where the subrosians steal your feather, they
// will steal it, regardless of whether you have it or what level it is. if
// you had the cape then you'll get it back as a cape, and if you didn't have
// it at all then you'll get it as a feather. it's impossible to get to this
// area without at least a L-1 feather, though.
//
// the main problem is that you lose your fool's ore when you reclaim your
// feather, and if fool's ore is a randomizable item then you could be relying
// on it. one solution would be to disable the event if you already have fool's
// ore. another would be not to randomize fool's ore at all, which is what i'm
// going to do for now.

var subrosiaPrenodes = map[string]*Prenode{
	"temple 1": And("rosa portal"),
	"temple 4": And("beach", "ribbon"),
	"temple 5": And("beach", "jump"),
	"temple 3": And("hide and seek", "pegasus jump L-2"),
	"temple 2": And("bridge", "jump"),

	"beach 1": And("swamp portal"),
	"beach 2": And("hide and seek", "bracelet", "feather L-2"),
	"beach 3": And("hide and seek", "jump", "bracelet", "magnet gloves"),
	"beach 4": And("furnace", "bracelet", "jump"),
	"beach 5": And("furnace", "feather L-2"),
	"beach 6": And("furnace", "jump", "magnet gloves"),
	"beach 7": And("temple", "jump"),

	"hide and seek 1": And("mountain portal"),
	"hide and seek 2": And("pirate house", "jump"),
	"hide and seek 3": And("temple", "pegasus jump L-2"),
	"hide and seek 4": And("bridge", "pegasus jump L-2"),

	"pirate house 1": And("village portal"),
	"pirate house 2": And("desert portal"),
	"pirate house 3": And("hide and seek", "jump"),

	"furnace 1": And("lake portal"),
	"furnace 2": And("beach", "feather L-2"),
	"furnace 3": And("beach", "magnet gloves"),

	"bridge 1": And("temple", "jump"),
	"bridge 2": And("remains portal", "bracelet", "feather L-2"),
	"bridge 3": And("hide and seek", "pegasus jump L-2"),

	"boomerang gift":     AndSlot("temple"),
	"star ore spot":      AndSlot("beach", "shovel"),
	"pirate's bell":      And("temple", "rusty bell"),
	"cross winter tower": Or("hit far switch", "jump"),
	"winter tower":       And("temple", "cross winter tower"),
	"summer tower":       And("beach", "ribbon"),
	"spring tower":       And("hide and seek", "jump"),
	"autumn tower":       And("temple", "jump", "bomb flower"),

	"eruption room": Or("remains portal"),
	"enter d8":      OrStep("d8 portal"),

	// a few places are unaccounted for, but they're irrelevant for now
}
