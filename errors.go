package main

type retCode int

const (
	Ok = retCode(iota)
	ErrNoData
	ErrServerUnavailable
)

type ret struct {
	Code int
	Msg  string
}

var statusCodes = map[retCode]ret{
	Ok: {
		Code: 200,
		Msg:  "ok",
	},
	ErrNoData: {
		Code: 400,
		Msg:  "no data",
	},
	ErrServerUnavailable: {
		Code: 503,
		Msg:  "failed",
	},
}
