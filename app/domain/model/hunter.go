package model

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
)

type Hunter struct {
	Id              int    `json:hunterId`
	Name            string `json:hunterName`
	Life            int    `json:life`
	DefencePower    int    `json:defencePower`
	OffensePower    int    `json:offensePower`
	HuntedMaterials []HuntedMonsterMaterial
	AttackDamage    int `json:attackDamage`
}

func (hunter *Hunter) Attack(monster *Monster, damage int) (*Monster, helpers.DomainError) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage int) (*Hunter, helpers.DomainError) {
	var err helpers.DomainError
	diff := hunter.Life - givenDamage

	if hunter.Life == 0 {
		err = helpers.NewDomainError("ハンターは既に倒れています")
	} else if diff >= 0 {
		hunter.Life = diff
	} else {
		hunter.Life = 0
	}

	return hunter, err
}

func (hunter *Hunter) GetMonsterMaterial(monster Monster) (*MonsterMaterial, helpers.DomainError) {
	result, err := monster.TakenMaterial()

	return result, err
}
