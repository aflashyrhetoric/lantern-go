package models

type RelationshipType string

const (
	Spouse   RelationshipType = "spouse"
	Friend   RelationshipType = "friend"
	Partner  RelationshipType = "partner"
	Coworker RelationshipType = "coworker"

	Colleague    RelationshipType = "colleague"
	Acquaintance RelationshipType = "acquaintance"
	Cousin       RelationshipType = "cousin"
	Family       RelationshipType = "family"

	Rival RelationshipType = "rival"
	Enemy RelationshipType = "enemy"

	// Aunt RelationshipType = "aunt"
	// Uncle RelationshipType = "uncle"
)

type Relationship struct {
	ID               int              `db:"id" json:"id"`
	PersonOneID      int              `db:"person_one_id" json:"person_one_id,omitempty"`
	PersonTwoID      int              `db:"person_two_id" json:"person_two_id,omitempty"`
	RelationshipType RelationshipType `db:"relationship_type" json:"relationship_type"`
}

type RelationshipHydrated struct {
	ID               int              `json:"id,omitempty"`
	PersonID         int              `json:"person_id,omitempty"`
	RelationshipType RelationshipType `json:"relationship_type"`
	Person           Person           `json:"person"`
}

type CreateRelationshipRequest struct {
	PersonOneID      int              `db:"person_one_id" json:"person_one_id"`
	PersonTwoID      int              `db:"person_two_id" json:"person_two_id"`
	RelationshipType RelationshipType `db:"relationship_type" json:"relationship_type"`
}
