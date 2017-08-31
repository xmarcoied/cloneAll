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
	res, _ := http.Get(url)
	defer res.Body.Close()

	var b bytes.Buffer
	io.Copy(&b, res.Body)

	json.Unmarshal([]byte(b.String()), &data)

	fmt.Println("Numbers of repositories:", len(data))

	for i := 0; i < len(data); i++ {
		fmt.Printf("Cloning %s repository\n", data[i].RepoName)
		repo_url := fmt.Sprintf("https://github.com/%s/%s.git", username, data[i].RepoName)
		fmt.Println(repo_url)
		cmd := exec.Command("git", "clone", repo_url)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Finished Cloning %s repository\n", data[i].RepoName)
	}

}
