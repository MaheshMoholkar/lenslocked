package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, We offer a free trail for 30 days.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff available for 24/7, though response times may be a bit slower.",
		},
		{
			Question: "",
			Answer:   "",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
