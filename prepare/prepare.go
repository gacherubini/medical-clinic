package prepare

import (
	"medical-clinic/models"
)

type AdminUser struct {
	Admin int          `json:"admin_id"`
	User  *models.User `json:"user"`
}

type PatientUser struct {
	Patient          int          `json:"patient_id"`
	HealthInsurences interface{}  `json:"health"`
	User             *models.User `json:"user"`
}

type DoctorUser struct {
	Doctor            int          `json:"doctor_id"`
	HealthInsurences  interface{}  `json:"health"`
	DoctorSpecialties string       `json:"doctor_specialties"`
	User              *models.User `json:"user"`
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

		responseData := &DoctorUser{
			Doctor:            doctor.DoctorID,
			HealthInsurences:  healthInsuranceName,
			DoctorSpecialties: doctor.Specialties,
			User:              user,
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

		responseData := &PatientUser{
			Patient:          patient.PatientID,
			HealthInsurences: healthInsuranceName,
			User:             user,
		}

		combinedData = append(combinedData, *responseData)

	}
	return combinedData
}

func PrepareAdmin(admins models.AdminSlice) []AdminUser {
	var combinedData []AdminUser

	for _, admin := range admins {
		user := admin.R.User

		responseData := &AdminUser{
			Admin: admin.AdminID,
			User:  user,
		}

		combinedData = append(combinedData, *responseData)
	}

	return combinedData
}
