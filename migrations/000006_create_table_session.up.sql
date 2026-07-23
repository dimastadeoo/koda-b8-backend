CREATE TABLE "sessions" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_user"     BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "token"       VARCHAR(255) NOT NULL UNIQUE,
    "status"      VARCHAR(20) NOT NULL DEFAULT 'active',
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "sessions"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();