-- name: GenerateTestData :exec
CREATE OR REPLACE FUNCTION generate_test_data(
    category_count INT DEFAULT 5,
    spu_per_category INT DEFAULT 10,
    sku_per_spu INT DEFAULT 3
) RETURNS VOID AS $$
DECLARE
    cat_id BIGINT;
    spu_id BIGINT;
    i INT;
    j INT;
    k INT;
    color_options TEXT[] := ARRAY['黑色', '白色', '蓝色', '红色', '金色'];
    storage_options TEXT[] := ARRAY['64GB', '128GB', '256GB', '512GB', '1TB'];
    memory_options TEXT[] := ARRAY['4GB', '6GB', '8GB', '12GB', '16GB'];
    category_names TEXT[] := ARRAY['手机', '电脑', '平板', '耳机', '手表', '相机', '电视', '音响'];
    product_names TEXT[] := ARRAY['旗舰', '专业', '青春版', '尊享版', '标准版', '至尊版'];
    brand_names TEXT[] := ARRAY['苹果', '小米', '华为', '三星', 'OPPO', 'VIVO'];
BEGIN
    TRUNCATE TABLE product_sku, product_spu, category RESTART IDENTITY CASCADE;
    
    -- 生成分类数据
    FOR i IN 1..category_count LOOP
        INSERT INTO category (name, parent_id, spec_template, sort)
        VALUES (
            category_names[(i-1) % array_length(category_names, 1) + 1] || '分类',
            CASE WHEN i > 3 THEN (random() * 2)::INT + 1 ELSE NULL END,
            jsonb_build_object(
                '颜色', color_options,
                '存储', storage_options,
                '内存', memory_options
            ),
            i
        ) RETURNING id INTO cat_id;
        
        -- 生成SPU数据
        FOR j IN 1..spu_per_category LOOP
            INSERT INTO product_spu (name, description, category_id)
            VALUES (
                brand_names[(j-1) % array_length(brand_names, 1) + 1] || ' ' || 
                product_names[(j-1) % array_length(product_names, 1) + 1] || ' ' ||
                (i * 10 + j),
                '这是' || brand_names[(j-1) % array_length(brand_names, 1) + 1] || 
                product_names[(j-1) % array_length(product_names, 1) + 1] || '的产品描述',
                cat_id
            ) RETURNING id INTO spu_id;
            
            -- 生成SKU数据
            FOR k IN 1..sku_per_spu LOOP
                INSERT INTO product_sku (spu_id, price, stock, specs)
                VALUES (
                    spu_id,
                    (random() * 5000 + 1000)::DECIMAL(10,2), -- 价格 1000-6000
                    (random() * 100)::INT, -- 库存 0-100
                    jsonb_build_object(
                        '颜色', color_options[(k-1) % array_length(color_options, 1) + 1],
                        '存储', storage_options[(k-1) % array_length(storage_options, 1) + 1],
                        '内存', memory_options[(k-1) % array_length(memory_options, 1) + 1]
                    )
                );
            END LOOP;
        END LOOP;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

-- name: CallGenerateTestData :exec
SELECT generate_test_data(
    @category_count::INT, 
    @spu_per_category::INT, 
    @sku_per_spu::INT
);