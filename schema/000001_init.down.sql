UPDATE schema_migrations SET dirty = false;

DROP TABLE IF EXISTS photo CASCADE;

DROP TABLE IF EXISTS advert CASCADE;
