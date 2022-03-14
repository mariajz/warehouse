CREATE TABLE "warehouses" (
                            id bigserial PRIMARY KEY,
                            item_name varchar NOT NULL,
                            quantity bigint NOT NULL,
                            city varchar NOT NULL
);