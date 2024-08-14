package test

import (
	"fmt"
	"gopportunities/internal/model"
	"testing"
)

func TestRequest(t *testing.T) {
	t.Run("Test Request", func(t *testing.T) {
		remote := true
		request_1 := model.CreateOpeningRequest{
			Role:     "Fullstack",
			Company:  "Dj Rogerin",
			Location: "São Paulo",
			Remote:   &remote,
			Link:     "https://save.test.com",
			Salary:   1000,
		}
		request_2 := model.CreateOpeningRequest{
			Role:     "Fullstack",
			Company:  "Dj Rogerin",
			Location: "São Paulo",
			Remote:   &remote,
			Link:     "https://save.test.com",
		}
		request_3 := model.CreateOpeningRequest{
			Role:     "Fullstack",
			Company:  "Dj Rogerin",
			Location: "São Paulo",
			Remote:   &remote,
			Salary:   1000,
		}
		request_4 := model.CreateOpeningRequest{
			Role:     "Fullstack",
			Company:  "Dj Rogerin",
			Location: "São Paulo",
			Link:     "https://save.test.com",
			Salary:   1000,
		}
		request_5 := model.CreateOpeningRequest{
			Role:    "Fullstack",
			Company: "Dj Rogerin",
			Remote:  &remote,
			Link:    "https://save.test.com",
			Salary:  1000,
		}
		request_6 := model.CreateOpeningRequest{
			Role:     "Fullstack",
			Location: "São Paulo",
			Remote:   &remote,
			Link:     "https://save.test.com",
			Salary:   1000,
		}
		request_7 := model.CreateOpeningRequest{
			Company:  "Dj Rogerin",
			Location: "São Paulo",
			Remote:   &remote,
			Link:     "https://save.test.com",
			Salary:   1000,
		}

		err_1_result := request_1.Validate()

		if err_1_result != nil {
			t.Fatalf("resultado recebido %s é diferente do esperado %f", "", err_1_result)
		}

		err_2_result := request_2.Validate()

		if err_2_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "salary", "int64"), err_2_result)
		}

		err_3_result := request_3.Validate()

		if err_3_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "link", "string"), err_3_result)
		}

		err_4_result := request_4.Validate()

		if err_4_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "remote", "null"), err_4_result)
		}

		err_5_result := request_5.Validate()

		if err_5_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "location", "string"), err_5_result)
		}

		err_6_result := request_6.Validate()

		if err_6_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "company", "string"), err_6_result)
		}

		err_7_result := request_7.Validate()

		if err_7_result == nil {
			t.Fatalf("resultado recebido %f é diferente do esperado %f", fmt.Errorf("param: %s (type: %s) is required", "role", "string"), err_7_result)
		}
	})
}
