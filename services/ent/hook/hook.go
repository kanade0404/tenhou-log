// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/kanade0404/tenhou-log/services/ent"
)

// The CallFunc type is an adapter to allow the use of ordinary
// function as Call mutator.
type CallFunc func(context.Context, *ent.CallMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CallFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CallMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CallMutation", m)
}

// The ChakanFunc type is an adapter to allow the use of ordinary
// function as Chakan mutator.
type ChakanFunc func(context.Context, *ent.ChakanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ChakanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ChakanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ChakanMutation", m)
}

// The ChiiFunc type is an adapter to allow the use of ordinary
// function as Chii mutator.
type ChiiFunc func(context.Context, *ent.ChiiMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ChiiFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ChiiMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ChiiMutation", m)
}

// The CompressedMJLogFunc type is an adapter to allow the use of ordinary
// function as CompressedMJLog mutator.
type CompressedMJLogFunc func(context.Context, *ent.CompressedMJLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompressedMJLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompressedMJLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompressedMJLogMutation", m)
}

// The ConcealedKanFunc type is an adapter to allow the use of ordinary
// function as ConcealedKan mutator.
type ConcealedKanFunc func(context.Context, *ent.ConcealedKanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ConcealedKanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ConcealedKanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ConcealedKanMutation", m)
}

// The DanFunc type is an adapter to allow the use of ordinary
// function as Dan mutator.
type DanFunc func(context.Context, *ent.DanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DanMutation", m)
}

// The DiscardFunc type is an adapter to allow the use of ordinary
// function as Discard mutator.
type DiscardFunc func(context.Context, *ent.DiscardMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiscardFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DiscardMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiscardMutation", m)
}

// The DrawnFunc type is an adapter to allow the use of ordinary
// function as Drawn mutator.
type DrawnFunc func(context.Context, *ent.DrawnMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DrawnFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DrawnMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DrawnMutation", m)
}

// The EventFunc type is an adapter to allow the use of ordinary
// function as Event mutator.
type EventFunc func(context.Context, *ent.EventMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EventFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.EventMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EventMutation", m)
}

// The GameFunc type is an adapter to allow the use of ordinary
// function as Game mutator.
type GameFunc func(context.Context, *ent.GameMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GameFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GameMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GameMutation", m)
}

// The GamePlayerFunc type is an adapter to allow the use of ordinary
// function as GamePlayer mutator.
type GamePlayerFunc func(context.Context, *ent.GamePlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GamePlayerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerMutation", m)
}

// The GamePlayerHandHaiFunc type is an adapter to allow the use of ordinary
// function as GamePlayerHandHai mutator.
type GamePlayerHandHaiFunc func(context.Context, *ent.GamePlayerHandHaiMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerHandHaiFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GamePlayerHandHaiMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerHandHaiMutation", m)
}

// The GamePlayerPointFunc type is an adapter to allow the use of ordinary
// function as GamePlayerPoint mutator.
type GamePlayerPointFunc func(context.Context, *ent.GamePlayerPointMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GamePlayerPointFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GamePlayerPointMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GamePlayerPointMutation", m)
}

// The HandFunc type is an adapter to allow the use of ordinary
// function as Hand mutator.
type HandFunc func(context.Context, *ent.HandMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HandFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.HandMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HandMutation", m)
}

// The MJLogFunc type is an adapter to allow the use of ordinary
// function as MJLog mutator.
type MJLogFunc func(context.Context, *ent.MJLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MJLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MJLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MJLogMutation", m)
}

// The MJLogFileFunc type is an adapter to allow the use of ordinary
// function as MJLogFile mutator.
type MJLogFileFunc func(context.Context, *ent.MJLogFileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MJLogFileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MJLogFileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MJLogFileMutation", m)
}

// The MeldedKanFunc type is an adapter to allow the use of ordinary
// function as MeldedKan mutator.
type MeldedKanFunc func(context.Context, *ent.MeldedKanMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MeldedKanFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MeldedKanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MeldedKanMutation", m)
}

// The PlayerFunc type is an adapter to allow the use of ordinary
// function as Player mutator.
type PlayerFunc func(context.Context, *ent.PlayerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PlayerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PlayerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PlayerMutation", m)
}

// The PonFunc type is an adapter to allow the use of ordinary
// function as Pon mutator.
type PonFunc func(context.Context, *ent.PonMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PonFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PonMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PonMutation", m)
}

// The ReachFunc type is an adapter to allow the use of ordinary
// function as Reach mutator.
type ReachFunc func(context.Context, *ent.ReachMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ReachFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ReachMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ReachMutation", m)
}

// The RoomFunc type is an adapter to allow the use of ordinary
// function as Room mutator.
type RoomFunc func(context.Context, *ent.RoomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoomMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoomMutation", m)
}

// The RoundFunc type is an adapter to allow the use of ordinary
// function as Round mutator.
type RoundFunc func(context.Context, *ent.RoundMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoundFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoundMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoundMutation", m)
}

// The TurnFunc type is an adapter to allow the use of ordinary
// function as Turn mutator.
type TurnFunc func(context.Context, *ent.TurnMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TurnFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TurnMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TurnMutation", m)
}

// The WinFunc type is an adapter to allow the use of ordinary
// function as Win mutator.
type WinFunc func(context.Context, *ent.WinMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WinFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.WinMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WinMutation", m)
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
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
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
