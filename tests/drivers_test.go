package tests

import (
	"bytes"
	"database/sql"
	"delivery-system/drivers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDrivers(t *testing.T) {
	router, sqlMock := setupRouter("api/v1/", "/drivers/fetch/:license_number", GET, drivers.GetDriverHandler)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/drivers/fetch/0DXAMBL3ZB", nil)
	if err != nil {
		t.Errorf("error while creating new request")
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"full_name",
		"license_number",
		"primary_phone_number",
		"primary_phone_country_code",
		"secondary_phone_number",
		"secondary_phone_country_code",
		"email",
		"status",
	}).AddRow(1, "Bbxndxqxw Qnyclf", "0DXAMBL3ZB", "4529235181", "+111", "1181291757", "+023", "dq35k@kyaex.fto", 5)
	sqlMock.ExpectQuery("^SELECT (.+?) FROM drivers WHERE license_number = \\?$").
		WithArgs("0DXAMBL3ZB").
		WillReturnRows(rows)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)
	var response drivers.DriverRequest
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	expectedResponse, err := json.Marshal(drivers.DriverRequest{
		ID:                        1,
		FullName:                  "Bbxndxqxw Qnyclf",
		LicenseNumber:             "0DXAMBL3ZB",
		PrimaryPhoneNumber:        "4529235181",
		PrimaryPhoneCountryCode:   "+111",
		SecondaryPhoneNumber:      "1181291757",
		SecondaryPhoneCountryCode: "+023",
		Email:                     "dq35k@kyaex.fto",
		Status:                    5,
	})
	assert.NoError(t, err)
	generatedResponse, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedResponse), string(generatedResponse))
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGetDriversMissingCol(t *testing.T) {
	router, sqlMock := setupRouter("api/v1/", "/drivers/fetch/:license_number", GET, drivers.GetDriverHandler)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/drivers/fetch/0DXAMBL3ZB", nil)
	if err != nil {
		t.Errorf("error while creating new request")
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"full_name",
		"license_number",
		"primary_phone_number",
		"primary_phone_country_code",
		"secondary_phone_number",
		"secondary_phone_country_code",
		"email",
		"status",
	}).AddRow(1, "Bbxndxqxw Qnyclf", "0DXAMBL3ZB", nil, "+111", nil, "+023", "dq35k@kyaex.fto", 5)
	sqlMock.ExpectQuery("^SELECT (.+?) FROM drivers WHERE license_number = \\?$").
		WithArgs("0DXAMBL3ZB").
		WillReturnRows(rows)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)
	var response drivers.DriverRequest
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	expectedResponse, err := json.Marshal(drivers.DriverRequest{
		ID:                        1,
		FullName:                  "Bbxndxqxw Qnyclf",
		LicenseNumber:             "0DXAMBL3ZB",
		PrimaryPhoneNumber:        "",
		PrimaryPhoneCountryCode:   "+111",
		SecondaryPhoneNumber:      "",
		SecondaryPhoneCountryCode: "+023",
		Email:                     "dq35k@kyaex.fto",
		Status:                    5,
	})
	assert.NoError(t, err)
	generatedResponse, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedResponse), string(generatedResponse))
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGetDriversNoEntry(t *testing.T) {
	router, sqlMock := setupRouter("api/v1/", "/drivers/fetch/:license_number", GET, drivers.GetDriverHandler)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/drivers/fetch/0DXAMBL3ZB", nil)
	if err != nil {
		t.Errorf("error while creating new request")
	}
	sqlMock.ExpectQuery("^SELECT (.+?) FROM drivers WHERE license_number = \\?$").
		WithArgs("0DXAMBL3ZB").
		WillReturnError(sql.ErrNoRows)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestGetDriversInternalError(t *testing.T) {
	router, sqlMock := setupRouter("api/v1/", "/drivers/fetch/:license_number", GET, drivers.GetDriverHandler)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/drivers/fetch/0DXAMBL3ZB", nil)
	if err != nil {
		t.Errorf("error while creating new request")
	}
	sqlMock.ExpectQuery("^SELECT (.+?) FROM drivers WHERE license_number = \\?$").
		WithArgs("0DXAMBL3ZB").
		WillReturnError(sql.ErrConnDone)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestRegisterDrivers(t *testing.T) {
	router, sqlMock := setupRouter("api/v1/", "/drivers/register", POST, drivers.RegisterDriverHandler)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/drivers/register", bytes.NewBuffer(
		[]byte(`{
			"full_name": "Bbxndxqxw Qnyclf",
			"license_number": "0DXAMBL3ZB",
			"primary_phone_number" : "4529235181",
			"primary_phone_country_code": "+111",
			"secondary_phone_number": "1181291757",
			"secondary_phone_country_code": "+023",
			"email": "dq35k@kyaex.fto",
			"status": 0
		}`),
	))
	if err != nil {
		t.Errorf("error while creating new request")
	}
	sqlMock.ExpectExec("^INSERT INTO drivers \\(full_name, license_number, primary_phone_country_code, primary_phone_number, secondary_phone_country_code, secondary_phone_number, email, status\\)VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)$").
		WithArgs("Bbxndxqxw Qnyclf", "0DXAMBL3ZB", "+111", "4529235181", "+023", "1181291757", "dq35k@kyaex.fto", 0).
		WillReturnResult(sqlmock.NewResult(0, 1))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
