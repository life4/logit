package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Parser func(meta toml.MetaData, primitive toml.Primitive) (*Handler, error)
type Parsers map[string]Parser

var parsers = Parsers{}

func RegisterParser(name string, parser Parser) {
	_, ok := parsers[name]
	if ok {
		panic(fmt.Errorf("parser already registered: %s", name))
	}
	parsers[name] = parser
}

func ParseHandler(meta toml.MetaData, primitive toml.Primitive) (*Handler, error) {
	var config CHandler
	err := meta.PrimitiveDecode(primitive, &config)
	if err != nil {
		return nil, fmt.Errorf("cannot parse format: %v", err)
	}
	parser, ok := parsers[config.Format]
	if !ok {
		return nil, fmt.Errorf("unknown format: %s", config.Format)
	}
	handler, err := parser(meta, primitive)
	if err != nil {
		return nil, fmt.Errorf("cannot parse %s: %v", config.Format, err)
	}
	return handler, nil
}
