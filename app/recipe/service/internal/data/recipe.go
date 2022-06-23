package data

import (
	"context"
	"strconv"

	v1 "github.com/exuan/ruber/api/recipe/service/v1"
	"github.com/exuan/ruber/app/recipe/service/internal/biz"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/category"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/favorite"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/recipe"
	"github.com/go-kratos/kratos/v2/log"
)

type recipeRepo struct {
	data *Data
	log  *log.Helper
}

func NewRecipeRepo(data *Data, logger log.Logger) biz.RecipeRepo {
	return &recipeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (rp *recipeRepo) Recipes(ctx context.Context, req *v1.RecipesRequest) (int64, []*ent.Recipe, error) {
	sess := rp.data.db.Debug().Recipe.Query().
		Where(recipe.StatusNEQ(-1))

	if req.Id > 0 {
		sess.Where(recipe.ID(req.Id))
	}
	if len(req.Title) > 0 {
		sess.Where(recipe.TitleContains(req.Title))
	}
	if req.CategoryId > 0 {
		sess.Where(recipe.CategoryID(req.CategoryId))
	}

	// custom tag
	if req.Weight > 0 {
		sess.Where(recipe.Weight(req.Weight))
	}
	if len(req.Ids) > 0 {
		sess.Where(recipe.IDIn(req.Ids...))
	}

	total, err := sess.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	res, err := sess.Order(ent.Desc(recipe.FieldID)).
		Limit(int(req.Size)).
		Offset(int((req.Page - 1) * req.Size)).
		All(ctx)

	return int64(total), res, err
}

func (rp *recipeRepo) Recipe(ctx context.Context, req *v1.RecipeRequest) (*ent.Recipe, error) {
	return rp.data.db.Recipe.Query().
		Where(
			recipe.ID(req.Id),
			recipe.StatusNEQ(-1),
		).
		Only(ctx)
}

func (rp *recipeRepo) SaveRecipe(ctx context.Context, req *v1.SaveRecipeRequest) (int64, error) {
	var err error
	var er *ent.Recipe

	if req.Id > 0 {
		er, err = rp.data.db.Recipe.UpdateOneID(req.Id).
			SetTitle(req.Title).
			SetImage(req.Image).
			SetCategoryID(req.CategoryId).
			SetContent(req.Content).
			SetWeight(req.Weight).
			SetSort(req.Sort).
			SetModifier(req.UserId).
			Save(ctx)

	} else {

		er, err = rp.data.db.Recipe.Create().
			SetTitle(req.Title).
			SetCategoryID(req.CategoryId).
			SetImage(req.Image).
			SetContent(req.Content).
			SetWeight(req.Weight).
			SetSort(req.Sort).
			SetStatus(0).
			SetModifier(req.UserId).
			SetCreator(req.UserId).
			Save(ctx)
	}

	if err != nil {
		return 0, err
	}

	return er.ID, err
}

func (rp *recipeRepo) DeleteRecipe(ctx context.Context, req *v1.DeleteRecipeRequest) (int64, error) {
	et, err := rp.data.db.Recipe.
		UpdateOneID(req.Id).
		SetStatus(-1).
		SetModifier(req.UserId).
		Save(ctx)

	return et.ID, err
}

func (rp *recipeRepo) Categories(ctx context.Context, req *v1.CategoriesRequest) (int64, []*ent.Category, error) {
	sess := rp.data.db.Category.Query().
		Where(category.StatusNEQ(-1))

	if req.Id > 0 {
		sess.Where(category.ID(req.Id))
	}
	if len(req.Name) > 0 {
		sess.Where(category.NameContains(req.Name))
	}
	if len(req.ParentId) > 0 {
		ParentId, err := strconv.ParseInt(req.ParentId, 10, 64)
		if err == nil && ParentId > -1 {
			sess.Where(category.ParentID(ParentId))
		}
	}
	if len(req.NotId) > 0 {
		notId, err := strconv.ParseInt(req.NotId, 10, 64)
		if err == nil && notId > -1 {
			sess.Where(category.IDNEQ(notId))
		}
	}

	total, err := sess.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	if req.All == 0 {
		sess.Limit(int(req.Size)).
			Offset(int((req.Page - 1) * req.Size))
	}
	res, err := sess.Order(ent.Desc(category.FieldSort)).
		All(ctx)

	return int64(total), res, err
}

func (rp *recipeRepo) Category(ctx context.Context, req *v1.CategoryRequest) (*ent.Category, error) {
	return rp.data.db.Category.Query().
		Where(
			category.ID(req.Id),
			category.StatusNEQ(-1),
		).
		Only(ctx)
}

func (rp *recipeRepo) SaveCategory(ctx context.Context, req *v1.SaveCategoryRequest) (int64, error) {
	var err error
	var ec *ent.Category

	if req.Id > 0 {
		ec, err = rp.data.db.Category.UpdateOneID(req.Id).
			SetName(req.Name).
			SetParentID(req.ParentId).
			SetSort(req.Sort).
			SetModifier(req.UserId).
			Save(ctx)

	} else {
		ec, err = rp.data.db.Category.Create().
			SetName(req.Name).
			SetParentID(req.ParentId).
			SetSort(req.Sort).
			SetStatus(0).
			SetModifier(req.UserId).
			SetCreator(req.UserId).
			Save(ctx)
	}

	if err != nil {
		return 0, err
	}

	return ec.ID, err
}

func (rp *recipeRepo) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) (int64, error) {
	et, err := rp.data.db.Category.
		UpdateOneID(req.Id).
		SetModifier(req.UserId).
		SetStatus(-1).
		Save(ctx)

	return et.ID, err
}

func (rp *recipeRepo) CountChildrenCategories(ctx context.Context, parentId int64) (int, error) {
	return rp.data.db.Category.Query().
		Where(category.StatusNEQ(-1)).
		Where(category.ParentID(parentId)).
		Count(ctx)
}

func (rp *recipeRepo) UpdateRecipeFavorite(ctx context.Context, id int64, action int32) (int64, error) {
	var i int64 = 1
	if action == 0 {
		i = -1
	}

	err := rp.data.db.Recipe.Update().Where(
		recipe.StatusNEQ(-1),
		recipe.IDEQ(id),
		recipe.FavoritesGT(-1),
	).AddFavorites(i).Exec(ctx)
	return id, err
}

func (rp *recipeRepo) Favorites(ctx context.Context, req *v1.FavoritesRequest) (int64, []*ent.Favorite, error) {
	sess := rp.data.db.Favorite.Query()

	if len(req.UserId) > 0 {
		sess.Where(favorite.UserID(req.UserId))
	}

	total, err := sess.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	res, err := sess.Order(ent.Desc(favorite.FieldRecipeID)).
		Limit(int(req.Size)).
		Offset(int((req.Page - 1) * req.Size)).
		All(ctx)
	return int64(total), res, err
}

func (rp *recipeRepo) Favorite(ctx context.Context, req *v1.FavoriteRequest) (*ent.Favorite, error) {
	return rp.data.db.Favorite.Query().
		Where(
			favorite.RecipeID(req.RecipeId),
			favorite.UserID(req.UserId),
		).
		Only(ctx)
}

func (rp *recipeRepo) SaveFavorites(ctx context.Context, req *v1.SaveFavoritesRequest) (int64, error) {
	if req.Action == 1 {
		id, err := rp.data.db.Favorite.
			Create().
			SetRecipeID(req.RecipeId).
			SetUserID(req.UserId).
			OnConflictColumns(favorite.FieldRecipeID, favorite.FieldUserID).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return 0, err
		}
		return id, nil
	}

	if req.Action == 0 {
		res, err := rp.data.db.Favorite.Delete().Where(
			favorite.RecipeID(req.RecipeId),
			favorite.UserID(req.UserId),
		).Exec(ctx)
		if err != nil {
			return 0, err
		}
		return int64(res), nil
	}

	return 0, nil
}
