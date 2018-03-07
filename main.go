package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	var (
		data []struct {
			RepoName string `json:"name"`
		}
		username string
	)
	fmt.Println("Enter github username:")
	fmt.Scanf("%s", &username)

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	fmt.Println("Fetching data from github ... ")
	res, _ := http.Get(url)

	var b bytes.Buffer
	io.Copy(&b, res.Body)

	json.Unmarshal([]byte(b.String()), &data)

	fmt.Println("Numbers of repositories:", len(data))

	for i := 0; i < len(data); i++ {
		fmt.Printf("Cloning %s repository... \n", data[i].RepoName)
		RepoURL := fmt.Sprintf("https://github.com/%s/%s.git", username, data[i].RepoName)
		FilePATH := fmt.Sprintf("%s/%s", username, data[i].RepoName)
		fmt.Println(RepoURL)
		cmd := exec.Command("git", "clone", RepoURL, FilePATH)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Finished Cloning %s repository\n", data[i].RepoName)

	}
	defer res.Body.Close()

}
