package hunter

import (
	"yu-croco/ddd_on_golang/app/domain/helpers"
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/infrastructure"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

func AttackMonsterUseCase(hunterId int, monsterId int) (*model.Monster, *helpers.DomainError) {
	db := infrastructure.GetDB()
	hunter := repositoryImpl.FindHunterById(db, hunterId)
	monster := repositoryImpl.FindMonsterBy(db, monsterId)

	hunterAttackDamage := service.CalculateAttackMonsterDamage(hunter, monster)
	damagedMonster, attackErr := hunter.Attack(monster, hunterAttackDamage)
	if attackErr.HasErrors() {
		return nil, &attackErr
	}

	updatedMonster := repositoryImpl.UpdateMonster(db, damagedMonster)

	return updatedMonster, nil
}
