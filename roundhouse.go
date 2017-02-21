package roundhouse

import (
	"github.com/whosonfirst/go-whosonfirst-uri"
	"net/url"
)

type WOFRoundhouse struct {
	Base string
}

func NewWOFRoundhouse() *WOFRoundhouse {
	r := WOFRoundhouse{"https://whosonfirst.mapzen.com/data/"}
	return &r
}

func (r *WOFRoundhouse) URL(wofid int) (*url.URL, error) {

	rel, err := uri.Id2RelPath(wofid)

	if err != nil {
		return nil, err
	}

	return url.Parse(r.Base + rel)
}
