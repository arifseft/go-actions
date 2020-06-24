package exception

func BadRequest(message string, flag string) {
	errors := map[string]interface{}{
		"message": message,
		"flag":    flag,
	}

	response := map[string]interface{}{
		"status": 400,
		"flag":   "BAD_REQUEST",
		"errors": errors,
	}

	panic(response)
}
