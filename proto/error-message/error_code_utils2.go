package errormsg

import (
	"context"

	comm "github.com/huydq/proto/gen-go/common/v2"
)

type KObject interface {
	String() string
}

func NewApiError(ctx context.Context, code32 comm.Code) *comm.ApiError {
	httpCode := int32(code32) / 1000
	if httpCode < 100 {
		httpCode = int32(code32)
	}

	if httpCode >= 1000 {
		httpCode = int32(code32) / 10
	}

	waitingFor := int32(10)
	if _, ok := comm.Code_name[int32(code32)]; ok {
		if code32 > comm.Code_ERR_OTHER {
			waitingFor = int32(60)
		}
	} else {
		code32 = comm.Code_ERR_INTERNAL
		waitingFor = int32(120)
	}

	lang, ok := ctx.Value("lang").(string)
	if !ok || lang == "" {
		lang = DEFAULT_LANG
	}

	message := DEFAULT_MESSAGE
	if val, ok := mapMessError[code32]; ok {
		if val2, ok2 := val[lang]; ok2 {
			message = val2
		}
	}

	return &comm.ApiError{
		Code:      httpCode,
		ErrorCode: int32(code32),
		Message:   message,
		WaitFor:   waitingFor,
	}
}

func ErrorMsgFromCode(ctx context.Context, code32 comm.Code) string {
	lang, ok := ctx.Value("lang").(string)
	if !ok {
		lang = DEFAULT_LANG
	}
	message := DEFAULT_MESSAGE
	if val, ok := mapMessError[code32]; ok {
		if val2, ok2 := val[lang]; ok2 {
			message = val2
		}
	}
	return message
}
