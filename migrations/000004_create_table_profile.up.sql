CREATE TABLE "profiles" (
    "id"              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_user"         BIGINT NOT NULL UNIQUE REFERENCES "users"("id") ON DELETE CASCADE,
    "name"            VARCHAR(150) NOT NULL,
    "gender"          VARCHAR(10),
    "place_birth"     VARCHAR(100),
    "date_birth"      DATE,
    "created_at"      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "profiles"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();