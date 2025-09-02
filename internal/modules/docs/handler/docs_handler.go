package handler

import (
	"docs-notify/internal/modules/docs/dto"
	"docs-notify/internal/modules/docs/service"
	"docs-notify/internal/utils/exceptions"
	numutils "docs-notify/internal/utils/num_utils"
	"docs-notify/internal/utils/util"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DocsHandler struct {
	docsService *service.DocsService
	// fcmService  *fcm.FCMService
}

func NewDocsHandler(docsService *service.DocsService) *DocsHandler {
	return &DocsHandler{docsService: docsService}
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

func (h *DocsHandler) GetDocById(c echo.Context) error {
	docId, err := numutils.GetUintParam(c, "doc_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	doc, err := h.docsService.GetDocById(docId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(doc))
}

func (h *DocsHandler) GetDocPermissions(c echo.Context) error {
	docId, err := numutils.GetUintParam(c, "doc_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	permissions, err := h.docsService.GetDocPermissions(docId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.WrapResponse(permissions))
}

func (h *DocsHandler) Delete(c echo.Context) error {
	// fmt.Println("Deleting doc with ID:")

	docId, err := numutils.GetUintParam(c, "doc_id")
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	err = h.docsService.DeleteDoc(docId)
	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *DocsHandler) Update(c echo.Context) error {

	var doc dto.DocUpdateDto
	if err := c.Bind(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}
	if err := c.Validate(&doc); err != nil {
		return exceptions.NewResponseError(exceptions.ErrBadRequest, err)
	}

	file, _ := c.FormFile("file")

	updatedDoc, err := h.docsService.UpdateDoc(&doc, file)

	if err != nil {
		return exceptions.NewResponseError(exceptions.ErrInternalServerError, err)
	}
	return c.JSON(http.StatusOK, util.WrapResponse(updatedDoc))
}
