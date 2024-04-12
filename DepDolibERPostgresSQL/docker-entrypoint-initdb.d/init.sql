-- Créer un utilisateur pour la base de données Dolibarr
CREATE USER dolibarrdbuser WITH PASSWORD 'doli_pass';

-- Créer une base de données pour Dolibarr si elle n'existe pas déjà
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'dolidb') THEN
        CREATE DATABASE dolidb;
        GRANT ALL PRIVILEGES ON DATABASE dolidb TO dolibarrdbuser;
    END IF;
END $$;

