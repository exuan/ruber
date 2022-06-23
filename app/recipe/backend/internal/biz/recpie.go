package biz

import (
	"context"

	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
)

func (b *Backend) RecipeIndex(ctx context.Context, req *rsV1.IndexRequest) (*rsV1.IndexReply, error) {
	return b.rc.Index(ctx, req)
}

func (b *Backend) SaveRecipeIndex(ctx context.Context, req *rsV1.SaveIndexRequest) (*rsV1.SaveIndexReply, error) {
	return b.rc.SaveIndex(ctx, req)
}

func (b *Backend) RecipeCategories(ctx context.Context, req *rsV1.CategoriesRequest) (*rsV1.CategoriesReply, error) {
	return b.rc.Categories(ctx, req)
}

func (b *Backend) RecipeCategory(ctx context.Context, req *rsV1.CategoryRequest) (*rsV1.CategoryReply, error) {
	return b.rc.Category(ctx, req)
}

func (b *Backend) SaveRecipeCategory(ctx context.Context, req *rsV1.SaveCategoryRequest) (*rsV1.SaveCategoryReply, error) {
	return b.rc.SaveCategory(ctx, req)
}

func (b *Backend) DeleteRecipeCategory(ctx context.Context, req *rsV1.DeleteCategoryRequest) (*rsV1.DeleteCategoryReply, error) {
	return b.rc.DeleteCategory(ctx, req)
}

func (b *Backend) Recipes(ctx context.Context, req *rsV1.RecipesRequest) (*rsV1.RecipesReply, error) {
	return b.rc.Recipes(ctx, req)
}

func (b *Backend) Recipe(ctx context.Context, req *rsV1.RecipeRequest) (*rsV1.RecipeReply, error) {
	return b.rc.Recipe(ctx, req)
}

func (b *Backend) SaveRecipe(ctx context.Context, req *rsV1.SaveRecipeRequest) (*rsV1.SaveRecipeReply, error) {
	return b.rc.SaveRecipe(ctx, req)
}

func (b *Backend) DeleteRecipe(ctx context.Context, req *rsV1.DeleteRecipeRequest) (*rsV1.DeleteRecipeReply, error) {
	return b.rc.DeleteRecipe(ctx, req)
}
