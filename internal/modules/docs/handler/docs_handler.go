package handler

import (
	"docs-notify/internal/modules/docs/dto"
	"docs-notify/internal/modules/docs/service"
	"docs-notify/internal/utils/exceptions"
	"docs-notify/internal/utils/util"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DocsHandler struct {
	docsService *service.DocsService
}

func NewDocsHandler(docsService *service.DocsService) *DocsHandler {
	return &DocsHandler{docsService: docsService}
}

func (h *DocsHandler) Login(c echo.Context) error {
	// Implement the login logic here
	return nil
}
func (h *DocsHandler) Create(c echo.Context) error {

	// Parse form fields
	var doc dto.DocCreateDto
	// Bind JSON fields first (DocName, DocNo, etc.)
	if err := c.Bind(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	doc.UserId = c.Get("user_id").(uint) // from middleware
	if err := c.Validate(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	// Expecting a single file with key "file"
	file, err := c.FormFile("file")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, errors.New("Dokument faýlyny ýükläň"))
	}

	createdDoc, err := h.docsService.CreateDoc(&doc, file)

	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, util.WrapResponse(createdDoc))

}

func (h *DocsHandler) GetDocs(c echo.Context) error {

	var getDocsDto dto.GetDocsDto
	if err := c.Bind(&getDocsDto); err != nil {
		return err
	}
	getDocsDto.UserId = c.Get("user_id").(uint)
	if getDocsDto.Limit == 0 {
		getDocsDto.Limit = 20
	}

	docs, err := h.docsService.GetDocs(getDocsDto)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(docs))
}
