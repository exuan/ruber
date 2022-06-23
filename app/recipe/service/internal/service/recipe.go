package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/service/v1"
	"github.com/exuan/ruber/app/recipe/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RecipeService struct {
	v1.UnimplementedRecipeServer

	b   *biz.Recipe
	log *log.Helper
}

func NewRecipeService(b *biz.Recipe, logger log.Logger) *RecipeService {
	return &RecipeService{b: b, log: log.NewHelper(logger)}
}

func (rs *RecipeService) Index(ctx context.Context, req *v1.IndexRequest) (*v1.IndexReply, error) {
	return rs.b.Index(ctx, req)
}

func (rs *RecipeService) SaveIndex(ctx context.Context, req *v1.SaveIndexRequest) (*v1.SaveIndexReply, error) {
	return rs.b.SaveIndex(ctx, req)
}

func (rs *RecipeService) Recipes(ctx context.Context, req *v1.RecipesRequest) (*v1.RecipesReply, error) {
	return rs.b.Recipes(ctx, req)
}

func (rs *RecipeService) Recipe(ctx context.Context, req *v1.RecipeRequest) (*v1.RecipeReply, error) {
	return rs.b.Recipe(ctx, req)
}

func (rs *RecipeService) SaveRecipe(ctx context.Context, req *v1.SaveRecipeRequest) (*v1.SaveRecipeReply, error) {
	return rs.b.SaveRecipe(ctx, req)
}

func (rs *RecipeService) DeleteRecipe(ctx context.Context, req *v1.DeleteRecipeRequest) (*v1.DeleteRecipeReply, error) {
	return rs.b.DeleteRecipe(ctx, req)
}

func (rs *RecipeService) Categories(ctx context.Context, req *v1.CategoriesRequest) (*v1.CategoriesReply, error) {
	return rs.b.Categories(ctx, req)
}

func (rs *RecipeService) Category(ctx context.Context, req *v1.CategoryRequest) (*v1.CategoryReply, error) {
	return rs.b.Category(ctx, req)
}

func (rs *RecipeService) SaveCategory(ctx context.Context, req *v1.SaveCategoryRequest) (*v1.SaveCategoryReply, error) {
	return rs.b.SaveCategory(ctx, req)
}

func (rs *RecipeService) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) (*v1.DeleteCategoryReply, error) {
	return rs.b.DeleteCategory(ctx, req)
}

func (rs *RecipeService) Favorites(ctx context.Context, req *v1.FavoritesRequest) (*v1.FavoritesReply, error) {
	return rs.b.Favorites(ctx, req)
}

func (rs *RecipeService) Favorite(ctx context.Context, req *v1.FavoriteRequest) (*v1.FavoriteReply, error) {
	return rs.b.Favorite(ctx, req)
}

func (rs *RecipeService) SaveFavorites(ctx context.Context, req *v1.SaveFavoritesRequest) (*v1.SaveFavoritesReply, error) {
	return rs.b.SaveFavorites(ctx, req)
}
