package validators

import (
	"fmt"
	"slices"
	"strings"
	"github.com/topengdev/svi_backend/interfaces"
)

func ValidateCreatePost(payload interfaces.ICreatePostDTO) (string, bool) {
	// Title validations
	if payload.Title == "" {
		return "Title is required.", true
	}
	if len([]byte(payload.Title)) < 20 {
		return "Title has to be at least 20 characters.", true
	}
	if len([]byte(payload.Title)) > 200 {
		return "Title can not be more than 200 characters.", true
	}

	// Content validations
	if payload.Content == "" {
		return "Content is required.", true
	}
	if len([]byte(payload.Content)) < 200 {
		return "Content has to be at least 200 characters.", true
	}

	// Category validations
	if payload.Category == "" {
		return "Category is required.", true
	}
	if len([]byte(payload.Category)) < 3 {
		return "Category has to be at least 3 characters.", true
	}
	if len([]byte(payload.Category)) > 100 {
		return "Category can not be more than 100 characters.", true
	}

	// Status validations
	if payload.Status == "" {
		return "Status is required.", true
	}
	if len([]byte(payload.Status)) < 3 {
		return "Status has to be at least 3 characters.", true
	}
	if len([]byte(payload.Status)) > 100 {
		return "Status can not be more than 100 characters.", true
	}
	statusEnum := []string{"Publish", "Draft", "Trash"}
	if !slices.Contains(statusEnum, payload.Status) {
		return fmt.Sprintf("Status has to be one of these ( %s ).", strings.Join(statusEnum, ", ")), true
	}

	return "", false
}

func ValidateUpdatePost(payload interfaces.IUpdatePostDTO) (string, bool) {

	// Id validations
	if payload.Id == 0 {
		return "Id is required.", true
	}

	if payload.Id < 0 {
		return "Id has to be positive integer.", true
	}

	// Title validations
	if payload.Title != "" {
		if len([]byte(payload.Title)) < 20 {
			return "Title has to be at least 20 characters.", true
		}
		if len([]byte(payload.Title)) > 200 {
			return "Title can not be more than 200 characters.", true
		}
	}

	// Content validations
	if payload.Content != "" {
		if len([]byte(payload.Content)) < 200 {
			return "Content has to be at least 200 characters.", true
		}
	}

	// Category validations
	if payload.Category != "" {
		if len([]byte(payload.Category)) < 3 {
			return "Category has to be at least 3 characters.", true
		}
		if len([]byte(payload.Category)) > 100 {
			return "Category can not be more than 100 characters.", true
		}
	}

	// Status validations
	if payload.Status != "" {
		if len([]byte(payload.Status)) < 3 {
			return "Status has to be at least 3 characters.", true
		}
		if len([]byte(payload.Status)) > 100 {
			return "Status can not be more than 100 characters.", true
		}
		statusEnum := []string{"Publish", "Draft", "Trash"}
		if !slices.Contains(statusEnum, payload.Status) {
			return fmt.Sprintf("Status has to be one of these ( %s ).", strings.Join(statusEnum, ", ")), true
		}
	}

	return "", false
}

