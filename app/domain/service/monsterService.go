package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateAttackHunterDamage(monster *model.Monster, hunter *model.Hunter) model.MonsterOffensePower {
	if int(hunter.DefencePower) >= int(monster.OffensePower) {
		return monster.OffensePower
	} else {
		return model.MonsterOffensePower(int(monster.OffensePower)*2 - int(hunter.DefencePower))
	}
}
