package game

var (
	chairCount = 6
	players    = Players{}
)

// CreateTable ..
func CreateTable() {

}

// SetChairCount ..
func SetChairCount(count int) {
	chairCount = count
}

// FreeChairs ..
func FreeChairs() int {
	return chairCount - len(players)
}

// AddPlayer ..
func AddPlayer(money int) bool {
	if FreeChairs() == 0 {
		return false
	}
	// TODO
	return true
}
