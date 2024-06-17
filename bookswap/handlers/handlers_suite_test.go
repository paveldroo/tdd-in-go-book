package handlers_test

import (
	"github.com/google/uuid"
	"github.com/paveldroo/tdd-in-go-book/bookswap/db"
	"github.com/paveldroo/tdd-in-go-book/bookswap/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var book db.Book
var svr *httptest.Server

var _ = Describe("Handlers integration", func() {
	var svr *httptest.Server
	BeforeEach(func() {
		book := db.Book{
			ID:     uuid.New().String(),
			Name:   "My first ginkgo test",
			Status: db.Available.String(),
		}
		bs := db.NewBookService([]db.Book{book}, nil)
		ha := handlers.NewHandler(bs, nil)
		svr = httptest.NewServer(http.HandlerFunc(ha.Index))
	})
	AfterEach(func() {
		svr.Close()
	})
	Describe("Index endpoint", func() {
		Context("with one existing book", func() {
			It("should return book", func() {
				r, err := http.Get(svr.URL)
				Expect(err).To(BeNil())
				Expect(r.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}
