package service

import (
	"context"
	"fmt"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
)

func (fs *FrontendService) Recipes(ctx context.Context, req *v1.RecipesRequest) (*v1.RecipesReply, error) {

	return fs.b.Recipes(ctx, req)
}

func (fs *FrontendService) Recipe(ctx context.Context, req *v1.RecipeRequest) (*v1.RecipeReply, error) {
	return fs.b.Recipe(ctx, req)
}

func (fs *FrontendService) RecipeCategories(ctx context.Context, req *v1.CategoriesRequest) (*v1.CategoriesReply, error) {
	return fs.b.RecipeCategories(ctx, req)

}

func (fs *FrontendService) SaveFavorites(ctx context.Context, req *v1.SaveFavoritesRequest) (*v1.SaveFavoritesReply, error) {
	fmt.Println("222")
	return fs.b.SaveFavorites(ctx, req)
}
