# AIP

[![CI](https://github.com/sraphs/aip/actions/workflows/ci.yml/badge.svg)](https://github.com/sraphs/aip/actions/workflows/ci.yml)

> Go SDK for implementing [Google API Improvement Proposals](https://aip.dev/) (AIP).

Modified from [https://github.com/einride/aip-go](https://github.com/einride/aip-go)

## Documentation
-------------

See [https://aip.dev](https://aip.dev/) for the full AIP documentation.

## Install

```bash
go get github.com/sraphs/aip
```

## Usage

### [AIP-132](https://google.aip.dev/132) (Standard method: List)

```go
package examplelibrary

import (
	"context"

	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sraphs/aip/pagination"
)

var generator = pagination.TokenGeneratorWithSalt("example.library.ListShelves")

func (s *Server) ListShelves(
	ctx context.Context,
	request *library.ListShelvesRequest,
) (*library.ListShelvesResponse, error) {
	// Handle request constraints.
	const (
		maxPageSize     = 1000
		defaultPageSize = 100
	)
	switch {
	case request.PageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size is negative")
	case request.PageSize == 0:
		request.PageSize = defaultPageSize
	case request.PageSize > maxPageSize:
		request.PageSize = maxPageSize
	}
	// Use pagination.PageToken for offset-based page tokens.
	offset, err := generator.GetIndex(request.PageToken)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}
	// Query the storage.
	result, err := s.Storage.ListShelves(ctx, &ListShelvesQuery{
		Offset:   int64(offset),
		PageSize: request.GetPageSize(),
	})
	if err != nil {
		return nil, err
	}
	// Build the response.
	response := &library.ListShelvesResponse{
		Shelves: result.Shelves,
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = generator.ForIndex(offset + int(request.PageSize))
	}
	// Respond.
	return response, nil
}

```

## Contributing

We alway welcome your contributions :clap:

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


## CHANGELOG
See [Releases](https://github.com/sraphs/aip/releases)

## License
[MIT © sraph.com](./LICENSE)
