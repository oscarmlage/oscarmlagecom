---
title: "Automating the media (hugo-medialog.go)"
date: 2024-10-15T17:18:37+02:00
draft: false
tags: []
---

As someone who consumes and tracks a some media items (books, series, movies, videogames, music...), I’ve always looked for ways to keep an updated record without adding too much friction to my workflow. After several iterations, I’ve created a *Go script* that automates this process using tools I already use daily: *Obsidian* and *Hugo*.

I’ll try to explain how this system works and how you can use it to keep track of any type of content with minimal effort. 

### The  workflow

1. **Obsidian** is my go-to note-taking tool, where I have a YAML file for each type of content (e.g., `books.yml.md`, `series.yml.md`, etc.). In these files, I record titles, progress, ratings, and more.
    
2. **Hugo** is my static site generator, where I publish this tracking as part of my blog. Each time I read a book or watch a movie, my goal is to generate a markdown file in Hugo automatically from the YAML file in Obsidian.
    
3. **The Script** in Go takes care of this transition, moving data from Obsidian to Hugo automatically. It reads the YAML files from Obsidian and generates corresponding markdown files in Hugo. Also it uses some APIs to get more information about the item (number of episodes/pages, release date, images, author/company...)

### The script in action

The script is designed to process different types of content (books, series, movies, etc.). Using a `--media` flag, you can specify which type of content you want to process. Additionally, I’ve added an optional `--update` flag that updates the frontmatter (structured data in the markdown file) without overwriting any personal content you may have added to the markdown, like notes or reviews. Example usage: `go run main.go --media books --update`

This will process books and only update the metadata, leaving your personal content in *Hugo* untouched.

### Project structure

Here's a simplified view of how the project is organized:

```sh
medialog-project/
├── internal/
│   ├── books/
│   │   └── model.go
│   │   └── controller.go
│   │   └── api_googlebooks.go
│   ├── games/
│   │   └── model.go
│   │   └── controller.go
│   │   └── api_igdb.go
│   ├── movies/
│   │   └── model.go
│   │   └── controller.go
│   │   └── api_trakt.go
│   ├── music/
│   │   └── model.go
│   │   └── controller.go
│   │   └── api_spotify.go
├── templates/
│   └── books.md.tpl
│   └── games.md.tpl
│   └── movies.md.tpl
│   └── music.md.tpl
├── utils/
│   └── config.go
│   └── markdown.go
│   └── utils.go
├── .env
├── main.go
```

### How does it work for books?

Let’s look at a concrete example. In the `books.yml.md` file in *Obsidian*, I have a simple record like this:

```yaml
- title: "Terraform: Up and Running"
  author: "Yevgeniy Brikman"
  year: 2017
  progress: 25%
  rate: 7.2
  date: 2024-10-15
```

The script reads this data and searches for the corresponding markdown file in Hugo. If it doesn’t exist, it creates it calling some APIs and using a template; if it does, it can update only the metadata without touching the content, so any personal notes remain intact.

Here’s a breakdown of how we handle **books**:

#### Model

The `Book` model represents the data structure for books. This model is responsible for parsing the YAML file from Obsidian and storing the information that will be used to generate the final markdown in Hugo.

Here’s an example of how the `Book` model might look in Go:

```go
type Book struct {
    Title    string  `yaml:"title"`
    Author   string  `yaml:"author"`
    Year     int     `yaml:"year"`
    Progress string  `yaml:"progress"`
    Rate     float64 `yaml:"rate"`
    Date     string  `yaml:"date"`
}
```

This structure matches the fields in the YAML file in Obsidian and allows us to easily read and write the data into markdown format.

#### Controller

The controller is responsible for handling the *logic* of reading the YAML file, updating or generating the corresponding markdown, and applying the correct template. The logic flow is as follows:

1. **Load YAML data:** The controller loads the book data from the YAML file (`books.yml.md`) into the `Book` model.
2. **Check for existing markdown:** It checks if a markdown file for the book already exists in Hugo. If it does, it updates only the frontmatter fields (like `progress`, `rate`, etc.), leaving the content untouched. If it does not exist the controller calls the proper APIs to get the information.
3. **Generate or update markdown:** If no markdown exists, it generates a new markdown file using a predefined template. Otherwise, it updates the existing frontmatter while preserving user content.

Here’s a simplified view of how this process looks in code:

```go
func ProcessBooks() {
    books := LoadBooksFromYAML()
    
    for _, book := range books {
        mdFile := fmt.Sprintf("%s/%s.md", MARKDOWN_OUTPUT_BOOKS_DIR, book.Slug)
        
        if fileExists(mdFile) {
            // Update only the frontmatter
            UpdateBookFrontmatter(mdFile, book)
        } else {
            // Get proper data from APIs
            SearchBookByTitle(book.Title, &book)
    		getBookImagesByTitle(book.Title, &book)
            // Generate a new markdown file using the template
            GenerateMarkdownFile(mdFile, book, "book_template.md")
            continue
        }
    }
}

```

#### Template

The templates are stored in Hugo’s `content` folder, and they define how the final markdown looks. Here’s an example of a simple template for books (`book_template.md`):

```md
---
title: "{{ .Title }}"
author: "{{ .Author }}"
year: {{ .Year }}
progress: {{ .Progress }}
rate: {{ .Rate }}
date: {{ .Date }}
---

# {{ .Title }}

## Overview


## My thoughts
```

This template uses Go’s text/template syntax to dynamically insert book data into the frontmatter and content sections of the markdown. For example, the generated or updated markdown file in Hugo might look like this:

```md
---
title: "Terraform: Up and Running"
author: "Yevgeniy Brikman"
year: 2017
progress: 25%
rate: 7.2
date: 2024-10-15
---

# Terraform: Up and Running

## Overview
<no value>

## My thoughts
This is where you can add your personal thoughts about the book.
```

### In the end

This system lets me keep my medialog updated with *minimal friction*, as I only need to update the YAML files in *Obsidian*, which is already part of my daily note-taking routine. 

From there, the entire process of generating content in my blog is automated. Additionally, by using the `--update` flag, I can ensure that any personal content added in Hugo remains untouched.

Keeping track of what we consume doesn’t have to be a tedious or manual process. With the right tools and automation, you can make it as seamless as possible. If you're interested in implementing something similar, you can check the [script repository](https://git.oscarmlage.com/oscarmlage/hugo-medialog), where you can see the full code and adapt/edit it to fit your needs.