CREATE TABLE "product_specification" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_product"  BIGINT NOT NULL REFERENCES "products"("id") ON DELETE CASCADE,
    "key"         VARCHAR(100) NOT NULL,
    "value"       VARCHAR(255) NOT NULL,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "product_specification"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();