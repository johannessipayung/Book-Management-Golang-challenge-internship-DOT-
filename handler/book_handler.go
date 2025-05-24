package handler

import (
	"challengeGO/model"
	"challengeGO/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateBookInput struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bs service.BookService) *BookHandler {
	return &BookHandler{bs}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDInterface := c.MustGet("userID")
	userIDStr, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID invalid type"})
		return
	}

	userIDUint64, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID conversion error"})
		return
	}
	userID := uint(userIDUint64)

	book := model.Book{
		Title:      input.Title,
		Author:     input.Author,
		UserID:     userID,
		CategoryID: input.CategoryID,
	}

	if err := h.bookService.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Struct input tanpa UserID
	type UpdateBookInput struct {
		Title      string `json:"title" binding:"required"`
		Author     string `json:"author" binding:"required"`
		CategoryID uint   `json:"categoryId" binding:"required"`
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update hanya field yang boleh diubah, jangan ubah UserID!
	book.Title = input.Title
	book.Author = input.Author
	book.CategoryID = input.CategoryID

	if err := h.bookService.Update(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.bookService.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
