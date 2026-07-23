-- Products
CREATE INDEX IF NOT EXISTS "idx_products_id_merk" ON "products"("id_merk");
CREATE INDEX IF NOT EXISTS "idx_products_name" ON "products"("name");

-- Product Specifications
CREATE INDEX IF NOT EXISTS "idx_product_spec_product" ON "product_specification"("id_product");

-- Images
CREATE INDEX IF NOT EXISTS "idx_img_product_product" ON "img_product"("id_product");

-- Reviews
CREATE INDEX IF NOT EXISTS "idx_reviews_product" ON "reviews"("id_product");

-- Carts
CREATE INDEX IF NOT EXISTS "idx_carts_list_cart" ON "carts_list"("id_cart");

-- Orders (perhatikan ada 2 index)
CREATE INDEX IF NOT EXISTS "idx_orders_cart" ON "orders"("id_cart");
CREATE INDEX IF NOT EXISTS "idx_orders_status" ON "orders"("status");

-- Order Items
CREATE INDEX IF NOT EXISTS "idx_order_items_order" ON "order_items"("id_order");

-- Payment
CREATE INDEX IF NOT EXISTS "idx_payment_trx_order" ON "payment_transactions"("id_order");

-- Logs & Sessions
CREATE INDEX IF NOT EXISTS "idx_users_logs_user" ON "users_logs"("id_user");
CREATE INDEX IF NOT EXISTS "idx_sessions_user" ON "sessions"("id_user");