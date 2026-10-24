USE gistbox;

DROP TABLE IF EXISTS gists;

-- Create a `gists` table
CREATE TABLE gists (
    id      INTEGER      NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_gists_created (created)
);
