package interfaces

import (
	"context"
)

// @tg http-prefix=v1
// @tg http-server log trace metrics
type MyService interface {
	// @tg http-path=/info/:id
	// @tg http-method=GET
	// @tg http-success=200
	// @tg 400=`неправильное тело или параметры запроса`
	Info(ctx context.Context, id string) (err error)
}
