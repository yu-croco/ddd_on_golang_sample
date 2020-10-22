package monster

import (
	"yu-croco/ddd_on_golang/app/domain/model"
	"yu-croco/ddd_on_golang/app/domain/repository"
	"yu-croco/ddd_on_golang/app/domain/service"
	"yu-croco/ddd_on_golang/app/errors"
)

type attackHunterUseCaseImpl struct {
	HunterId          model.HunterId
	MonsterId         model.MonsterId
	HunterRepository  repository.HunterRepository
	MonsterRepository repository.MonsterRepository
}

type attackHunterUseCase interface {
	Exec() (*model.Hunter, *errors.AppError)
}

func NewAttackHunterUseCaseImpl(hunterId model.HunterId, monsterId model.MonsterId,
	hunterRepository repository.HunterRepository,
	monsterRepository repository.MonsterRepository) attackHunterUseCase {

	return attackHunterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl attackHunterUseCaseImpl) Exec() (*model.Hunter, *errors.AppError) {
	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	monsterAttackDamage := service.CalculateAttackHunterDamage(monster, hunter)
	damagedHunter, attackErr := monster.Attack(hunter, monsterAttackDamage)
	if attackErr.HasErrors() {
		return nil, attackErr
	}
	updatedHunter, updateErr := impl.HunterRepository.Update(damagedHunter)
	if updateErr.HasErrors() {
		return nil, updateErr
	}
	return updatedHunter, nil
}
