package mockmock

import (
	"net/http"

	"mockapi/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r gin.IRouter) {
	api := r.Group("/api/v1/mockmocks")
	api.GET("", h.list)
	api.POST("", h.create)
	api.PUT("/:id", h.update)
	api.POST("/:id/duplicate", h.duplicate)
	api.DELETE("/:id", h.delete)

	r.Any("/mock/*path", h.serve)
}

func (h *Handler) list(c *gin.Context) {
	mockmocks, err := h.svc.List(c.Request.Context())
	if err != nil {
		httputil.RespondErr(c, err)
		return
	}
	httputil.RespondOK(c, mockmocks)
}

func (h *Handler) create(c *gin.Context) {
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		httputil.RespondAPIError(c, http.StatusBadRequest, "VALIDATION_ERROR", "invalid request body")
		return
	}
	m, err := h.svc.Create(c.Request.Context(), in)
	if err != nil {
		httputil.RespondErr(c, err)
		return
	}
	httputil.RespondCreated(c, m)
}

func (h *Handler) update(c *gin.Context) {
	var in Input
	if err := c.ShouldBindJSON(&in); err != nil {
		httputil.RespondAPIError(c, http.StatusBadRequest, "VALIDATION_ERROR", "invalid request body")
		return
	}
	m, err := h.svc.Update(c.Request.Context(), c.Param("id"), in)
	if err != nil {
		httputil.RespondErr(c, err)
		return
	}
	httputil.RespondOK(c, m)
}

func (h *Handler) duplicate(c *gin.Context) {
	m, err := h.svc.Duplicate(c.Request.Context(), c.Param("id"))
	if err != nil {
		httputil.RespondErr(c, err)
		return
	}
	httputil.RespondCreated(c, m)
}

func (h *Handler) delete(c *gin.Context) {
	if err := h.svc.Delete(c.Request.Context(), c.Param("id")); err != nil {
		httputil.RespondErr(c, err)
		return
	}
	httputil.RespondOK(c, map[string]bool{"deleted": true})
}

func (h *Handler) serve(c *gin.Context) {
	path := c.Param("path")
	m, err := h.svc.Match(c.Request.Context(), c.Request.Method, path)
	if err != nil {
		httputil.RespondErr(c, err)
		return
	}
	c.Data(m.Status, "application/json", m.Response)
}
