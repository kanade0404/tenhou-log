// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/kanade0404/tenhou-log/services/ent"
)

// The CompressedMJLogFunc type is an adapter to allow the use of ordinary
// function as CompressedMJLog mutator.
type CompressedMJLogFunc func(context.Context, *ent.CompressedMJLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompressedMJLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompressedMJLogMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompressedMJLogMutation", m)
	}
	return f(ctx, mv)
}

// The DanFunc type is an adapter to allow the use of ordinary
// function as Dan mutator.
type DanFunc func(context.Context, *ent.DanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DanMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DanMutation", m)
	}
	return f(ctx, mv)
}

// The GameFunc type is an adapter to allow the use of ordinary
// function as Game mutator.
type GameFunc func(context.Context, *ent.GameMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GameFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GameMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GameMutation", m)
	}
	return f(ctx, mv)
}

// The GamePlayerFunc type is an adapter to allow the use of ordinary
// function as GamePlayer mutator.
type GamePlayerFunc func(context.Context, *ent.GamePlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GamePlayerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerMutation", m)
	}
	return f(ctx, mv)
}

// The GamePlayerHandHaiFunc type is an adapter to allow the use of ordinary
// function as GamePlayerHandHai mutator.
type GamePlayerHandHaiFunc func(context.Context, *ent.GamePlayerHandHaiMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerHandHaiFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GamePlayerHandHaiMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerHandHaiMutation", m)
	}
	return f(ctx, mv)
}

// The GamePlayerPointFunc type is an adapter to allow the use of ordinary
// function as GamePlayerPoint mutator.
type GamePlayerPointFunc func(context.Context, *ent.GamePlayerPointMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerPointFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GamePlayerPointMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerPointMutation", m)
	}
	return f(ctx, mv)
}

// The GoAroundFunc type is an adapter to allow the use of ordinary
// function as GoAround mutator.
type GoAroundFunc func(context.Context, *ent.GoAroundMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoAroundFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoAroundMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoAroundMutation", m)
	}
	return f(ctx, mv)
}

// The HandFunc type is an adapter to allow the use of ordinary
// function as Hand mutator.
type HandFunc func(context.Context, *ent.HandMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HandFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HandMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HandMutation", m)
	}
	return f(ctx, mv)
}

// The MJLogFunc type is an adapter to allow the use of ordinary
// function as MJLog mutator.
type MJLogFunc func(context.Context, *ent.MJLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MJLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MJLogMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MJLogMutation", m)
	}
	return f(ctx, mv)
}

// The MJLogFileFunc type is an adapter to allow the use of ordinary
// function as MJLogFile mutator.
type MJLogFileFunc func(context.Context, *ent.MJLogFileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MJLogFileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MJLogFileMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MJLogFileMutation", m)
	}
	return f(ctx, mv)
}

// The MJLogFileCompressedFunc type is an adapter to allow the use of ordinary
// function as MJLogFileCompressed mutator.
type MJLogFileCompressedFunc func(context.Context, *ent.MJLogFileCompressedMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MJLogFileCompressedFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MJLogFileCompressedMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MJLogFileCompressedMutation", m)
	}
	return f(ctx, mv)
}

// The PlayerFunc type is an adapter to allow the use of ordinary
// function as Player mutator.
type PlayerFunc func(context.Context, *ent.PlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PlayerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerMutation", m)
	}
	return f(ctx, mv)
}

// The RoomFunc type is an adapter to allow the use of ordinary
// function as Room mutator.
type RoomFunc func(context.Context, *ent.RoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoomMutation", m)
	}
	return f(ctx, mv)
}

// The RoundFunc type is an adapter to allow the use of ordinary
// function as Round mutator.
type RoundFunc func(context.Context, *ent.RoundMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoundFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoundMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoundMutation", m)
	}
	return f(ctx, mv)
}

// The WindFunc type is an adapter to allow the use of ordinary
// function as Wind mutator.
type WindFunc func(context.Context, *ent.WindMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WindFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WindMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WindMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
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
	return func(ctx context.Context, m ent.Mutation) bool {
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
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
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
	return func(_ context.Context, m ent.Mutation) bool {
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
	return func(_ context.Context, m ent.Mutation) bool {
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
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
