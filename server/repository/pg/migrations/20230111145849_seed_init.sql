-- +goose Up
-- +goose StatementBegin
insert into "users" ("id","age","name","is_admin") values 
('6c00e395-0656-4f22-8fda-01addc593097',12,'name-1',false),
('0a11366c-c429-4c3a-a103-4f0c8fe2b38f',13,'name-1',false);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate "users" 
-- +goose StatementEnd
