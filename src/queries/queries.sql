SELECT * FROM get_user_populated(2);

DELETE FROM songs;
DELETE FROM albums_bags;
delete FROM albums;


CALL create_user(
'u3132',
'carlos',
'carlos@gmail.com',
'https://cdn1.epicgames.com/salesEvent/salesEvent/EGS_KIDAMNESIAEXHIBITION_namethemachinexAGP_S1_2560x1440-01dfad2110f87a6348843484091a2741',
'1283812'
);

CALL insert_album(
    '123dsakd',
    'In Rainbow',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);

CALL insert_album(
    'asdadkasdjk',
    'KID A',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);

CALL insert_album(
    '123124dasdk',
    'KID A',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);

CALL insert_album(
    'asdjaksd',
    'KID A',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);

CALL insert_album(
    '12312nda',
    'KID A',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);

CALL insert_album(
    '1231axcmczkxc',
    'KID A',
    'https://upload.wikimedia.org/wikipedia/en/1/14/Inrainbowscover.png',
    109,
    74,
    50,
    1
);


CALL insert_song(
    'asdj3enfa',
    'Idioteque',
    2
)

SELECT * FROM albums;
DELETE from songs;
SELECT * FROM songs;

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


SELECT * FROM get_user(2);
SELECT * FROM get_user_populated(3);

SELECT * FROM users;


DELETE FROM users;
DROP DATABASE wimb;