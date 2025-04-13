package generic

import "fmt"

type APIResponse struct {
	Response string
}

func ParseResponse(response APIResponse) {
	fmt.Println("parsing...")
}
