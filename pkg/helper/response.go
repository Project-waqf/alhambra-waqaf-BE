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

func SuccessGetAll(msg string, data interface{}, countOnline, countDraft, countArchive int) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data": data,
		"total_online": countOnline,
		"total_draft": countDraft,
		"total_archive": countArchive,
	}
}