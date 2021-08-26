CREATE USER dateguess_api@'%' IDENTIFIED BY 'password';

-- Reflect live permissions.
GRANT SELECT, INSERT, UPDATE, DELETE ON history.* TO dateguess_api@'%';

-- These permissions are for dev / testing purposes only
GRANT INSERT, UPDATE, DELETE, DROP ON *.* TO dateguess_api@'%';
