package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
	"yu-croco/ddd_on_golang/app/domain/model/monster"
)

type Hunter struct {
	Id              Id           `json:hunterId`
	Name            Name         `json:hunterName`
	Life            Life         `json:life`
	DefencePower    DefencePower `json:defencePower`
	OffensePower    OffensePower `json:offensePower`
	HuntedMaterials []HuntedMonsterMaterial
	AttackDamage    AttackDamage `json:attackDamage`
}

type Id int
type Name string
type Life int
type DefencePower int
type OffensePower int
type AttackDamage int

func (hunter *Hunter) Attack(monster monster.Monster, damage AttackDamage) (*monster.Monster, helpers.DomainError) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage monster.AttackDamage) (*Hunter, helpers.DomainError) {
	var err helpers.DomainError

	if hunter.Life >= 0 {
		hunter.Life = Life(givenDamage)
	} else {
		err = helpers.NewDomainError("すでに倒しています")
	}

	return hunter, err
}

func (hunter *Hunter) GetMonsterMaterial(monster monster.Monster) (*monster.MonsterMaterial, helpers.DomainError) {
	result, err := monster.TakenMaterial()

	return result, err
}
