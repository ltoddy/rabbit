package rabbit

import (
	"github.com/ltoddy/rabbit/handler"
	"path"
)

type Blueprint struct {
	Prefix string

	get     []*entryHandler
	head    []*entryHandler
	post    []*entryHandler
	put     []*entryHandler
	patch   []*entryHandler
	delete  []*entryHandler
	connect []*entryHandler
	options []*entryHandler
	trace   []*entryHandler
}

func NewBlueprint(prefix string) *Blueprint {
	return &Blueprint{
		Prefix: prefix,
	}
}

func (bp *Blueprint) Get(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.get = append(bp.get, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Head(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.head = append(bp.head, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Post(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.post = append(bp.post, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Put(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.put = append(bp.put, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Patch(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.patch = append(bp.patch, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Delete(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.delete = append(bp.delete, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Connect(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.connect = append(bp.connect, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Options(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.options = append(bp.options, &entryHandler{fullpath, fn})
}

func (bp *Blueprint) Trace(p string, fn handler.HandlerFunction) {
	fullpath := path.Join(bp.Prefix, p)
	bp.trace = append(bp.trace, &entryHandler{fullpath, fn})
}
