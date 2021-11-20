package monster

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repository2 "yu-croco/ddd_on_golang/pkg/domain/repository"
	service2 "yu-croco/ddd_on_golang/pkg/domain/service"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type attackHunterUseCaseImpl struct {
	HunterId          model2.HunterId
	MonsterId         model2.MonsterId
	HunterRepository  repository2.HunterRepository
	MonsterRepository repository2.MonsterRepository
}

type attackHunterUseCase interface {
	Exec() (*model2.Hunter, *errors2.AppError)
}

func NewAttackHunterUseCaseImpl(hunterId model2.HunterId, monsterId model2.MonsterId,
	hunterRepository repository2.HunterRepository,
	monsterRepository repository2.MonsterRepository) attackHunterUseCase {

	return attackHunterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl attackHunterUseCaseImpl) Exec() (*model2.Hunter, *errors2.AppError) {
	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	monsterAttackDamage := service2.CalculateAttackHunterDamage(monster, hunter)
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
