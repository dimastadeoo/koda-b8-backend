CREATE TABLE "categories" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(100) NOT NULL,
    "url_img"     VARCHAR(255),
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "categories"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();