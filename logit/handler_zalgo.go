// +build h_zalgo !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/aybabtme/logzalgo"
)

type ZalgoHandler struct {
	BaseHandler
}

func NewZalgoHandler() ZalgoHandler {
	return ZalgoHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config ZalgoHandler) Parse() (Handler, error) {
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.SetFormatter(logzalgo.NewZalgoFormatterrrrrr())
	return h, nil
}

func init() {
	RegisterParser("zalgo", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewZalgoHandler()
		err := meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		handler, err := fconf.Parse()
		if err != nil {
			return nil, fmt.Errorf("init: %v", err)
		}
		return handler, nil
	})
}
