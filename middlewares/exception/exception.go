package exception

func BadRequest(message string, flag string) {
	errors := map[string]interface{}{
		"message": message,
		"flag":    flag,
	}

	response := map[string]interface{}{
		"status": 400,
		"flag":   "INVALID_BODY",
		"errors": errors,
	}

	panic(response)
}
