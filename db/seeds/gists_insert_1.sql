USE gistbox;

-- Add some dummy records
INSERT INTO gists(title, content, created, expires) VALUES(
    "I've just seen a face",
    "I've just seen a face.\nI cant forget the time or place",
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO gists(title, content, created, expires) VALUES(
    "I'm in the corner",
    "I'm in the corner like a tired boxer",
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO gists(title, content, created, expires) VALUES(
    "Last one ale",
    "Last one ale for the see.",
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
