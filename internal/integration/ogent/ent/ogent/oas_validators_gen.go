// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func (s CreateUserReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s CreateUserReqFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s CreateUserReqFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s CreateUserReqFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s ListCategoryOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListCategoryPetsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetCategoriesOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetFriendsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListPetOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}
func (s ListUserOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range s {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s ListUserPetsOKApplicationJSON) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}

func (s PetCreate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "owner",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetCreateOwner) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetCreateOwnerFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetCreateOwnerFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetCreateOwnerFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetOwnerRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetOwnerReadFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetOwnerReadFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetOwnerReadFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UpdateUserReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.FavoriteCatBreed.Set {
			if err := func() error {
				if err := s.FavoriteCatBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UpdateUserReqFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UpdateUserReqFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UpdateUserReqFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserBestFriendRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserBestFriendReadFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserBestFriendReadFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserBestFriendReadFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserCreate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserCreateFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserCreateFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserCreateFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserListFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserListFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserListFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserReadFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserReadFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserReadFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserUpdate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.FavoriteCatBreed.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_cat_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteDogBreed.Set {
			if err := func() error {
				if err := s.FavoriteDogBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_dog_breed",
			Error: err,
		})
	}
	if err := func() error {
		if s.FavoriteFishBreed.Set {
			if err := func() error {
				if err := s.FavoriteFishBreed.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "favorite_fish_breed",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s UserUpdateFavoriteCatBreed) Validate() error {
	switch s {
	case "siamese":
		return nil
	case "bengal":
		return nil
	case "lion":
		return nil
	case "tiger":
		return nil
	case "leopard":
		return nil
	case "other":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserUpdateFavoriteDogBreed) Validate() error {
	switch s {
	case "Kuro":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s UserUpdateFavoriteFishBreed) Validate() error {
	switch s {
	case "gold":
		return nil
	case "koi":
		return nil
	case "shark":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
