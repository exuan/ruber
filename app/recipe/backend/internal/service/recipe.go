package service

import (
	"context"

	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
)

func (bs *BackendService) RecipeIndex(ctx context.Context, req *rsV1.IndexRequest) (*rsV1.IndexReply, error) {
	return bs.b.RecipeIndex(ctx, req)
}

func (bs *BackendService) SaveRecipeIndex(ctx context.Context, req *rsV1.SaveIndexRequest) (*rsV1.SaveIndexReply, error) {
	return bs.b.SaveRecipeIndex(ctx, req)
}

func (bs *BackendService) RecipeCategories(ctx context.Context, req *rsV1.CategoriesRequest) (*rsV1.CategoriesReply, error) {
	return bs.b.RecipeCategories(ctx, req)
}

func (bs *BackendService) RecipeCategory(ctx context.Context, req *rsV1.CategoryRequest) (*rsV1.CategoryReply, error) {
	return bs.b.RecipeCategory(ctx, req)
}

func (bs *BackendService) SaveRecipeCategory(ctx context.Context, req *rsV1.SaveCategoryRequest) (*rsV1.SaveCategoryReply, error) {
	return bs.b.SaveRecipeCategory(ctx, req)
}

func (bs *BackendService) DeleteRecipeCategory(ctx context.Context, req *rsV1.DeleteCategoryRequest) (*rsV1.DeleteCategoryReply, error) {
	return bs.b.DeleteRecipeCategory(ctx, req)
}

func (bs *BackendService) Recipes(ctx context.Context, req *rsV1.RecipesRequest) (*rsV1.RecipesReply, error) {
	return bs.b.Recipes(ctx, req)
}

func (bs *BackendService) Recipe(ctx context.Context, req *rsV1.RecipeRequest) (*rsV1.RecipeReply, error) {
	return bs.b.Recipe(ctx, req)
}
func (bs *BackendService) SaveRecipe(ctx context.Context, req *rsV1.SaveRecipeRequest) (*rsV1.SaveRecipeReply, error) {
	return bs.b.SaveRecipe(ctx, req)
}

func (bs *BackendService) DeleteRecipe(ctx context.Context, req *rsV1.DeleteRecipeRequest) (*rsV1.DeleteRecipeReply, error) {
	return bs.b.DeleteRecipe(ctx, req)
}
