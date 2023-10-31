use admin;

DROP DATABASE wimb;
CREATE DATABASE wimb WITH OWNER = admin;

DROP TYPE IF EXISTS user_bag_album;

CREATE TYPE user_bag_album AS (
    id INT,
    uuid CHAR(16) ,
    username VARCHAR(25),
    profile_img VARCHAR(250),
    shirt_color VARCHAR(15),
    show_album_names BOOLEAN,
    album_spotify_id VARCHAR(20),
    name VARCHAR(40),
    cover VARCHAR(250),
    r_avg INT,
    g_avg INT,
    b_avg INT,
    width INT,
    height iNT,
    song_spotify_id VARCHAR(20),
    song_name VARCHAR(300)
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY  NOT NULL,
    uuid CHAR(16) UNIQUE NOT NULL,
    username VARCHAR(25) NOT NULL,
    email VARCHAR(40) NOT NULL,
    profile_img VARCHAR(250),
    spotify_id VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bags(
    id SERIAL PRIMARY KEY NOT NULL,
    shirt_color VARCHAR(15) DEFAULT 'black' NOT NULL,
    show_album_names BOOLEAN DEFAULT TRUE NOT NULL,
    user_id INT REFERENCES users (id)  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE albums(
    id serial primary key not null,
    spotify_id varchar(20) unique  not null,
    name VARCHAR(40) NOT NULL,
    cover VARCHAR(250) NOT NULL,
    width INT NOT NULL,
    height INT NOT NULL,
    r_avg INT DEFAULT 0 NOT NULL,
    g_avg INT DEFAULT 0 NOT NULL,
    b_avg INT DEFAULT 0 NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY NOT NULL,
    spotify_id varchar(20) unique  NOT NULL,
    name VARCHAR(300) NOT NULL,
    album_id INT REFERENCES albums (id) NOT NULL

);



CREATE TABLE albums_bags(
  bag_id INT REFERENCES bags(id), 
  album_id integer REFERENCES albums(id),
  PRIMARY KEY (bag_id, album_id)
);


DROP FUNCTION IF EXISTS check_email;

CREATE FUNCTION check_email(p_email VARCHAR(40))
RETURNS BOOLEAN AS $$
    BEGIN
        RETURN EXISTS (SELECT 1 FROM users WHERE email = p_email);
    END

$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS check_uuid;

CREATE FUNCTION check_uuid(p_uuid CHAR(4))
RETURNS BOOLEAN AS $$

    BEGIN
        RETURN EXISTS (SELECT 1 FROM users WHERE uuid = p_uuid);
    END

$$ LANGUAGE plpgsql;


DROP FUNCTION IF EXISTS get_bag_by_user_id;

CREATE  FUNCTION get_bag_by_user_id(_user_id INTEGER)
RETURNS INTEGER AS $$
DECLARE
    bag_id INT;
BEGIN
        SELECT id INTO bag_id FROM bags
        WHERE user_id = _user_id ORDER BY id
        DESC LIMIT 1; 

        RETURN bag_id;  
END

$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS get_user();

CREATE FUNCTION get_user(
    user_id INT
)
RETURNS SETOF users AS $$

    BEGIN
        RETURN QUERY SELECT *  FROM users WHERE id =  user_id;
    END

$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS get_user_populated;
 
CREATE FUNCTION get_user_populated(
    p_user_uuid CHAR(16)
)
RETURNS SETOF user_bag_album AS $$
    BEGIN
        RETURN QUERY SELECT 
            u.id,
            u.uuid,
            u.username, 
            u.profile_img, 
            b.shirt_color, 
            b.show_album_names,
            a.spotify_id AS album_spotify_id,
            a.name,
            a.cover,
            a.r_avg,
            a.g_avg,
            a.b_avg,
            a.width,
            a.height,
            s.spotify_id AS song_spotify_id,
            s.name AS song_name
          FROM users u
          JOIN (
            SELECT
                id,
                shirt_color,
                show_album_names,
                user_id
            FROM bags
            WHERE (bags.user_id, bags.id) IN (
                SELECT bags.user_id, MAX(id)
                FROM bags GROUP BY bags.user_id
            )

          ) b ON u.id = b.user_id
            LEFT JOIN albums_bags ba ON b.id = ba.bag_id                
            LEFT JOIN albums a ON ba.album_id = a.id
            LEFT JOIN songs s ON a.id = s.album_id
            WHERE u.uuid = p_user_uuid; 

    END

$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS get_bag_populated(_bag_id INT) ;

CREATE FUNCTION get_bag_populated(_bag_id INT) 
RETURNS SETOF user_bag_album AS $$
BEGIN
    RETURN QUERY SELECT 
            u.id,
            u.uuid,
            u.username, 
            u.profile_img, 
            b.shirt_color, 
            b.show_album_names,
            a.spotify_id AS album_spotify_id,
            a.name,
            a.cover,
            a.r_avg,
            a.g_avg,
            a.b_avg,
            a.width,
            a.height,
            s.spotify_id AS song_spotify_id,
            s.name AS song_name
        FROM bags b 
        JOIN users u ON b.user_id = u.id
        LEFT JOIN albums_bags ba ON b.id = ba.bag_id                
        LEFT JOIN albums a ON ba.album_id = a.id
        LEFT JOIN songs s ON a.id = s.album_id
        WHERE b.id = _bag_id;
END
$$ LANGUAGE plpgsql;

DROP PROCEDURE IF EXISTS create_user;

CREATE FUNCTION create_user(
    p_uuid CHAR(25),
    p_username VARCHAR(25), 
    p_email VARCHAR(40), 
    p_profile_img VARCHAR(250) , 
    p_spotify_id VARCHAR(20)
)
RETURNS INT AS $$
    DECLARE
        user_id INTEGER;
        bag_id INTEGER;
    BEGIN
    
        IF check_uuid(p_uuid) THEN
            RAISE EXCEPTION 'Uuid % already exist', p_uuid;
        ELSEIF check_email(p_email) THEN
            RAISE EXCEPTION 'email % already exist', p_email;
        END IF;

        INSERT INTO users (uuid, username, email, profile_img, spotify_id) 
        VALUES (p_uuid, p_username, p_email, p_profile_img, p_spotify_id) RETURNING id into user_id;

        INSERT INTO bags (user_id) VALUES (user_id) RETURNING id INTO bag_id;

        RETURN bag_id;

    END
$$ LANGUAGE plpgsql;

DROP PROCEDURE IF EXISTS insert_album;

CREATE FUNCTION insert_album(
    p_spotify_id VARCHAR(20),
    p_name VARCHAR(40),
    p_cover VARCHAR(250),
    p_r_avg INT,
    p_g_avg INT,
    p_b_avg INT,
    p_width INT,
    p_height INT,
    p_user_id INT
) 
RETURNS INTEGER AS $$
    DECLARE 
        _bag_id INTEGER;
        _album_id INTEGER;
    BEGIN

        SELECT * INTO _bag_id FROM get_bag_by_user_id(p_user_id) ;

        IF _bag_id IS NULL THEN
            RAISE EXCEPTION 'Invalid User';
        END IF;

        SELECT id INTO _album_id FROM albums WHERE spotify_id = p_spotify_id;

        IF _album_id IS NULL THEN

            INSERT INTO albums (spotify_id, name, cover, r_avg, g_avg, b_avg, width, height) 
            VALUES (p_spotify_id, p_name, p_cover, p_r_avg, p_g_avg, p_b_avg, p_width, p_height)
            RETURNING id INTO _album_id;

        ELSE
            RAISE NOTICE 'album already exist';
        END IF;


        IF NOT EXISTS (SELECT 1 FROM albums_bags WHERE album_id = _album_id AND bag_id = _bag_id) THEN

            INSERT INTO albums_bags (bag_id , album_id) VALUES (_bag_id, _album_id);
        ELSE 
            RAISE NOTICE 'this album has been already assigned to bag with id: %', _bag_id;
        END IF;

        RETURN _album_id;
    END
$$ LANGUAGE plpgsql ;


DROP PROCEDURE IF EXISTS insert_song;

CREATE OR REPLACE PROCEDURE insert_song(
    _spotify_id VARCHAR(20),
    _name VARCHAR(300),
    _album_id INTEGER

) LANGUAGE plpgsql AS $$
    BEGIN

        IF NOT EXISTS (SELECT 1 FROM albums WHERE id = _album_id ) THEN
            RAISE EXCEPTION 'The album id is invalid ';
        END IF;

        IF NOT EXISTS (SELECT 1 FROM songs WHERE album_id = _album_id AND spotify_id = _spotify_id) THEN

            INSERT INTO songs (spotify_id, album_id, name) VALUES (_spotify_id, _album_id, _name);

        ELSE
            RAISE NOTICE 'this song has been already created %', _spotify_id;

        END IF;
    END
$$;


DROP PROCEDURE IF EXISTS update_bag;

CREATE OR REPLACE PROCEDURE update_bag(
    p_bag_id INTEGER,
    p_shirt_color VARCHAR(8),
    p_show_album_names BOOLEAN
) 
LANGUAGE plpgsql AS $$
    BEGIN
        IF p_bag_id IS NULL THEN
            RAISE EXCEPTION 'Invalid User';
        END IF;

        UPDATE bags SET 
        shirt_color = p_shirt_color, show_album_names = p_show_album_names
        WHERE id = p_bag_id;

    END
$$;


DROP PROCEDURE IF EXISTS update_latest_bag;

CREATE OR REPLACE PROCEDURE update_latest_bag(
    p_user_id INTEGER,
    p_shirt_color VARCHAR(8),
    p_show_album_names BOOLEAN
) 
LANGUAGE plpgsql AS $$
    DECLARE
        _bag_id INTEGER;
    BEGIN

        SELECT * FROM get_bag_by_user_id(p_user_id) INTO _bag_id;
        CALL update_bag(_bag_id, p_shirt_color, p_show_album_names);

    END
$$;

DROP PROCEDURE IF EXISTS delete_bag;

CREATE OR REPLACE PROCEDURE delete_bag(
    _bag_id INTEGER
) 
LANGUAGE plpgsql AS $$
    BEGIN

        DELETE from albums_bags WHERE bag_id = _bag_id;
        DELETE from bags WHERE id = _bag_id;

    END
$$;

DROP PROCEDURE IF EXISTS delete_latest_bag;

CREATE OR REPLACE PROCEDURE delete_latest_bag(
    user_id INTEGER
) 
LANGUAGE plpgsql AS $$
    DECLARE
        _bag_id INTEGER;
    BEGIN

        SELECT get_bag_by_user_id(user_id) INTO _bag_id;
        DELETE from albums_bags WHERE bag_id = _bag_id;
        DELETE from bags WHERE id = _bag_id;

    END
$$;

DROP FUNCTION IF EXISTS check_album_exists;

CREATE FUNCTION check_album_exists(
    _spotify_id VARCHAR(20)
) RETURNS BOOLEAN AS $$
    BEGIN
        RETURN EXISTS (SELECT 1 FROM albums WHERE spotify_id = _spotify_id);
    END
$$ LANGUAGE plpgsql;

