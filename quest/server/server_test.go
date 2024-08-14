package questserver

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"log/slog"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	quest "github.com/lafetz/quest-micro/quest/core"
// 	//"github.com/stretchr/testify/assert"
// 	///om/stretchr/testify/assert"
// )

// type stubQuestService struct{}

// func (s *stubQuestService) AddQuest(ctx context.Context, q quest.Quest) (*quest.Quest, error) {
// 	return &q, nil
// }

// func (s *stubQuestService) GetAssignedQuests(ctx context.Context, username string) ([]*quest.Quest, error) {
// 	return []*quest.Quest{}, nil
// }

// func (s *stubQuestService) GetQuest(ctx context.Context, id uuid.UUID) (*quest.Quest, error) {
// 	return &quest.Quest{}, nil
// }

// func (s *stubQuestService) CompleteQuest(ctx context.Context, id uuid.UUID) error {
// 	return nil
// }
// func TestAddQuest(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	stubQuestService := &stubQuestService{}
// 	logger := slog.New(slog.Default().Handler())

// 	app := NewApp(stubQuestService, 8080, logger)

// 	questReq := addQuestReq{
// 		KntUsername: "knight1",
// 		Name:        "Quest Name",
// 		Owner:       "Owner Name",
// 		Description: "Quest Description",
// 	}

// 	body, _ := json.Marshal(questReq)
// 	req, err := http.NewRequest(http.MethodPost, "/quest", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatalf("could not create request: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	app.gin.ServeHTTP(w, req)

// 	if w.Code != http.StatusCreated {
// 		t.Errorf("expected status %v; got %v", http.StatusCreated, w.Code)
// 	}

// 	var resp map[string]interface{}
// 	err = json.Unmarshal(w.Body.Bytes(), &resp)
// 	if err != nil {
// 		t.Fatalf("could not parse response: %v", err)
// 	}

// 	if resp["msg"] != "Quest added" {
// 		t.Errorf("expected message %v; got %v", "Quest added", resp["msg"])
// 	}

// 	if resp["quest"] == nil {
// 		t.Errorf("expected quest to be not nil")
// 	}
// }
