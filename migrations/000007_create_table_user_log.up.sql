CREATE TABLE "users_logs" (
    "id_user"         BIGINT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "id_session"      BIGINT REFERENCES "sessions"("id") ON DELETE SET NULL,
    "activity_detail" TEXT NOT NULL,
    "ip_address"      VARCHAR(45),
    "created_at"      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);