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
}

func main(){
  var URL string = "https://api.github.com/users/jonathon-chew/repos"

  req, err := http.Get(URL)
  if err != nil{
    log.Fatal(err)
  }

  defer req.Body.Close()

  var repos []Repo
  if err := json.NewDecoder(req.Body).Decode(&repos); err != nil {
    log.Fatalf("Error unmarshalling JSON: %v", err)
  }

  for _, repo := range repos {
    fmt.Printf("Name: %s\n", repo.Name)
    fmt.Printf("Description: %s\n", repo.Description)
    fmt.Printf("URL: %s\n\n", repo.Url)
  }

  file, err := os.Create("README.md")
  if err != nil {
    log.Fatalf("Error creating file: %v", err)
  }
 
  defer file.Close()
  
  // Write header
  fmt.Fprintln(file, "# GitHub Repositories\n")

  // Write each repo in Markdown format
  for _, repo := range repos {
    fmt.Fprintf(file, "## [%s](%s)\n", repo.Name, repo.Url)
    fmt.Fprintf(file, "%s\n\n", repo.Description)
  }

  fmt.Println("Markdown file created: repos.md")
}
