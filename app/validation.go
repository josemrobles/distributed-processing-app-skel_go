package main

func validateAuthToken(token string) bool {
	if token == "valid_token" {
		return true
	}
	return false
}

func validatePayload(payload Payload) bool {
	return true
}
