package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int("course_name").
			Positive(),
		field.String("user_id").
			Default("unknown"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return nil
}
