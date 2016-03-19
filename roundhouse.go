package roundhouse

import (
	"github.com/whosonfirst/go-whosonfirst-utils"
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

	rel := utils.Id2RelPath(wofid)
	return url.Parse(r.Base + rel)
}
