package handlers

import (
	"net/http"

	"github.com/Jonnay101/icon/pkg/glitch"
)

// HandlerBlogCreatePost -
func (h *Handlers) HandlerBlogCreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		blogPost, err := h.Blog.BindRequestBody(w, r)
		if err != nil {
			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.Blog.DB.StoreBlogPost(blogPost); err != nil {

			if err == glitch.ErrItemAlreadyExists {
				h.respond(w, r, err.Error(), http.StatusConflict)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		h.respond(w, r, blogPost, http.StatusOK)
	}
}

// HandlerBlogGetPost -
func (h *Handlers) HandlerBlogGetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := h.Blog.BindRequestParams(w, r)
		if err != nil {
			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		blogPost, err := h.Blog.DB.FindBlogPostByKey(reqParams)
		if err != nil {

			if err == glitch.ErrRecordNotFound {
				h.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		h.respond(w, r, blogPost, http.StatusOK)
	}
}

// HandlerBlogGetAllPosts -
func (h *Handlers) HandlerBlogGetAllPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := h.Blog.BindRequestParams(w, r)
		if err != nil {
			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		blogPosts, err := h.Blog.DB.FindAllBlogPosts(reqParams)
		if err != nil {

			if err == glitch.ErrRecordNotFound {
				h.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		h.respond(w, r, blogPosts, http.StatusOK)
	}
}

// HandlerBlogUpdatePost -
func (h *Handlers) HandlerBlogUpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := h.Blog.BindRequestParams(w, r)
		if err != nil {
			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		oldBlogPost, err := h.Blog.DB.FindBlogPostByKey(reqParams)
		if err != nil {

			if err == glitch.ErrRecordNotFound {
				h.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		newBlogPost, err := h.Blog.BindRequestBody(w, r)
		if err != nil {
			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		h.Blog.PopulateZeroValueFieldsWithOldData(oldBlogPost, newBlogPost)

		if err := h.Blog.DB.UpdateBlogPost(newBlogPost); err != nil {

			if err == glitch.ErrRecordNotFound {
				h.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		h.respond(w, r, newBlogPost, http.StatusOK)
	}
}

//HandlerBlogDeletePost -
func (h *Handlers) HandlerBlogDeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqParams, err := h.Blog.BindRequestParams(w, r)
		if err != nil {

			h.respond(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.Blog.DB.RemoveBlogPost(reqParams); err != nil {

			if err == glitch.ErrRecordNotFound {
				h.respond(w, r, err.Error(), http.StatusNotFound)
				return
			}

			h.respond(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		h.respond(w, r, "OK", http.StatusOK)
	}
}
