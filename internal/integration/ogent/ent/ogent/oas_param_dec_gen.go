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

func decodeDeleteCategoryParams(args [1]string, r *http.Request) (DeleteCategoryParams, error) {
	var (
		params DeleteCategoryParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeDeletePetParams(args [1]string, r *http.Request) (DeletePetParams, error) {
	var (
		params DeletePetParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeDeleteUserParams(args [1]string, r *http.Request) (DeleteUserParams, error) {
	var (
		params DeleteUserParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeListCategoryParams(args [0]string, r *http.Request) (ListCategoryParams, error) {
	var (
		params    ListCategoryParams
		queryArgs = r.URL.Query()
	)
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListCategoryPetsParams(args [1]string, r *http.Request) (ListCategoryPetsParams, error) {
	var (
		params    ListCategoryPetsParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListPetParams(args [0]string, r *http.Request) (ListPetParams, error) {
	var (
		params    ListPetParams
		queryArgs = r.URL.Query()
	)
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListPetCategoriesParams(args [1]string, r *http.Request) (ListPetCategoriesParams, error) {
	var (
		params    ListPetCategoriesParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListPetFriendsParams(args [1]string, r *http.Request) (ListPetFriendsParams, error) {
	var (
		params    ListPetFriendsParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListUserParams(args [0]string, r *http.Request) (ListUserParams, error) {
	var (
		params    ListUserParams
		queryArgs = r.URL.Query()
	)
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeListUserPetsParams(args [1]string, r *http.Request) (ListUserPetsParams, error) {
	var (
		params    ListUserPetsParams
		queryArgs = r.URL.Query()
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	// Decode query: itemsPerPage.
	{
		values, ok := queryArgs["itemsPerPage"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsDotItemsPerPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotItemsPerPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ItemsPerPage.SetTo(paramsDotItemsPerPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "query: itemsPerPage: parse")
			}
		}
	}
	return params, nil
}

func decodeReadCategoryParams(args [1]string, r *http.Request) (ReadCategoryParams, error) {
	var (
		params ReadCategoryParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeReadPetParams(args [1]string, r *http.Request) (ReadPetParams, error) {
	var (
		params ReadPetParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeReadPetOwnerParams(args [1]string, r *http.Request) (ReadPetOwnerParams, error) {
	var (
		params ReadPetOwnerParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeReadUserParams(args [1]string, r *http.Request) (ReadUserParams, error) {
	var (
		params ReadUserParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeReadUserBestFriendParams(args [1]string, r *http.Request) (ReadUserBestFriendParams, error) {
	var (
		params ReadUserBestFriendParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeUpdateCategoryParams(args [1]string, r *http.Request) (UpdateCategoryParams, error) {
	var (
		params UpdateCategoryParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeUpdatePetParams(args [1]string, r *http.Request) (UpdatePetParams, error) {
	var (
		params UpdatePetParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}

func decodeUpdateUserParams(args [1]string, r *http.Request) (UpdateUserParams, error) {
	var (
		params UpdateUserParams
	)
	// Decode path: id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: id: not specified")
		}
	}
	return params, nil
}
