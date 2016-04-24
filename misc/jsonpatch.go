package main

import (
	"fmt"

	"github.com/evanphx/json-patch"
)

func main() {
	p, _ := jsonpatch.DecodePatch([]byte(`[{"op":"replace","path":"/foo/bar","value":"hello"}]`))
	doc, _ := p.Apply([]byte(`{"foo":{"bar":"hai"}}`))
	fmt.Printf("%s", doc)
}
