package prepare

import (
	"medical-clinic/models"
)

func PrepareData(doctors models.DoctorSlice) []map[string]interface{} {

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User

		responseData := map[string]interface{}{
			"doctor": doctor,
			"user":   user,
		}

		combinedData = append(combinedData, responseData)
	}
	return combinedData
}
