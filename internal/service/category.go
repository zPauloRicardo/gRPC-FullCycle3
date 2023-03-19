package service

import (
	"FcCeS/internal/database"
	"FcCeS/internal/pb"
	"golang.org/x/net/context"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(_ context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}
	return categoryResponse, nil
}
