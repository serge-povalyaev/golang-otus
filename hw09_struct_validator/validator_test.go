package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			Response{201, "test"},
			ValidationErrors{ValidationError{"Code", inIntError}},
		},
		{
			Response{200, "test"},
			nil,
		},
		{
			Token{[]byte{65, 66}, []byte{67, 226}, []byte{130, 172}},
			nil,
		},
		{
			App{"test"},
			ValidationErrors{ValidationError{"Version", lenError}},
		},
		{
			App{"5.0.4"},
			nil,
		},
		{
			User{
				ID:     "123",
				Name:   "Петя",
				Age:    55,
				Email:  "petr_petr.ru",
				Role:   "user",
				Phones: []string{"123456", "654321"},
				meta:   []byte("{}"),
			},
			ValidationErrors{
				ValidationError{"ID", lenError},
				ValidationError{"Age", maxError},
				ValidationError{"Email", regexpError},
				ValidationError{"Role", inStringError},
				ValidationError{"Phones[0]", lenError},
				ValidationError{"Phones[1]", lenError},
			},
		},
		{
			User{
				ID:     "123123-123123-123123-123123-123456-1",
				Name:   "Сергей",
				Age:    33,
				Email:  "serge_povalyaev@x5.ru",
				Role:   "admin",
				Phones: []string{"79040913976", "12345678901"},
				meta:   []byte("{}"),
			},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, err, tt.expectedErr)
			_ = tt
		})
	}
}
