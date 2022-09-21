CREATE TABLE public.transactions (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,
    code VARCHAR(150) NOT NULL,
    payment_url VARCHAR(150) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);