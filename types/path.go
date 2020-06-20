package types

import "path"

// url path
type Path string

func (p *Path) Push(subpath Path) {
	*p = Path(path.Join(string(*p), string(subpath)))
}

func (p *Path) Append(subpath Path) Path {
	return Path(path.Join(string(*p), string(subpath)))
}
