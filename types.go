package rabbit

import (
	"path"
	"strings"
)

type Any = interface{}

type J = map[string]Any

// url path
type Path string

func (p *Path) Push(subpath Path) {
	*p = Path(path.Join(string(*p), string(subpath)))
}

func (p *Path) Append(subpath Path) Path {
	return Path(path.Join(string(*p), string(subpath)))
}

func (p *Path) Split() []Path {
	ss := strings.Split(string(*p), "/")
	parts := make([]Path, 0, len(ss))

	for _, s := range ss {
		parts = append(parts, Path(s))
	}

	return parts
}
