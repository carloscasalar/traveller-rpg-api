package apirest

import (
	"strings"
)

var allCitizenCategories = map[CitizenCategory]bool{
	BelowAverage: true,
	Average:      true,
	AboveAverage: true,
	Exceptional:  true,
}

func IsInvalidCitizenCategory(category CitizenCategory) bool {
	return !allCitizenCategories[category]
}

func AllCitizenCategoriesString() string {
	return strings.Join([]string{
		string(BelowAverage),
		string(Average),
		string(AboveAverage),
		string(Exceptional),
	}, ", ")
}

var allExperiences = map[Experience]bool{
	Recruit:      true,
	Rookie:       true,
	Intermediate: true,
	Regular:      true,
	Veteran:      true,
	Elite:        true,
}

func IsInvalidExperience(experience Experience) bool {
	return !allExperiences[experience]
}

func AllExperiencesString() string {
	return strings.Join([]string{
		string(Recruit),
		string(Rookie),
		string(Intermediate),
		string(Regular),
		string(Veteran),
		string(Elite),
	}, ", ")
}

var allGenders = map[Gender]bool{
	Female:      true,
	Male:        true,
	Unspecified: true,
}

func IsInvalidGender(gender Gender) bool {
	return !allGenders[gender]
}

func AllGendersString() string {
	return strings.Join([]string{
		string(Female),
		string(Male),
		string(Unspecified),
	}, ", ")
}

var allRoles = map[Role]bool{
	Diplomat:    true,
	Engineer:    true,
	Entertainer: true,
	Gunner:      true,
	Leader:      true,
	Marine:      true,
	Medic:       true,
	Navigator:   true,
	Pilot:       true,
	Scout:       true,
	Steward:     true,
	Technician:  true,
	Thug:        true,
	Trader:      true,
}

func IsInvalidRole(role Role) bool {
	return !allRoles[role]
}

func AllRolesString() string {
	return strings.Join([]string{
		string(Diplomat),
		string(Engineer),
		string(Entertainer),
		string(Gunner),
		string(Leader),
		string(Marine),
		string(Medic),
		string(Navigator),
		string(Pilot),
		string(Scout),
		string(Steward),
		string(Technician),
		string(Thug),
		string(Trader),
	}, ", ")
}
