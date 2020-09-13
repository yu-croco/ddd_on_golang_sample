package hunter

import (
	"errors"
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

func (hunter *Hunter) Attack(monster monster.Monster, damage AttackDamage) (*monster.Monster, error) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage monster.AttackDamage) (*Hunter, error) {
	var err error

	if hunter.Life >= 0 {
		hunter.Life = Life(givenDamage)
	} else {
		err = errors.New("すでに倒しています")
	}

	return hunter, err
}

func (hunter *Hunter) GetMonsterMaterial(monster monster.Monster) (*monster.MonsterMaterial, error) {
	result, err := monster.TakenMaterial()
	if err != nil {
		return result, err
	}
	return result, nil
}
