// Package pagination provides primitives for implementing AIP pagination.
//
// See: https://google.aip.dev/158 (Pagination).
package pagination

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var Generator = NewTokenGenerator()

func GetIndex(s string) (int32, error) {
	return Generator.GetIndex(s)
}

func ForIndex(i int32) string {
	return Generator.ForIndex(i)
}

// NewTokenGenerator provides a new instance of a TokenGenerator.
func NewTokenGenerator() TokenGenerator {
	return &tokenGenerator{salt: strconv.FormatInt(time.Now().Unix(), 10)}
}

// TokenGeneratorWithSalt provieds an instance of a TokenGenerator which
// uses the given salt.
func TokenGeneratorWithSalt(salt string) TokenGenerator {
	return &tokenGenerator{salt}
}

// TokenGenerator generates a page token for a given index.
type TokenGenerator interface {
	ForIndex(int32) string
	GetIndex(string) (int32, error)
}

// InvalidTokenErr is the error returned if the token provided is not
// parseable by the TokenGenerator.
var InvalidTokenErr = status.Errorf(
	codes.InvalidArgument,
	"The field `page_token` is invalid.")

type tokenGenerator struct {
	salt string
}

func (t *tokenGenerator) ForIndex(i int32) string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s%d", t.salt, i)))
}

func (t *tokenGenerator) GetIndex(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}

	bs, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		return -1, InvalidTokenErr
	}

	if !strings.HasPrefix(string(bs), t.salt) {
		return -1, InvalidTokenErr
	}

	i, err := strconv.Atoi(strings.TrimPrefix(string(bs), t.salt))
	if err != nil {
		return -1, InvalidTokenErr
	}
	return int32(i), nil
}
