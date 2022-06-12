package testdata

//go:generate protoc -I=. --go_out=paths=source_relative:. test_import.proto test_all_types.proto
