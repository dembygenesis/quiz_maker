package case_access

import "github.com/dembygenesis/quiz_maker/src/v3/api/db"

func Initialize() Handler {
	repository := NewRepository(db.Handle)
	service := NewService(repository)
	handler := NewHandler(service)

	return handler
}