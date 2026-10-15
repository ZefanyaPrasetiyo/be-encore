package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// AppError represents application-level errors with status code
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Predefined application errors
var (
	ErrBadRequest        = NewAppError(fiber.StatusBadRequest, "request tidak valid")
	ErrNotFound          = NewAppError(fiber.StatusNotFound, "data tidak ditemukan")
	ErrConflict          = NewAppError(fiber.StatusConflict, "data sudah ada")
	ErrInternalServer    = NewAppError(fiber.StatusInternalServerError, "terjadi kesalahan pada server")
	ErrUnauthorized      = NewAppError(fiber.StatusUnauthorized, "tidak memiliki akses")
	ErrForbidden         = NewAppError(fiber.StatusForbidden, "akses ditolak")
	ErrInvalidUUID       = NewAppError(fiber.StatusBadRequest, "format UUID tidak valid")
	ErrValidationFailed  = NewAppError(fiber.StatusBadRequest, "validasi request gagal")
	ErrEmailExists       = NewAppError(fiber.StatusConflict, "email sudah digunakan")
	ErrArtistNotFound    = NewAppError(fiber.StatusBadRequest, "artist tidak ditemukan")
	ErrVenueNotFound     = NewAppError(fiber.StatusBadRequest, "venue tidak ditemukan")
	ErrEmptyField        = NewAppError(fiber.StatusBadRequest, "minimal satu field harus diisi")
)

// ErrorHandler is a middleware to handle errors consistently
func ErrorHandler(c fiber.Ctx, err error) error {
	// Handle Fiber's built-in errors
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"message": fiberErr.Message,
		})
	}

	// Handle AppError
	var appErr *AppError
	if errors.As(err, &appErr) {
		return c.Status(appErr.Code).JSON(fiber.Map{
			"message": appErr.Message,
		})
	}

	// Handle GORM errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	// Handle validation errors
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		details := make([]fiber.Map, 0, len(validationErrors))
		for _, fe := range validationErrors {
			details = append(details, fiber.Map{
				"field":   strings.ToLower(fe.Field()),
				"message": getValidationMessage(fe),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validasi request gagal",
			"errors":  details,
		})
	}

	// Handle specific error messages from service layer
	message := err.Error()
	switch message {
	case "email sudah digunakan":
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": message,
		})
	case "artist tidak ditemukan", "venue tidak ditemukan", "minimal satu field harus diisi":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": message,
		})
	}

	// Default: internal server error
	// Log the actual error here if logging is available
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "terjadi kesalahan pada server",
	})
}

// getValidationMessage returns a human-readable validation error message
func getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "wajib diisi"
	case "email":
		return "format email tidak valid"
	case "min":
		return "minimal " + fe.Param() + " karakter"
	case "max":
		return "maksimal " + fe.Param() + " karakter"
	case "len":
		return "harus tepat " + fe.Param() + " karakter"
	case "uuid":
		return "format UUID tidak valid"
	case "oneof":
		return "nilai tidak valid"
	default:
		return "nilai tidak valid: " + fe.Tag()
	}
}