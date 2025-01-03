package apiserver

import (
	"encoding/json"
	"mdl/internal/app/model"
	"mdl/internal/app/store"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	// logger zap.Logger
	store store.Store
}

func newServer(store store.Store) *server {

	s := &server{
		// logger: *,
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	admin := s.router.PathPrefix("/admin").Subrouter()

	admin.HandleFunc("/seller", s.HandleSellerCreate()).Methods(http.MethodPost)
	admin.HandleFunc("/sellers", s.HandleSellersGet()).Methods(http.MethodGet)

	admin.HandleFunc("/measureunits", s.HandleMeasureUnitsGet()).Methods(http.MethodGet)

	admin.HandleFunc("/category", s.HandleCategoryCreate()).Methods(http.MethodPost)
	admin.HandleFunc("/categories", s.HandleCategoriesGet()).Methods(http.MethodGet)
	admin.HandleFunc("/category/{id}", s.HandleCategoryGet()).Methods(http.MethodGet)

	admin.HandleFunc("/product", s.HandleProductCreate()).Methods(http.MethodPost)

	// seller := s.router.PathPrefix("seller").Subrouter()
	// seller.HandleFunc("/goods", s.HandleGoodsGet).Methods(http.MethodGet)

}

func (s *server) HandleSellerCreate() http.HandlerFunc {
	type request struct {
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		SecondName  string `json:"second_name"`
		Password    string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		seller := &model.Seller{
			PhoneNumber: req.PhoneNumber,
			FirstName:   req.FirstName,
			SecondName:  req.SecondName,
			Password:    req.Password,
		}
		seller.AdminID = 1

		if err := s.store.Seller().Create(seller); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		seller.Sanitize()

		s.respond(w, r, http.StatusCreated, seller)

	}
}

func (s *server) HandleSellersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminid := 1
		sellers, err := s.store.Seller().GetAll(adminid)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		sexyjson := make(map[string][]*model.Seller)
		sexyjson["data"] = sellers
		s.respond(w, r, http.StatusOK, sexyjson)
	}
}

func (s *server) HandleMeasureUnitsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		measureUnits, err := s.store.MeasureUnits().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		sexyjson := make(map[string][]*model.MeasureUnits)
		sexyjson["data"] = measureUnits
		s.respond(w, r, http.StatusOK, sexyjson)
	}
}

func (s *server) HandleCategoryCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		category := &model.Category{
			Name: req.Name,
		}

		category.AdminID = 1
		if err := s.store.Category().Create(category); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, category)
	}
}

func (s *server) HandleCategoriesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adminid := 1
		categories, err := s.store.Category().GetAll(adminid)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		sexyjson := make(map[string][]*model.Category)
		sexyjson["data"] = categories
		s.respond(w, r, http.StatusOK, sexyjson)
	}
}

func (s *server) HandleCategoryGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		adminID := 1
		products, err := s.store.Product().Get(adminID, id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		sexyjson := make(map[string][]*model.Product)
		sexyjson["data"] = products
		s.respond(w, r, http.StatusOK, sexyjson)

	}
}

func (s *server) HandleProductCreate() http.HandlerFunc {
	type request struct {
		Name           string `json:"name"`
		Price          int    `json:"price"`
		MeasureUnitsID int    `json:"measure_units_id"`
		CategoryID     int    `json:"category_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		product := &model.Product{
			Name:           req.Name,
			Price:          req.Price,
			MeasureUnitsID: req.MeasureUnitsID,
			CategoryID:     req.CategoryID,
		}

		if err := s.store.Product().Create(product); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, product)

	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
