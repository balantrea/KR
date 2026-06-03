CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

    CREATE TABLE sports (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    sport_id INT REFERENCES sports(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date DATE NOT NULL,
    team_id INT REFERENCES teams(id) ON DELETE SET NULL,
    position VARCHAR(50) NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    home_team_id INT REFERENCES teams(id) ON DELETE CASCADE,
    away_team_id INT REFERENCES teams(id) ON DELETE CASCADE,
    match_date TIMESTAMP NOT NULL,
    home_score INT DEFAULT 0,
    away_score INT DEFAULT 0,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE player_archive (
    id INT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    deleted_at TIMESTAMP,
    user_id INT
);

CREATE VIEW vw_player_profiles AS
SELECT p.id, p.first_name, p.last_name, p.position, t.name AS team_name, s.name AS sport_name, p.user_id
FROM players p
LEFT JOIN teams t ON p.team_id = t.id
LEFT JOIN sports s ON t.sport_id = s.id;

CREATE VIEW vw_match_results AS
SELECT m.id, m.match_date, t1.name AS home_team, t2.name AS away_team, m.home_score, m.away_score, m.user_id
FROM matches m
JOIN teams t1 ON m.home_team_id = t1.id
JOIN teams t2 ON m.away_team_id = t2.id;

CREATE VIEW vw_team_roster_count AS
SELECT t.id, t.name, COUNT(p.id) AS total_players, t.user_id
FROM teams t
LEFT JOIN players p ON t.id = p.team_id
GROUP BY t.id, t.name, t.user_id;

CREATE FUNCTION fn_calculate_age(bdate DATE)
RETURNS INT AS $$
BEGIN
    RETURN EXTRACT(YEAR FROM AGE(bdate));
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION fn_get_total_matches(t_id INT, u_id INT)
RETURNS INT AS $$
DECLARE
    total INT;
BEGIN
    SELECT COUNT(*) INTO total FROM matches WHERE (home_team_id = t_id OR away_team_id = t_id) AND user_id = u_id;
    RETURN total;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION fn_get_sport_team_count(s_id INT, u_id INT)
RETURNS INT AS $$
DECLARE
    total INT;
BEGIN
    SELECT COUNT(*) INTO total FROM teams WHERE sport_id = s_id AND user_id = u_id;
    RETURN total;
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE sp_insert_player(f_name VARCHAR, l_name VARCHAR, b_date DATE, t_id INT, pos VARCHAR, u_id INT)
AS $$
BEGIN
    INSERT INTO players (first_name, last_name, birth_date, team_id, position, user_id)
    VALUES (f_name, l_name, b_date, t_id, pos, u_id);
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE sp_update_player_team(p_id INT, new_t_id INT, u_id INT)
AS $$
BEGIN
    UPDATE players SET team_id = new_t_id WHERE id = p_id AND user_id = u_id;
END;
$$ LANGUAGE plpgsql;

CREATE PROCEDURE sp_delete_player(p_id INT, u_id INT)
AS $$
BEGIN
    DELETE FROM players WHERE id = p_id AND user_id = u_id;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION fn_tg_check_teams()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.home_team_id = NEW.away_team_id THEN
        RAISE EXCEPTION 'Error: Same teams';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tg_check_teams
BEFORE INSERT OR UPDATE ON matches
FOR EACH ROW EXECUTE FUNCTION fn_tg_check_teams();

CREATE FUNCTION fn_tg_archive_player()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO player_archive (id, first_name, last_name, deleted_at, user_id)
    VALUES (OLD.id, OLD.first_name, OLD.last_name, NOW(), OLD.user_id);
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tg_archive_player
AFTER DELETE ON players
FOR EACH ROW EXECUTE FUNCTION fn_tg_archive_player();

CREATE FUNCTION fn_tg_check_scores()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.home_score < 0 OR NEW.away_score < 0 THEN
        RAISE EXCEPTION 'Error: Negative score';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER tg_check_scores
BEFORE INSERT OR UPDATE ON matches
FOR EACH ROW EXECUTE FUNCTION fn_tg_check_scores();