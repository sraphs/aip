package types

import (
	"math"
	"testing"

	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/types/known/anypb"

	"github.com/sraphs/aip/filtering/internal/testdata"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    proto.Message
		b    proto.Message
		out  bool
	}{
		{
			name: "EqualEmptyInstances",
			a:    &testdata.TestAllTypes{},
			b:    &testdata.TestAllTypes{},
			out:  true,
		},
		{
			name: "NotEqualEmptyInstances",
			a:    &testdata.TestAllTypes{},
			b:    &testdata.NestedTestAllTypes{},
			out:  false,
		},
		{
			name: "EqualScalarFields",
			a: &testdata.TestAllTypes{
				SingleBool:   true,
				SingleBytes:  []byte("world"),
				SingleDouble: 3.0,
				SingleFloat:  1.5,
				SingleInt32:  1,
				SingleUint64: 1,
				SingleString: "hello",
			},
			b: &testdata.TestAllTypes{
				SingleBool:   true,
				SingleBytes:  []byte("world"),
				SingleDouble: 3.0,
				SingleFloat:  1.5,
				SingleInt32:  1,
				SingleUint64: 1,
				SingleString: "hello",
			},
			out: true,
		},
		{
			name: "NotEqualFloatNan",
			a: &testdata.TestAllTypes{
				SingleFloat: float32(math.NaN()),
			},
			b: &testdata.TestAllTypes{
				SingleFloat: float32(math.NaN()),
			},
			out: false,
		},
		{
			name: "NotEqualDifferentFieldsSet",
			a: &testdata.TestAllTypes{
				SingleInt32: 1,
			},
			b:   &testdata.TestAllTypes{},
			out: false,
		},
		{
			name: "NotEqualDifferentFieldsSetReverse",
			a:    &testdata.TestAllTypes{},
			b: &testdata.TestAllTypes{
				SingleInt32: 1,
			},
			out: false,
		},
		{
			name: "EqualListField",
			a: &testdata.TestAllTypes{
				RepeatedInt32: []int32{1, 2, 3, 4},
			},
			b: &testdata.TestAllTypes{
				RepeatedInt32: []int32{1, 2, 3, 4},
			},
			out: true,
		},
		{
			name: "NotEqualListFieldDifferentLength",
			a: &testdata.TestAllTypes{
				RepeatedInt32: []int32{1, 2, 3},
			},
			b: &testdata.TestAllTypes{
				RepeatedInt32: []int32{1, 2, 3, 4},
			},
			out: false,
		},
		{
			name: "NotEqualListFieldDifferentContent",
			a: &testdata.TestAllTypes{
				RepeatedInt32: []int32{2, 1},
			},
			b: &testdata.TestAllTypes{
				RepeatedInt32: []int32{1, 2},
			},
			out: false,
		},
		{
			name: "EqualMapField",
			a: &testdata.TestAllTypes{
				MapInt64NestedType: map[int64]*testdata.NestedTestAllTypes{
					1: {
						Child: &testdata.NestedTestAllTypes{
							Payload: &testdata.TestAllTypes{
								StandaloneEnum: testdata.TestAllTypes_BAR,
							},
						},
					},
					2: {
						Payload: &testdata.TestAllTypes{},
					},
				},
			},
			b: &testdata.TestAllTypes{
				MapInt64NestedType: map[int64]*testdata.NestedTestAllTypes{
					1: {
						Child: &testdata.NestedTestAllTypes{
							Payload: &testdata.TestAllTypes{
								StandaloneEnum: testdata.TestAllTypes_BAR,
							},
						},
					},
					2: {
						Payload: &testdata.TestAllTypes{},
					},
				},
			},
			out: true,
		},
		{
			name: "NotEqualMapFieldDifferentLength",
			a: &testdata.TestAllTypes{
				MapInt64NestedType: map[int64]*testdata.NestedTestAllTypes{
					1: {
						Child: &testdata.NestedTestAllTypes{},
					},
					2: {
						Payload: &testdata.TestAllTypes{},
					},
				},
			},
			b: &testdata.TestAllTypes{
				MapInt64NestedType: map[int64]*testdata.NestedTestAllTypes{
					1: {
						Child: &testdata.NestedTestAllTypes{},
					},
				},
			},
			out: false,
		},
		{
			name: "EqualAnyBytes",
			a: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			out: true,
		},
		{
			name: "NotEqualDoublePackedAny",
			a: &testdata.TestAllTypes{
				SingleAny: doublePackAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: doublePackAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3, 4},
				}),
			},
			out: false,
		},
		{
			name: "NotEqualAnyTypeURL",
			a: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.NestedTestAllTypes{}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.TestAllTypes{}),
			},
			out: false,
		},
		{
			name: "NotEqualAnyFields",
			a: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: packAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			out: false,
		},
		{
			name: "NotEqualAnyDeserializeA",
			a: &testdata.TestAllTypes{
				SingleAny: badPackAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: badPackAny(t, &testdata.TestAllTypes{
					SingleInt32:   1,
					SingleUint32:  2,
					SingleString:  "three",
					RepeatedInt32: []int32{1, 2, 3},
				}),
			},
			out: false,
		},
		{
			name: "EqualUnknownFields",
			a: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt32: 1,
						},
					},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt32: 1,
						},
					},
				}),
			},
			out: true,
		},
		{
			name: "NotEqualUnknownFieldsCount",
			a: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt32: 1,
							SingleFloat: 2.0,
						},
					},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt32: 1,
						},
					},
				}),
			},
			out: false,
		},
		{
			name: "NotEqualUnknownFields",
			a: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt64: 2,
						},
					},
				}),
			},
			b: &testdata.TestAllTypes{
				SingleAny: misPackAny(t, &testdata.NestedTestAllTypes{
					Child: &testdata.NestedTestAllTypes{
						Payload: &testdata.TestAllTypes{
							SingleInt32: 1,
						},
					},
				}),
			},
			out: false,
		},
		{
			name: "NotEqualOneNil",
			a:    nil,
			b:    &testdata.TestAllTypes{},
			out:  false,
		},
		{
			name: "EqualBothNil",
			a:    nil,
			b:    nil,
			out:  true,
		},
	}

	for _, tst := range tests {
		tc := tst
		t.Run(tc.name, func(t *testing.T) {
			got := Equal(tc.a, tc.b)
			if got != tc.out {
				t.Errorf("Equal(%v, %v) got %v, wanted %v", tc.a, tc.b, got, tc.out)
			}
		})
	}
}

func packAny(t *testing.T, m proto.Message) *anypb.Any {
	t.Helper()
	any, err := anypb.New(m)
	if err != nil {
		t.Fatalf("anypb.New(%v) failed with error: %v", m, err)
	}
	return any
}

func doublePackAny(t *testing.T, m proto.Message) *anypb.Any {
	t.Helper()
	any, err := anypb.New(m)
	if err != nil {
		t.Fatalf("anypb.New(%v) failed with error: %v", m, err)
	}
	any, err = anypb.New(any)
	if err != nil {
		t.Fatalf("anypb.New(%v) failed with error: %v", any, err)
	}
	return any
}

func badPackAny(t *testing.T, m proto.Message) *anypb.Any {
	t.Helper()
	any, err := anypb.New(m)
	if err != nil {
		t.Fatalf("anypb.New(%v) failed with error: %v", m, err)
	}
	any.TypeUrl = "type.googleapis.com/BadType"
	return any
}

func misPackAny(t *testing.T, m proto.Message) *anypb.Any {
	t.Helper()
	any, err := anypb.New(m)
	if err != nil {
		t.Fatalf("anypb.New(%v) failed with error: %v", m, err)
	}
	any.TypeUrl = "type.googleapis.com/google.expr.proto3.test.TestAllTypes"
	return any
}
