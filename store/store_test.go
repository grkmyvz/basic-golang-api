package store

import (
	"fmt"
	"randgo/database"
	"randgo/utils"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSetup(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to connect to the database:", err)
	}

	err = db.AutoMigrate(
		&utils.Login{},
		&utils.User{},
		&utils.Company{},
		&utils.CompanyService{},
		&utils.Appointment{},
		&utils.Comment{})
	if err != nil {
		panic(err)
	}

	database.Connection = db
	DBStore = NewStore()
}

func TestLogin(t *testing.T) {
	TestSetup(t)

	for _, login := range logins {
		database.Connection.Create(&login)
	}

	for _, expectedLogin := range logins {
		var loginData utils.Login
		if err := database.Connection.Model(&utils.Login{}).Where("ID = ?", expectedLogin.ID).First(&loginData).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				t.Errorf("No rows were returned for user ID: %d", expectedLogin.ID)
			} else {
				t.Errorf("Error while getting user ID: %d", expectedLogin.ID)
			}
		}

		if loginData.ID != expectedLogin.ID || loginData.Password != expectedLogin.Password {
			fmt.Printf("Mismatch in login data for user ID: %d\n", expectedLogin.ID)
			t.Fail()
		}
	}

	t.Logf("TestLogin passed")
}

func TestUser(t *testing.T) {
	TestSetup(t)

	for _, user := range users {
		database.Connection.Create(&user)
	}

	for _, expectedUser := range users {
		userData := DBStore.GetUser(expectedUser.ID)
		if userData == nil {
			t.Errorf("User not found: %d", expectedUser.ID)
		} else {
			if userData.ID != expectedUser.ID || userData.Mail != expectedUser.Mail || userData.Name != expectedUser.Name || userData.Surname != expectedUser.Surname || userData.Phone != expectedUser.Phone || userData.BirthDate != expectedUser.BirthDate || userData.IsVerified != expectedUser.IsVerified || userData.CreatedAt != expectedUser.CreatedAt {
				t.Errorf("Mismatch in user data for user ID: %d", expectedUser.ID)
			}
		}
	}

	t.Logf("TestUser passed")
}

func TestCompany(t *testing.T) {
	TestSetup(t)

	for _, company := range companies {
		database.Connection.Create(&company)
	}

	for _, expectedCompany := range companies {
		companyData := DBStore.GetCompany(expectedCompany.ID)
		if companyData == nil {
			t.Errorf("Company not found: %d", expectedCompany.ID)
		} else {
			if companyData.ID != expectedCompany.ID || companyData.UserID != expectedCompany.UserID || companyData.Name != expectedCompany.Name || companyData.Sector != expectedCompany.Sector || companyData.Phone != expectedCompany.Phone || companyData.IsVerified != expectedCompany.IsVerified || companyData.CompanyCreatedAt != expectedCompany.CompanyCreatedAt || companyData.CreatedAt != expectedCompany.CreatedAt {
				t.Errorf("Mismatch in company data for company ID: %d", expectedCompany.ID)
			}
		}
	}

	t.Logf("TestCompany passed")
}

func TestCompanyService(t *testing.T) {
	TestSetup(t)

	for _, service := range services {
		database.Connection.Create(&service)
	}

	for _, expectedService := range services {
		serviceData := DBStore.GetCompanyService(expectedService.ID)
		if serviceData == nil {
			t.Errorf("Company service not found: %d", expectedService.ID)
		} else {
			if serviceData.ID != expectedService.ID || serviceData.CompanyID != expectedService.CompanyID || serviceData.ServiceName != expectedService.ServiceName || serviceData.StartAt != expectedService.StartAt || serviceData.EndAt != expectedService.EndAt || serviceData.TimePerService != expectedService.TimePerService || serviceData.Price != expectedService.Price || serviceData.CreatedAt != expectedService.CreatedAt {
				t.Errorf("Mismatch in company service data for company service ID: %d", expectedService.ID)
			}
		}
	}

	t.Logf("TestCompanyService passed")
}

func TestAppointment(t *testing.T) {
	TestSetup(t)

	for _, appointment := range appointments {
		database.Connection.Create(&appointment)
	}

	for _, expectedAppointment := range appointments {
		appointmentData := DBStore.GetAppointment(expectedAppointment.ID)
		if appointmentData == nil {
			t.Errorf("Appointment not found: %d", expectedAppointment.ID)
		} else {
			if appointmentData.ID != expectedAppointment.ID || appointmentData.CompanyID != expectedAppointment.CompanyID || appointmentData.UserID != expectedAppointment.UserID || appointmentData.ServiceID != expectedAppointment.ServiceID || appointmentData.AppointmentAt != expectedAppointment.AppointmentAt || appointmentData.CreatedAt != expectedAppointment.CreatedAt {
				t.Errorf("Mismatch in appointment data for appointment ID: %d", expectedAppointment.ID)
			}
		}
	}

	t.Logf("TestAppointment passed")
}

func TestComment(t *testing.T) {
	TestSetup(t)

	for _, comment := range comments {
		database.Connection.Create(&comment)
	}

	for _, expectedComment := range comments {
		commentData := DBStore.GetComment(expectedComment.ID)
		if commentData == nil {
			t.Errorf("Comment not found: %d", expectedComment.ID)
		} else {
			if commentData.ID != expectedComment.ID || commentData.CompanyServiceID != expectedComment.CompanyServiceID || commentData.UserID != expectedComment.UserID || commentData.Comment != expectedComment.Comment || commentData.CreatedAt != expectedComment.CreatedAt {
				t.Errorf("Mismatch in comment data for comment ID: %d", expectedComment.ID)
			}
		}
	}

	t.Logf("TestComment passed")
}

var logins = []utils.Login{
	{
		ID:       1234567891,
		Password: "password1",
	},
	{
		ID:       1234567892,
		Password: "password2",
	},
	{
		ID:       1234567893,
		Password: "password3",
	},
	{
		ID:       1234567894,
		Password: "password4",
	},
	{
		ID:       1234567895,
		Password: "password5",
	},
}

var users = []utils.User{
	{
		ID:         1234567891,
		Mail:       "user1@mail.com",
		Name:       "name1",
		Surname:    "surname1",
		Phone:      "phone1",
		BirthDate:  "birthDate1",
		IsVerified: true,
		CreatedAt:  "11-02-2021",
	},
	{
		ID:         1234567892,
		Mail:       "user2@mail.com",
		Name:       "name2",
		Surname:    "surname2",
		Phone:      "phone2",
		BirthDate:  "birthDate2",
		IsVerified: true,
		CreatedAt:  "12-02-2021",
	},
	{
		ID:         1234567893,
		Mail:       "user3@mail.com",
		Name:       "name3",
		Surname:    "surname3",
		Phone:      "phone3",
		BirthDate:  "birthDate3",
		IsVerified: true,
		CreatedAt:  "13-02-2021",
	},
	{
		ID:         1234567894,
		Mail:       "user4@mail.com",
		Name:       "name4",
		Surname:    "surname4",
		Phone:      "phone4",
		BirthDate:  "birthDate4",
		IsVerified: true,
		CreatedAt:  "14-02-2021",
	},
	{
		ID:         1234567895,
		Mail:       "user5@mail.com",
		Name:       "name5",
		Surname:    "surname5",
		Phone:      "phone5",
		BirthDate:  "birthDate5",
		IsVerified: true,
		CreatedAt:  "15-02-2021",
	},
}

var companies = []utils.Company{
	{
		ID:               9234567891,
		UserID:           1234567891,
		Name:             "company1",
		Sector:           "sector1",
		Phone:            "phone1",
		IsVerified:       true,
		CompanyCreatedAt: "05-05-2021",
		CreatedAt:        "11-02-2023",
	},
	{
		ID:               9234567892,
		UserID:           1234567892,
		Name:             "company2",
		Sector:           "sector2",
		Phone:            "phone2",
		IsVerified:       true,
		CompanyCreatedAt: "05-05-2022",
		CreatedAt:        "12-02-2023",
	},
}

var services = []utils.CompanyService{
	{
		ID:             8234567891,
		CompanyID:      9234567891,
		ServiceName:    "service1",
		StartAt:        "10:00 AM",
		EndAt:          "12:00 PM",
		TimePerService: 30,
		Price:          50.0,
		CreatedAt:      "11-02-2023",
	},
	{
		ID:             8234567892,
		CompanyID:      9234567892,
		ServiceName:    "service2",
		StartAt:        "02:00 PM",
		EndAt:          "04:00 PM",
		TimePerService: 45,
		Price:          75.0,
		CreatedAt:      "12-02-2023",
	},
}

var appointments = []utils.Appointment{
	{
		ID:            7234567891,
		CompanyID:     9234567891,
		UserID:        1234567891,
		ServiceID:     1234567891,
		AppointmentAt: "11-03-2023 10:30 AM",
		CreatedAt:     "11-02-2023",
	},
	{
		ID:            7234567892,
		CompanyID:     9234567892,
		UserID:        1234567892,
		ServiceID:     1234567892,
		AppointmentAt: "11-03-2023 02:30 PM",
		CreatedAt:     "12-02-2023",
	},
}

var comments = []utils.Comment{
	{
		ID:               6234567891,
		CompanyServiceID: 8234567891,
		UserID:           1234567891,
		Comment:          "Great service!",
		CreatedAt:        "11-02-2023",
	},
	{
		ID:               6234567892,
		CompanyServiceID: 8234567892,
		UserID:           1234567892,
		Comment:          "Very professional.",
		CreatedAt:        "12-02-2023",
	},
}
