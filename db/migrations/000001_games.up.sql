CREATE TABLE IF NOT EXISTS games (
    id serial PRIMARY KEY,
    match_start TIMESTAMP WITH TIME ZONE NOT NULL ,
    cancelation_deadline TIMESTAMP WITH TIME ZONE NOT NULL,
    request_rsvp TIMESTAMP WITH TIME ZONE NOT NULL,
    minimal_attendees INTEGER DEFAULT 8,
    place VARCHAR (255) NOT NULL,
    is_canceled BOOLEAN
);
