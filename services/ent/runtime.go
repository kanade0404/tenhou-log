//go:build tools
// +build tools

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/dan"
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerpoint"
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
	"github.com/kanade0404/tenhou-log/services/ent/player"
	"github.com/kanade0404/tenhou-log/services/ent/room"
	"github.com/kanade0404/tenhou-log/services/ent/round"
	"github.com/kanade0404/tenhou-log/services/ent/schema"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	compressedmjlogFields := schema.CompressedMJLog{}.Fields()
	_ = compressedmjlogFields
	// compressedmjlogDescID is the schema descriptor for id field.
	compressedmjlogDescID := compressedmjlogFields[0].Descriptor()
	// compressedmjlog.DefaultID holds the default value on creation for the id field.
	compressedmjlog.DefaultID = compressedmjlogDescID.Default.(func() uuid.UUID)
	danFields := schema.Dan{}.Fields()
	_ = danFields
	// danDescID is the schema descriptor for id field.
	danDescID := danFields[0].Descriptor()
	// dan.DefaultID holds the default value on creation for the id field.
	dan.DefaultID = danDescID.Default.(func() uuid.UUID)
	gameFields := schema.Game{}.Fields()
	_ = gameFields
	// gameDescID is the schema descriptor for id field.
	gameDescID := gameFields[0].Descriptor()
	// game.DefaultID holds the default value on creation for the id field.
	game.DefaultID = gameDescID.Default.(func() uuid.UUID)
	gameplayerFields := schema.GamePlayer{}.Fields()
	_ = gameplayerFields
	// gameplayerDescStartPosition is the schema descriptor for start_position field.
	gameplayerDescStartPosition := gameplayerFields[2].Descriptor()
	// gameplayer.StartPositionValidator is a validator for the "start_position" field. It is called by the builders before save.
	gameplayer.StartPositionValidator = gameplayerDescStartPosition.Validators[0].(func(string) error)
	// gameplayerDescID is the schema descriptor for id field.
	gameplayerDescID := gameplayerFields[0].Descriptor()
	// gameplayer.DefaultID holds the default value on creation for the id field.
	gameplayer.DefaultID = gameplayerDescID.Default.(func() uuid.UUID)
	gameplayerpointFields := schema.GamePlayerPoint{}.Fields()
	_ = gameplayerpointFields
	// gameplayerpointDescPoint is the schema descriptor for point field.
	gameplayerpointDescPoint := gameplayerpointFields[1].Descriptor()
	// gameplayerpoint.PointValidator is a validator for the "point" field. It is called by the builders before save.
	gameplayerpoint.PointValidator = gameplayerpointDescPoint.Validators[0].(func(uint) error)
	// gameplayerpointDescID is the schema descriptor for id field.
	gameplayerpointDescID := gameplayerpointFields[0].Descriptor()
	// gameplayerpoint.DefaultID holds the default value on creation for the id field.
	gameplayerpoint.DefaultID = gameplayerpointDescID.Default.(func() uuid.UUID)
	handFields := schema.Hand{}.Fields()
	_ = handFields
	// handDescContinuePoint is the schema descriptor for continue_point field.
	handDescContinuePoint := handFields[2].Descriptor()
	// hand.ContinuePointValidator is a validator for the "continue_point" field. It is called by the builders before save.
	hand.ContinuePointValidator = handDescContinuePoint.Validators[0].(func(uint) error)
	// handDescReachPoint is the schema descriptor for reach_point field.
	handDescReachPoint := handFields[3].Descriptor()
	// hand.ReachPointValidator is a validator for the "reach_point" field. It is called by the builders before save.
	hand.ReachPointValidator = handDescReachPoint.Validators[0].(func(uint) error)
	// handDescID is the schema descriptor for id field.
	handDescID := handFields[0].Descriptor()
	// hand.DefaultID holds the default value on creation for the id field.
	hand.DefaultID = handDescID.Default.(func() uuid.UUID)
	mjlogFields := schema.MJLog{}.Fields()
	_ = mjlogFields
	// mjlogDescInsertedAt is the schema descriptor for inserted_at field.
	mjlogDescInsertedAt := mjlogFields[4].Descriptor()
	// mjlog.DefaultInsertedAt holds the default value on creation for the inserted_at field.
	mjlog.DefaultInsertedAt = mjlogDescInsertedAt.Default.(func() time.Time)
	// mjlogDescID is the schema descriptor for id field.
	mjlogDescID := mjlogFields[0].Descriptor()
	// mjlog.DefaultID holds the default value on creation for the id field.
	mjlog.DefaultID = mjlogDescID.Default.(func() uuid.UUID)
	mjlogfileFields := schema.MJLogFile{}.Fields()
	_ = mjlogfileFields
	// mjlogfileDescID is the schema descriptor for id field.
	mjlogfileDescID := mjlogfileFields[0].Descriptor()
	// mjlogfile.DefaultID holds the default value on creation for the id field.
	mjlogfile.DefaultID = mjlogfileDescID.Default.(func() uuid.UUID)
	playerFields := schema.Player{}.Fields()
	_ = playerFields
	// playerDescSex is the schema descriptor for sex field.
	playerDescSex := playerFields[2].Descriptor()
	// player.SexValidator is a validator for the "sex" field. It is called by the builders before save.
	player.SexValidator = playerDescSex.Validators[0].(func(string) error)
	// playerDescID is the schema descriptor for id field.
	playerDescID := playerFields[0].Descriptor()
	// player.DefaultID holds the default value on creation for the id field.
	player.DefaultID = playerDescID.Default.(func() uuid.UUID)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() uuid.UUID)
	roundFields := schema.Round{}.Fields()
	_ = roundFields
	// roundDescWind is the schema descriptor for wind field.
	roundDescWind := roundFields[1].Descriptor()
	// round.WindValidator is a validator for the "wind" field. It is called by the builders before save.
	round.WindValidator = roundDescWind.Validators[0].(func(string) error)
	// roundDescID is the schema descriptor for id field.
	roundDescID := roundFields[0].Descriptor()
	// round.DefaultID holds the default value on creation for the id field.
	round.DefaultID = roundDescID.Default.(func() uuid.UUID)
	turnFields := schema.Turn{}.Fields()
	_ = turnFields
	// turnDescID is the schema descriptor for id field.
	turnDescID := turnFields[0].Descriptor()
	// turn.DefaultID holds the default value on creation for the id field.
	turn.DefaultID = turnDescID.Default.(func() uuid.UUID)
}
