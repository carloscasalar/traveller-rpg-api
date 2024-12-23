package npc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carloscasalar/traveller-npc-generator/pkg/generator"
	"github.com/carloscasalar/traveller-rpg-api/pkg/apirest"
)

// SingleHandler is the handler to generate a single NPC
func SingleHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != http.MethodPost {
		setError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var npcRequest apirest.NPCRequest
	if err := json.NewDecoder(req.Body).Decode(&npcRequest); err != nil {
		setError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	npcGenerator, err := generator.NewNpcGeneratorBuilder().Build()
	if err != nil {
		setError(w, "failed to create NPC generator", http.StatusInternalServerError)
		return
	}

	request, err := buildCharacterRequest(npcRequest)
	if err != nil {
		setError(w, err.Error(), http.StatusBadRequest)
		return
	}

	generated, err := npcGenerator.Generate(*request)
	if err != nil {
		setError(w, fmt.Sprintf("unable to generate NPC: %s", err.Error()), http.StatusBadRequest)
		return
	}
	npc := apirest.NPC{
		FirstName:       generated.FirstName(),
		Surname:         generated.Surname(),
		Role:            toRestRole(generated.Role()),
		CitizenCategory: toRestCitizenCategory(generated.CitizenCategory()),
		Experience:      toRestExperience(generated.Experience()),
		Skills:          generated.Skills(), // Example skills
		Characteristics: toRestCharacteristics(generated.Characteristics()),
	}

	if err := json.NewEncoder(w).Encode(npc); err != nil {
		setError(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func buildCharacterRequest(npcRequest apirest.NPCRequest) (*generator.GenerateCharacterRequest, error) {
	if npcRequest.Role == "" {
		return nil, fmt.Errorf("role is required")
	}

	if apirest.IsInvalidRole(npcRequest.Role) {
		return nil, fmt.Errorf("role is invalid, must be one of: %s", apirest.AllRolesString())
	}

	if npcRequest.CitizenCategory != nil && apirest.IsInvalidCitizenCategory(*npcRequest.CitizenCategory) {
		return nil, fmt.Errorf("citizen_category is invalid, must be one of: %s", apirest.AllCitizenCategoriesString())
	}

	if npcRequest.Experience != nil && apirest.IsInvalidExperience(*npcRequest.Experience) {
		return nil, fmt.Errorf("experience is invalid, must be one of: %s", apirest.AllExperiencesString())
	}

	if npcRequest.Gender != nil && apirest.IsInvalidGender(*npcRequest.Gender) {
		return nil, fmt.Errorf("gender is invalid, must be one of: %s", apirest.AllGendersString())
	}

	request := generator.NewGenerateCharacterRequestBuilder().
		Role(toRole(npcRequest.Role)).
		Experience(toExperience(npcRequest.Experience)).
		CitizenCategory(toCitizenCategory(npcRequest.CitizenCategory)).
		Gender(toGender(npcRequest.Gender)).
		Build()

	return request, nil
}

func toRestCharacteristics(characteristics map[generator.Characteristic]int) apirest.Characteristics {
	return apirest.Characteristics{
		STR: characteristics[generator.STR],
		DEX: characteristics[generator.DEX],
		END: characteristics[generator.END],
		INT: characteristics[generator.INT],
		EDU: characteristics[generator.EDU],
		SOC: characteristics[generator.SOC],
	}
}

func toRestExperience(experience generator.Experience) apirest.Experience {
	switch experience {
	case generator.ExperienceRecruit:
		return apirest.Recruit
	case generator.ExperienceRookie:
		return apirest.Rookie
	case generator.ExperienceIntermediate:
		return apirest.Intermediate
	case generator.ExperienceRegular:
		return apirest.Regular
	case generator.ExperienceVeteran:
		return apirest.Veteran
	case generator.ExperienceElite:
		return apirest.Elite
	default:
		return apirest.Regular
	}
}

func toRestRole(role generator.Role) apirest.Role {
	switch role {
	case generator.RolePilot:
		return apirest.Pilot
	case generator.RoleNavigator:
		return apirest.Navigator
	case generator.RoleEngineer:
		return apirest.Engineer
	case generator.RoleSteward:
		return apirest.Steward
	case generator.RoleMedic:
		return apirest.Medic
	case generator.RoleMarine:
		return apirest.Marine
	case generator.RoleGunner:
		return apirest.Gunner
	case generator.RoleScout:
		return apirest.Scout
	case generator.RoleTechnician:
		return apirest.Technician
	case generator.RoleLeader:
		return apirest.Leader
	case generator.RoleDiplomat:
		return apirest.Diplomat
	case generator.RoleEntertainer:
		return apirest.Entertainer
	case generator.RoleTrader:
		return apirest.Trader
	case generator.RoleThug:
		return apirest.Thug
	default:
		return apirest.Pilot
	}
}

func toRestCitizenCategory(category generator.CitizenCategory) apirest.CitizenCategory {
	switch category {
	case generator.CitizenCategoryBelowAverage:
		return apirest.BelowAverage
	case generator.CitizenCategoryAverage:
		return apirest.Average
	case generator.CitizenCategoryAboveAverage:
		return apirest.AboveAverage
	case generator.CitizenCategoryExceptional:
		return apirest.Exceptional
	default:
		return apirest.Average
	}
}

func toGender(gender *apirest.Gender) generator.Gender {
	if gender == nil {
		return generator.GenderUnspecified
	}

	switch *gender {
	case apirest.Female:
		return generator.GenderFemale
	case apirest.Male:
		return generator.GenderMale
	default:
		return generator.GenderUnspecified
	}
}

func toRole(role apirest.Role) generator.Role {
	switch role {
	case apirest.Pilot:
		return generator.RolePilot
	case apirest.Navigator:
		return generator.RoleNavigator
	case apirest.Engineer:
		return generator.RoleEngineer
	case apirest.Steward:
		return generator.RoleSteward
	case apirest.Medic:
		return generator.RoleMedic
	case apirest.Marine:
		return generator.RoleMarine
	case apirest.Gunner:
		return generator.RoleGunner
	case apirest.Scout:
		return generator.RoleScout
	case apirest.Technician:
		return generator.RoleTechnician
	case apirest.Leader:
		return generator.RoleLeader
	case apirest.Diplomat:
		return generator.RoleDiplomat
	case apirest.Entertainer:
		return generator.RoleEntertainer
	case apirest.Trader:
		return generator.RoleTrader
	case apirest.Thug:
		return generator.RoleThug
	default:
		return generator.RolePilot
	}
}

func toExperience(experience *apirest.Experience) generator.Experience {
	if experience == nil {
		return generator.ExperienceRegular
	}
	switch *experience {
	case apirest.Recruit:
		return generator.ExperienceRecruit
	case apirest.Rookie:
		return generator.ExperienceRookie
	case apirest.Intermediate:
		return generator.ExperienceIntermediate
	case apirest.Regular:
		return generator.ExperienceRegular
	case apirest.Veteran:
		return generator.ExperienceVeteran
	case apirest.Elite:
		return generator.ExperienceElite
	default:
		return generator.ExperienceRegular
	}
}

func toCitizenCategory(category *apirest.CitizenCategory) generator.CitizenCategory {
	if category == nil {
		return generator.CitizenCategoryAverage
	}
	switch *category {
	case apirest.BelowAverage:
		return generator.CitizenCategoryBelowAverage
	case apirest.Average:
		return generator.CitizenCategoryAverage
	case apirest.AboveAverage:
		return generator.CitizenCategoryAboveAverage
	case apirest.Exceptional:
		return generator.CitizenCategoryExceptional
	default:
		return generator.CitizenCategoryAverage
	}
}

func setError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}
