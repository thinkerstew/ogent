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

type DeleteCategoryParams struct {
	// ID of the Category.
	ID int
}

type DeletePetParams struct {
	// ID of the Pet.
	ID int
}

type DeleteUserParams struct {
	// ID of the User.
	ID int
}

type ListCategoryParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListCategoryPetsParams struct {
	// ID of the Category.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListPetParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListPetCategoriesParams struct {
	// ID of the Pet.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListPetFriendsParams struct {
	// ID of the Pet.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListUserParams struct {
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ListUserPetsParams struct {
	// ID of the User.
	ID int
	// What page to render.
	Page OptInt
	// Item count to render per page.
	ItemsPerPage OptInt
}

type ReadCategoryParams struct {
	// ID of the Category.
	ID int
}

type ReadPetParams struct {
	// ID of the Pet.
	ID int
}

type ReadPetOwnerParams struct {
	// ID of the Pet.
	ID int
}

type ReadUserParams struct {
	// ID of the User.
	ID int
}

type ReadUserBestFriendParams struct {
	// ID of the User.
	ID int
}

type UpdateCategoryParams struct {
	// ID of the Category.
	ID int
}

type UpdatePetParams struct {
	// ID of the Pet.
	ID int
}

type UpdateUserParams struct {
	// ID of the User.
	ID int
}
