CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE departments IS 'Список отделов университета';
COMMENT ON COLUMN departments.name IS 'Название отдела';

CREATE OR REPLACE TRIGGER update_departments_modtime
BEFORE UPDATE ON departments
FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();