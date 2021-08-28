package handlers

import (
	"database/sql"
	"github.com/albukhary/monolith/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	db *sql.DB
	users []models.User
}

func NewHandler(db *sql.DB) *Handler{
	return &Handler{
		db: db,
	}
}

func (h *Handler) Hello( c *gin.Context) {
	c.String(http.StatusOK, "salam")
}

func (h *Handler) CreateUser (c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
	}
	res, err := h.db.Exec(`INSERT INTO users VALUES ($1, $2, $3)`, user.Name, user.Email, user.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if n, _ := res.RowsAffected(); n == 0 {
		c.String(http.StatusBadRequest, "already exists")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUsers(c *gin.Context) {
	pageValue := c.Query("page")
	if pageValue == "" {
		pageValue = "1"
	}
	limitValue := c.Query("limit")
	if limitValue == "" {
		limitValue = "10"
	}

	page, err := strconv.Atoi(pageValue)
	if err != nil {
		c.String(http.StatusBadRequest, "page value should be integer")
	}

	limit, err := strconv.Atoi(limitValue)
	if err != nil {
		c.String(http.StatusBadRequest, "limit value should be integer")
	}

	var (
		users = []models.User{}
		offset = (page - 1)*limit
	)
	rows, err := h.db.Query(`SELECT name, email, password FROM users OFFSET $1 LIMIT $2`, offset, limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	for  rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.Email, &user.Password)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByEmail (c *gin.Context) {
	email := c.Param("email")
	var user models.User
	row := h.db.QueryRow(`SELECT name, email, password FROM users WHERE email = $1`, email)
	err := row.Scan(
		&user.Name,
		&user.Email,
		&user.Password,
		)
	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "this user does not exist")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)a
}
