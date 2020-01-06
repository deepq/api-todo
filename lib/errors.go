package lib

import "errors"

var (
	NotFound = errors.New("Specified ID is not found")
	BadItem  = errors.New("Item is corrupted or of wrong type")
	NotValid = errors.New("Parameter 'ID' is not specified")
)
