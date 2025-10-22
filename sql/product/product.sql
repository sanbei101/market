-- name: CreateProductSPU :one
INSERT INTO product_spu (
    name, 
    description, 
    category_id
) VALUES (
    @name, @description, @category_id
) RETURNING *;

-- name: GetProductSPU :one
SELECT * FROM product_spu WHERE id = @id;

-- name: ListProductSPUs :many
SELECT * FROM product_spu 
ORDER BY created_at DESC;

-- name: ListProductSPUsByCategory :many
SELECT * FROM product_spu 
WHERE category_id = @category_id
ORDER BY created_at DESC;

-- name: UpdateProductSPU :one
UPDATE product_spu 
SET 
    name = @name,
    description = @description,
    category_id = @category_id
WHERE id = @id
RETURNING *;

-- name: DeleteProductSPU :exec
DELETE FROM product_spu WHERE id = @id;

-----------------------------------------------------------
-- name: CreateProductSKU :one
INSERT INTO product_sku (
    spu_id, 
    price, 
    stock, 
    specs
) VALUES (
    @spu_id, @price, @stock, @specs
) RETURNING *;

-- name: GetProductSKU :one
SELECT * FROM product_sku WHERE id = @id;

-- name: ListProductSKUsBySPU :many
SELECT * FROM product_sku 
WHERE spu_id = @spu_id
ORDER BY created_at DESC;

-- name: UpdateProductSKU :one
UPDATE product_sku 
SET 
    price = @price,
    stock = @stock,
    specs = @specs
WHERE id = @id
RETURNING *;

-- name: DeleteProductSKU :exec
DELETE FROM product_sku WHERE id = @id;

-- name: DeleteProductSKUsBySPU :exec
DELETE FROM product_sku WHERE spu_id = @spu_id;

-----------------------------------------------------------
-- name: SearchProductsBySpecs :many
SELECT 
    sku.*,
    spu.name as spu_name,
    spu.description as spu_description
FROM product_sku sku
JOIN product_spu spu ON sku.spu_id = spu.id
WHERE sku.specs @> @specs_filter
AND spu.category_id = @category_id;

-- name: GetCategorySpecValues :many
SELECT 
    jsonb_object_keys(specs) as spec_key,
    specs->>jsonb_object_keys(specs) as spec_value,
    COUNT(*) as sku_count
FROM product_sku sku
JOIN product_spu spu ON sku.spu_id = spu.id
WHERE spu.category_id = @category_id
GROUP BY spec_key, spec_value
ORDER BY spec_key, sku_count DESC;

-- name: GetSPUWithSKUs :one
SELECT 
    spu.*,
    jsonb_agg(sku.*) as skus
FROM product_spu spu
LEFT JOIN product_sku sku ON spu.id = sku.spu_id
WHERE spu.id = @id
GROUP BY spu.id;