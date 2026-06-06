package content

import "html/template"

type BlogPost struct {
	Slug    string
	Title   string
	Date    string
	Author  string
	Tags    []string
	Summary string
	Content template.HTML
}

var BlogPosts = []BlogPost{}

func FindPost(slug string) *BlogPost {
	for i := range BlogPosts {
		if BlogPosts[i].Slug == slug {
			return &BlogPosts[i]
		}
	}
	return nil
}
