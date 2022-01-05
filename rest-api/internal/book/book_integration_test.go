//go:build integration
// +build integration

package book_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "gorm.io/driver/sqlite"

	"github.com/wfen/go-testing-bible/rest-api/internal/book"
	"github.com/wfen/go-testing-bible/rest-api/internal/database"
	"github.com/wfen/go-testing-bible/rest-api/internal/transport"
)

type BookTestSuite struct {
	suite.Suite
	h *transport.Handler
}

func (suite *BookTestSuite) SetupSuite() {
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}
	err = database.MigrateDB(db)
	if err != nil {
		panic("Failed to perform database migration")
	}

	bookSvc := book.NewService(db)
	suite.h = transport.Setup(bookSvc)
}

func (suite *BookTestSuite) TearDownSuite() {
	err := os.Remove("books.db")
	if err != nil {
		panic("Failed to remove database")
	}
}

func (suite *BookTestSuite) TestCreateBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": 12345, "title":"Test Book", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	// Perform the request plain with the *fiber.App,
	// the second argument is a request latency (set to -1 for no latency)
	res, err := suite.h.App.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var bookTest book.Book
	suite.h.Svc.DB.Where("title = ?", "Test Book").First(&bookTest)
	fmt.Println(bookTest)
	assert.Equal(suite.T(), bookTest.Title, "Test Book")
}

func (suite *BookTestSuite) TestReadBook() {
	req := httptest.NewRequest(
		"GET",
		"/api/v1/book/1",
		nil,
	)
	res, err := suite.h.App.Test(req, -1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var testbook book.Book
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &testbook)

	assert.Equal(suite.T(), "Test Book", testbook.Title)
}

func (suite *BookTestSuite) TestDeleteBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": 45678, "title":"Test Book 2", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.h.App.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var bookTest book.Book
	suite.h.Svc.DB.Where("title = ?", "Test Book 2").First(&bookTest)
	fmt.Println(bookTest)

	req = httptest.NewRequest(
		"DELETE",
		"/api/v1/book/45678",
		nil,
	)
	res, err = suite.h.App.Test(req, -1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
