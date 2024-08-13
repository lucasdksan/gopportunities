package model

import "fmt"

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (u *UpdateOpeningRequest) Validate() error {
	if u.Role != "" || u.Company != "" || u.Link != "" || u.Location != "" || u.Remote != nil || u.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
