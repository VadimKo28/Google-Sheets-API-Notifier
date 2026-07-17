CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  event_date TEXT NOT NULL, 
  name TEXT NOT NULL,
  complete BOOLEAN NOT NULL DEFAULT false
);
  
CREATE UNIQUE INDEX IF NOT EXISTS uniq_events_event_date 
ON events (event_date);