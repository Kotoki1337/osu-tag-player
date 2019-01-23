package objects

import (
	om "danser/bmath"
	"danser/settings"
	"strconv"
	"strings"
	"github.com/go-gl/mathgl/mgl32"
	"danser/render"
	. "danser/osuconst"
)

type BaseObject interface {
	GetBasicData() *basicData
	Update(time int64) bool
	GetPosition() om.Vector2d
	SetDifficulty(preempt, fadeIn float64)
}

type Renderable interface {
	Draw(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) bool
	DrawApproach(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch)
}

type basicData struct {
	StartPos, EndPos   om.Vector2d
	StartTime, EndTime int64
	StackOffset        om.Vector2d
	StackIndex         int64
	Number             int64
	SliderPoint        bool

	sampleSet    int
	additionSet  int
	customIndex  int
	customVolume float64
}

func commonParse(data []string) *basicData {
	x, _ := strconv.ParseFloat(data[0], 64)
	y, _ := strconv.ParseFloat(data[1], 64)
	if settings.VSplayer.Mods.EnableHR {
		y = PLAYFIELD_HEIGHT - y
	}
	time, _ := strconv.ParseInt(data[2], 10, 64)
	return &basicData{StartPos: om.NewVec2d(x, y), StartTime: time, Number: -1}
}

func (bData *basicData) parseExtras(data []string, extraIndex int) {
	if extraIndex < len(data) {
		extras := strings.Split(data[extraIndex], ":")
		sampleSet, _ := strconv.ParseInt(extras[0], 10, 64)
		additionSet, _ := strconv.ParseInt(extras[1], 10, 64)
		index, _ := strconv.ParseInt(extras[2], 10, 64)
		if len(extras) > 3 {
			volume, _ := strconv.ParseInt(extras[3], 10, 64)
			bData.customVolume = float64(volume) / 100.0
		} else {
			bData.customVolume = 0
		}


		bData.sampleSet = int(sampleSet)
		bData.additionSet = int(additionSet)
		bData.customIndex = int(index)
	}
}
