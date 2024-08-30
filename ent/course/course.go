// Code generated by ent, DO NOT EDIT.

package course

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCourseName holds the string denoting the course_name field in the database.
	FieldCourseName = "course_name"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the course in the database.
	Table = "courses"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldCourseName,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// CourseNameValidator is a validator for the "course_name" field. It is called by the builders before save.
	CourseNameValidator func(int) error
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID string
)

// OrderOption defines the ordering options for the Course queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCourseName orders the results by the course_name field.
func ByCourseName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseName, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}
