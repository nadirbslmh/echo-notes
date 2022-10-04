package categories

import (
	"echo-notes/businesses/categories"
	"echo-notes/controller/categories/request"
	"echo-notes/controller/categories/response"
	"echo-notes/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAllCategories(c echo.Context) error {
	categoriesData := ctrl.categoryUseCase.GetAll()

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return c.JSON(http.StatusOK, model.Response[[]response.Category]{
		Status:  "success",
		Message: "all categories",
		Data:    categories,
	})
}

func (ctrl *CategoryController) CreateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	category := ctrl.categoryUseCase.Create(input.ToDomain())

	return c.JSON(http.StatusCreated, model.Response[response.Category]{
		Status:  "success",
		Message: "category created",
		Data:    response.FromDomain(category),
	})
}

func (ctrl *CategoryController) UpdateCategory(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	var id string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    nil,
		})
	}

	category := ctrl.categoryUseCase.Update(id, input.ToDomain())

	return c.JSON(http.StatusOK, model.Response[response.Category]{
		Status:  "success",
		Message: "category updated",
		Data:    response.FromDomain(category),
	})
}

func (ctrl *CategoryController) DeleteCategory(c echo.Context) error {
	var id string = c.Param("id")

	isDeleted := ctrl.categoryUseCase.Delete(id)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, model.Response[bool]{
			Status:  "failed",
			Message: "category not found",
			Data:    isDeleted,
		})
	}

	return c.JSON(http.StatusOK, model.Response[bool]{
		Status:  "success",
		Message: "category deleted",
		Data:    isDeleted,
	})
}
