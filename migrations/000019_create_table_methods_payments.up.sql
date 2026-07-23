CREATE TABLE "methods_payments" (
    "id"          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"        VARCHAR(100) NOT NULL,
    "created_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- trigger update
CREATE TRIGGER "set_updated_at"
BEFORE UPDATE ON "methods_payments"
FOR EACH ROW
EXECUTE FUNCTION "trigger_set_updated_at"();