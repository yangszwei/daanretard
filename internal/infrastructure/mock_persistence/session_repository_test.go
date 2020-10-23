package mock_persistence_test

import (
	"daanretard/internal/infrastructure/mock_persistence"
	entity "daanretard/internal/service/session"
	"testing"
)

var (
	sessions *mock_persistence.SessionRepository
	testSession = entity.Session{
		UserID: 1,
	}
)

func TestNewSessionRepository(t *testing.T) {
	sessions = mock_persistence.NewSessionRepository()
}

func TestSessionRepository_InsertOne(t *testing.T) {
	err := sessions.InsertOne(&testSession)
	if err != nil {
		t.Error(err)
	}
}

func TestSessionRepository_FindOneByID(t *testing.T) {
	t.Run("find local", func(t *testing.T) {
		_, err := sessions.FindOneByID(testSession.ID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("not found", func(t *testing.T) {
		_, err := mock_persistence.NewSessionRepository().FindOneByID(testSession.ID)
		if err == nil || err.Error() != "record not found" {
			t.Error(err)
		}
	})
}

func TestSessionRepository_FindAllByUserID(t *testing.T) {
	sessions, err := sessions.FindAllByUserID(testSession.UserID)
	if err != nil {
		t.Error(err)
	}
	if sessions[0].UserID != testSession.UserID {
		t.Error(sessions)
	}
}

func TestSessionRepository_SaveOne(t *testing.T) {
	err := sessions.SaveOne(&testSession)
	if err != nil {
		t.Error(err)
	}
}

func TestSessionRepository_DeleteOne(t *testing.T) {
	err := sessions.DeleteOne(&testSession)
	if err != nil {
		t.Error(err)
	}
}