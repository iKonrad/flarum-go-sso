package go_flarum

type Error struct {
	ERROR_GENERAL string
	ERROR_INVALID_CREDENTIALS string
}

var flarumErrors Error

func init() {
	flarumErrors = Error {
		ERROR_GENERAL: "An error occurred while logging in",
		ERROR_INVALID_CREDENTIALS: "Invalid username or password",
	}
}