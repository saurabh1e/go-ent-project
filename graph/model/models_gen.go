// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"go-ent-project/internal/ent"
	"io"
	"strconv"
)

type Car struct {
	ID           string `json:"id"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int32  `json:"year"`
	Registration string `json:"registration"`
	Color        string `json:"color"`
}

func (Car) IsNode() {}

// CarWhereInput is used for filtering Car objects.
// Input was generated by ent.
type CarWhereInput struct {
	Not *CarWhereInput   `json:"not,omitempty"`
	And []*CarWhereInput `json:"and,omitempty"`
	Or  []*CarWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *string  `json:"id,omitempty"`
	IDNeq   *string  `json:"idNEQ,omitempty"`
	IDIn    []string `json:"idIn,omitempty"`
	IDNotIn []string `json:"idNotIn,omitempty"`
	IDGt    *string  `json:"idGT,omitempty"`
	IDGte   *string  `json:"idGTE,omitempty"`
	IDLt    *string  `json:"idLT,omitempty"`
	IDLte   *string  `json:"idLTE,omitempty"`
	// make field predicates
	Make             *string  `json:"make,omitempty"`
	MakeNeq          *string  `json:"makeNEQ,omitempty"`
	MakeIn           []string `json:"makeIn,omitempty"`
	MakeNotIn        []string `json:"makeNotIn,omitempty"`
	MakeGt           *string  `json:"makeGT,omitempty"`
	MakeGte          *string  `json:"makeGTE,omitempty"`
	MakeLt           *string  `json:"makeLT,omitempty"`
	MakeLte          *string  `json:"makeLTE,omitempty"`
	MakeContains     *string  `json:"makeContains,omitempty"`
	MakeHasPrefix    *string  `json:"makeHasPrefix,omitempty"`
	MakeHasSuffix    *string  `json:"makeHasSuffix,omitempty"`
	MakeEqualFold    *string  `json:"makeEqualFold,omitempty"`
	MakeContainsFold *string  `json:"makeContainsFold,omitempty"`
	// model field predicates
	Model             *string  `json:"model,omitempty"`
	ModelNeq          *string  `json:"modelNEQ,omitempty"`
	ModelIn           []string `json:"modelIn,omitempty"`
	ModelNotIn        []string `json:"modelNotIn,omitempty"`
	ModelGt           *string  `json:"modelGT,omitempty"`
	ModelGte          *string  `json:"modelGTE,omitempty"`
	ModelLt           *string  `json:"modelLT,omitempty"`
	ModelLte          *string  `json:"modelLTE,omitempty"`
	ModelContains     *string  `json:"modelContains,omitempty"`
	ModelHasPrefix    *string  `json:"modelHasPrefix,omitempty"`
	ModelHasSuffix    *string  `json:"modelHasSuffix,omitempty"`
	ModelEqualFold    *string  `json:"modelEqualFold,omitempty"`
	ModelContainsFold *string  `json:"modelContainsFold,omitempty"`
	// year field predicates
	Year      *int32  `json:"year,omitempty"`
	YearNeq   *int32  `json:"yearNEQ,omitempty"`
	YearIn    []int32 `json:"yearIn,omitempty"`
	YearNotIn []int32 `json:"yearNotIn,omitempty"`
	YearGt    *int32  `json:"yearGT,omitempty"`
	YearGte   *int32  `json:"yearGTE,omitempty"`
	YearLt    *int32  `json:"yearLT,omitempty"`
	YearLte   *int32  `json:"yearLTE,omitempty"`
	// registration field predicates
	Registration             *string  `json:"registration,omitempty"`
	RegistrationNeq          *string  `json:"registrationNEQ,omitempty"`
	RegistrationIn           []string `json:"registrationIn,omitempty"`
	RegistrationNotIn        []string `json:"registrationNotIn,omitempty"`
	RegistrationGt           *string  `json:"registrationGT,omitempty"`
	RegistrationGte          *string  `json:"registrationGTE,omitempty"`
	RegistrationLt           *string  `json:"registrationLT,omitempty"`
	RegistrationLte          *string  `json:"registrationLTE,omitempty"`
	RegistrationContains     *string  `json:"registrationContains,omitempty"`
	RegistrationHasPrefix    *string  `json:"registrationHasPrefix,omitempty"`
	RegistrationHasSuffix    *string  `json:"registrationHasSuffix,omitempty"`
	RegistrationEqualFold    *string  `json:"registrationEqualFold,omitempty"`
	RegistrationContainsFold *string  `json:"registrationContainsFold,omitempty"`
	// color field predicates
	Color             *string  `json:"color,omitempty"`
	ColorNeq          *string  `json:"colorNEQ,omitempty"`
	ColorIn           []string `json:"colorIn,omitempty"`
	ColorNotIn        []string `json:"colorNotIn,omitempty"`
	ColorGt           *string  `json:"colorGT,omitempty"`
	ColorGte          *string  `json:"colorGTE,omitempty"`
	ColorLt           *string  `json:"colorLT,omitempty"`
	ColorLte          *string  `json:"colorLTE,omitempty"`
	ColorContains     *string  `json:"colorContains,omitempty"`
	ColorHasPrefix    *string  `json:"colorHasPrefix,omitempty"`
	ColorHasSuffix    *string  `json:"colorHasSuffix,omitempty"`
	ColorEqualFold    *string  `json:"colorEqualFold,omitempty"`
	ColorContainsFold *string  `json:"colorContainsFold,omitempty"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type PoliceStationConnection struct {
	PageInfo *PageInfo            `json:"pageInfo"`
	Edges    []*ent.PoliceStation `json:"edges"`
}

type Query struct {
}

// Properties by which Camera connections can be ordered.
type CameraOrderField string

const (
	CameraOrderFieldName CameraOrderField = "NAME"
)

var AllCameraOrderField = []CameraOrderField{
	CameraOrderFieldName,
}

func (e CameraOrderField) IsValid() bool {
	switch e {
	case CameraOrderFieldName:
		return true
	}
	return false
}

func (e CameraOrderField) String() string {
	return string(e)
}

func (e *CameraOrderField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CameraOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CameraOrderField", str)
	}
	return nil
}

func (e CameraOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which PoliceStation connections can be ordered.
type PoliceStationOrderField string

const (
	PoliceStationOrderFieldName PoliceStationOrderField = "NAME"
)

var AllPoliceStationOrderField = []PoliceStationOrderField{
	PoliceStationOrderFieldName,
}

func (e PoliceStationOrderField) IsValid() bool {
	switch e {
	case PoliceStationOrderFieldName:
		return true
	}
	return false
}

func (e PoliceStationOrderField) String() string {
	return string(e)
}

func (e *PoliceStationOrderField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PoliceStationOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PoliceStationOrderField", str)
	}
	return nil
}

func (e PoliceStationOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which User connections can be ordered.
type UserOrderField string

const (
	UserOrderFieldName      UserOrderField = "NAME"
	UserOrderFieldEmail     UserOrderField = "EMAIL"
	UserOrderFieldCreatedAt UserOrderField = "CREATED_AT"
)

var AllUserOrderField = []UserOrderField{
	UserOrderFieldName,
	UserOrderFieldEmail,
	UserOrderFieldCreatedAt,
}

func (e UserOrderField) IsValid() bool {
	switch e {
	case UserOrderFieldName, UserOrderFieldEmail, UserOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e UserOrderField) String() string {
	return string(e)
}

func (e *UserOrderField) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

func (e UserOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
