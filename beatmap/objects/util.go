package objects

import (
	"danser/bmath"
	"danser/render"
	"strconv"
)

func GetObject(data []string, number int64) (BaseObject, int64) {
	objType, _ := strconv.ParseInt(data[3], 10, 64)
	newnumber := number
	if (objType & CIRCLE) > 0 {
		if objType == NEWCIRCLE {
			// 新的combo
			newnumber = 1
		}else {
			// 继续combo
			newnumber += 1
		}
		return NewCircle(data, newnumber), newnumber
	} else if (objType & SLIDER) > 0 {
		if objType == NEWSLIDER {
			// 新的combo
			newnumber = 1
		}else {
			// 继续combo
			newnumber += 1
		}
		sl := NewSlider(data, newnumber)
		if sl == nil {
			return nil, newnumber
		} else {
			return sl, newnumber
		}
	} else if (objType & SPINNNER) > 0 {
		if objType == NEWSPINNNER{
			// 新的combo
			newnumber = 1
		}else {
			// 继续combo
			newnumber += 1
		}
		return NewSpinner(data, newnumber), newnumber
	}
	return nil, newnumber
}

// 绘制圈内数字
func DrawHitCircleNumber(number int64, position bmath.Vector2d, batch *render.SpriteBatch) {
	switch number {
	case 0:
		batch.DrawUnitN(*render.Circle0, position)
		break
	case 1:
		batch.DrawUnitN(*render.Circle1, position)
		break
	case 2:
		batch.DrawUnitN(*render.Circle2, position)
		break
	case 3:
		batch.DrawUnitN(*render.Circle3, position)
		break
	case 4:
		batch.DrawUnitN(*render.Circle4, position)
		break
	case 5:
		batch.DrawUnitN(*render.Circle5, position)
		break
	case 6:
		batch.DrawUnitN(*render.Circle6, position)
		break
	case 7:
		batch.DrawUnitN(*render.Circle7, position)
		break
	case 8:
		batch.DrawUnitN(*render.Circle8, position)
		break
	case 9:
		batch.DrawUnitN(*render.Circle9, position)
		break
	}
}

// 获取圈内数字宽度
func GetHitCircleNumberWidth(number int64) int32{
	switch number {
	case 0:
		return render.Circle0.Width
	case 1:
		return render.Circle0.Width
	case 2:
		return render.Circle0.Width
	case 3:
		return render.Circle0.Width
	case 4:
		return render.Circle0.Width
	case 5:
		return render.Circle0.Width
	case 6:
		return render.Circle0.Width
	case 7:
		return render.Circle0.Width
	case 8:
		return render.Circle0.Width
	case 9:
		return render.Circle0.Width
	}
	return 0
}

const (
	CIRCLE int64 = 1
	SLIDER int64 = 2
	SPINNNER int64 = 8
	NEWCIRCLE int64 = 5
	NEWSLIDER int64 = 6
	NEWSPINNNER int64 = 12
)
