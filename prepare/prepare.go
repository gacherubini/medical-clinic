package prepare

import (
	"errors"
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

func PrepareInsurence(doctors models.DoctorSlice) ([]map[string]interface{}, error) {

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User
		healthInsurance := doctor.R.Healthinsurance

		if healthInsurance == nil {
			return nil, errors.New("health insurance is nil")
		}

		responseData := map[string]interface{}{
			"health_insurance": healthInsurance.Name,
			"specialties":      doctor.Specialties,
			"user":             user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData, nil
}
