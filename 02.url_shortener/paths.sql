DROP TABLE IF EXISTS paths;
CREATE TABLE paths(path VARCHAR(255), url VARCHAR(255));
INSERT INTO paths(path, url) VALUES('/mysql', 'https://zetcode.com/golang/mysql/');
INSERT INTO paths(path, url) VALUES('/go-memdb', 'https://pkg.go.dev/github.com/hashicorp/go-memdb');