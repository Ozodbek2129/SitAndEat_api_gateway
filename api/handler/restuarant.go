package handler

import (
	pb "api_gateway/genproto/restaurant"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary restaurant olish
// @Description Id bilan restaurant olinyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body restaurant.RestuanantId true "Restaurant so'rovi"
// @Success 200 {object} restaurant.GetRes
// @Failure 400 {object} models.Error
// @Router /restaurant [get]
func (h *Handler) GetRestaurant(c *gin.Context) {
	req := pb.RestuanantId{
		Id: c.Param("id"),
	}

	resp, err := h.RestaurantClient.GetRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("GetRestuarant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant yangilash
// @Description Bodydagi ma'luotlar asosida restaurant yangilanyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body restaurant.RestuarantUpdate true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant [put]
func (h *Handler) UpdateRestaurant(c *gin.Context) {
	req := pb.RestuarantUpdate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Update restaurant error : %v", err)
		return
	}

	resp, err := h.RestaurantClient.UpdateRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("UpdateRestaurant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant o'chirish
// @Description Id orqali restaurant o'chirilmoqda
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body restaurant.RestuanantId true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant [delete]
func (h *Handler) DeleteRestaurant(c *gin.Context) {
	req := pb.RestuanantId{
		Id: c.Param("id"),
	}

	resp, err := h.RestaurantClient.DeleteRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("DeleteRestuarant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant yaratish
// @Description Bodydagi ma'luotlar asosida restaurant yaratilyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body restaurant.Restuarant true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant [post]
func (h *Handler) CreateRestaurant(c *gin.Context) {
	req := pb.Restuarant{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("CreateRestaurant  error : %v", err)
		return
	}

	resp, err := h.RestaurantClient.CreateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("CreateRestaurant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Barcha restaurantni olish
// @Description Istalgancha restaurantni filterlab olish mumkin
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant body restaurant.Filter true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Restuanants
// @Failure 400 {object} models.Error
// @Router /restaurant [get]
func (h *Handler) GetAllRestaurants(c *gin.Context) {
	query := `
				SELECT 
					* 
				FROM 
					Restaurants
				WHERE 
					deleted_at is null`
	param := []string{}
	arr := []string{}
	req := pb.GetRes{
		Id:          c.Query("id"),
		Name:        c.Query("name"),
		Address:     c.Query("address"),
		Phone:       c.Param("phone"),
		Description: c.Param("description"),
		CreatedAt:   c.Param("createdAt"),
		UpdatedAt:   c.Param("updatedAt"),
	}
	limit := c.Query("limit")
	offset := c.Query("offset")

	if len(req.Id) > 0 {
		query += " and id = :id"
		param = append(param, ":id")
		arr = append(arr, req.Id)
	}
	if len(req.Name) > 0 {
		query += " and name = :name"
		param = append(param, ":name")
		arr = append(arr, req.Name)
	}

	if len(req.Address) > 0 {
		query += " and address = :address"
		param = append(param, ":address")
		arr = append(arr, req.Address)
	}

	if len(req.Phone) > 0 {
		query += " and phone = :phone"
		param = append(param, ":phone")
		arr = append(arr, req.Phone)
	}

	if len(req.Description) > 0 {
		query += " and description = :description"
		param = append(param, ":description")
		arr = append(arr, req.Description)
	}

	if len(req.CreatedAt) > 0 {
		data := strings.Split(req.CreatedAt, "-")
		query += " and created_at BETWEEN :created_at1 and :created_at2"
		param = append(param, ":created_at1", ":created_at2")
		arr = append(arr, data[0], data[1])
	}

	if len(req.UpdatedAt) > 0 {
		data := strings.Split(req.UpdatedAt, "-")
		query += " and updated_at BETWEEN :updated_at1 and :updated_at2"
		param = append(param, ":updated_at1", ":updated_at2")
		arr = append(arr, data[0], data[1])
	}

	if len(limit) > 0 {
		query += " limit " + limit
	}

	if len(offset) > 0 {
		query += " offset " + offset
	}

	for k, v := range param {
		query = strings.Replace(query, v, "$"+strconv.Itoa(k), 1)
	}

	resp, err := h.RestaurantClient.GetAllRestaurants(context.Background(), &pb.Filter{Query: query, Arr: arr})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("GetAll request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
