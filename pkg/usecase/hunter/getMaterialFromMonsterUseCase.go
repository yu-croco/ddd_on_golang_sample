package hunter

import (
	model2 "yu-croco/ddd_on_golang/pkg/domain/model"
	repository2 "yu-croco/ddd_on_golang/pkg/domain/repository"
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type getMaterialFromMonsterUseCaseImpl struct {
	HunterId          model2.HunterId
	MonsterId         model2.MonsterId
	HunterRepository  repository2.HunterRepository
	MonsterRepository repository2.MonsterRepository
}
type getMaterialFromMonsterUseCase interface {
	Exec() (*model2.MonsterMaterial, *errors2.AppError)
}

func NewGetMaterialFromMonsterUseCaseImpl(hunterId model2.HunterId, monsterId model2.MonsterId,
	hunterRepository repository2.HunterRepository,
	monsterRepository repository2.MonsterRepository) getMaterialFromMonsterUseCase {

	return getMaterialFromMonsterUseCaseImpl{
		HunterId:          hunterId,
		MonsterId:         monsterId,
		HunterRepository:  hunterRepository,
		MonsterRepository: monsterRepository,
	}
}

func (impl getMaterialFromMonsterUseCaseImpl) Exec() (*model2.MonsterMaterial, *errors2.AppError) {

	hunter, hunterFindErr := impl.HunterRepository.FindById(impl.HunterId)
	if hunterFindErr.HasErrors() {
		return nil, hunterFindErr
	}

	monster, monsterFindErr := impl.MonsterRepository.FindById(impl.MonsterId)
	if monsterFindErr.HasErrors() {
		return nil, monsterFindErr
	}

	takenMaterial, takeErr := hunter.GetMonsterMaterial(monster)
	if takeErr.HasErrors() {
		return nil, takeErr
	}

	impl.HunterRepository.AddMonsterMaterial(hunter, takenMaterial)

	return takenMaterial, nil
}
