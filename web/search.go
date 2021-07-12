package web

import (
	"github.com/bouncepaw/mycorrhiza/shroom"
	"github.com/bouncepaw/mycorrhiza/user"
	"github.com/bouncepaw/mycorrhiza/views"
	"io"
	"net/http"

	"github.com/bouncepaw/mycorrhiza/util"
)

func initSearch() {
	http.HandleFunc("/primitive-search/", handlerPrimitiveSearch)
}

func handlerPrimitiveSearch(w http.ResponseWriter, rq *http.Request) {
	util.PrepareRq(rq)
	_ = rq.ParseForm()
	var (
		query = rq.FormValue("q")
		u     = user.FromRequest(rq)
	)
	_, _ = io.WriteString(
		w,
		views.BaseHTML(
			"Search: "+query,
			views.PrimitiveSearchHTML(query, shroom.YieldHyphaNamesContainingString),
			u,
		),
	)
}