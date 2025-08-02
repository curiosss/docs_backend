package validators

import (
	"docs-notify/internal/dto"
)

func ValidateCreateDoc(req *dto.CreateDocRequest) error {
	return validate.Struct(req)
}
