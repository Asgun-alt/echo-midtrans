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