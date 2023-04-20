CREATE TABLE IF NOT EXISTS videos
(
    id     uuid         NOT NULL,
    name   VARCHAR(250) NOT NULL,
    url    VARCHAR(250) NOT NULL UNIQUE,
    length VARCHAR(9)   NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS annotations
(
    id         uuid         NOT NULL,
    video_id   uuid         NOT NULL,
    type       VARCHAR(100) NOT NULL,
    notes      TEXT,
    start_time VARCHAR(9),
    end_time   VARCHAR(9),
    PRIMARY KEY (id),
    UNIQUE (video_id, start_time, end_time),

    CONSTRAINT FK_annotation_video FOREIGN KEY (video_id)
        REFERENCES videos (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users
(
    id uuid NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(10) NOT NULL,
    PRIMARY KEY (id)
);
