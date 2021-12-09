package handlerr

import (
	"fmt"

	"go.uber.org/zap"
)

type Handlerr struct {
	Logger *zap.Logger
}

func (h Handlerr) Err(s string, e error) {
	eStr := fmt.Sprintf("hdlr :: %s", s)
	h.Logger.Error(eStr, zap.Error(e))
}
