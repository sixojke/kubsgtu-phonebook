DROP TRIGGER IF EXISTS update_contacts_modtime ON contacts;

DROP INDEX IF EXISTS idx_contacts_fullname;
DROP INDEX IF EXISTS idx_contacts_phone;
DROP INDEX IF EXISTS idx_contacts_department;

COMMENT ON TABLE contacts IS NULL;
COMMENT ON COLUMN contacts.name IS NULL;
COMMENT ON COLUMN contacts.position IS NULL;
COMMENT ON COLUMN contacts.phone IS NULL;
COMMENT ON COLUMN contacts.email IS NULL;
COMMENT ON COLUMN contacts.room IS NULL;
COMMENT ON COLUMN contacts.department_id IS NULL;

DROP TABLE IF EXISTS contacts;