package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"
	"go-ent-project/graph/model"
	"go-ent-project/internal/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Direction is the resolver for the direction field.
func (r *cameraOrderResolver) Direction(ctx context.Context, obj *ent.CameraOrder, data model.OrderDirection) error {
	panic(fmt.Errorf("not implemented: Direction - direction"))
}

// Field is the resolver for the field field.
func (r *cameraOrderResolver) Field(ctx context.Context, obj *ent.CameraOrder, data model.CameraOrderField) error {
	panic(fmt.Errorf("not implemented: Field - field"))
}

// Direction is the resolver for the direction field.
func (r *permissionOrderResolver) Direction(ctx context.Context, obj *ent.PermissionOrder, data model.OrderDirection) error {
	panic(fmt.Errorf("not implemented: Direction - direction"))
}

// Field is the resolver for the field field.
func (r *permissionOrderResolver) Field(ctx context.Context, obj *ent.PermissionOrder, data model.PermissionOrderField) error {
	panic(fmt.Errorf("not implemented: Field - field"))
}

// Direction is the resolver for the direction field.
func (r *policeStationOrderResolver) Direction(ctx context.Context, obj *ent.PoliceStationOrder, data model.OrderDirection) error {
	panic(fmt.Errorf("not implemented: Direction - direction"))
}

// Field is the resolver for the field field.
func (r *policeStationOrderResolver) Field(ctx context.Context, obj *ent.PoliceStationOrder, data model.PoliceStationOrderField) error {
	panic(fmt.Errorf("not implemented: Field - field"))
}

// Direction is the resolver for the direction field.
func (r *roleOrderResolver) Direction(ctx context.Context, obj *ent.RoleOrder, data model.OrderDirection) error {
	panic(fmt.Errorf("not implemented: Direction - direction"))
}

// Field is the resolver for the field field.
func (r *roleOrderResolver) Field(ctx context.Context, obj *ent.RoleOrder, data model.RoleOrderField) error {
	panic(fmt.Errorf("not implemented: Field - field"))
}

// Direction is the resolver for the direction field.
func (r *userOrderResolver) Direction(ctx context.Context, obj *ent.UserOrder, data model.OrderDirection) error {
	panic(fmt.Errorf("not implemented: Direction - direction"))
}

// Field is the resolver for the field field.
func (r *userOrderResolver) Field(ctx context.Context, obj *ent.UserOrder, data model.UserOrderField) error {
	panic(fmt.Errorf("not implemented: Field - field"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CameraOrder returns CameraOrderResolver implementation.
func (r *Resolver) CameraOrder() CameraOrderResolver { return &cameraOrderResolver{r} }

// PermissionOrder returns PermissionOrderResolver implementation.
func (r *Resolver) PermissionOrder() PermissionOrderResolver { return &permissionOrderResolver{r} }

// PoliceStationOrder returns PoliceStationOrderResolver implementation.
func (r *Resolver) PoliceStationOrder() PoliceStationOrderResolver {
	return &policeStationOrderResolver{r}
}

// RoleOrder returns RoleOrderResolver implementation.
func (r *Resolver) RoleOrder() RoleOrderResolver { return &roleOrderResolver{r} }

// UserOrder returns UserOrderResolver implementation.
func (r *Resolver) UserOrder() UserOrderResolver { return &userOrderResolver{r} }

type queryResolver struct{ *Resolver }
type cameraOrderResolver struct{ *Resolver }
type permissionOrderResolver struct{ *Resolver }
type policeStationOrderResolver struct{ *Resolver }
type roleOrderResolver struct{ *Resolver }
type userOrderResolver struct{ *Resolver }
