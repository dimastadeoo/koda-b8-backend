CREATE TABLE "whishlists" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_profile"  BIGINT NOT NULL REFERENCES "profiles"("id") ON DELETE CASCADE,
    "id_product"  BIGINT NOT NULL REFERENCES "products"("id") ON DELETE CASCADE,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE ("id_profile", "id_product")
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "whishlists"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();