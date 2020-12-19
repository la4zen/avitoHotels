CREATE TABLE IF NOT EXISTS rooms(
    id INT PRIMARY KEY,
    discription TEXT,
    price INT,
    creationdate DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS booking(
    id INT PRIMARY KEY,
    roomid INT,
    datestart DATETIME,
    dateend DATETIME
);