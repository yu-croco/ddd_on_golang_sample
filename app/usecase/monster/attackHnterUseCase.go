package monster

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/errors"
	"yu-croco/ddd_on_golang/app/infrastructure/repositoryImpl"
)

var hunterRepository = repositoryImpl.NewHunterRepositoryImpl()
var monsterRepository = repositoryImpl.NewMonsterRepositoryImpl()

func AttackHunterUseCase(monsterId model.MonsterId, hunterId model.HunterId) (*model.Hunter, *errors.AppError) {
	hunter, hunterFindErr := hunterRepository.FindById(hunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := monsterRepository.FindById(monsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	monsterAttackDamage := service.CalculateAttackHunterDamage(monster, hunter)
	damagedHunter, attackErr := monster.Attack(hunter, monsterAttackDamage)
	if attackErr.HasErrors() {
		return nil, attackErr
	}
	updatedHunter, updateErr := hunterRepository.Update(damagedHunter)
	if updateErr.HasErrors() {
		return nil, updateErr
	}
	return updatedHunter, nil
}
