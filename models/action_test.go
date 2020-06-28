package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAction struct {
	KeyValue string
	Func     func(*State) bool
}

func (m *mockAction) Key() string {
	return m.KeyValue
}

func (m *mockAction) Do(s *State) bool {
	return m.Func(s)
}

func TestActions_Keys(t *testing.T) {
	tests := []struct {
		name string
		give Actions
		want []string
	}{
		{
			name: "empty actions will not give any keys",
			give: nil,
			want: nil,
		},
		{
			name: "a single action will return the right key",
			give: Actions{&mockAction{"Foo", nil}},
			want: []string{"Foo"},
		},
		{
			name: "multiple actions will return all the right keys",
			give: Actions{
				&mockAction{"Foo", nil},
				&mockAction{"Bar", nil},
				&mockAction{"Baz", nil},
			},
			want: []string{"Foo", "Bar", "Baz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := tt.give.Keys()
			assert.Equal(t, tt.want, keys)
		})
	}
}

func TestActions_Do(t *testing.T) {
	var called int

	tests := []struct {
		name        string
		setup       func()
		giveActions Actions
		giveState   *State
		giveKey     string
		wantCalled  int
	}{
		{
			name: "with one action that matches the key it will be called",
			giveActions: Actions{&mockAction{"Foo", func(_ *State) bool {
				called += 1
				return false
			}}},
			giveState:  &State{},
			giveKey:    "Foo",
			wantCalled: 1,
		},
		{
			name: "with two action that both match the key only the first one will be called",
			giveActions: Actions{
				&mockAction{"Foo", func(_ *State) bool {
					called += 1
					return false
				}},
				&mockAction{"Foo", func(_ *State) bool {
					assert.Fail(t, "function should not have been called")
					return false
				}},
			},
			giveState:  &State{},
			giveKey:    "Foo",
			wantCalled: 1,
		},
		{
			name: "with multiple actions and one matching the right key this one shoudl be called",
			giveActions: Actions{
				&mockAction{"Foo", func(_ *State) bool {
					assert.Fail(t, "function should not have been called")
					return false
				}},
				&mockAction{"Bar", func(_ *State) bool {
					called += 1
					return false
				}},
				&mockAction{"Baz", func(_ *State) bool {
					assert.Fail(t, "function should not have been called")
					return false
				}},
			},
			giveState:  &State{},
			giveKey:    "Bar",
			wantCalled: 1,
		},
		{
			name: "with multiple actions and none matching nothing will be called",
			giveActions: Actions{
				&mockAction{"Foo", func(_ *State) bool {
					assert.Fail(t, "function should not have been called")
					return false
				}},
				&mockAction{"Bar", func(_ *State) bool {
					assert.Fail(t, "function should not have been called")
					return false
				}},
			},
			giveState:  &State{},
			giveKey:    "Baz",
			wantCalled: 0,
		},
		{
			name:        "with no actions nothing will be called",
			giveActions: Actions{},
			giveState:   &State{},
			giveKey:     "Foo",
			wantCalled:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called = 0

			tt.giveActions.Do(tt.giveState, tt.giveKey)
			assert.Equal(t, tt.wantCalled, called, "unexpected amount of function calls")
		})
	}
}
