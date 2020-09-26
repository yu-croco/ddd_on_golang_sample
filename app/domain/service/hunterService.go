package service

import "yu-croco/ddd_on_golang/app/domain/model"

func CalculateAttackMonsterDamage(hunter *model.Hunter, monster *model.Monster) model.HunterOffensePower {
	if int(monster.DefencePower) >= int(hunter.OffensePower) {
		return hunter.OffensePower
	} else {
		return model.HunterOffensePower(int(hunter.OffensePower)*2 - int(monster.DefencePower))
	}
}
