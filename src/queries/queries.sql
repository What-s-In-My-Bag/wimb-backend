SELECT * FROM get_user_populated(2);

DELETE FROM songs;
DELETE FROM albums_bags;
delete FROM albums;
SELECT * FROM albums;
DELETE from songs;
SELECT * FROM songs;
SELECT * FROM users;
SELECT * FROM bags;

SELECT * FROM get_user(2);
SELECT * FROM get_user_populated('u3132');
SELECT * FROM get_bag_populated(1);



-- Insert TEST USER

SELECT create_user(
'u3132',
'carlos',
'carlos@gmail.com',
'https://cdn1.epicgames.com/salesEvent/salesEvent/EGS_KIDAMNESIAEXHIBITION_namethemachinexAGP_S1_2560x1440-01dfad2110f87a6348843484091a2741',
'1283812'
);

SELECT insert_album(
    '123dsakd',
    'In Rainbows',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    300,
    300,
    1
);


SELECT insert_album(
    'kida123',
    'kID A',
    'https://upload.wikimedia.org/wikipedia/en/0/02/Radioheadkida.png',
    124,
    123,
    123,    
    300,
    300,
    1
);

SELECT insert_album(
    'okcomp123',
    'Ok Computer',
    'https://upload.wikimedia.org/wikipedia/en/b/ba/Radioheadokcomputer.png',
    200,
    220,
    228,    
    300,
    300,
    1
);


CALL insert_song(
    'wf123',
    'Weird Fishes',
    1
);

CALL insert_song(
    'ain',
    'All I need',
    1
);

CALL insert_song(
    'jfip',
    'Jigsaw Falling Into Place',
    1
);

CALL insert_song(
    'inlimbo123',
    'In Limbo',
    2
);


CALL insert_song(
    'htdc123',
    'How to disappear completely',
    2
);

CALL insert_song(
    'optimistic123',
    'Optimistic',
    2
);

CALL insert_song(
    'cbutw123',
    'Climbing Up The Walls',
    3
);

CALL insert_song(
    'pa123',
    'Paranoid Android',
    3
);

CALL insert_song(
    'ns133',
    'No Surprises',
    3
);

SELECT * FROM insert_song(
    'fh133',
    'Fittier Happier',
    3
);


-- END OF INSERTING TEST USER

CALL update_bag(
    3,
    '#12833',
    FALSE
);

CALL update_latest_bag(
    3,
    '#12834',
    FALSE
);

CALL delete_bag(
    2
);

CALL delete_latest_bag(
    3
);

select * from bags;



SELECT * FROM users;


DELETE FROM users;
DROP DATABASE wimb;