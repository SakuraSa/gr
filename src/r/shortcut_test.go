package r

import (
	"context"
	"reflect"
	"testing"

	"github.com/SakuraSa/gr/src/container"
)

type testAction interface {
	Do(ctx context.Context, t *testing.T) error
}

type testSet[T any] struct {
	v T
}

func (s *testSet[T]) Do(ctx context.Context, t *testing.T) error {
	return Set[T](ctx, s.v)
}

func S[T any](v T) *testSet[T] {
	return &testSet[T]{v: v}
}

type testCheck[T any] struct {
	want T
}

func (c *testCheck[T]) Do(ctx context.Context, t *testing.T) error {
	got, err := Get[T](ctx)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(c.want, got) {
		t.Errorf("want %v, got %v", c.want, got)
	}
	return nil
}

func C[T any](want T) *testCheck[T] {
	return &testCheck[T]{want: want}
}

type testMiss[T any] struct {
}

func (m *testMiss[T]) Do(ctx context.Context, t *testing.T) error {
	_, err := Get[T](ctx)
	if err != container.ErrKeyNotFound {
		t.Errorf("want ErrKeyNotFound, got %v", err)
	}
	return nil
}

func M[T any]() *testMiss[T] {
	return &testMiss[T]{}
}

func TestGetSet(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		actions []testAction
		wantErr bool
	}{
		{
			name: "normal",
			ctx:  New(context.Background()),
			actions: []testAction{
				S("hello"),
				C("hello"),
			},
			wantErr: false,
		},
		{
			name: "missing",
			ctx:  New(context.Background()),
			actions: []testAction{
				M[string](),
			},
			wantErr: false,
		},
		{
			name: "duplicate",
			ctx:  New(context.Background()),
			actions: []testAction{
				S("hello"),
				S("world"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			for _, action := range tt.actions {
				err = action.Do(tt.ctx, t)
				if err != nil {
					break
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("action.Do() error = %v wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
