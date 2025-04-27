package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
)

type Repo struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Url string `json:"html_url"`
	Star int `json:"stargazers_count"`
}

func main(){
  var userName string = "jonathon-chew"
  var URL string = fmt.Sprintf("https://api.github.com/users/%s/repos", userName)

  req, err := http.Get(URL)
  if err != nil{
    log.Fatal(err)
  }

  defer req.Body.Close()

  var repos []Repo
  if err := json.NewDecoder(req.Body).Decode(&repos); err != nil {
    log.Fatalf("Error unmarshalling JSON: %v", err)
  }

	var fileName string = "README.md"

  file, err := os.Create(fileName)
  if err != nil {
    log.Fatalf("Error creating file: %v", err)
  }
 
  defer file.Close()
  
  // Write header
  fmt.Fprintln(file, "# GitHub Repositories\n")

  // Write each repo in Markdown format
  for _, repo := range repos {
    if repo.Name != userName{
      fmt.Fprintf(file, "## [%s](%s) :star: %d\n", repo.Name, repo.Url, repo.Star)
      fmt.Fprintf(file, "%s\n\n", repo.Description)
      fmt.Printf("Name: %s\n", repo.Name)
      fmt.Printf("Description: %s\n", repo.Description)
      fmt.Printf("URL: %s\n\n", repo.Url)
    }
  }

  fmt.Printf("Markdown file created: %s", fileName)
}
