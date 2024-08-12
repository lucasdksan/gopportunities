package model

import (
	"fmt"
	"gopportunities/internal/tools"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (m *CreateOpeningRequest) Validate() error {
	if m.Role == "" && m.Company == "" && m.Link == "" && m.Location == "" && m.Remote != nil && m.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}

	if m.Role == "" {
		return tools.ErrParamIsRequired("role", "string")
	}

	if m.Company == "" {
		return tools.ErrParamIsRequired("company", "string")
	}

	if m.Link == "" {
		return tools.ErrParamIsRequired("link", "string")
	}

	if m.Location == "" {
		return tools.ErrParamIsRequired("location", "string")
	}

	if m.Remote == nil {
		return tools.ErrParamIsRequired("remote", "null")
	}

	if m.Salary <= 0 {
		return tools.ErrParamIsRequired("salary", "int64")
	}

	return nil
}
