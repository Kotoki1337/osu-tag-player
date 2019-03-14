package objects

import (
	"danser/audio"
	"danser/bmath"
	"danser/render"
	"danser/settings"
	. "danser/osuconst"
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
)

type Circle struct {
	objData *basicData
	sample  int
	Timings *Timings
}

func NewCircle(data []string, number int64) *Circle {
	circle := &Circle{}
	circle.objData = commonParse(data)
	f, _ := strconv.ParseInt(data[4], 10, 64)
	circle.sample = int(f)
	circle.objData.EndTime = circle.objData.StartTime
	circle.objData.EndPos = circle.objData.StartPos
	circle.objData.parseExtras(data, 5)
	circle.objData.Number = number
	return circle
}

func DummyCircle(pos bmath.Vector2d, time int64) *Circle {
	return DummyCircleInherit(pos, time, false)
}

func DummyCircleInherit(pos bmath.Vector2d, time int64, inherit bool) *Circle {
	circle := &Circle{objData: &basicData{}}
	circle.objData.StartPos = pos
	circle.objData.EndPos = pos
	circle.objData.StartTime = time
	circle.objData.EndTime = time
	circle.objData.EndPos = circle.objData.StartPos
	circle.objData.SliderPoint = inherit
	return circle
}

func (self Circle) GetBasicData() *basicData {
	return self.objData
}

func (self *Circle) Update(time int64) bool {

	index := self.objData.customIndex

	if index == 0 {
		index = self.Timings.Current.SampleIndex
	}

	if self.objData.sampleSet == 0 {
		audio.PlaySample(self.Timings.Current.SampleSet, self.objData.additionSet, self.sample, index, self.Timings.Current.SampleVolume)
	} else {
		audio.PlaySample(self.objData.sampleSet, self.objData.additionSet, self.sample, index, self.Timings.Current.SampleVolume)
	}

	return true
}

func (self *Circle) SetTiming(timings *Timings) {
	self.Timings = timings
}

func (self *Circle) GetPosition() bmath.Vector2d {
	return self.objData.StartPos
}

func (self *Circle) Draw(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) bool {

	alpha := 1.0

	if settings.VSplayer.Mods.EnableHD {
		fadein := preempt * FADE_IN_DURATION_MULTIPLIER
		fadeoutstarttime := float64(self.objData.StartTime) - preempt + fadein
		fadeoutduration := preempt * FADE_OUT_DURATION_MULTIPLIER
		if time < self.objData.StartTime - int64(fadein) {
			alpha = (float64(time)- fadeoutstarttime) / fadein
		} else if time >= self.objData.StartTime {
			alpha = 0.0
		} else {
			alpha = float64(color[3]) * float64(self.objData.EndTime-time) / fadeoutduration
		}
	}else {
		if time < self.objData.StartTime-int64(preempt) {
			alpha = float64(time-(self.objData.StartTime-int64(preempt))) / preempt
		} else if time >= self.objData.StartTime {
			alpha = 1.0 - float64(time-self.objData.StartTime)/(preempt/2)
		} else {
			alpha = float64(color[3])
		}
	}

	batch.SetTranslation(self.objData.StartPos)

	if time >= self.objData.StartTime {
		batch.SetSubScale(1+(1.0-alpha)*0.5, 1+(1.0-alpha)*0.5)
	}

	if settings.DIVIDES >= settings.Objects.MandalaTexturesTrigger {
		alpha *= settings.Objects.MandalaTexturesAlpha
	}

	batch.SetColor(float64(color[0]), float64(color[1]), float64(color[2]), alpha)
	if settings.DIVIDES >= settings.Objects.MandalaTexturesTrigger {
		batch.DrawUnit(*render.CircleFull)
	} else {
		batch.DrawUnit(*render.Circle)
	}

	if settings.DIVIDES < settings.Objects.MandalaTexturesTrigger {
		batch.SetColor(1, 1, 1, alpha)
		batch.DrawUnit(*render.CircleOverlay)
	}

	batch.SetSubScale(1, 1)

	if time >= self.objData.StartTime+int64(preempt/2) {
		return true
	}
	return false
}

func (self *Circle) SetDifficulty(preempt, fadeIn float64) {

}

func (self *Circle) DrawApproach(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) {

	alpha := 1.0
	arr := float64(self.objData.StartTime-time) / preempt

	if time < self.objData.StartTime-int64(preempt)/2 {
		alpha = float64(time-(self.objData.StartTime-int64(preempt))) / (preempt / 2)
	} else if time >= self.objData.StartTime {
		alpha = 1.0 - float64(time-self.objData.StartTime)/(preempt/2)
	} else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	if settings.Objects.DrawApproachCircles && time <= self.objData.StartTime {
		batch.SetColor(float64(color[0]), float64(color[1]), float64(color[2]), alpha)
		batch.SetSubScale(1.0+arr*2, 1.0+arr*2)
		batch.DrawUnit(*render.ApproachCircle)
	}

	batch.SetSubScale(1, 1)
}
