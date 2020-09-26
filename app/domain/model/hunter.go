package model

import (
	"yu-croco/ddd_on_golang/app/errors"
)

type Hunter struct {
	Id              HunterId           `json:"hunterId"`
	Name            HunterName         `json:"hunterName"`
	Life            HunterLife         `json:"life"`
	DefencePower    HunterDefencePower `json:"defencePower"`
	OffensePower    HunterOffensePower `json:"offensePower"`
	HuntedMaterials HuntedMonsterMaterials
	AttackDamage    HunterAttackDamage `json:"attackDamage"`
}

type HunterId int
type HunterName string
type HunterLife int
type HunterDefencePower int
type HunterOffensePower int
type HunterAttackDamage int

type Hunters []Hunter

func (hunter *Hunter) Attack(monster *Monster, damage HunterOffensePower) (*Monster, *errors.AppError) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage MonsterOffensePower) (*Hunter, *errors.AppError) {
	var err errors.AppError
	diff := int(hunter.Life) - int(givenDamage)

	if hunter.Life == 0 {
		err = errors.NewAppError("ハンターは既に倒れています")
	} else if diff >= 0 {
		hunter.Life = HunterLife(diff)
	} else {
		hunter.Life = HunterLife(0)
	}

	return hunter, &err
}

func (hunter *Hunter) GetMonsterMaterial(monster *Monster) (*MonsterMaterial, *errors.AppError) {
	result, err := monster.TakenMaterial()

	return result, &err
}
