-- Passwords are 'password'
INSERT INTO users (username, display_name, hashed_password)
VALUES
  ('gretel', 'Gretel', '$2a$10$QRZ9dpL2JRGTK7hNG/UXF.w5tyBtGmDj2fKu8Ce4v2pYrcjGUJKtK'),
  ('bobby', 'Bobby Bob', '$2a$10$QRZ9dpL2JRGTK7hNG/UXF.w5tyBtGmDj2fKu8Ce4v2pYrcjGUJKtK'),
  ('tommy', 'Tommy Tom', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('robby', 'Robby Rob', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('poppy', 'Poppy Pop', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('benny', 'Benny Ben', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('ronny', 'Ronny Ron', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('danny', 'Danny Dan', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('fanny', 'Fanny Fan', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('hanny', 'Hanny Han', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe'),
  ('jenny', 'Jenny Jen', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe');

INSERT INTO gems (message, latitude, longitude, creator_id, receiver_id, color)
VALUES
  ('hello tommy', 1.2966, 103.7764, 2, 3, 0),
  ('hello tommy', 1.3483, 103.6831, 2, 3, 1),
  ('hello tommy', 1.3644, 103.9915, 2, 3, 2),
  ('hello tommy', 1.4382, 103.7891, 2, 3, 3),
  ('hello tommy', 1.3109, 103.7952, 2, 3, 4),
  ('hello tommy', 1.3309, 103.8752, 2, 3, 5),
  ('hello bobby', 1.2966, 103.7764, 3, 2, 0);

INSERT INTO friend_requests (requester_id, receiver_id)
VALUES
  (3, 2),
  (4, 2),
  (5, 2),
  (6, 2),
  (7, 2),
  (8, 2);
