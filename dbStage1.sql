DROP DATABASE IF EXISTS `db`;
CREATE DATABASE `db`;

-- Erstelle die Tabelle User
CREATE TABLE User (
  id INT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  key VARCHAR(255) NOT NULL
);

-- Erstelle die Tabelle Personchats
CREATE TABLE Personchats (
  id INT PRIMARY KEY,
  user1 INT NOT NULL,
  user2 INT NOT NULL,
  FOREIGN KEY (user1) REFERENCES User(id),
  FOREIGN KEY (user2) REFERENCES User(id)
);

-- Erstelle die Tabelle Message
CREATE TABLE Message (
  id INT PRIMARY KEY,
  message VARCHAR(255) NOT NULL,
  status VARCHAR(255) NOT NULL,
  time DATETIME NOT NULL,
  id_Personchats INT NOT NULL,
  sender INT NOT NULL,
  FOREIGN KEY (id_Personchats) REFERENCES Personchats(id),
  FOREIGN KEY (sender) REFERENCES User(id)
);