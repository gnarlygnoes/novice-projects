package game

type CId uint32

var idIncrement CId = 0
var deletedIds = make([]CId, 0)

func NextId() CId {
	l := len(deletedIds)

	if l > 0 {
		id := deletedIds[l-1]
		deletedIds = deletedIds[:l-1]
		return id
	}

	id := idIncrement
	idIncrement++

	return id
}
