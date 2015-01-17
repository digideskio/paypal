package paypal

type (
	Error struct {
		Name            string `json:"name"`
		Message         string `json:"message"`
		InformationLink string `json:"information_link"`
		Details         string `json:"details"`
	}
)
