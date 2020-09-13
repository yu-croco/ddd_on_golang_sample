package monster

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
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

func (monster *Monster) Attack(hunter hunter.Hunter, damage AttackDamage) (*hunter.Hunter, helpers.DomainError) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage hunter.AttackDamage) (*Monster, helpers.DomainError) {
	var err helpers.DomainError

	if monster.Life >= 0 {
		monster.Life = Life(givenDamage)
	} else {
		err = helpers.NewDomainError("すでに倒しています")
	}

	return monster, err
}

func (monster *Monster) TakenMaterial() (*MonsterMaterial, helpers.DomainError) {
	if monster.Life != 0 {
		return &MonsterMaterial{}, helpers.NewDomainError("まだ生きてるよ")
	} else {
		return &monster.Materials[0], helpers.DomainError{}
	}
}
