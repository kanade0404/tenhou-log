// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompressedMjLogsColumns holds the columns for the "compressed_mj_logs" table.
	CompressedMjLogsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "size", Type: field.TypeUint},
	}
	// CompressedMjLogsTable holds the schema information for the "compressed_mj_logs" table.
	CompressedMjLogsTable = &schema.Table{
		Name:       "compressed_mj_logs",
		Columns:    CompressedMjLogsColumns,
		PrimaryKey: []*schema.Column{CompressedMjLogsColumns[0]},
	}
	// DansColumns holds the columns for the "dans" table.
	DansColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// DansTable holds the schema information for the "dans" table.
	DansTable = &schema.Table{
		Name:       "dans",
		Columns:    DansColumns,
		PrimaryKey: []*schema.Column{DansColumns[0]},
	}
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "room_games", Type: field.TypeUUID, Nullable: true},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_rooms_games",
				Columns:    []*schema.Column{GamesColumns[3]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GamePlayersColumns holds the columns for the "game_players" table.
	GamePlayersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "rate", Type: field.TypeFloat64},
		{Name: "start_position", Type: field.TypeString},
		{Name: "dan_game_players", Type: field.TypeUUID, Nullable: true},
		{Name: "player_game_players", Type: field.TypeUUID, Nullable: true},
	}
	// GamePlayersTable holds the schema information for the "game_players" table.
	GamePlayersTable = &schema.Table{
		Name:       "game_players",
		Columns:    GamePlayersColumns,
		PrimaryKey: []*schema.Column{GamePlayersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_players_dans_game_players",
				Columns:    []*schema.Column{GamePlayersColumns[3]},
				RefColumns: []*schema.Column{DansColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "game_players_players_game_players",
				Columns:    []*schema.Column{GamePlayersColumns[4]},
				RefColumns: []*schema.Column{PlayersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GamePlayerHandHaisColumns holds the columns for the "game_player_hand_hais" table.
	GamePlayerHandHaisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GamePlayerHandHaisTable holds the schema information for the "game_player_hand_hais" table.
	GamePlayerHandHaisTable = &schema.Table{
		Name:       "game_player_hand_hais",
		Columns:    GamePlayerHandHaisColumns,
		PrimaryKey: []*schema.Column{GamePlayerHandHaisColumns[0]},
	}
	// GamePlayerPointsColumns holds the columns for the "game_player_points" table.
	GamePlayerPointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GamePlayerPointsTable holds the schema information for the "game_player_points" table.
	GamePlayerPointsTable = &schema.Table{
		Name:       "game_player_points",
		Columns:    GamePlayerPointsColumns,
		PrimaryKey: []*schema.Column{GamePlayerPointsColumns[0]},
	}
	// GoAroundsColumns holds the columns for the "go_arounds" table.
	GoAroundsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
	}
	// GoAroundsTable holds the schema information for the "go_arounds" table.
	GoAroundsTable = &schema.Table{
		Name:       "go_arounds",
		Columns:    GoAroundsColumns,
		PrimaryKey: []*schema.Column{GoAroundsColumns[0]},
	}
	// HandsColumns holds the columns for the "hands" table.
	HandsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
	}
	// HandsTable holds the schema information for the "hands" table.
	HandsTable = &schema.Table{
		Name:       "hands",
		Columns:    HandsColumns,
		PrimaryKey: []*schema.Column{HandsColumns[0]},
	}
	// MjLogsColumns holds the columns for the "mj_logs" table.
	MjLogsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "version", Type: field.TypeFloat64},
		{Name: "seed", Type: field.TypeString},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "inserted_at", Type: field.TypeTime},
		{Name: "game_mjlogs", Type: field.TypeUUID, Unique: true},
		{Name: "mj_log_file_mjlogs", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// MjLogsTable holds the schema information for the "mj_logs" table.
	MjLogsTable = &schema.Table{
		Name:       "mj_logs",
		Columns:    MjLogsColumns,
		PrimaryKey: []*schema.Column{MjLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mj_logs_games_mjlogs",
				Columns:    []*schema.Column{MjLogsColumns[5]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "mj_logs_mj_log_files_mjlogs",
				Columns:    []*schema.Column{MjLogsColumns[6]},
				RefColumns: []*schema.Column{MjLogFilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MjLogFilesColumns holds the columns for the "mj_log_files" table.
	MjLogFilesColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "compressed_mj_log_mjlog_files", Type: field.TypeUUID, Unique: true},
	}
	// MjLogFilesTable holds the schema information for the "mj_log_files" table.
	MjLogFilesTable = &schema.Table{
		Name:       "mj_log_files",
		Columns:    MjLogFilesColumns,
		PrimaryKey: []*schema.Column{MjLogFilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "mj_log_files_compressed_mj_logs_mjlog_files",
				Columns:    []*schema.Column{MjLogFilesColumns[2]},
				RefColumns: []*schema.Column{CompressedMjLogsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PlayersColumns holds the columns for the "players" table.
	PlayersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "sex", Type: field.TypeString},
	}
	// PlayersTable holds the schema information for the "players" table.
	PlayersTable = &schema.Table{
		Name:       "players",
		Columns:    PlayersColumns,
		PrimaryKey: []*schema.Column{PlayersColumns[0]},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
	}
	// RoundsColumns holds the columns for the "rounds" table.
	RoundsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "wind_rounds", Type: field.TypeUUID, Nullable: true},
	}
	// RoundsTable holds the schema information for the "rounds" table.
	RoundsTable = &schema.Table{
		Name:       "rounds",
		Columns:    RoundsColumns,
		PrimaryKey: []*schema.Column{RoundsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rounds_winds_rounds",
				Columns:    []*schema.Column{RoundsColumns[1]},
				RefColumns: []*schema.Column{WindsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WindsColumns holds the columns for the "winds" table.
	WindsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// WindsTable holds the schema information for the "winds" table.
	WindsTable = &schema.Table{
		Name:       "winds",
		Columns:    WindsColumns,
		PrimaryKey: []*schema.Column{WindsColumns[0]},
	}
	// GameGamePlayersColumns holds the columns for the "game_game_players" table.
	GameGamePlayersColumns = []*schema.Column{
		{Name: "game_id", Type: field.TypeUUID},
		{Name: "game_player_id", Type: field.TypeUUID},
	}
	// GameGamePlayersTable holds the schema information for the "game_game_players" table.
	GameGamePlayersTable = &schema.Table{
		Name:       "game_game_players",
		Columns:    GameGamePlayersColumns,
		PrimaryKey: []*schema.Column{GameGamePlayersColumns[0], GameGamePlayersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_game_players_game_id",
				Columns:    []*schema.Column{GameGamePlayersColumns[0]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "game_game_players_game_player_id",
				Columns:    []*schema.Column{GameGamePlayersColumns[1]},
				RefColumns: []*schema.Column{GamePlayersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompressedMjLogsTable,
		DansTable,
		GamesTable,
		GamePlayersTable,
		GamePlayerHandHaisTable,
		GamePlayerPointsTable,
		GoAroundsTable,
		HandsTable,
		MjLogsTable,
		MjLogFilesTable,
		PlayersTable,
		RoomsTable,
		RoundsTable,
		WindsTable,
		GameGamePlayersTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = RoomsTable
	GamePlayersTable.ForeignKeys[0].RefTable = DansTable
	GamePlayersTable.ForeignKeys[1].RefTable = PlayersTable
	MjLogsTable.ForeignKeys[0].RefTable = GamesTable
	MjLogsTable.ForeignKeys[1].RefTable = MjLogFilesTable
	MjLogFilesTable.ForeignKeys[0].RefTable = CompressedMjLogsTable
	RoundsTable.ForeignKeys[0].RefTable = WindsTable
	GameGamePlayersTable.ForeignKeys[0].RefTable = GamesTable
	GameGamePlayersTable.ForeignKeys[1].RefTable = GamePlayersTable
}
