CREATE TABLE "method_shippings" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(100) NOT NULL,
    "price"       INT NOT NULL CHECK ("price" >= 0),
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "method_shippings"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();