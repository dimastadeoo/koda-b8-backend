CREATE TABLE "roles" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(50) NOT NULL UNIQUE, -- admin, customer, staff
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "roles"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();