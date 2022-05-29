package init

import (
	"fmt"
	"strconv"
	"strings"
)

type Round string

const (
	East    Round = "東"
	South   Round = "南"
	West    Round = "西"
	North   Round = "北"
	UNKNOWN Round = "UNKNOWN"
)

func NewRound(round uint) Round {
	switch round {
	case 0:
		return East
	case 1:
		return South
	case 2:
		return West
	case 3:
		return North
	default:
		return UNKNOWN
	}
}
func Unmarshal(round Round) (string, error) {
	switch round {
	case East:
		return "0", nil
	case South:
		return "1", nil
	case West:
		return "2", nil
	case North:
		return "3", nil
	default:
		return "", fmt.Errorf("unexpected value :%s", round)

	}
}

type Hand uint

const (
	First Hand = iota
	Second
	Third
	Fourth
)

type Honba uint

func (h Honba) String() string {
	return fmt.Sprintf("%d本場", h)
}

type RiichPoints uint

type DiceValue uint

const (
	one DiceValue = iota + 1
	two
	three
	four
	five
	six
)

func (d DiceValue) UInt() uint {
	return uint(d)
}

func NewDiceValue(diceValue uint) (DiceValue, error) {
	var isDiceValue bool
	var list = []DiceValue{one, two, three, four, five, six}
	for _, value := range list {
		if value.UInt()-1 == diceValue {
			isDiceValue = true
			break
		}
	}
	if isDiceValue {
		return DiceValue(diceValue + 1), nil
	} else {
		return 0, fmt.Errorf("unexpected value %d", diceValue)
	}
}

type Seed struct {
	Round
	Hand
	Honba
	RiichPoints
	DiceValue1 DiceValue
	DiceValue2 DiceValue
	_          struct{}
}

func (s *Seed) Unmarshal(seed string) error {
	seeds := strings.Split(seed, ",")
	var intseeds []uint
	for _, s := range seeds {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		intseeds = append(intseeds, uint(i))
	}
	diceValue1, err := NewDiceValue(intseeds[4])
	if err != nil {
		return err
	}
	diceValue2, err := NewDiceValue(intseeds[5])
	if err != nil {
		return err
	}
	s.Round = NewRound(intseeds[0])
	s.Hand = Hand(intseeds[1])
	s.Honba = Honba(intseeds[2])
	s.RiichPoints = RiichPoints(intseeds[3])
	s.DiceValue1 = diceValue1
	s.DiceValue2 = diceValue2
	return nil
}
func (s Seed) Marshal() (seed string, err error) {
	round, err := Unmarshal(s.Round)
	if err != nil {
		return "", err
	}
	hand := strconv.Itoa(int(s.Hand))
	honba := strconv.Itoa(int(s.Honba))
	riichPoints := strconv.Itoa(int(s.RiichPoints))
	diceValue1 := strconv.Itoa(int(s.DiceValue1) - 1)
	diceValue2 := strconv.Itoa(int(s.DiceValue2) - 1)
	return strings.Join([]string{round, hand, honba, riichPoints, diceValue1, diceValue2}, ","), nil
}
