package datastorage

import "errors"

var (
	ErrDatabase           = databaseError()
	ErrDuplicate          = duplicateError()
	ErrNotFound           = notFoundError()
	ErrViolatesForeignKey = violatesForeignKey()
)

func databaseError() error {
	return errors.New("произошла ошибка при работе с базой данных")
}

func duplicateError() error {
	return errors.New("запись уже существует")
}

func notFoundError() error {
	return errors.New("запись не найдена")
}

func violatesForeignKey() error {
	return errors.New("указан несущестующий идентификатор родителя")
}
