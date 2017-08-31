package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	var (
		data []struct {
			RepoName string `json:"name"`
		}
		username string
	)
	fmt.Println("Enter username:")
	fmt.Scanf("%s", &username)

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	res, _ := http.Get(url)
	defer res.Body.Close()

	var b bytes.Buffer
	io.Copy(&b, res.Body)

	json.Unmarshal([]byte(b.String()), &data)

	fmt.Println("Numbers of repositories:", len(data))
	fmt.Println(data)
}
