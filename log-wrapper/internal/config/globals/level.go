package globals

var LevelMap map[string]int

func init() {
	LevelMap = make(map[string]int)

	LevelMap["FATAL"] = 2
	LevelMap["ERROR"] = 3
	LevelMap["WARN"] = 4
	LevelMap["INFO"] = 6
	LevelMap["DEBUG"] = 7
}
