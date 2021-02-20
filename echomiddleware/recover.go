package echomiddleware

import (
	"fmt"
	"github.com/caoshuyu/kit/dlog"
	"github.com/labstack/echo/v4"
)

// Recover 捕获 panic.
func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					err = v
				default:
					err = fmt.Errorf("%v", v)
				}
				stack := stackString(callers(4))
				dlog.ERROR("panic", err, "stack", stack)
			}
		}()
		return next(ctx)
	}
}
