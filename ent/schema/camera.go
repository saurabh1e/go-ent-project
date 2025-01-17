package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Camera holds the schema definition for the Camera entity.
type Camera struct {
	ent.Schema
}

// Fields of the Camera.
func (Camera) Fields() []ent.Field {
	return []ent.Field{
		// Name of the camera
		field.String("name").
			NotEmpty().
			Annotations(

				entgql.OrderField("NAME"),
			),
		// Model of the camera
		field.String("model").
			NotEmpty().
			Annotations(),

		// IMEI number of the camera
		field.String("imei").
			Unique().
			NotEmpty().
			Annotations(),

		// Location stored as geometry point in PostgreSQL
		field.JSON("location", map[string]interface{}{}).
			Annotations().
			SchemaType(map[string]string{
				"postgres": "geometry(Point)",
			}),

		// Whether the camera is active
		field.Bool("active").
			Default(true).
			Annotations(),
	}
}

// Edges of the Camera.
func (Camera) Edges() []ent.Edge {
	return nil
}
