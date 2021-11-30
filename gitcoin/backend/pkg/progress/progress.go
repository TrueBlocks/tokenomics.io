// Copyright 2021 The TrueBlocks Authors. All rights reserved.
// Use of this source code is governed by a license that can
// be found in the LICENSE file.

package progress

type Event uint

const (
	BeforeStart Event = iota
	Start
	Update
	Done
	Error
	Finished
)

type Progress struct {
	Event   Event
	Message string
	Payload interface{}
}

func MakeChan() chan *Progress {
	return make(chan *Progress, 100)
}

func StartMsg(msg string, obj interface{}) *Progress {
	return &Progress{
		Payload: obj,
		Event:   Start,
		Message: msg,
	}
}

func UpdateMsg(msg string, obj interface{}) *Progress {
	return &Progress{
		Payload: obj,
		Event:   Update,
		Message: msg,
	}
}

func DoneMsg(msg string, obj interface{}) *Progress {
	return &Progress{
		Payload: obj,
		Event:   Done,
		Message: msg,
	}
}

func ErrorMsg(msg string, obj interface{}) *Progress {
	return &Progress{
		Payload: obj,
		Event:   Error,
		Message: msg,
	}
}

func FinishedMsg(msg string, obj interface{}) *Progress {
	return &Progress{
		Payload: obj,
		Event:   Finished,
		Message: msg,
	}
}
