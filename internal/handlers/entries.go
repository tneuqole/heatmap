package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tneuqole/habitmap/internal/model"
)

type EntryHandler struct {
	queries *model.Queries
}

func NewEntryHandler(queries *model.Queries) *EntryHandler {
	return &EntryHandler{
		queries: queries,
	}
}

type DeleteEntryForm struct {
	EntryID int64 `form:"id"`
}

func (h *EntryHandler) DeleteEntry(c echo.Context) error {
	form := DeleteEntryForm{}
	if err := c.Bind(&form); err != nil {
		return err
	}

	err := h.queries.DeleteEntry(c.Request().Context(), form.EntryID)
	if err != nil {
		return err
	}

	c.Response().Header().Add("Hx-Redirect", "/habits") // TODO

	return nil
}

type CreateEntryForm struct {
	HabitID   int64 `form:"habit_id" validate:"required,notblank`
	EntryDate int64 `form:"entry_date" validate:"required,notblank`
}

func (h *EntryHandler) PostEntry(c echo.Context) error {
	form := CreateEntryForm{}
	if err := c.Bind(&form); err != nil {
		return err
	}

	err := validate.Struct(&form)
	if err != nil {
		return err
	}

	params := model.CreateEntryParams{EntryDate: form.EntryDate, HabitID: form.HabitID}
	entry, err := h.queries.CreateEntry(c.Request().Context(), params)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/habits/%d", entry.HabitID))
}
