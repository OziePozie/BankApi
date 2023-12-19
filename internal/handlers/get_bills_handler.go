package handlers

import (
	"BankApi/internal/service"
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"strconv"
)

type GETBillsHandler struct {
	useCase *service.GetBillUseCase
}

func NewGETBillsHandler(useCase *service.GetBillUseCase) *GETBillsHandler {
	return &GETBillsHandler{useCase: useCase}
}

type GETBillRequest struct {
}

type GETBillResponse struct {
	id       uuid.UUID
	name     string
	isOpened bool
	userID   uuid.UUID
}

type GETBillParameters struct {
	Page  string `json:"page,omitempty"`
	Items string `json:"itemsPerPage,omitempty"`
	Name  string `json:"name,omitempty"`
}

func (response *GETBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		UserID   uuid.UUID `json:"user"`
		IsOpened bool      `json:"isOpened"`
	}{
		ID:       response.id,
		Name:     response.name,
		IsOpened: response.isOpened,
		UserID:   response.userID,
	})
}

func (handler *GETBillsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		return
	}

	param := GETBillParameters{
		Page:  request.Form.Get("page"),
		Name:  request.Form.Get("name"),
		Items: request.Form.Get("itemsPerPage"),
	}

	page := param.validatePage()
	name := param.validateName()
	items := param.validateItemsPerPage()

	bills, err := handler.useCase.Handle(request.Context(), service.GetBillCommand{
		Page:  page,
		Items: items,
		Name:  name,
	})

	log.Print(err)

	if err != nil {

		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var responses []GETBillResponse

	for _, bill := range *bills {
		response := &GETBillResponse{
			id:       bill.ID(),
			name:     bill.Name(),
			isOpened: bill.IsOpen(),
			userID:   bill.UserID,
		}
		responses = append(responses, *response)
	}
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(responses)

}

func (p GETBillParameters) validateName() string {
	if len(p.Name) > 0 {
		return p.Name
	} else {
		return ""
	}
}

func (p GETBillParameters) validatePage() int {
	page, err := strconv.Atoi(p.Page)

	if err != nil {

		return 0
	} else {
		return page
	}
}

func (p GETBillParameters) validateItemsPerPage() int {

	items, err := strconv.Atoi(p.Items)

	if err != nil {

		return 1
	} else {
		return items
	}
}

//
//var ErrEmptyField = errors.New("field is empty")
//
//var ErrPageParse = errors.New("page parse is failed")
//
//var ErrItemsPerPageParse = errors.New("itemsPerPage parse is failed")
