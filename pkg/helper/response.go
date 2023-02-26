package helper

func Success(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
	}
}

func Failed(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func SuccessGetAll(msg string, data interface{}, count int) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
		"total_data": count,
	}
}