package helper

func SetResponse(message string, data any) map[string]any {
	var response = make(map[string]any)
	response["message"] = message
	if data != nil {
		response["data"] = data
	}

	return response
}
