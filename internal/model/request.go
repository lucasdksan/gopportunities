package model

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name, value string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, value)
}

func (m *CreateOpeningRequest) Validate() error {
	if m.Role == "" && m.Company == "" && m.Link == "" && m.Location == "" && m.Remote != nil && m.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}

	if m.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if m.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if m.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if m.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if m.Remote == nil {
		return errParamIsRequired("remote", "null")
	}

	if m.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}
