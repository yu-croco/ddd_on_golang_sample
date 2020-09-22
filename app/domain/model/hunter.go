package model

import (
	"yu-croco/ddd_on_golang/app/errors"
)

type Hunter struct {
	Id              int    `json:"hunterId"`
	Name            string `json:"hunterName"`
	Life            int    `json:"life"`
	DefencePower    int    `json:"defencePower"`
	OffensePower    int    `json:"offensePower"`
	HuntedMaterials HuntedMonsterMaterials
	AttackDamage    int `json:"attackDamage"`
}

type Hunters []Hunter

func (hunter *Hunter) Attack(monster *Monster, damage int) (*Monster, *errors.AppError) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage int) (*Hunter, *errors.AppError) {
	var err errors.AppError
	diff := hunter.Life - givenDamage

	if hunter.Life == 0 {
		err = errors.NewAppError("ハンターは既に倒れています")
	} else if diff >= 0 {
		hunter.Life = diff
	} else {
		hunter.Life = 0
	}

	return hunter, &err
}

func (hunter *Hunter) GetMonsterMaterial(monster *Monster) (*MonsterMaterial, *errors.AppError) {
	result, err := monster.TakenMaterial()

	return result, &err
}
