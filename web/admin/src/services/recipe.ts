// @ts-ignore
/* eslint-disable */
import {request} from 'umi';

export async function recipeIndex() {
  return request<API.Default<API.RecipeIndex>>('/v1/recipe/index', {
    method: 'GET'
  });
}

export async function saveRecipeIndex(body: API.RecipeIndex, options?: Record<string, any>) {
  return request<API.Default<any>>('/v1/recipe/index/save', {
    method: 'POST',
    data  : body,
    ...(options || {}),
  });
}


export async function recipeCategories(page: number, size: number, options?: { all: number, parent_id: string, not_id: string }) {
  const res = await request<API.DefaultList<API.RecipeCategoryItem>>('/v1/recipe/categories', {
    method: 'GET',
    params: {
      page,
      size,
      ...options
    }
  });

  return {
    data   : res.data.items ? res.data.items : [],
    total  : res.data.total,
    success: res.code === 0,
  };
}

export async function saveRecipeCategory(body: API.RecipeCategoryItem, options?: Record<string, any>) {
  return request<API.Default<number>>('/v1/recipe/category/save', {
    method: 'POST',
    data  : body,
    ...(options || {}),
  });
}

export async function recipeCategory(id: number, options?: object) {
  return request<API.Default<API.RecipeCategoryItem>>('/v1/recipe/category', {
    method: 'GET',
    params: {
      id,
      ...options
    }
  });
}

export async function deleteRecipeCategory(id: number, options?: object) {
  return request('/v1/recipe/category/delete', {
    method: 'POST',
    data  : {
      id,
    },
    ...(options || {}),
  });
}

export async function recipes(page: number, size: number, options?: {id: number, category_id: number, title: string}) {
  const res = await request<API.DefaultList<API.RecipeItem>>('/v1/recipes', {
    method: 'GET',
    params: {
      page,
      size,
      ...options
    }
  });

  return {
    data   : res.data.items ? res.data.items : [],
    total  : res.data.total,
    success: res.code === 0,
  };
}

export async function saveRecipe(body: API.RecipeItem, options?: Record<string, any>) {
  return request<API.Default<number>>('/v1/recipe/save', {
    method: 'POST',
    data  : body,
    ...(options || {}),
  });
}

export async function recipe(id: number, options?: object) {
  return request<API.Default<API.RecipeItem>>('/v1/recipe', {
    method: 'GET',
    params: {
      id,
      ...options
    }
  });
}

export async function deleteRecipe(id: number, options?: object) {
  return request('/v1/recipe/delete', {
    method: 'POST',
    data  : {
      id,
    },
    ...(options || {}),
  });
}
