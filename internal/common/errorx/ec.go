package errorx

import (
	"fmt"
	"strconv"

	"github.com/gogf/gf/i18n/gi18n"
	"google.golang.org/grpc/status"
)

// ErrorConf 错误配置
type ErrorConf struct {
	Lang      string //语言
	Path      string //语言翻译目录:i18n 一般存于根目录下
	HeaderKey string //http头部关键字. 用于获取Lang值
}

type CodeError struct {
	Code int    `json:"code"` //错误码
	Msg  string `json:"msg"`  //错误码对应key
	Lang string `json:"-"`
	Err  error  `json:"-"`
}

type CodeErrorResp struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// T 本地化
func (e *CodeError) T(t *gi18n.Manager) *CodeErrorResp {
	ctx := gi18n.WithLanguage(nil, e.Lang)
	return &CodeErrorResp{
		Status: e.Code,
		Error:  t.T(ctx, e.Msg),
	}
}

// Wrap 包装rpc error, 不建议在api层直接使用. 除非在i18n里对应翻译
func Wrap(err error) *CodeError {
	s := status.Convert(err)
	return &CodeError{
		Code: int(s.Code()),
		Msg:  s.Message(),
	}
}

// SetGi18n 设置本地化语言
func (e *CodeError) SetGi18n(lang string) {
	e.Lang = lang
}

//错误输出
func (e *CodeError) Error() string {
	return fmt.Sprintf("%s: %s", strconv.Itoa(e.Code), e.Msg)
}

func New(code int, msg string, err error) *CodeError {
	return &CodeError{Code: code, Msg: msg, Err: err}
}
