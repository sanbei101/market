CREATE TABLE category (
    id             BIGSERIAL PRIMARY KEY,
    name           VARCHAR(100) NOT NULL,
    parent_id      BIGINT REFERENCES category(id) DEFAULT NULL,
    spec_template  JSONB NOT NULL,
    sort           INT NOT NULL DEFAULT 0,
    created_at     TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE product_spu (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    category_id BIGINT NOT NULL REFERENCES category(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE product_sku (
    id          BIGSERIAL PRIMARY KEY,
    spu_id      BIGINT NOT NULL REFERENCES product_spu(id),
    price       DECIMAL(10,2) NOT NULL,
    stock       INT NOT NULL DEFAULT 0,
    specs       JSONB NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

