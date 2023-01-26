// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/chakan"
	"github.com/kanade0404/tenhou-log/services/ent/chii"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
	"github.com/kanade0404/tenhou-log/services/ent/dan"
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/drawn"
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerhandhai"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerpoint"
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
	"github.com/kanade0404/tenhou-log/services/ent/player"
	"github.com/kanade0404/tenhou-log/services/ent/pon"
	"github.com/kanade0404/tenhou-log/services/ent/reach"
	"github.com/kanade0404/tenhou-log/services/ent/room"
	"github.com/kanade0404/tenhou-log/services/ent/round"
	"github.com/kanade0404/tenhou-log/services/ent/schema"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
	"github.com/kanade0404/tenhou-log/services/ent/win"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	callFields := schema.Call{}.Fields()
	_ = callFields
	// callDescID is the schema descriptor for id field.
	callDescID := callFields[0].Descriptor()
	// call.DefaultID holds the default value on creation for the id field.
	call.DefaultID = callDescID.Default.(func() uuid.UUID)
	chakanFields := schema.Chakan{}.Fields()
	_ = chakanFields
	// chakanDescID is the schema descriptor for id field.
	chakanDescID := chakanFields[0].Descriptor()
	// chakan.DefaultID holds the default value on creation for the id field.
	chakan.DefaultID = chakanDescID.Default.(func() uuid.UUID)
	chiiFields := schema.Chii{}.Fields()
	_ = chiiFields
	// chiiDescID is the schema descriptor for id field.
	chiiDescID := chiiFields[0].Descriptor()
	// chii.DefaultID holds the default value on creation for the id field.
	chii.DefaultID = chiiDescID.Default.(func() uuid.UUID)
	compressedmjlogFields := schema.CompressedMJLog{}.Fields()
	_ = compressedmjlogFields
	// compressedmjlogDescInsertedAt is the schema descriptor for inserted_at field.
	compressedmjlogDescInsertedAt := compressedmjlogFields[3].Descriptor()
	// compressedmjlog.DefaultInsertedAt holds the default value on creation for the inserted_at field.
	compressedmjlog.DefaultInsertedAt = compressedmjlogDescInsertedAt.Default.(func() time.Time)
	// compressedmjlogDescID is the schema descriptor for id field.
	compressedmjlogDescID := compressedmjlogFields[0].Descriptor()
	// compressedmjlog.DefaultID holds the default value on creation for the id field.
	compressedmjlog.DefaultID = compressedmjlogDescID.Default.(func() uuid.UUID)
	concealedkanFields := schema.ConcealedKan{}.Fields()
	_ = concealedkanFields
	// concealedkanDescID is the schema descriptor for id field.
	concealedkanDescID := concealedkanFields[0].Descriptor()
	// concealedkan.DefaultID holds the default value on creation for the id field.
	concealedkan.DefaultID = concealedkanDescID.Default.(func() uuid.UUID)
	danFields := schema.Dan{}.Fields()
	_ = danFields
	// danDescID is the schema descriptor for id field.
	danDescID := danFields[0].Descriptor()
	// dan.DefaultID holds the default value on creation for the id field.
	dan.DefaultID = danDescID.Default.(func() uuid.UUID)
	discardFields := schema.Discard{}.Fields()
	_ = discardFields
	// discardDescID is the schema descriptor for id field.
	discardDescID := discardFields[0].Descriptor()
	// discard.DefaultID holds the default value on creation for the id field.
	discard.DefaultID = discardDescID.Default.(func() uuid.UUID)
	drawnFields := schema.Drawn{}.Fields()
	_ = drawnFields
	// drawnDescID is the schema descriptor for id field.
	drawnDescID := drawnFields[0].Descriptor()
	// drawn.DefaultID holds the default value on creation for the id field.
	drawn.DefaultID = drawnDescID.Default.(func() uuid.UUID)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
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
	gameplayerhandhaiFields := schema.GamePlayerHandHai{}.Fields()
	_ = gameplayerhandhaiFields
	// gameplayerhandhaiDescID is the schema descriptor for id field.
	gameplayerhandhaiDescID := gameplayerhandhaiFields[0].Descriptor()
	// gameplayerhandhai.DefaultID holds the default value on creation for the id field.
	gameplayerhandhai.DefaultID = gameplayerhandhaiDescID.Default.(func() uuid.UUID)
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
	meldedkanFields := schema.MeldedKan{}.Fields()
	_ = meldedkanFields
	// meldedkanDescID is the schema descriptor for id field.
	meldedkanDescID := meldedkanFields[0].Descriptor()
	// meldedkan.DefaultID holds the default value on creation for the id field.
	meldedkan.DefaultID = meldedkanDescID.Default.(func() uuid.UUID)
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
	ponFields := schema.Pon{}.Fields()
	_ = ponFields
	// ponDescID is the schema descriptor for id field.
	ponDescID := ponFields[0].Descriptor()
	// pon.DefaultID holds the default value on creation for the id field.
	pon.DefaultID = ponDescID.Default.(func() uuid.UUID)
	reachFields := schema.Reach{}.Fields()
	_ = reachFields
	// reachDescID is the schema descriptor for id field.
	reachDescID := reachFields[0].Descriptor()
	// reach.DefaultID holds the default value on creation for the id field.
	reach.DefaultID = reachDescID.Default.(func() uuid.UUID)
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
	winFields := schema.Win{}.Fields()
	_ = winFields
	// winDescID is the schema descriptor for id field.
	winDescID := winFields[0].Descriptor()
	// win.DefaultID holds the default value on creation for the id field.
	win.DefaultID = winDescID.Default.(func() uuid.UUID)
}