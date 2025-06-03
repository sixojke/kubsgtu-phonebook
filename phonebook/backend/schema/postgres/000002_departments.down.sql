DROP TRIGGER IF EXISTS update_departments_modtime ON departments;

COMMENT ON TABLE departments IS NULL;
COMMENT ON COLUMN departments.name IS NULL;

DROP TABLE IF EXISTS departments;