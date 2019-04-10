package objects

import (
	"danser/audio"
	"danser/bmath"
	. "danser/osuconst"
	"danser/render"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"strconv"
)

type Spinner struct {
	objData *basicData
	pos     bmath.Vector2d
	Timings *Timings
	sample  int
	renderStartTime int64
}

func NewSpinner(data []string, number int64) *Spinner {
	spinner := &Spinner{}
	spinner.objData = commonParse(data, number)
	endtime, _ := strconv.ParseInt(data[5], 10, 64)
	spinner.objData.EndTime = int64(endtime)
	spinner.pos = bmath.Vector2d{PLAYFIELD_WIDTH / 2,PLAYFIELD_HEIGHT / 2}

	sample, _ := strconv.ParseInt(data[4], 10, 64)
	spinner.sample = int(sample)

	spinner.renderStartTime = -12345
	return spinner
}

func (self Spinner) GetBasicData() *basicData {
	return self.objData
}

func (self *Spinner) SetTiming(timings *Timings) {
	self.Timings = timings
}

func (self *Spinner) GetPosition() bmath.Vector2d {
	return self.pos
}

func (self *Spinner) Update(time int64) bool {
	if time < self.objData.EndTime {
		return false
	}

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

func (self *Spinner) Draw(time int64, preempt float64, fadeIn float64, color mgl32.Vec4, batch *render.SpriteBatch) bool {
	if self.renderStartTime == -12345 {
		self.renderStartTime = time
	}

	alpha := 1.0

	var angle float64

	// 1秒之内rpm从0到300
	if time - self.renderStartTime <= 1000 {
		rpm := float64(time - self.renderStartTime)*0.3
		angle = float64(time - self.renderStartTime) * (rpm * math.Pi / 30000)
	}else {
		angle = float64(time - self.renderStartTime - 1000) * math.Pi / 100
	}

	if time < self.renderStartTime - int64(preempt) {
		return false
	} else if time < self.renderStartTime {
		alpha = float64(color[3]) / preempt
	}else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	batch.SetColor(1, 1, 1, alpha)
	// 绘制Spinner转圈
	batch.DrawUnitSR(*render.SpinnerCircle, bmath.Vector2d{float64(render.SpinnerCircle.Width) / 4.75, float64(render.SpinnerCircle.Height) / 4.75}, angle)
	batch.DrawUnitSR(*render.SpinnerMiddle, bmath.Vector2d{float64(render.SpinnerMiddle.Width) / 2, float64(render.SpinnerMiddle.Height) / 2}, angle)
	batch.DrawUnitSR(*render.SpinnerBottom, bmath.Vector2d{float64(render.SpinnerBottom.Width) / 4, float64(render.SpinnerBottom.Height) / 4}, angle)

	batch.SetSubScale(1, 1)

	if time >= self.objData.EndTime+int64(preempt/4) {
		return true
	}
	return false
}

func (self *Spinner) SetDifficulty(preempt, fadeIn float64) {

}

func (self *Spinner) DrawApproach(time int64, preempt float64, fadeIn float64, color mgl32.Vec4, batch *render.SpriteBatch) {
	// 记录第一次渲染转盘的时间，第一次渲染时，转盘正好撑满整个屏幕，随后逐渐变小
	if self.renderStartTime == -12345 {
		self.renderStartTime = time
	}

	alpha := 1.0
	// 计算AR
	fake_preempt := 2 * float64(self.objData.EndTime - self.renderStartTime) / PLAYFIELD_HEIGHT
	arr := float64(self.objData.EndTime - time) / fake_preempt

	if time < self.renderStartTime - int64(preempt){
		alpha = 0
	} else if time < self.renderStartTime{
		alpha = float64(color[3]) / preempt
	}else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	if time <= self.objData.EndTime {
		batch.SetColor(1, 1, 1, alpha)
		batch.DrawUnitS(*render.SpinnerApproachCircle, bmath.Vector2d{arr, arr})
	}
}

func (self *Spinner) GetObjectNumber() int64 {
	return self.objData.ObjectNumber
}