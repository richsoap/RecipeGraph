// idl/api.thrift
namespace go recipe.api

// Common response structure
struct BaseResponse {
    1: i32 code;
    2: string message;
}

// Item related structures
struct Item {
    1: i64 id;
    2: string name;
    3: string created_at;
    4: string updated_at;
}

struct CreateItemReq {
    1: string name (api.body="name");
}

struct CreateItemResp {
    1: BaseResponse base;
    2: Item data;
}

struct GetItemReq {
    1: i64 id (api.path="id");
}

struct GetItemResp {
    1: BaseResponse base;
    2: Item data;
}

struct UpdateItemReq {
    1: i64 id (api.path="id");
    2: string name (api.body="name");
}

struct UpdateItemResp {
    1: BaseResponse base;
    2: Item data;
}

struct DeleteItemReq {
    1: i64 id (api.path="id");
}

struct DeleteItemResp {
    1: BaseResponse base;
}

struct ListItemsReq {
    1: optional i32 page (api.query="page");
    2: optional i32 limit (api.query="limit");
    3: optional string name (api.query="name");
}

struct ListItemsResp {
    1: BaseResponse base;
    2: list<Item> data;
    3: i32 total;
}

// Recipe related structures
struct Recipe {
    1: i64 id;
    2: i64 space;
    3: i64 item_id;
    4: double efficiency;
    5: string created_at;
    6: string updated_at;
}

struct RecipeItem {
    1: i64 id;
    2: i64 recipe_id;
    3: i64 space;
    4: i32 type;
    5: i64 item_id;
    6: i32 count;
    7: string created_at;
    8: string updated_at;
}

struct CreateRecipeReq {
    1: i64 space (api.body="space");
    2: i64 item_id (api.body="item_id");
    3: double efficiency (api.body="efficiency");
    4: list<RecipeItem> items (api.body="items");
}

struct CreateRecipeResp {
    1: BaseResponse base;
    2: Recipe data;
}

struct GetRecipeReq {
    1: i64 id (api.path="id");
}

struct GetRecipeResp {
    1: BaseResponse base;
    2: Recipe data;
    3: list<RecipeItem> items;
}

struct UpdateRecipeReq {
    1: i64 id (api.path="id");
    2: optional i64 space (api.body="space");
    3: optional i64 item_id (api.body="item_id");
    4: optional double efficiency (api.body="efficiency");
    5: optional list<RecipeItem> items (api.body="items");
}

struct UpdateRecipeResp {
    1: BaseResponse base;
    2: Recipe data;
}

struct DeleteRecipeReq {
    1: i64 id (api.path="id");
}

struct DeleteRecipeResp {
    1: BaseResponse base;
}

struct ListRecipesReq {
    1: optional i32 page (api.query="page");
    2: optional i32 limit (api.query="limit");
    3: optional i64 space (api.query="space");
    4: optional i64 item_id (api.query="item_id");
}

struct ListRecipesResp {
    1: BaseResponse base;
    2: list<Recipe> data;
    3: i32 total;
}

// Service definitions
service ItemService {
    CreateItemResp CreateItem(1: CreateItemReq request) (api.post="/api/v1/items");
    GetItemResp GetItem(1: GetItemReq request) (api.get="/api/v1/items/:id");
    UpdateItemResp UpdateItem(1: UpdateItemReq request) (api.put="/api/v1/items/:id");
    DeleteItemResp DeleteItem(1: DeleteItemReq request) (api.delete="/api/v1/items/:id");
    ListItemsResp ListItems(1: ListItemsReq request) (api.get="/api/v1/items");
}

service RecipeService {
    CreateRecipeResp CreateRecipe(1: CreateRecipeReq request) (api.post="/api/v1/recipes");
    GetRecipeResp GetRecipe(1: GetRecipeReq request) (api.get="/api/v1/recipes/:id");
    UpdateRecipeResp UpdateRecipe(1: UpdateRecipeReq request) (api.put="/api/v1/recipes/:id");
    DeleteRecipeResp DeleteRecipe(1: DeleteRecipeReq request) (api.delete="/api/v1/recipes/:id");
    ListRecipesResp ListRecipes(1: ListRecipesReq request) (api.get="/api/v1/recipes");
}