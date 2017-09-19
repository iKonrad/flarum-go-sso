package go_flarum

type RequestType struct {
	GET string
	POST string
	PATCH string
	DELETE string
}

var request RequestType

func init() {
	request = RequestType{
		GET:   "GET",
		POST:  "POST",
		PATCH: "PATCH",
		DELETE: "DELETE",
	}
}