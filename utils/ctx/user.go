package ctx_util

import (
	"context"

	"github.com/kbiits/dealls-take-home-test/domain/entity"
)

func GetUserIDFromCtx(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(entity.CtxUserID).(string)
	return val, ok
}
