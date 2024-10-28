-- +goose Up

CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    userType ENUM('Applicant', 'Admin') NOT NULL,
    profileHeadline VARCHAR(255) NOT NULL
);

CREATE TABLE Profile(
    id SERIAL PRIMARY KEY,
    resumeFileAddress TEXT,
    skills TEXT,
    education TEXT,
    experience TEXT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    phone VARCHAR(20),

    FOREIGN KEY (email) REFERENCES User(email) ON DELETE CASCADE
);

CREATE TABLE Job(
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    companyName VARCHAR(150) NOT NULL,
    postedOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    postedBy INT NOT NULL,
    totalApplicants INT DEFAULT 0,

    FOREIGN KEY (postedBy) REFERENCES User(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Profile;
DROP TABLE IF EXISTS Job;