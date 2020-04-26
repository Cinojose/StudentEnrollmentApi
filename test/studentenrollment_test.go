package student

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/viper"

	resty "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type StudentEnrollMentTest struct {
	suite.Suite
	APIClient *resty.Client
	host      string
	port      string
	scheme    string
}

// Setup testing
func (suite *StudentEnrollMentTest) SetupTest() {
	suite.APIClient = resty.New()
	viper.BindEnv("API_HOST")
	viper.BindEnv("API_PORT")

	suite.host = viper.GetString("API_HOST")
	suite.port = viper.GetString("API_PORT")
	suite.scheme = "http"
}

// Function to test if the  endpoint is up
func (suite *StudentEnrollMentTest) TestEndpontRunning() {

	url := fmt.Sprintf("%s://%s:%s/internal/apidocs", suite.scheme, suite.host, suite.port)
	resp, _ := suite.APIClient.R().Get(url)

	// Check if it return 200
	checkResponseCode(suite.T(), http.StatusOK, resp.StatusCode())

}

//Test for a non existing ID
func (suite *StudentEnrollMentTest) TestStudentFetchEnpoint() {

	url := fmt.Sprintf("%s://%s:%s/v2/fetchStudents", suite.scheme, suite.host, suite.port)
	resp, _ := suite.APIClient.R().SetQueryParams(
		map[string]string{
			"Id": "222",
		}).SetHeader("accept", "application/json").
		Get(url)
	//response := executeRequest(req)

	checkResponseCode(suite.T(), http.StatusOK, resp.StatusCode())

	if body := resp.String(); body == "[]" {
		suite.T().Logf("Got an empty array, test pass %s", body)
	}
}

//Test creation of a new student
func (suite *StudentEnrollMentTest) TestCreateStudentEndpoint() {
	url := fmt.Sprintf("%s://%s:%s/v2/student", suite.scheme, suite.host, suite.port)
	resp, _ := suite.APIClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
			"class": "3 A",
			"firstName": "mike",
			"id": 3333,
			"lastName": "wong",
			"nationality": "Singapore"
		  }`)).
		Post(url)
	checkResponseCode(suite.T(), http.StatusOK, resp.StatusCode())
}

//Test delete student endpoint
func (suite *StudentEnrollMentTest) TestDeleteStudentEndpoint() {
	url := fmt.Sprintf("%s://%s:%s/v2/student", suite.scheme, suite.host, suite.port)
	resp, _ := suite.APIClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
			"id": 3333
		  }`)).Delete(url)
	checkResponseCode(suite.T(), http.StatusOK, resp.StatusCode())
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestStudentEnrollMentSuite(t *testing.T) {
	suite.Run(t, new(StudentEnrollMentTest))
}
