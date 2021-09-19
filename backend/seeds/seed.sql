-- Passwords are 'password'
INSERT INTO users (username, email, hashed_password)
VALUES
  ('bobby', 'bobby@example.com', '$2a$10$QRZ9dpL2JRGTK7hNG/UXF.w5tyBtGmDj2fKu8Ce4v2pYrcjGUJKtK'),
  ('tommy', 'tommy@example.com', '$2a$10$j6Lj7bgLoS1nAk5304kZbu.cQLYtYhf0/Zueu.3wi5lKt/NmF/cBe');

INSERT INTO gems (message, latitude, longitude, creator_id, receiver_id)
VALUES
  ('hello tommy', 1.2966, 103.7764, 1, 2),
  ('hello tommy', 1.3483, 103.6831, 1, 2),
  ('hello tommy', 1.3644, 103.9915, 1, 2),
  ('hello tommy', 1.4382, 103.7891, 1, 2),
  ('hello tommy', 1.3109, 103.7952, 1, 2),
  ('hello tommy', 1.3309, 103.8752, 1, 2);
