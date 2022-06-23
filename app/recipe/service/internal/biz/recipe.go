package biz

import (
	"context"
	"fmt"

	v1 "github.com/exuan/ruber/api/recipe/service/v1"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
)

type RecipeRepo interface {
	Recipes(ctx context.Context, req *v1.RecipesRequest) (int64, []*ent.Recipe, error)

	Recipe(ctx context.Context, req *v1.RecipeRequest) (*ent.Recipe, error)

	SaveRecipe(context.Context, *v1.SaveRecipeRequest) (int64, error)

	DeleteRecipe(context.Context, *v1.DeleteRecipeRequest) (int64, error)

	Categories(context.Context, *v1.CategoriesRequest) (int64, []*ent.Category, error)

	Category(context.Context, *v1.CategoryRequest) (*ent.Category, error)

	SaveCategory(context.Context, *v1.SaveCategoryRequest) (int64, error)

	DeleteCategory(context.Context, *v1.DeleteCategoryRequest) (int64, error)

	CountChildrenCategories(context.Context, int64) (int, error)

	SaveFavorites(context.Context, *v1.SaveFavoritesRequest) (int64, error)

	Favorites(context.Context, *v1.FavoritesRequest) (int64, []*ent.Favorite, error)

	Favorite(context.Context, *v1.FavoriteRequest) (*ent.Favorite, error)

	UpdateRecipeFavorite(context.Context, int64, int32) (int64, error)
}

type Recipe struct {
	repo RecipeRepo
	log  *log.Helper
	r    *redis.Client
}

const (
	cIndex = "recipe_index"
)

func NewRecipe(repo RecipeRepo, logger log.Logger, r *redis.Client) *Recipe {
	return &Recipe{repo: repo, log: log.NewHelper(logger), r: r}
}

func (r *Recipe) Index(ctx context.Context, req *v1.IndexRequest) (*v1.IndexReply, error) {

	str, err := r.r.Get(ctx, cIndex).Result()
	if err == redis.Nil {
		err = nil
	}
	if err != nil {
		return nil, err
	}

	res := new(v1.IndexReply)
	if err := jsoniter.Unmarshal([]byte(str), res); err != nil {
		return nil, err
	}

	return res, err
}

func (r *Recipe) SaveIndex(ctx context.Context, req *v1.SaveIndexRequest) (*v1.SaveIndexReply, error) {
	str, err := jsoniter.MarshalToString(req)
	if err != nil {
		return nil, err
	}
	err = r.r.Set(ctx, cIndex, str, 0).Err()
	return &v1.SaveIndexReply{}, err
}

func (r *Recipe) Recipes(ctx context.Context, req *v1.RecipesRequest) (*v1.RecipesReply, error) {
	total, res, err := r.repo.Recipes(ctx, req)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.RecipeReply, 0)
	for _, er := range res {
		items = append(items, &v1.RecipeReply{
			Id:         er.ID,
			Title:      er.Title,
			CategoryId: er.CategoryID,
			Image:      er.Image,
			Content:    er.Content,
			Weight:     er.Weight,
			Sort:       er.Sort,
			Favorites:  er.Favorites,
			Status:     int32(er.Status),
			Creator:    er.Creator,
			Modifier:   er.Modifier,
			CreateTime: er.CreateTime,
			UpdateTime: er.UpdateTime,
		})
	}

	return &v1.RecipesReply{Total: total, Items: items}, err
}
func (r *Recipe) Recipe(ctx context.Context, req *v1.RecipeRequest) (*v1.RecipeReply, error) {
	res, err := r.repo.Recipe(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.RecipeReply{
		Id:         res.ID,
		Title:      res.Title,
		CategoryId: res.CategoryID,
		Image:      res.Image,
		Content:    res.Content,
		Weight:     res.Weight,
		Sort:       res.Sort,
		Favorites:  res.Favorites,
		Status:     int32(res.Status),
		Creator:    res.Creator,
		Modifier:   res.Modifier,
		CreateTime: res.CreateTime,
		UpdateTime: res.UpdateTime,
	}, nil
}

func (r *Recipe) SaveRecipe(ctx context.Context, req *v1.SaveRecipeRequest) (*v1.SaveRecipeReply, error) {
	id, err := r.repo.SaveRecipe(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.SaveRecipeReply{Id: id}, nil
}

func (r *Recipe) DeleteRecipe(ctx context.Context, req *v1.DeleteRecipeRequest) (*v1.DeleteRecipeReply, error) {
	id, err := r.repo.DeleteRecipe(ctx, req)

	return &v1.DeleteRecipeReply{Id: id}, err
}

func (r *Recipe) Categories(ctx context.Context, req *v1.CategoriesRequest) (*v1.CategoriesReply, error) {
	total, res, err := r.repo.Categories(ctx, req)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.CategoryReply, 0)
	for _, ec := range res {
		items = append(items, &v1.CategoryReply{
			Id:         ec.ID,
			Name:       ec.Name,
			ParentId:   ec.ParentID,
			Sort:       ec.Sort,
			Status:     int32(ec.Status),
			Creator:    ec.Creator,
			Modifier:   ec.Modifier,
			CreateTime: ec.CreateTime,
			UpdateTime: ec.UpdateTime,
		})
	}

	return &v1.CategoriesReply{Total: total, Items: items}, err
}
func (r *Recipe) Category(ctx context.Context, req *v1.CategoryRequest) (*v1.CategoryReply, error) {
	res, err := r.repo.Category(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.CategoryReply{
		Id:         res.ID,
		Name:       res.Name,
		ParentId:   res.ParentID,
		Sort:       res.Sort,
		Status:     int32(res.Status),
		Creator:    res.Creator,
		Modifier:   res.Modifier,
		CreateTime: res.CreateTime,
		UpdateTime: res.UpdateTime,
	}, nil
}

func (r *Recipe) SaveCategory(ctx context.Context, req *v1.SaveCategoryRequest) (*v1.SaveCategoryReply, error) {
	id, err := r.repo.SaveCategory(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.SaveCategoryReply{Id: id}, nil
}

func (r *Recipe) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) (*v1.DeleteCategoryReply, error) {
	count, err := r.repo.CountChildrenCategories(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("分类下有子分类, 不能删除")
	}

	id, err := r.repo.DeleteCategory(ctx, req)

	return &v1.DeleteCategoryReply{Id: id}, err
}

func (r *Recipe) Favorites(ctx context.Context, req *v1.FavoritesRequest) (*v1.FavoritesReply, error) {
	_, res, err := r.repo.Favorites(ctx, req)
	if err != nil {
		return nil, err
	}

	items := make([]*v1.FavoriteReply, 0)
	for _, et := range res {
		items = append(items, &v1.FavoriteReply{
			Id:       et.ID,
			RecipeId: et.RecipeID,
			UserId:   et.UserID,
		})
	}

	return &v1.FavoritesReply{Items: items}, nil
}

func (r *Recipe) Favorite(ctx context.Context, req *v1.FavoriteRequest) (*v1.FavoriteReply, error) {
	res, err := r.repo.Favorite(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.FavoriteReply{
		Id:       res.ID,
		RecipeId: res.RecipeID,
		UserId:   res.UserID,
	}, nil
}

func (r *Recipe) SaveFavorites(ctx context.Context, req *v1.SaveFavoritesRequest) (*v1.SaveFavoritesReply, error) {
	_, err := r.repo.SaveFavorites(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = r.repo.UpdateRecipeFavorite(ctx, req.RecipeId, req.Action)
	if err != nil {
		return nil, err
	}

	return &v1.SaveFavoritesReply{}, nil
}
