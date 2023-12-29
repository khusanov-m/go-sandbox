package maps

import "fmt"

func main() {
	// maps are unordered lists of key-value pairs
	// maps are reference types
	// maps are dynamically resizable
	// maps are created with make
	// maps are indexed with keys
	// maps are used to look up values by keys
	// maps are nilable
	// maps are comparable
	// maps are not iterable with range
	// maps are not safe for concurrent use
	// maps are usually used to represent a collection of related properties

	// websites := []string{"http://www.google.com", "http://www.facebook.com", "http://www.youtube.com"}
	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
