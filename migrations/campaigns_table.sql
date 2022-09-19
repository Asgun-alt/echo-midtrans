CREATE TABLE public.campaigns (
    id SERIAL PRIMARY KEY,
    campaign_name VARCHAR(150) NOT NULL,
    description VARCHAR(350) NOT NULL,
    perks VARCHAR(100) NOT NULL,
    backer_count INTEGER NOT NULL,
    goal_amount INTEGER NOT NULL,
    current_amount INTEGER NOT NULL,
    slug VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE public.campaign_images (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER NOT NULL,
    is_primary BOOLEAN NOT NULL,
    file_name VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
);