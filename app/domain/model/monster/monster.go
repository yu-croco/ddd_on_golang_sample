package monster

import (
	"errors"
	"yu-croco/ddd_on_golang/app/domain/model/hunter"
)

type Monster struct {
	Id           Id           `json:monsterId`
	Name         Name         `json:monsterName`
	Life         Life         `json:life`
	DefencePower DefencePower `json:defencePower`
	OffensePower OffensePower `json:offensePower`
	Materials    []MonsterMaterial
	AttackDamage AttackDamage `json:attackDamage`
}

type Id int
type Name string
type Life int
type DefencePower int
type OffensePower int
type AttackDamage int

func (monster *Monster) Attack(hunter hunter.Hunter, damage AttackDamage) (*hunter.Hunter, error) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage hunter.AttackDamage) (*Monster, error) {
	var err error

	if monster.Life >= 0 {
		monster.Life = Life(givenDamage)
	} else {
		err = errors.New("すでに倒しています")
	}

	return monster, err
}

func (monster *Monster) TakenMaterial() (*MonsterMaterial, error) {
	if monster.Life != 0 {
		return &MonsterMaterial{}, errors.New("まだ生きてるよ")
	}
	return &monster.Materials[0], nil
}
