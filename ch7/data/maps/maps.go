package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
  websites["LinkedIn"] = "https//linkedIn"
	fmt.Println(websites["Amazon Web Services"])

  delete(websites,"Google")
	fmt.Println(websites)
}
