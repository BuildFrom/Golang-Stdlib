package todoapp

import (
	"encoding/json"
	"time"

	"github.com/shohinsan/nestly/internal/sdk/errs"
)

type Todo struct {
	ID        int        `json:"id,omitempty"`
	Title     string     `json:"title" validate:"required, min=1,max=25" `
	Status    string     `json:"status" validate:"required"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

type Todos []Todo

// -----------------------------------------------------------------------------

func (t Todo) Encode() ([]byte, string, error) {
	data, err := json.Marshal(t)
	return data, "application/json", err
}

func (ts Todos) Encode() ([]byte, string, error) {
	data, err := json.Marshal(ts)
	return data, "application/json", err
}

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// -----------------------------------------------------------------------------

// Decode implements the decoder interface.
func (app *Todo) Decode(data []byte) error {
	return json.Unmarshal(data, app)
}

type Decoder interface {
	Decode(data []byte) error
}

func (app Todo) Validate() error {
	if err := errs.Check(app); err != nil {
		return errs.Newf(
			errs.InvalidArgument, "invalidArgument: validate: %s", err)
	}

	return nil
}
