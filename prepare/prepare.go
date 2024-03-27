package prepare

import (
	"medical-clinic/models"
)

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Email  string `json:"email"`
}

type AdminUser struct {
	Admin int   `json:"admin_id"`
	User  *User `json:"user"`
}

type PatientUser struct {
	Patient          int         `json:"patient_id"`
	HealthInsurences interface{} `json:"health"`
	User             *User       `json:"user"`
}

type DoctorUser struct {
	Doctor            int         `json:"doctor_id"`
	HealthInsurences  interface{} `json:"health"`
	DoctorSpecialties string      `json:"doctor_specialties"`
	User              *User       `json:"user"`
}

func PrepareDoctor(doctors models.DoctorSlice) []DoctorUser {
	var combinedData []DoctorUser

	for _, doctor := range doctors {
		user := doctor.R.User
		healthInsurance := doctor.R.Healthinsurance

		var healthInsuranceName interface{}
		if healthInsurance != nil {
			healthInsuranceName = healthInsurance.Name
		} else {
			healthInsuranceName = nil
		}

		User := &User{
			UserID: user.UserID,
			Name:   user.Name,
			Age:    user.Age,
			Gender: user.Gender,
			Phone:  user.Phone,
			Role:   user.Role,
			Email:  user.Email,
		}

		responseData := &DoctorUser{
			Doctor:            doctor.DoctorID,
			HealthInsurences:  healthInsuranceName,
			DoctorSpecialties: doctor.Specialties,
			User:              User,
		}

		combinedData = append(combinedData, *responseData)
	}
	return combinedData
}

func PreparePatient(patients models.PatientSlice) []PatientUser {
	var combinedData []PatientUser

	for _, patient := range patients {
		user := patient.R.User
		healthInsurance := patient.R.Healthinsurance

		var healthInsuranceName interface{}
		if healthInsurance != nil {
			healthInsuranceName = healthInsurance.Name
		} else {
			healthInsuranceName = nil
		}

		User := &User{
			UserID: user.UserID,
			Name:   user.Name,
			Age:    user.Age,
			Gender: user.Gender,
			Phone:  user.Phone,
			Role:   user.Role,
			Email:  user.Email,
		}

		responseData := &PatientUser{
			Patient:          patient.PatientID,
			HealthInsurences: healthInsuranceName,
			User:             User,
		}

		combinedData = append(combinedData, *responseData)

	}
	return combinedData
}

func PrepareAdmin(admins models.AdminSlice) []AdminUser {
	var combinedData []AdminUser

	for _, admin := range admins {
		user := admin.R.User

		User := &User{
			UserID: user.UserID,
			Name:   user.Name,
			Age:    user.Age,
			Gender: user.Gender,
			Phone:  user.Phone,
			Role:   user.Role,
			Email:  user.Email,
		}

		responseData := &AdminUser{
			Admin: admin.AdminID,
			User:  User,
		}

		combinedData = append(combinedData, *responseData)
	}

	return combinedData
}
