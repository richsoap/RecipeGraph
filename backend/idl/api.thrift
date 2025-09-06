// idl/api.thrift
namespace go recipe.api

// Common response structure
struct BaseResponse {
    1: optional i32 code,
    2: optional string message,
}

// Item related structures
struct Item {
    1: optional i64 id,
    2: optional string name,
    3: optional string created_at,
    4: optional string updated_at,
}

struct CreateItemReq {
    1: optional string name (api.body = "name"),
}

struct CreateItemResp {
    1: optional BaseResponse base,
    2: optional Item data,
}

struct GetItemReq {
    1: optional i64 id (api.path = "id"),
}

struct GetItemResp {
    1: optional BaseResponse base,
    2: optional Item data,
}

struct UpdateItemReq {
    1: optional i64 id (api.path = "id"),
    2: optional string name (api.body = "name"),
}

struct UpdateItemResp {
    1: optional BaseResponse base,
    2: optional Item data,
}

struct DeleteItemReq {
    1: optional i64 id (api.path = "id"),
}

struct DeleteItemResp {
    1: optional BaseResponse base,
}

struct ListItemsReq {
    1: optional i32 page (api.query = "page"),
    2: optional i32 limit (api.query = "limit"),
    3: optional string name (api.query = "name"),
}

struct ListItemsResp {
    1: optional BaseResponse base,
    2: optional list<Item> data,
    3: optional i32 total,
}

// Recipe related structures
struct Recipe {
    1: optional i64 id,
    2: optional i64 space,
    3: optional i64 item_id,
    4: optional string item_name,
    5: optional double efficiency,
    16: optional list<RecipeItem> items,
    32: optional string created_at,
    33: optional string updated_at,
}

enum RecipeItemType {
    INPUT = 1,
    OUTPUT = 2,
}

struct RecipeItem {
    1: optional i64 id,
    2: optional i64 recipe_id,
    3: optional i64 space,
    4: optional RecipeItemType type,
    5: optional i32 count,
    6: optional Item item,
    7: optional string created_at,
    8: optional string updated_at,
}

struct CreateRecipeReq {
    1: optional i64 space (api.body = "space"),
    2: optional i64 item_id (api.body = "item_id"),
    3: optional double efficiency (api.body = "efficiency"),
    4: optional list<RecipeItem> items (api.body = "items"),
}

struct CreateRecipeResp {
    1: optional BaseResponse base,
    2: optional Recipe data,
}

struct GetRecipeReq {
    1: optional i64 id (api.path = "id"),
}

struct GetRecipeResp {
    1: optional BaseResponse base,
    2: optional Recipe data,
}

struct UpdateRecipeReq {
    1: optional Recipe data,
}

struct UpdateRecipeResp {
    1: optional BaseResponse base,
    2: optional Recipe data,
}

struct DeleteRecipeReq {
    1: optional i64 id (api.path = "id"),
}

struct DeleteRecipeResp {
    1: optional BaseResponse base,
}

struct ListRecipesReq {
    1: optional i32 page (api.query = "page"),
    2: optional i32 limit (api.query = "limit"),
    3: optional i64 space (api.query = "space"),
    4: optional i64 item_id (api.query = "item_id"),
}

struct ListRecipesResp {
    1: optional BaseResponse base,
    2: optional list<Recipe> data,
    3: optional i32 total,
}

// Service definitions
service ItemService {
    CreateItemResp CreateItem(1: CreateItemReq request) (api.post = "/api/v1/items"),
    GetItemResp GetItem(1: GetItemReq request) (api.get = "/api/v1/items/:id"),
    UpdateItemResp UpdateItem(1: UpdateItemReq request) (api.put = "/api/v1/items/:id"),
    DeleteItemResp DeleteItem(1: DeleteItemReq request) (api.delete = "/api/v1/items/:id"),
    ListItemsResp ListItems(1: ListItemsReq request) (api.get = "/api/v1/items"),
}

service RecipeService {
    CreateRecipeResp CreateRecipe(1: CreateRecipeReq request) (api.post = "/api/v1/recipes"),
    GetRecipeResp GetRecipe(1: GetRecipeReq request) (api.get = "/api/v1/recipes/:id"),
    UpdateRecipeResp UpdateRecipe(1: UpdateRecipeReq request) (api.put = "/api/v1/recipes/:id"),
    DeleteRecipeResp DeleteRecipe(1: DeleteRecipeReq request) (api.delete = "/api/v1/recipes/:id"),
    ListRecipesResp ListRecipes(1: ListRecipesReq request) (api.get = "/api/v1/recipes"),
}