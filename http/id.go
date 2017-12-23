package http

import (
	"github.com/whosonfirst/go-whosonfirst-roundhouse"
	gohttp "net/http"
	"regexp"
	"strconv"
)

func IDHandler(rh *roundhouse.WOFRoundhouse) (gohttp.Handler, error) {

	re, err := regexp.Compile("/([0-9]+)/?$")

	if err != nil {
		return nil, err
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		url := req.URL
		uri := url.RequestURI()

		m := re.FindStringSubmatch(uri)

		if len(m) == 0 {
			gohttp.Error(rsp, "E_INSUFFICIENT_WOFID", gohttp.StatusBadRequest)
			return
		}

		id := m[1]

		wofid, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			gohttp.Error(rsp, "E_INVALID_WOFID", gohttp.StatusBadRequest)
			return
		}

		u, err := rh.URL(wofid)

		if err != nil {
			gohttp.Error(rsp, "E_IMPOSSIBLE_WOFID", gohttp.StatusBadRequest)
			return
		}

		rsp.Header().Set("Access-Control-Allow-Origin", "*")
		
		gohttp.Redirect(rsp, req, u.String(), gohttp.StatusSeeOther)
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
