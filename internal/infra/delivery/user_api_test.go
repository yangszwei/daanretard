package delivery_test

import (
	"daanretard/internal/entity/user"
	"daanretard/internal/infra/delivery"
	"daanretard/internal/infra/mock_fbgraph"
	"daanretard/internal/infra/persistence"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	userFBToken    string
	userMockFB     = mock_fbgraph.New("/test", "test_id", "test_secret")
	testUserServer *httptest.Server
	testUserClient *http.Client
	testUserAPI    *delivery.Server
)

func init() {
	if err := godotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
}

func TestSetupUserAPI(t *testing.T) {
	testUserAPI = createTestServer()
	db, err := persistence.Open(os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
	u := user.NewService(persistence.NewUserRepo(db), userMockFB)
	delivery.SetupUserAPI(testUserAPI, u)
	userFBToken = userMockFB.NewUser("test user")
	testUserServer = httptest.NewServer(testUserAPI.Engine)
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	testUserClient = &http.Client{Jar: jar}
}

func TestUserAPI_GET_Auth(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		reader := strings.NewReader("access_token=" + userFBToken)
		req, err := http.NewRequest("POST", testUserServer.URL+"/api/user/auth", reader)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := testUserClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			data := make(map[string]interface{})
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				t.Error(resp)
			}
			t.Error(resp, data)
		}
	})
}

func TestUserAPI_GET(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		req, err := http.NewRequest("GET", testUserServer.URL+"/api/user?fields=id,name", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := testUserClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error(resp)
		}
	})
}

func TestUserAPI_Delete(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		reader := strings.NewReader("access_token=" + userFBToken)
		req, err := http.NewRequest("POST", testUserServer.URL+"/api/user/delete", reader)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := testUserClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error(resp)
		}
	})
}

func TestUserAPI_Delete_Auth(t *testing.T) {
	defer testUserServer.Close()
	t.Run("should succeed", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", testUserServer.URL+"/api/user/auth", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := testUserClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error(resp)
		}
	})
}
