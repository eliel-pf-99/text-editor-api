package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	InsertNote(ctx *gin.Context)
	UpdateNote(ctx *gin.Context)
	DeleteNote(ctx *gin.Context)
	FindNoteById(ctx *gin.Context)
	GetNotes(ctx *gin.Context)
}

type handler struct {
	service Service
}

// DeleteNote implements Handler.
func (h *handler) DeleteNote(ctx *gin.Context) {
	var req NoteReq
	if ctx.BindJSON(&req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	err := h.service.DeleteNote(ctx, req.NoteID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Note deleted with success!",
	})
}

// FindNoteById implements Handler.
func (h *handler) FindNoteById(ctx *gin.Context) {
	var req NoteReq
	if ctx.BindJSON(&req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	note, err := h.service.FindNoteById(ctx, req.NoteID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"note": note,
	})
}

// GetNotes implements Handler.
func (h *handler) GetNotes(ctx *gin.Context) {
	user := ctx.GetString("user")
	notes, _ := h.service.GetNotes(ctx, user)
	ctx.JSON(http.StatusOK, gin.H{"notes": notes})
}

// InsertNote implements Handler.
func (h *handler) InsertNote(ctx *gin.Context) {
	var noteCreate NoteCreate
	if ctx.BindJSON(&noteCreate) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}
	note, err := h.service.InsertNote(ctx, AddUserId(noteCreate, ctx.GetString("user")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create note",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Note create with success",
		"note":    note,
	})

}

// UpdateNote implements Handler.
func (h *handler) UpdateNote(ctx *gin.Context) {
	var note Note
	if ctx.BindJSON(&note) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	_, err := h.service.UpdateNote(ctx, note)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Note update with success!",
	})
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
