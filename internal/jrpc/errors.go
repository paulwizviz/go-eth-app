package jrpc

import "errors"

var (
	ErrMarshalRequest       = errors.New("marshal request error")
	ErrUmarshalResponse     = errors.New("unmarshal respond error")
	ErrUnmarshalBlock       = errors.New("unmarshal block error")
	ErrUnmarshalBlockNumber = errors.New("unmarshal block number error")
	ErrSendingRequest       = errors.New("sending request error")
	ErrMismatchResponse     = errors.New("mismatch response error")
)
