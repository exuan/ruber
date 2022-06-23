package biz

import (
	"context"
	"fmt"
	"strconv"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
	"github.com/exuan/ruber/pkg/errors"
	"github.com/go-kratos/kratos/v2/metadata"
)

const (
	RecipeView = "recipe:view"
)

func (f *Frontend) RecipeCategories(ctx context.Context, req *v1.CategoriesRequest) (*v1.CategoriesReply, error) {
	res, err := f.rc.Categories(ctx, &rsV1.CategoriesRequest{
		Id:   req.Id,
		Name: req.Name,
		All:  req.All,
	})

	if err != nil {
		return nil, err
	}

	items := make([]*v1.CategoryReply, 0)

	for _, c := range res.Items {
		items = append(items, &v1.CategoryReply{
			Id:         c.Id,
			Name:       c.Name,
			ParentId:   c.ParentId,
			Sort:       c.Sort,
			CreateTime: c.CreateTime,
			UpdateTime: c.UpdateTime,
		})
	}

	return &v1.CategoriesReply{Items: items}, nil
}

func (f *Frontend) Recipes(ctx context.Context, req *v1.RecipesRequest) (*v1.RecipesReply, error) {
	ids := make([]int64, 0)
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.InternalServerError
	}

	userId := md.Get("user_id")
	userFavorites := make(map[int64]bool)
	if len(userId) > 0 {
		res, err := f.rc.Favorites(ctx, &rsV1.FavoritesRequest{
			UserId: userId,
			All:    1,
		})
		if err != nil {
			return nil, errors.InternalServerError
		}

		for _, favorite := range res.Items {
			ids = append(ids, favorite.RecipeId)
			userFavorites[favorite.RecipeId] = true
		}
	}

	if req.IsFavorite == 0 {
		ids = make([]int64, 0)
	}

	res, err := f.rc.Recipes(ctx, &rsV1.RecipesRequest{
		Id:         req.Id,
		CategoryId: req.CategoryId,
		Weight:     req.Weight,
		Page:       req.Page,
		Size:       req.Size,
		Ids:        ids,
	})

	if err != nil {
		return nil, errors.InternalServerError
	}

	// views
	vIds := make([]string, 0)
	for _, recipe := range res.Items {
		vIds = append(vIds, fmt.Sprintf("%d", recipe.Id))
	}
	if len(vIds) == 0 {
		return &v1.RecipesReply{Items: make([]*v1.RecipeReply, 0)}, nil
	}

	rvs, err := f.r.HMGet(ctx, RecipeView, vIds...).Result()
	if err != nil {
		return nil, errors.InternalServerError
	}

	items := make([]*v1.RecipeReply, 0)
	for i, recipe := range res.Items {
		//favorite
		var isFavorite int64
		if _, ok := userFavorites[recipe.Id]; ok {
			isFavorite = 1
		}

		//views
		var views int64
		if rvs[i] != nil {
			if v, err := strconv.ParseInt(rvs[i].(string), 10, 64); err == nil {
				views = v
			}
		}

		items = append(items, &v1.RecipeReply{
			Id:         recipe.Id,
			Title:      recipe.Title,
			CategoryId: recipe.CategoryId,
			Content:    recipe.Content,
			Weight:     recipe.Weight,
			Image:      recipe.Image,
			Sort:       recipe.Sort,
			IsFavorite: isFavorite,
			Views:      views,
			CreateTime: recipe.CreateTime,
			UpdateTime: recipe.UpdateTime,
		})
	}

	return &v1.RecipesReply{Items: items}, nil
}

func (f *Frontend) Recipe(ctx context.Context, req *v1.RecipeRequest) (*v1.RecipeReply, error) {
	recipe, err := f.rc.Recipe(ctx, &rsV1.RecipeRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, errors.UnFound
	}
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.InternalServerError
	}

	var isFavorite int64
	if userId := md.Get("user_id"); len(userId) > 0 {
		res, err := f.rc.Favorite(ctx, &rsV1.FavoriteRequest{
			RecipeId: req.Id,
			UserId:   userId,
		})

		if err != nil {
			return nil, errors.UnFound
		}

		if res != nil && res.Id > 0 {
			isFavorite = 1
		}
	}

	views, err := f.r.HIncrBy(ctx, RecipeView, fmt.Sprintf("%d", recipe.Id), 1).Result()
	if err != nil {
		return nil, errors.ServiceUnavailable
	}

	return &v1.RecipeReply{
		Id:         recipe.Id,
		Title:      recipe.Title,
		CategoryId: recipe.CategoryId,
		Content:    recipe.Content,
		Weight:     recipe.Weight,
		Image:      recipe.Image,
		Sort:       recipe.Sort,
		IsFavorite: isFavorite,
		Views:      views,
		CreateTime: recipe.CreateTime,
		UpdateTime: recipe.UpdateTime,
	}, nil
}

func (f *Frontend) SaveFavorites(ctx context.Context, req *v1.SaveFavoritesRequest) (*v1.SaveFavoritesReply, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.Unauthorized
	}

	recipe, err := f.rc.Recipe(ctx, &rsV1.RecipeRequest{
		Id: req.RecipeId,
	})

	if recipe == nil || err != nil {
		return nil, errors.UnFound
	}

	_, err = f.rc.SaveFavorites(ctx, &rsV1.SaveFavoritesRequest{
		RecipeId: req.RecipeId,
		UserId:   md.Get("user_id"),
		Action:   req.Action,
	})

	if err != nil {
		return nil, errors.InternalServerError
	}

	return &v1.SaveFavoritesReply{}, nil
}
