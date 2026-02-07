package domain

import "strconv"

type UserMarker struct {
	// Внешний Айди из Вью Системы - В данном случае телеграмма
	ID string
	//Метка Вью системы, пока только телеграм
	Tag MakerTag
}

func (um *UserMarker) IdInInt64() (int64, bool) {
	intId, err := strconv.Atoi(um.ID)
	if err != nil {
		return 0, false
	}
	return int64(intId), true
}

func (um *UserMarker) IdString() (string, bool) {
	return um.ID, true
}
