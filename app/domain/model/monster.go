package model

import (
	"yu-croco/ddd_on_golang/app/errors"
)

type Monster struct {
	Id           MonsterId           `json:"monsterId"`
	Name         MonsterName         `json:"monsterName"`
	Life         MonsterLife         `json:"life"`
	DefencePower MonsterDefencePower `json:"defencePower"`
	OffensePower MonsterOffensePower `json:"offensePower"`
	Materials    MonsterMaterials
	AttackDamage MonsterAttackDamage `json:"attackDamage"`
}

type MonsterId int
type MonsterName string
type MonsterLife int
type MonsterDefencePower int
type MonsterOffensePower int
type MonsterAttackDamage int

type Monsters []Monster

func (monster *Monster) Attack(hunter *Hunter, damage MonsterOffensePower) (*Hunter, *errors.AppError) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage HunterOffensePower) (*Monster, *errors.AppError) {
	var err errors.AppError
	diff := int(monster.Life) - int(givenDamage)

	if monster.Life == 0 {
		err = errors.NewAppError("このモンスターはすでに倒しています")
	} else if diff >= 0 {
		monster.Life = MonsterLife(diff)
	} else {
		monster.Life = MonsterLife(0)
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
