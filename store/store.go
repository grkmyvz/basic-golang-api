package store

import (
	"fmt"
	"randgo/database"
	"randgo/utils"
	"time"

	"gorm.io/gorm"
)

type Store struct {
	Users           []utils.User
	Companies       []utils.Company
	CompanyServices []utils.CompanyService
	Appointments    []utils.Appointment
	Comments        []utils.Comment
}

var DBStore *Store

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUser(ID uint64) *utils.User {

	for i := range s.Users {
		if s.Users[i].ID == ID {
			return &s.Users[i]
		}
	}

	var user utils.User
	if err := database.Connection.Model(&utils.User{}).Where("ID = ?", ID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for user ID: %d\n", ID)
			return nil
		} else {
			fmt.Printf("Error while getting user ID: %d\n", ID)
			panic(err)
		}
	}
	s.Users = append(s.Users, user)

	return &user
}

func (s *Store) GetCompany(ID uint64) *utils.Company {
	for i := range s.Companies {
		if s.Companies[i].ID == ID {
			return &s.Companies[i]
		}
	}

	var company utils.Company
	if err := database.Connection.Model(&utils.Company{}).Where("ID = ?", ID).First(&company).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for company ID: %d\n", ID)
			return nil
		} else {
			fmt.Printf("Error while getting company ID: %d\n", ID)
			panic(err)
		}
	}

	s.Companies = append(s.Companies, company)

	return &company
}

func (s *Store) GetCompanyService(ID uint64) *utils.CompanyService {
	for i := range s.CompanyServices {
		if s.CompanyServices[i].ID == ID {
			return &s.CompanyServices[i]
		}
	}

	var companyService utils.CompanyService
	if err := database.Connection.Model(&utils.CompanyService{}).Where("ID = ?", ID).First(&companyService).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for companyService ID: %d\n", ID)
			return nil
		} else {
			fmt.Printf("Error while getting companyService ID: %d\n", ID)
			panic(err)
		}
	}

	s.CompanyServices = append(s.CompanyServices, companyService)

	return &companyService
}

func (s *Store) GetAppointment(ID uint64) *utils.Appointment {
	for i := range s.Appointments {
		if s.Appointments[i].ID == ID {
			return &s.Appointments[i]
		}
	}

	var appointment utils.Appointment
	if err := database.Connection.Model(&utils.Appointment{}).Where("ID = ?", ID).First(&appointment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for appointment ID: %d\n", ID)
			return nil
		} else {
			fmt.Printf("Error while getting appointment ID: %d\n", ID)
			panic(err)
		}
	}

	s.Appointments = append(s.Appointments, appointment)

	return &appointment
}

func (s *Store) GetComment(ID uint64) *utils.Comment {
	for i := range s.Comments {
		if s.Comments[i].ID == ID {
			return &s.Comments[i]
		}
	}

	var comment utils.Comment
	if err := database.Connection.Model(&utils.Comment{}).Where("ID = ?", ID).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("No rows were returned for comment ID: %d\n", ID)
			return nil
		} else {
			fmt.Printf("Error while getting comment ID: %d\n", ID)
			panic(err)
		}
	}

	s.Comments = append(s.Comments, comment)

	return &comment
}

func (s *Store) InsertUser(user utils.User) error {
	if err := database.Connection.Create(&user).Error; err != nil {
		return err
	}

	s.Users = append(s.Users, user)

	return nil
}

func (s *Store) InsertCompany(company utils.Company) error {
	if err := database.Connection.Create(&company).Error; err != nil {
		return err
	}

	s.Companies = append(s.Companies, company)

	return nil
}

func (s *Store) InsertCompanyService(companyService utils.CompanyService) error {
	if err := database.Connection.Create(&companyService).Error; err != nil {
		return err
	}

	s.CompanyServices = append(s.CompanyServices, companyService)

	return nil
}

func (s *Store) InsertAppointment(appointment utils.Appointment) error {
	if err := database.Connection.Create(&appointment).Error; err != nil {
		return err
	}

	s.Appointments = append(s.Appointments, appointment)

	return nil
}

func (s *Store) InsertComment(comment utils.Comment) error {
	if err := database.Connection.Create(&comment).Error; err != nil {
		return err
	}

	s.Comments = append(s.Comments, comment)

	return nil
}

func ClearStore() {
	clearCount := 100
	for {
		time.Sleep(10 * time.Second)
		if len(DBStore.Users) > clearCount {
			DBStore.Users = DBStore.Users[:clearCount]
		}

		if len(DBStore.Companies) > clearCount {
			DBStore.Companies = DBStore.Companies[:clearCount]
		}

		if len(DBStore.CompanyServices) > clearCount {
			DBStore.CompanyServices = DBStore.CompanyServices[:clearCount]
		}

		if len(DBStore.Appointments) > clearCount {
			DBStore.Appointments = DBStore.Appointments[:clearCount]
		}

		if len(DBStore.Comments) > clearCount {
			DBStore.Comments = DBStore.Comments[:clearCount]
		}
	}
}
