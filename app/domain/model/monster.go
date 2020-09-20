package model

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
)

type Monster struct {
	Id           int    `json:monsterId`
	Name         string `json:monsterName`
	Life         int    `json:life`
	DefencePower int    `json:defencePower`
	OffensePower int    `json:offensePower`
	Materials    []MonsterMaterial
	AttackDamage int `json:attackDamage`
}

func (monster *Monster) Attack(hunter Hunter, damage int) (*Hunter, helpers.DomainError) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage int) (*Monster, helpers.DomainError) {
	var err helpers.DomainError
	diff := monster.Life - givenDamage

	if monster.Life == 0 {
		err = helpers.NewDomainError("このモンスターはすでに倒しています")
	} else if diff >= 0 {
		monster.Life = diff
	} else {
		monster.Life = 0
	}

	return monster, err
}

func (monster *Monster) TakenMaterial() (*MonsterMaterial, helpers.DomainError) {
	if monster.Life != 0 {
		return &MonsterMaterial{}, helpers.NewDomainError("このモンスターはまだ生きてるので剥ぎ取れません")
	} else {
		return &monster.Materials[0], helpers.DomainError{}
	}
}
