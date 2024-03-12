package prepare

import (
	"medical-clinic/models"
)

func PrepareDoctor(doctors models.DoctorSlice) []map[string]interface{} {

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User
		responseData := map[string]interface{}{
			"doctor_id":       doctor.DoctorID,
			"specialties":     doctor.Specialties,
			"healthInsurence": doctor.HealthinsuranceID,
			"user":            user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData
}

func PrepareInsurence(doctors models.DoctorSlice) []map[string]interface{} {

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User
		healthInsurance := doctor.R.Healthinsurance

		responseData := map[string]interface{}{
			"health_insurance": healthInsurance.Name,
			"specialties":      doctor.Specialties,
			"user":             user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData
}
