CREATE TABLE "forgot_pin" (
    "id"                  BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "id_session"          BIGINT REFERENCES "sessions"("id") ON DELETE SET NULL,
    "id_user"             BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "reset_token"         VARCHAR(255) NOT NULL UNIQUE,
    "token_expired_at"    TIMESTAMPTZ NOT NULL,
    "ip_address"          VARCHAR(45),
    "created_at"          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);