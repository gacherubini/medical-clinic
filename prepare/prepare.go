package prepare

import (
	"medical-clinic/models"
)

func PrepareDoctor(doctors models.DoctorSlice) []map[string]interface{} {

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User
		healthInsurance := doctor.R.Healthinsurance

		var healthInsuranceName interface{}
		if healthInsurance != nil {
			healthInsuranceName = healthInsurance.Name
		} else {
			healthInsuranceName = nil
		}

		responseData := map[string]interface{}{
			"doctor_id":        doctor.DoctorID,
			"health_insurance": healthInsuranceName,
			"specialties":      doctor.Specialties,
			"user":             user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData
}

func PreparePatient(patients models.PatientSlice) []map[string]interface{} {

	var combinedData []map[string]interface{}

	for _, patient := range patients {
		user := patient.R.User
		healthInsurance := patient.R.Healthinsurance

		var healthInsuranceName interface{}
		if healthInsurance != nil {
			healthInsuranceName = healthInsurance.Name
		} else {
			healthInsuranceName = nil
		}

		responseData := map[string]interface{}{
			"patient_id":      patient.PatientID,
			"healthInsurence": healthInsuranceName,
			"user":            user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData
}
