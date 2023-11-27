package utils

type Login struct {
	ID       uint64 `json:"id"`
	Password string `json:"password"`
}

type User struct {
	ID         uint64 `json:"id"`
	Mail       string `json:"mail"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Phone      string `json:"phone"`
	BirthDate  string `json:"birthDate"`
	IsVerified bool   `json:"isVerified"`
	CreatedAt  string `json:"createdAt"`
}

type Company struct {
	ID               uint64 `json:"id"`
	UserID           uint64 `json:"userId"`
	Name             string `json:"name"`
	Sector           string `json:"sector"`
	Phone            string `json:"phone"`
	IsVerified       bool   `json:"isVerified"`
	CompanyCreatedAt string `json:"companyCreatedAt"`
	CreatedAt        string `json:"createdAt"`
}

type CompanyService struct {
	ID             uint64  `json:"id"`
	CompanyID      uint64  `json:"companyId"`
	ServiceName    string  `json:"serviceName"`
	StartAt        string  `json:"StartAt"`
	EndAt          string  `json:"EndAt"`
	TimePerService uint64  `json:"timePerService"`
	Price          float64 `json:"price"`
	CreatedAt      string  `json:"createdAt"`
}

type Appointment struct {
	ID            uint64 `json:"id"`
	CompanyID     uint64 `json:"companyId"`
	UserID        uint64 `json:"userId"`
	ServiceID     uint64 `json:"serviceId"`
	AppointmentAt string `json:"appointmentAt"`
	CreatedAt     string `json:"createdAt"`
}

type Comment struct {
	ID               uint64 `json:"id"`
	CompanyServiceID uint64 `json:"companyServiceId"`
	UserID           uint64 `json:"userId"`
	Comment          string `json:"comment"`
	CreatedAt        string `json:"createdAt"`
}
