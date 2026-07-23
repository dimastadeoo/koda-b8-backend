CREATE TABLE "users" (
    "id"              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_role"         BIGINT NOT NULL REFERENCES "roles"("id"),
    "password"        VARCHAR(255) NOT NULL,
    "email"           VARCHAR(150) NOT NULL UNIQUE,
    "hp_number"       VARCHAR(20) NOT NULL UNIQUE,
    "status_account"  VARCHAR(20) NOT NULL DEFAULT 'active'
                      CHECK ("status_account" IN ('active', 'suspend', 'block')),
    "created_at"      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();