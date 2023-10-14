// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen"
)

// The OriginalVideoFunc type is an adapter to allow the use of ordinary
// function as OriginalVideo mutator.
type OriginalVideoFunc func(context.Context, *videodbgen.OriginalVideoMutation) (videodbgen.Value, error)

// Mutate calls f(ctx, m).
func (f OriginalVideoFunc) Mutate(ctx context.Context, m videodbgen.Mutation) (videodbgen.Value, error) {
	if mv, ok := m.(*videodbgen.OriginalVideoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *videodbgen.OriginalVideoMutation", m)
}

// The SubtitledVideoFunc type is an adapter to allow the use of ordinary
// function as SubtitledVideo mutator.
type SubtitledVideoFunc func(context.Context, *videodbgen.SubtitledVideoMutation) (videodbgen.Value, error)

// Mutate calls f(ctx, m).
func (f SubtitledVideoFunc) Mutate(ctx context.Context, m videodbgen.Mutation) (videodbgen.Value, error) {
	if mv, ok := m.(*videodbgen.SubtitledVideoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *videodbgen.SubtitledVideoMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, videodbgen.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m videodbgen.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m videodbgen.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m videodbgen.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op videodbgen.Op) Condition {
	return func(_ context.Context, m videodbgen.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m videodbgen.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m videodbgen.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m videodbgen.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk videodbgen.Hook, cond Condition) videodbgen.Hook {
	return func(next videodbgen.Mutator) videodbgen.Mutator {
		return videodbgen.MutateFunc(func(ctx context.Context, m videodbgen.Mutation) (videodbgen.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, videodbgen.Delete|videodbgen.Create)
func On(hk videodbgen.Hook, op videodbgen.Op) videodbgen.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, videodbgen.Update|videodbgen.UpdateOne)
func Unless(hk videodbgen.Hook, op videodbgen.Op) videodbgen.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) videodbgen.Hook {
	return func(videodbgen.Mutator) videodbgen.Mutator {
		return videodbgen.MutateFunc(func(context.Context, videodbgen.Mutation) (videodbgen.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []videodbgen.Hook {
//		return []videodbgen.Hook{
//			Reject(videodbgen.Delete|videodbgen.Update),
//		}
//	}
func Reject(op videodbgen.Op) videodbgen.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []videodbgen.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...videodbgen.Hook) Chain {
	return Chain{append([]videodbgen.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() videodbgen.Hook {
	return func(mutator videodbgen.Mutator) videodbgen.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...videodbgen.Hook) Chain {
	newHooks := make([]videodbgen.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}