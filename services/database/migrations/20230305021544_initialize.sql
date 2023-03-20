-- create "rooms" table
CREATE TABLE "rooms" ("id" uuid NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "rooms_name_key" to table: "rooms"
CREATE UNIQUE INDEX "rooms_name_key" ON "rooms" ("name");
-- create "games" table
CREATE TABLE "games" ("id" uuid NOT NULL, "name" character varying NOT NULL, "started_at" timestamptz NOT NULL, "room_games" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "games_rooms_games" FOREIGN KEY ("room_games") REFERENCES "rooms" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "games_name_key" to table: "games"
CREATE UNIQUE INDEX "games_name_key" ON "games" ("name");
-- create "rounds" table
CREATE TABLE "rounds" ("id" uuid NOT NULL, "wind" character varying NOT NULL, "game_rounds" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "rounds_games_rounds" FOREIGN KEY ("game_rounds") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "hands" table
CREATE TABLE "hands" ("id" uuid NOT NULL, "num" bigint NOT NULL, "continue_point" bigint NOT NULL, "reach_point" bigint NOT NULL, "round_hands" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "hands_rounds_hands" FOREIGN KEY ("round_hands") REFERENCES "rounds" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "compressed_mj_logs" table
CREATE TABLE "compressed_mj_logs" ("id" uuid NOT NULL, "name" character varying NOT NULL, "size" bigint NOT NULL, "inserted_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create "mj_log_files" table
CREATE TABLE "mj_log_files" ("id" uuid NOT NULL, "name" character varying NOT NULL, "compressed_mj_log_mjlog_files" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "mj_log_files_compressed_mj_logs_mjlog_files" FOREIGN KEY ("compressed_mj_log_mjlog_files") REFERENCES "compressed_mj_logs" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "mj_log_files_compressed_mj_log_mjlog_files_key" to table: "mj_log_files"
CREATE UNIQUE INDEX "mj_log_files_compressed_mj_log_mjlog_files_key" ON "mj_log_files" ("compressed_mj_log_mjlog_files");
-- create index "mj_log_files_name_key" to table: "mj_log_files"
CREATE UNIQUE INDEX "mj_log_files_name_key" ON "mj_log_files" ("name");
-- create "mj_logs" table
CREATE TABLE "mj_logs" ("id" uuid NOT NULL, "version" double precision NOT NULL, "seed" character varying NOT NULL, "started_at" timestamptz NOT NULL, "inserted_at" timestamptz NOT NULL, "game_mjlogs" uuid NOT NULL, "mj_log_file_mjlogs" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "mj_logs_games_mjlogs" FOREIGN KEY ("game_mjlogs") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "mj_logs_mj_log_files_mjlogs" FOREIGN KEY ("mj_log_file_mjlogs") REFERENCES "mj_log_files" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create index "mj_logs_game_mjlogs_key" to table: "mj_logs"
CREATE UNIQUE INDEX "mj_logs_game_mjlogs_key" ON "mj_logs" ("game_mjlogs");
-- create index "mj_logs_mj_log_file_mjlogs_key" to table: "mj_logs"
CREATE UNIQUE INDEX "mj_logs_mj_log_file_mjlogs_key" ON "mj_logs" ("mj_log_file_mjlogs");
-- create "turns" table
CREATE TABLE "turns" ("id" uuid NOT NULL, "num" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "events" table
CREATE TABLE "events" ("id" uuid NOT NULL, "turn_event" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "events_turns_event" FOREIGN KEY ("turn_event") REFERENCES "turns" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "events_turn_event_key" to table: "events"
CREATE UNIQUE INDEX "events_turn_event_key" ON "events" ("turn_event");
-- create "discards" table
CREATE TABLE "discards" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "reaches" table
CREATE TABLE "reaches" ("id" uuid NOT NULL, "discard_reach" uuid NOT NULL, "event_reach" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "reaches_discards_reach" FOREIGN KEY ("discard_reach") REFERENCES "discards" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "reaches_events_reach" FOREIGN KEY ("event_reach") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "reaches_discard_reach_key" to table: "reaches"
CREATE UNIQUE INDEX "reaches_discard_reach_key" ON "reaches" ("discard_reach");
-- create "hand_turns" table
CREATE TABLE "hand_turns" ("hand_id" uuid NOT NULL, "turn_id" uuid NOT NULL, PRIMARY KEY ("hand_id", "turn_id"), CONSTRAINT "hand_turns_hand_id" FOREIGN KEY ("hand_id") REFERENCES "hands" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "hand_turns_turn_id" FOREIGN KEY ("turn_id") REFERENCES "turns" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- create "drawns" table
CREATE TABLE "drawns" ("id" uuid NOT NULL, "discard_draw" uuid NOT NULL, "event_draw" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "drawns_discards_draw" FOREIGN KEY ("discard_draw") REFERENCES "discards" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "drawns_events_draw" FOREIGN KEY ("event_draw") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "drawns_discard_draw_key" to table: "drawns"
CREATE UNIQUE INDEX "drawns_discard_draw_key" ON "drawns" ("discard_draw");
-- create "wins" table
CREATE TABLE "wins" ("id" uuid NOT NULL, "event_win" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "wins_events_win" FOREIGN KEY ("event_win") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create "dans" table
CREATE TABLE "dans" ("id" uuid NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "dans_name_key" to table: "dans"
CREATE UNIQUE INDEX "dans_name_key" ON "dans" ("name");
-- create "players" table
CREATE TABLE "players" ("id" uuid NOT NULL, "name" character varying NOT NULL, "sex" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "players_name_key" to table: "players"
CREATE UNIQUE INDEX "players_name_key" ON "players" ("name");
-- create "game_players" table
CREATE TABLE "game_players" ("id" uuid NOT NULL, "rate" double precision NOT NULL, "start_position" character varying NOT NULL, "dan_game_players" uuid NULL, "player_game_players" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "game_players_dans_game_players" FOREIGN KEY ("dan_game_players") REFERENCES "dans" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "game_players_players_game_players" FOREIGN KEY ("player_game_players") REFERENCES "players" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- create "game_game_players" table
CREATE TABLE "game_game_players" ("game_id" uuid NOT NULL, "game_player_id" uuid NOT NULL, PRIMARY KEY ("game_id", "game_player_id"), CONSTRAINT "game_game_players_game_id" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "game_game_players_game_player_id" FOREIGN KEY ("game_player_id") REFERENCES "game_players" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- create "chakans" table
CREATE TABLE "chakans" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "chiis" table
CREATE TABLE "chiis" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "concealed_kans" table
CREATE TABLE "concealed_kans" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "melded_kans" table
CREATE TABLE "melded_kans" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "pons" table
CREATE TABLE "pons" ("id" uuid NOT NULL, PRIMARY KEY ("id"));
-- create "calls" table
CREATE TABLE "calls" ("id" uuid NOT NULL, "chakan_call" uuid NOT NULL, "chii_call" uuid NOT NULL, "concealed_kan_call" uuid NOT NULL, "discard_call" uuid NOT NULL, "event_call" uuid NOT NULL, "melded_kan_call" uuid NOT NULL, "pon_call" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "calls_chakans_call" FOREIGN KEY ("chakan_call") REFERENCES "chakans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_chiis_call" FOREIGN KEY ("chii_call") REFERENCES "chiis" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_concealed_kans_call" FOREIGN KEY ("concealed_kan_call") REFERENCES "concealed_kans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_discards_call" FOREIGN KEY ("discard_call") REFERENCES "discards" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_events_call" FOREIGN KEY ("event_call") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_melded_kans_call" FOREIGN KEY ("melded_kan_call") REFERENCES "melded_kans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "calls_pons_call" FOREIGN KEY ("pon_call") REFERENCES "pons" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "calls_chakan_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_chakan_call_key" ON "calls" ("chakan_call");
-- create index "calls_chii_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_chii_call_key" ON "calls" ("chii_call");
-- create index "calls_concealed_kan_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_concealed_kan_call_key" ON "calls" ("concealed_kan_call");
-- create index "calls_discard_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_discard_call_key" ON "calls" ("discard_call");
-- create index "calls_melded_kan_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_melded_kan_call_key" ON "calls" ("melded_kan_call");
-- create index "calls_pon_call_key" to table: "calls"
CREATE UNIQUE INDEX "calls_pon_call_key" ON "calls" ("pon_call");
-- create "game_player_hand_hais" table
CREATE TABLE "game_player_hand_hais" ("id" uuid NOT NULL, "hais" jsonb NOT NULL, "turn_gameplayerhandhai" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "game_player_hand_hais_turns_gameplayerhandhai" FOREIGN KEY ("turn_gameplayerhandhai") REFERENCES "turns" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "game_player_hand_hais_turn_gameplayerhandhai_key" to table: "game_player_hand_hais"
CREATE UNIQUE INDEX "game_player_hand_hais_turn_gameplayerhandhai_key" ON "game_player_hand_hais" ("turn_gameplayerhandhai");
-- create "game_player_points" table
CREATE TABLE "game_player_points" ("id" uuid NOT NULL, "point" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "turn_game_player_points" table
CREATE TABLE "turn_game_player_points" ("turn_id" uuid NOT NULL, "game_player_point_id" uuid NOT NULL, PRIMARY KEY ("turn_id", "game_player_point_id"), CONSTRAINT "turn_game_player_points_game_player_point_id" FOREIGN KEY ("game_player_point_id") REFERENCES "game_player_points" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "turn_game_player_points_turn_id" FOREIGN KEY ("turn_id") REFERENCES "turns" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
