package npc_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
			assert.NotEmpty(t, generatedNPC.FirstsName)
			assert.NotEmpty(t, generatedNPC.Surname)
			assert.Len(t, generatedNPC.Characteristics, 6)
			assert.Equal(t, tt.expectedCitizenCategory, generatedNPC.CitizenCategory)
			assert.Equal(t, tt.expectedExperience, generatedNPC.Experience)
			assert.Equal(t, tt.expectedRole, generatedNPC.Role)
			assert.NotEmpty(t, generatedNPC.Skills)
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
			name:   "Role is required",
			method: http.MethodPost,
			body: []byte(`{
				"citizen_category": "average",
				"experience":       "regular",
				"gender":           "unspecified"
			}`),
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "role is required",
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
	FirstsName      string         `json:"firsts_name"`
	Surname         string         `json:"surname"`
	Characteristics map[string]int `json:"characteristics"`
	CitizenCategory string         `json:"citizen_category"`
	Experience      string         `json:"experience"`
	Role            string         `json:"role"`
	Skills          []string       `json:"skills"`
}
