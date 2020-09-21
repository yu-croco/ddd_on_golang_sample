package model

import (
	"yu-croco/ddd_on_golang/app/errors"
)

type Monster struct {
	Id           MonsterId `json:monsterId`
	Name         string    `json:monsterName`
	Life         int       `json:life`
	DefencePower int       `json:defencePower`
	OffensePower int       `json:offensePower`
	Materials    []MonsterMaterial
	AttackDamage int `json:attackDamage`
}

type MonsterId int

func (id MonsterId) ToInt() int {
	return int(id)
}

func NewMonsterId(value int) (MonsterId, *errors.AppError) {
	var monsterId MonsterId
	var err errors.AppError

	if value <= 0 {
		err = errors.NewAppError("monsterIdは1以上の数字で指定してください")
	}
	monsterId = MonsterId(value)

	return monsterId, &err
}

func (monster *Monster) Attack(hunter *Hunter, damage int) (*Hunter, *errors.AppError) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage int) (*Monster, *errors.AppError) {
	var err errors.AppError
	diff := monster.Life - givenDamage

	if monster.Life == 0 {
		err = errors.NewAppError("このモンスターはすでに倒しています")
	} else if diff >= 0 {
		monster.Life = diff
	} else {
		monster.Life = 0
	}

	return monster, &err
}

func (monster *Monster) TakenMaterial() (*MonsterMaterial, errors.AppError) {
	if monster.Life != 0 {
		return nil, errors.NewAppError("このモンスターはまだ生きてるので剥ぎ取れません")
	} else {
		return &monster.Materials[0], errors.AppError{}
	}
}
