package npc_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carloscasalar/traveller-rpg-api/pkg/apirest"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/carloscasalar/traveller-rpg-api/internal/npc"
)

func TestNPCSingleHandler_when_request_is_valid(t *testing.T) {
	tests := []struct {
		name                    string
		body                    []byte
		expectedRole            string
		expectedCitizenCategory string
		expectedExperience      string
		expectedGender          string
	}{
		{
			name: "and all fields are provided, generated NPC should have the specified values",
			body: []byte(`{
				"role":             "pilot",
				"citizen_category": "above_average",
				"experience":       "rookie",
				"gender":           "unspecified"
			}`),
			expectedRole:            "pilot",
			expectedCitizenCategory: "above_average",
			expectedExperience:      "rookie",
		},
		{
			name: "and citizen category is not provided, generated NPC should have average citizen category",
			body: []byte(`{
				"role":             "pilot",
				"experience":       "rookie",
				"gender":           "unspecified"
			}`),
			expectedRole:            "pilot",
			expectedCitizenCategory: "average",
			expectedExperience:      "rookie",
		},
		{
			name: "and experience is not provided, generated NPC should have regular experience",
			body: []byte(`{
				"role":             "pilot",
				"citizen_category": "above_average",
				"gender":           "unspecified"
			}`),
			expectedRole:            "pilot",
			expectedCitizenCategory: "above_average",
			expectedExperience:      "regular",
		},
		{
			name: "and gender is not provided, it should generate a NPC anyway",
			body: []byte(`{
				"role":             "pilot",
				"citizen_category": "above_average",
				"experience":       "rookie"
			}`),
			expectedRole:            "pilot",
			expectedCitizenCategory: "above_average",
			expectedExperience:      "rookie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/npc/single", bytes.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(npc.SingleHandler)
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
			generatedNPC := extractNPC(t, rr.Body.String())
			assert.NotEmpty(t, generatedNPC.FirstName)
			assert.NotEmpty(t, generatedNPC.Surname)
			assert.Len(t, generatedNPC.Characteristics, 6)
			assert.Equal(t, tt.expectedCitizenCategory, generatedNPC.CitizenCategory)
			assert.Equal(t, tt.expectedExperience, generatedNPC.Experience)
			assert.Equal(t, tt.expectedRole, generatedNPC.Role)
			assert.NotEmpty(t, generatedNPC.Skills)
		})
	}
}

func TestNPCSingleHandler_should_properly_map_citizen_category(t *testing.T) {
	tests := []struct {
		requestedCategory apirest.CitizenCategory
		expectedCategory  string
	}{
		{apirest.BelowAverage, "below_average"},
		{apirest.Average, "average"},
		{apirest.AboveAverage, "above_average"},
		{apirest.Exceptional, "exceptional"},
	}

	for _, tt := range tests {
		t.Run(string(tt.requestedCategory), func(t *testing.T) {
			payload := []byte(`{
				"citizen_category": "` + tt.requestedCategory + `",
				"role":             "pilot"
            }`)
			req := httptest.NewRequest(http.MethodPost, "/npc/single", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(npc.SingleHandler)
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)
			generatedNPC := extractNPC(t, rr.Body.String())
			assert.Equal(t, tt.expectedCategory, generatedNPC.CitizenCategory)
		})
	}
}

func TestNPCSingleHandler_should_properly_map_experience(t *testing.T) {
	tests := []struct {
		requestedExperience apirest.Experience
		expectedExperience  string
	}{
		{apirest.Recruit, "recruit"},
		{apirest.Rookie, "rookie"},
		{apirest.Intermediate, "intermediate"},
		{apirest.Regular, "regular"},
		{apirest.Veteran, "veteran"},
		{apirest.Elite, "elite"},
	}

	for _, tt := range tests {
		t.Run(string(tt.requestedExperience), func(t *testing.T) {
			payload := []byte(`{
				"experience": "` + tt.requestedExperience + `",
				"role":             "pilot"
            }`)
			req := httptest.NewRequest(http.MethodPost, "/npc/single", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(npc.SingleHandler)
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)
			generatedNPC := extractNPC(t, rr.Body.String())
			assert.Equal(t, tt.expectedExperience, generatedNPC.Experience)
		})
	}
}

func TestNPCSingleHandler_should_properly_map_role(t *testing.T) {
	tests := []struct {
		requestedRole apirest.Role
		expectedRole  string
	}{
		{apirest.Diplomat, "diplomat"},
		{apirest.Engineer, "engineer"},
		{apirest.Entertainer, "entertainer"},
		{apirest.Gunner, "gunner"},
		{apirest.Leader, "leader"},
		{apirest.Marine, "marine"},
		{apirest.Medic, "medic"},
		{apirest.Navigator, "navigator"},
		{apirest.Pilot, "pilot"},
		{apirest.Scout, "scout"},
		{apirest.Steward, "steward"},
		{apirest.Technician, "technician"},
		{apirest.Thug, "thug"},
		{apirest.Trader, "trader"},
	}

	for _, tt := range tests {
		t.Run(string(tt.requestedRole), func(t *testing.T) {
			payload := []byte(`{
				"role": "` + tt.requestedRole + `"
            }`)
			req := httptest.NewRequest(http.MethodPost, "/npc/single", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(npc.SingleHandler)
			handler.ServeHTTP(rr, req)

			require.Equal(t, http.StatusOK, rr.Code)
			generatedNPC := extractNPC(t, rr.Body.String())
			assert.Equal(t, tt.expectedRole, generatedNPC.Role)
		})
	}
}

func TestNPCSingleHandler_when_request_is_invalid(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		body            []byte
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:            "because method is not allowed message should be method not allowed",
			method:          http.MethodGet,
			body:            []byte{},
			expectedStatus:  http.StatusMethodNotAllowed,
			expectedMessage: "method not allowed",
		},
		{
			name:   "because role is not provided message should be role is required",
			method: http.MethodPost,
			body: []byte(`{
				"citizen_category": "average",
				"experience":       "regular",
				"gender":           "unspecified"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "role is required",
		},
		{
			name:   "because citizen category is invalid message should properly say so",
			method: http.MethodPost,
			body: []byte(`{
				"citizen_category": "unknown category",

				"role":             "diplomat",
				"experience":       "regular",
				"gender":           "unspecified"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "citizen_category is invalid, must be one of: below_average, average, above_average, exceptional",
		},
		{
			name:   "because experience is invalid message should properly say so",
			method: http.MethodPost,
			body: []byte(`{
				"experience":       "unknown experience",

				"citizen_category": "average",
				"role":             "diplomat",
				"gender":           "unspecified"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "experience is invalid, must be one of: recruit, rookie, intermediate, regular, veteran, elite",
		},
		{
			name:   "because gender is invalid message should properly say so",
			method: http.MethodPost,
			body: []byte(`{
				"gender":           "unknown gender",

				"citizen_category": "average",
				"experience":       "regular",
				"role":             "diplomat"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "gender is invalid, must be one of: female, male, unspecified",
		},
		{
			name:   "because role is invalid message should properly say so",
			method: http.MethodPost,
			body: []byte(`{
				"role":             "unknown role",

				"citizen_category": "average",
				"experience":       "regular",
				"gender":           "unspecified"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "role is invalid, must be one of: diplomat, engineer, entertainer, gunner, leader, marine, medic, navigator, pilot, scout, steward, technician, thug, trader",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/npc/single", bytes.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(npc.SingleHandler)
			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
			assert.Equal(t, tt.expectedMessage, extractMessage(t, rr.Body.String()))
		})
	}
}

// extractMessage is a helper function to extract the message from the error JSON body in a payload in the form of {"message": "error message"}
func extractMessage(t *testing.T, errorJSONBody string) string {
	var payload errorPayload
	err := json.Unmarshal([]byte(errorJSONBody), &payload)
	require.NoError(t, err)
	return payload.Message
}

func extractNPC(t *testing.T, jsonBody string) npcPayload {
	var payload npcPayload
	err := json.Unmarshal([]byte(jsonBody), &payload)
	require.NoError(t, err)
	return payload
}

type errorPayload struct {
	Message string `json:"message"`
}
type npcPayload struct {
	FirstName       string         `json:"first_name"`
	Surname         string         `json:"surname"`
	Characteristics map[string]int `json:"characteristics"`
	CitizenCategory string         `json:"citizen_category"`
	Experience      string         `json:"experience"`
	Role            string         `json:"role"`
	Skills          []string       `json:"skills"`
}
