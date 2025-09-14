package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

type JsonRetriever[T any] struct{}

func (s JsonRetriever[T]) Init() {}

// RetrieveBytes retrieves a []byte slice from the specified URL.
func (s JsonRetriever[T]) RetrieveBytes(url string) (*[]byte, error) {

	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

// HydrateModel takes a []byte slice and unmarshals it into a model struct.
// The model struct type is defined in the generic type T.
func (s JsonRetriever[T]) HydrateModel(payload []byte) (*T, error) {
	var model T

	err := json.Unmarshal(payload, &model)
	if err != nil {
		return nil, err
	}

	if err := setBytesField(&model, "Raw", payload); err != nil {
		log.Printf("Error setting Raw field: %v\n", err)
	}

	return &model, nil
}

// RetrieveModel retrieves a model struct from the specified URL.
func (s JsonRetriever[T]) RetrieveModel(url string) (*T, error) {
	body, err := s.RetrieveBytes(url)
	if err != nil {
		return nil, err
	}
	return s.HydrateModel(*body)
}

// setBytesField sets a struct field with a given []byte slice, if it exists and is valid.
// This is useful for setting the Raw field on a model if it's present, without forcing an interface on the model struct.
// I tried the interface approach, but it was more verbose and challenging to deal with an embedded pointer struct.
// Could be improved in the future.
func setBytesField(s interface{}, fieldName string, data []byte) error {
	// Must pass a pointer to the struct to allow modification.
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("expected a non-nil pointer to a struct, got %T", s)
	}

	// Dereference the pointer to get the underlying struct value.
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct, got %T", s)
	}

	// Find the field by name.
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field '%s' not found in struct", fieldName)
	}

	// Check if the field is settable and of type []byte.
	if !field.CanSet() {
		return fmt.Errorf("field '%s' is not settable (must be exported)", fieldName)
	}
	if field.Kind() != reflect.Slice || field.Type().Elem().Kind() != reflect.Uint8 {
		return fmt.Errorf("field '%s' is not of type []byte, got %s", fieldName, field.Type())
	}

	// Create a reflect.Value for the []byte data and set the field.
	byteSliceValue := reflect.ValueOf(data)
	field.Set(byteSliceValue)

	return nil
}
