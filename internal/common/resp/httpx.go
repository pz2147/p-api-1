package resp

import (
	"encoding/json"
	"github.com/pz2147/p-api-1/internal/common/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type RespOk struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type RespErr struct {
	Status int              `json:"status"`
	Errors errorx.CodeError `json:"errors"`
}

// Ok writes HTTP 200 OK into w.
func Ok(w http.ResponseWriter, r ...*http.Request) {
	w.WriteHeader(http.StatusOK)
}

// OkJson writes v into w with 200 OK.
func OkJson(w http.ResponseWriter, v interface{}, r ...*http.Request) {
	res := &RespOk{
		Status: 0,
		Data:   v,
	}

	if len(r) == 0 {
		WriteJson(w, http.StatusOK, res, nil)
	} else {
		WriteJson(w, http.StatusOK, res, r[0])
	}
}

//Error writes err into w.
func Error(w http.ResponseWriter, err error, r ...*http.Request) {
	lange := w.Header().Get("Language")
	if lange == "" {
		lange = "en"
	}

	if ec, ok := err.(*errorx.CodeError); ok {
		ec.SetGi18n(lange)
	}

	delete(w.Header(), "Language")

	httpx.Error(w, err)
}

// WriteJson writes v as json string into w with code.
func WriteJson(w http.ResponseWriter, code int, v interface{}, r *http.Request) {
	w.Header().Set(httpx.ContentType, "application/json")
	w.WriteHeader(code)

	if bs, err := json.Marshal(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout && r != nil {
			logx.WithContext(r.Context()).Errorf("write response failed, error: %s", err)
		} else if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		if r != nil {
			logx.WithContext(r.Context()).Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
		} else {
			logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
		}

	} else if r != nil {
		logx.WithContext(r.Context()).Infof("【API-SRV-RESP】%+v", string(bs))
	}
}
