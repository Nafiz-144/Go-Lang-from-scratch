package product

import (
	"fmt"
	"nafiz/utill"
	"net/http"
	"strconv"
	"sync"
)

var cnt int64

func (h *Handler) Getproducts(w http.ResponseWriter, r *http.Request) {

	//get query parameters
	//get page from query params
	//get limit from query params
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)
	if page <= 0 {

		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		cnt1, err := h.svc.Count()
		if err != nil {
			utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return

		}
		cnt = cnt1

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt2, err := h.svc.Count()
		if err != nil {
			utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return

		}
		fmt.Println(cnt2)

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt3, err := h.svc.Count()
		if err != nil {
			utill.SendError(w, http.StatusInternalServerError, "Internal Server Error")
			return

		}
		fmt.Println(cnt3)

	}()

	//	time.Sleep(8 * time.Second)
	wg.Wait()

	utill.SendPage(w, productList, page, limit, cnt)
}
