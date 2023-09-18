-- Mock database for testing
-- User table
CREATE TABLE users (
    User_ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,

    Username VARCHAR(50),
    Password VARCHAR(255),
    Created BIGINT
);

-- test data
INSERT INTO users (Username, Password) VALUES
    ('TestUser', 'TestPw');

-- Different levels
CREATE TABLE levels (
    LevelID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,

    LevelName VARCHAR(255),
    MaxEntries INT NOT NULL
);

-- User progress
CREATE TABLE level_progress (
    ProgressID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,

    UserID INT NOT NULL,
    LevelID INT NOT NULL,
    Progress INT DEFAULT 0,
    
    FOREIGN KEY (UserID) REFERENCES users(User_ID),
    FOREIGN KEY (LevelID) REFERENCES levels(LevelID)
);

-- Level data
INSERT INTO levels (LevelName, MaxEntries) VALUES
    ('Hiragana Basic', 5),
    ('Hiragana Extended', 5);

-- Hiragana table
CREATE TABLE hiragana (
    HiraID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,

    Symbol VARCHAR(5),
    Translation VARCHAR(5)
);

-- load Hiragana characters data
INSERT INTO hiragana (Symbol, Translation) VALUES
-- base
    ('あ', 'a'),
    ('い', 'i'),
    ('う', 'u'),
    ('え', 'e'),
    ('お', 'o');

-- more hiragana characters
CREATE TABLE hiragana_extended (
    HiraExtID INT NOT NULL AUTO_INCREMENT,
    Symbol VARCHAR(5),
    Translation VARCHAR(5),

    PRIMARY KEY(HiraExtID)
);

-- add extended hiragana characters
INSERT INTO hiragana_extended (Symbol, Translation) VALUES
-- g
    ('が', 'ga'),
    ('ぎ', 'gi'),
    ('ぐ', 'gu'),
    ('げ', 'ge'),
    ('ご', 'go');