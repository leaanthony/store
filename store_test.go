package store_test

import (
	"github.com/matryer/is"
	"github.com/wailsapp/wails/v2/pkg/store"
	"testing"
)

func TestNew(t *testing.T) {
	i := is.New(t)
	tests := []struct {
		name         string
		initialValue any
	}{
		{
			name:         "nil value",
			initialValue: nil,
		},
		{
			name:         "int value",
			initialValue: 1,
		},
		{
			name:         "string value",
			initialValue: "test",
		},
		{
			name:         "struct value",
			initialValue: struct{ test string }{test: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := store.New(tt.initialValue)
			i.True(got != nil)
			i.Equal(got.Get(), tt.initialValue)
		})
	}
}

func TestStore_Set_Get(t *testing.T) {
	i := is.New(t)

	tests := []struct {
		name         string
		initialValue any
		setValue     any
	}{
		{
			name:         "nil value",
			initialValue: nil,
			setValue:     &struct{ test string }{test: "test"},
		},
		{
			name:         "int value",
			initialValue: 1,
			setValue:     2,
		},
		{
			name:         "string value",
			initialValue: "test",
			setValue:     "test2",
		},
		{
			name:         "struct value",
			initialValue: struct{ test string }{test: "test"},
			setValue:     struct{ test string }{test: "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.New(tt.initialValue)
			s.Set(tt.setValue)
			i.Equal(s.Get(), tt.setValue)
		})
	}
}

func TestStore_Update_Subscribe(t *testing.T) {
	i := is.New(t)
	tests := []struct {
		name         string
		initialValue any
		updatedValue any
		subscribers  []func(any)
	}{
		{
			name:         "no subscribers",
			initialValue: nil,
			updatedValue: nil,
		},
		{
			name:         "single subscriber",
			initialValue: 1,
			updatedValue: 2,
			subscribers: []func(any){
				func(value any) {
					i.Equal(value, 2)
				},
			},
		},
		{
			name:         "multiple subscribers",
			initialValue: 1,
			updatedValue: 2,
			subscribers: []func(value any){
				func(value any) {
					i.Equal(value, 2)
					println("subscriber 1 called")
				},
				func(value any) {
					i.Equal(value, 2)
					println("subscriber 2 called")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.New(tt.initialValue)
			if tt.subscribers != nil {
				for _, subscriber := range tt.subscribers {
					s.Subscribe(subscriber)
				}
			}
			s.Update(func(value any) any {
				return tt.updatedValue
			})
			i.Equal(s.Get(), tt.updatedValue)
		})
	}
}

func TestStore_UnSubscribe(t *testing.T) {
	i := is.New(t)
	s := store.New(1)
	callCount := 0
	unsub1 := s.Subscribe(func(value int) {
		println("subscriber 1 called")
		callCount++
	})
	unsub2 := s.Subscribe(func(value int) {
		println("subscriber 2 called")
		callCount++
	})
	s.Update(func(value int) int {
		return 2
	})
	i.Equal(callCount, 2)

	callCount = 0
	unsub1()
	s.Update(func(value int) int {
		return 0
	})
	i.Equal(callCount, 1)

	callCount = 0
	unsub2()
	s.Update(func(value int) int {
		return 0
	})
	i.Equal(callCount, 0)
}
