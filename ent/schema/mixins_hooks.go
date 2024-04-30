//go:build !skiphooks

package schema

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"entgo.io/ent"
// 	"entgo.io/ent/dialect/sql"
// 	gen "github.com/orbit-ops/launchpad-core/ent"
// 	"github.com/orbit-ops/launchpad-core/ent/hook"
// 	"github.com/orbit-ops/launchpad-core/ent/intercept"
// )

// // Hooks of the SoftDeleteMixin.
// func (d SoftDeleteMixin) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(
// 			func(next ent.Mutator) ent.Mutator {
// 				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 					// Skip soft-delete, means delete the entity permanently.
// 					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
// 						return next.Mutate(ctx, m)
// 					}
// 					mx, ok := m.(interface {
// 						SetOp(ent.Op)
// 						Client() *gen.Client
// 						SetDeleteTime(time.Time)
// 						WhereP(...func(*sql.Selector))
// 					})
// 					if !ok {
// 						return nil, fmt.Errorf("unexpected mutation type %T", m)
// 					}
// 					d.P(mx)
// 					mx.SetOp(ent.OpUpdate)
// 					mx.SetDeleteTime(time.Now())
// 					return mx.Client().Mutate(ctx, m)
// 				})
// 			},
// 			ent.OpDeleteOne|ent.OpDelete,
// 		),
// 	}
// }

// // Interceptors of the SoftDeleteMixin.
// func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
// 	return []ent.Interceptor{
// 		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
// 			// Skip soft-delete, means include soft-deleted entities.
// 			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
// 				return nil
// 			}
// 			d.P(q)
// 			return nil
// 		}),
// 	}
// }
