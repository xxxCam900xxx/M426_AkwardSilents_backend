CREATE TABLE User (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  key TEXT NOT NULL
);

-- Erstelle die Tabelle Personchats
CREATE TABLE Personchats (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user1 INTEGER NOT NULL,
  user2 INTEGER NOT NULL,
  FOREIGN KEY (user1) REFERENCES User(id),
  FOREIGN KEY (user2) REFERENCES User(id)
);

-- Erstelle die Tabelle Message
CREATE TABLE Message (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  message TEXT NOT NULL,
  status TEXT NOT NULL,
  time INTEGER NOT NULL,
  id_Personchats INTEGER NOT NULL,
  sender INTEGER NOT NULL,
  FOREIGN KEY (id_Personchats) REFERENCES Personchats(id),
  FOREIGN KEY (sender) REFERENCES User(id)
);
