package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"nfact-site/content"
)

// PageData holds data passed to templates
type PageData struct {
	Title      string
	ActivePage string
	Post       *content.BlogPost
}

// parseTemplate parses the layout together with a specific page template
func parseTemplate(page string) *template.Template {
	return template.Must(template.ParseFiles(
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", page),
	))
}

func init() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Parse templates: each page + layout
	indexTmpl := parseTemplate("index.html")
	blogTmpl := parseTemplate("blog.html")
	blogPostTmpl := parseTemplate("blog-post.html")
	contactTmpl := parseTemplate("contact.html")

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		data := struct {
			PageData
			Integrantes []content.Integrante
		}{
			PageData: PageData{
				Title:      "n! — Comunidad de Código y Matemáticas",
				ActivePage: "inicio",
			},
			Integrantes: content.Integrantes,
		}
		err := indexTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Blog listing
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			PageData
			Posts []content.BlogPost
		}{
			PageData: PageData{
				Title:      "n! — Blog",
				ActivePage: "blog",
			},
			Posts: content.BlogPosts,
		}
		err := blogTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Blog post detail
	http.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		slug := r.URL.Path[len("/blog/"):]
		if slug == "" {
			http.Redirect(w, r, "/blog", http.StatusFound)
			return
		}
		post := content.FindPost(slug)
		if post == nil {
			http.NotFound(w, r)
			return
		}
		data := PageData{
			Title:      "n! — " + post.Title,
			ActivePage: "blog",
			Post:       post,
		}
		err := blogPostTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})

	// Contact page
	http.HandleFunc("/contacto", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:      "n! — Contacto",
			ActivePage: "contacto",
		}
		err := contactTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, "Internal Server Error", 500)
		}
	})
}

// Handler is the entry point for Vercel Serverless Functions.
func Handler(w http.ResponseWriter, r *http.Request) {
	http.DefaultServeMux.ServeHTTP(w, r)
}

func main() {
	log.Println("Server started on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
