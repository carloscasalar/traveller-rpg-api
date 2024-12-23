// Package apirest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package apirest

// Defines values for CitizenCategory.
const (
	AboveAverage CitizenCategory = "above_average"
	Average      CitizenCategory = "average"
	BelowAverage CitizenCategory = "below_average"
	Exceptional  CitizenCategory = "exceptional"
)

// Defines values for Experience.
const (
	Elite        Experience = "elite"
	Intermediate Experience = "intermediate"
	Recruit      Experience = "recruit"
	Regular      Experience = "regular"
	Rookie       Experience = "rookie"
	Veteran      Experience = "veteran"
)

// Defines values for Gender.
const (
	Female      Gender = "female"
	Male        Gender = "male"
	Unspecified Gender = "unspecified"
)

// Defines values for Role.
const (
	Diplomat    Role = "diplomat"
	Engineer    Role = "engineer"
	Entertainer Role = "entertainer"
	Gunner      Role = "gunner"
	Leader      Role = "leader"
	Marine      Role = "marine"
	Medic       Role = "medic"
	Navigator   Role = "navigator"
	Pilot       Role = "pilot"
	Scout       Role = "scout"
	Steward     Role = "steward"
	Technician  Role = "technician"
	Thug        Role = "thug"
	Trader      Role = "trader"
)

// Characteristics defines model for Characteristics.
type Characteristics struct {
	// DEX Dexterity
	DEX int `json:"DEX"`

	// EDU Education
	EDU int `json:"EDU"`

	// END Endurance
	END int `json:"END"`

	// INT Intelligence
	INT int `json:"INT"`

	// SOC Social
	SOC int `json:"SOC"`

	// STR Strength
	STR int `json:"STR"`
}

// CitizenCategory How exceptional are the characteristics of the NPC
type CitizenCategory string

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Experience defines model for Experience.
type Experience string

// Gender Gender. If you don't care, just omit it or use unspecified
type Gender string

// NPC defines model for NPC.
type NPC struct {
	Characteristics Characteristics `json:"characteristics"`

	// CitizenCategory How exceptional are the characteristics of the NPC
	CitizenCategory CitizenCategory `json:"citizen_category"`
	Experience      Experience      `json:"experience"`
	FirstName       string          `json:"first_name"`

	// Role Role of the NPC
	Role    Role     `json:"role"`
	Skills  []string `json:"skills"`
	Surname string   `json:"surname"`
}

// NPCRequest defines model for NPCRequest.
type NPCRequest struct {
	// CitizenCategory How exceptional are the characteristics of the NPC
	CitizenCategory *CitizenCategory `json:"citizen_category,omitempty"`
	Experience      *Experience      `json:"experience,omitempty"`

	// Gender Gender. If you don't care, just omit it or use unspecified
	Gender *Gender `json:"gender,omitempty"`

	// Role Role of the NPC
	Role Role `json:"role"`
}

// Role Role of the NPC
type Role string

// GenerateNPCJSONRequestBody defines body for GenerateNPC for application/json ContentType.
type GenerateNPCJSONRequestBody = NPCRequest
