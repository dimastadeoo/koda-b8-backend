CREATE TABLE "merks" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(150) NOT NULL,
    "info"        TEXT,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "merks"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();