// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CallsColumns holds the columns for the "calls" table.
	CallsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CallsTable holds the schema information for the "calls" table.
	CallsTable = &schema.Table{
		Name:       "calls",
		Columns:    CallsColumns,
		PrimaryKey: []*schema.Column{CallsColumns[0]},
	}
	// ChakansColumns holds the columns for the "chakans" table.
	ChakansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ChakansTable holds the schema information for the "chakans" table.
	ChakansTable = &schema.Table{
		Name:       "chakans",
		Columns:    ChakansColumns,
		PrimaryKey: []*schema.Column{ChakansColumns[0]},
	}
	// ChiisColumns holds the columns for the "chiis" table.
	ChiisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ChiisTable holds the schema information for the "chiis" table.
	ChiisTable = &schema.Table{
		Name:       "chiis",
		Columns:    ChiisColumns,
		PrimaryKey: []*schema.Column{ChiisColumns[0]},
	}
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
	// ConcealedKansColumns holds the columns for the "concealed_kans" table.
	ConcealedKansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ConcealedKansTable holds the schema information for the "concealed_kans" table.
	ConcealedKansTable = &schema.Table{
		Name:       "concealed_kans",
		Columns:    ConcealedKansColumns,
		PrimaryKey: []*schema.Column{ConcealedKansColumns[0]},
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
	// DrawnsColumns holds the columns for the "drawns" table.
	DrawnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DrawnsTable holds the schema information for the "drawns" table.
	DrawnsTable = &schema.Table{
		Name:       "drawns",
		Columns:    DrawnsColumns,
		PrimaryKey: []*schema.Column{DrawnsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
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
		{Name: "oid", Type: field.TypeUUID},
		{Name: "point", Type: field.TypeUint},
	}
	// GamePlayerPointsTable holds the schema information for the "game_player_points" table.
	GamePlayerPointsTable = &schema.Table{
		Name:       "game_player_points",
		Columns:    GamePlayerPointsColumns,
		PrimaryKey: []*schema.Column{GamePlayerPointsColumns[0]},
	}
	// HandsColumns holds the columns for the "hands" table.
	HandsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "num", Type: field.TypeUint},
		{Name: "continue_point", Type: field.TypeUint},
		{Name: "reach_point", Type: field.TypeUint},
		{Name: "round_hands", Type: field.TypeUUID, Nullable: true},
	}
	// HandsTable holds the schema information for the "hands" table.
	HandsTable = &schema.Table{
		Name:       "hands",
		Columns:    HandsColumns,
		PrimaryKey: []*schema.Column{HandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hands_rounds_hands",
				Columns:    []*schema.Column{HandsColumns[4]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
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
	// MeldedKansColumns holds the columns for the "melded_kans" table.
	MeldedKansColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// MeldedKansTable holds the schema information for the "melded_kans" table.
	MeldedKansTable = &schema.Table{
		Name:       "melded_kans",
		Columns:    MeldedKansColumns,
		PrimaryKey: []*schema.Column{MeldedKansColumns[0]},
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
	// PonsColumns holds the columns for the "pons" table.
	PonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PonsTable holds the schema information for the "pons" table.
	PonsTable = &schema.Table{
		Name:       "pons",
		Columns:    PonsColumns,
		PrimaryKey: []*schema.Column{PonsColumns[0]},
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
		{Name: "wind", Type: field.TypeString},
		{Name: "game_rounds", Type: field.TypeUUID, Nullable: true},
	}
	// RoundsTable holds the schema information for the "rounds" table.
	RoundsTable = &schema.Table{
		Name:       "rounds",
		Columns:    RoundsColumns,
		PrimaryKey: []*schema.Column{RoundsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rounds_games_rounds",
				Columns:    []*schema.Column{RoundsColumns[2]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TurnsColumns holds the columns for the "turns" table.
	TurnsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "num", Type: field.TypeUint},
	}
	// TurnsTable holds the schema information for the "turns" table.
	TurnsTable = &schema.Table{
		Name:       "turns",
		Columns:    TurnsColumns,
		PrimaryKey: []*schema.Column{TurnsColumns[0]},
	}
	// WinsColumns holds the columns for the "wins" table.
	WinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// WinsTable holds the schema information for the "wins" table.
	WinsTable = &schema.Table{
		Name:       "wins",
		Columns:    WinsColumns,
		PrimaryKey: []*schema.Column{WinsColumns[0]},
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
	// HandTurnsColumns holds the columns for the "hand_turns" table.
	HandTurnsColumns = []*schema.Column{
		{Name: "hand_id", Type: field.TypeUUID},
		{Name: "turn_id", Type: field.TypeUUID},
	}
	// HandTurnsTable holds the schema information for the "hand_turns" table.
	HandTurnsTable = &schema.Table{
		Name:       "hand_turns",
		Columns:    HandTurnsColumns,
		PrimaryKey: []*schema.Column{HandTurnsColumns[0], HandTurnsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hand_turns_hand_id",
				Columns:    []*schema.Column{HandTurnsColumns[0]},
				RefColumns: []*schema.Column{HandsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "hand_turns_turn_id",
				Columns:    []*schema.Column{HandTurnsColumns[1]},
				RefColumns: []*schema.Column{TurnsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TurnGamePlayerPointsColumns holds the columns for the "turn_game_player_points" table.
	TurnGamePlayerPointsColumns = []*schema.Column{
		{Name: "turn_id", Type: field.TypeUUID},
		{Name: "game_player_point_id", Type: field.TypeUUID},
	}
	// TurnGamePlayerPointsTable holds the schema information for the "turn_game_player_points" table.
	TurnGamePlayerPointsTable = &schema.Table{
		Name:       "turn_game_player_points",
		Columns:    TurnGamePlayerPointsColumns,
		PrimaryKey: []*schema.Column{TurnGamePlayerPointsColumns[0], TurnGamePlayerPointsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "turn_game_player_points_turn_id",
				Columns:    []*schema.Column{TurnGamePlayerPointsColumns[0]},
				RefColumns: []*schema.Column{TurnsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "turn_game_player_points_game_player_point_id",
				Columns:    []*schema.Column{TurnGamePlayerPointsColumns[1]},
				RefColumns: []*schema.Column{GamePlayerPointsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CallsTable,
		ChakansTable,
		ChiisTable,
		CompressedMjLogsTable,
		ConcealedKansTable,
		DansTable,
		DrawnsTable,
		EventsTable,
		GamesTable,
		GamePlayersTable,
		GamePlayerHandHaisTable,
		GamePlayerPointsTable,
		HandsTable,
		MjLogsTable,
		MjLogFilesTable,
		MeldedKansTable,
		PlayersTable,
		PonsTable,
		RoomsTable,
		RoundsTable,
		TurnsTable,
		WinsTable,
		GameGamePlayersTable,
		HandTurnsTable,
		TurnGamePlayerPointsTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = RoomsTable
	GamePlayersTable.ForeignKeys[0].RefTable = DansTable
	GamePlayersTable.ForeignKeys[1].RefTable = PlayersTable
	HandsTable.ForeignKeys[0].RefTable = RoundsTable
	MjLogsTable.ForeignKeys[0].RefTable = GamesTable
	MjLogsTable.ForeignKeys[1].RefTable = MjLogFilesTable
	MjLogFilesTable.ForeignKeys[0].RefTable = CompressedMjLogsTable
	RoundsTable.ForeignKeys[0].RefTable = GamesTable
	GameGamePlayersTable.ForeignKeys[0].RefTable = GamesTable
	GameGamePlayersTable.ForeignKeys[1].RefTable = GamePlayersTable
	HandTurnsTable.ForeignKeys[0].RefTable = HandsTable
	HandTurnsTable.ForeignKeys[1].RefTable = TurnsTable
	TurnGamePlayerPointsTable.ForeignKeys[0].RefTable = TurnsTable
	TurnGamePlayerPointsTable.ForeignKeys[1].RefTable = GamePlayerPointsTable
}
