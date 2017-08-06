package stats

import (
	"net/http"
	"fmt"
	"log"
)

type Handler struct {
	r *Repo
}

func NewHandler(r *Repo) *Handler {
	return &Handler{r: r}
}

func (h *Handler) HandleListCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.r.MemberCount()
	// Error talking to mailchimp via rest api, so we'll send the error along.
	if err != nil {
		log.Println("Error loading member count:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error loading member count.")
		return
	}

	c := Count{count: count}
	out, outErr := c.Json()

	// Error generating JSON
	if outErr != nil {
		log.Println("Error generating json:", outErr)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,"Error loading member count")
		return
	}

	w.Write(out)

}