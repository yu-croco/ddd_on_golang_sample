package model

import (
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type Hunter struct {
	Id              HunterId           `json:"hunterId"`
	Name            HunterName         `json:"hunterName"`
	Life            HunterLife         `json:"life"`
	DefencePower    HunterDefencePower `json:"defencePower"`
	OffensePower    HunterOffensePower `json:"offensePower"`
	HuntedMaterials HuntedMonsterMaterials
	AttackDamage    HunterAttackDamage `json:"attackDamage"`
}

// 完全コンストラクタのための初期化処理サンプル
func NewHunterId(id int) (*HunterId, *errors2.AppError) {
	if id <= 0 {
		err := errors2.NewAppError("HunterIdは1以上の値にしてください")
		return nil, &err
	}

	hunterId := HunterId(id)
	return &hunterId, nil
}

type HunterId int
type HunterName string
type HunterLife int

func (life HunterLife) Minus(damage MonsterOffensePower) HunterLife {
	return HunterLife(int(life) - int(damage))
}

type HunterDefencePower int

func (defence HunterDefencePower) BiggerOrSameThan(offense MonsterOffensePower) bool {
	return int(defence) >= int(offense)
}

type HunterOffensePower int

func (offense HunterOffensePower) Twice() HunterOffensePower {
	return HunterOffensePower(int(offense) * 2)
}
func (offense HunterOffensePower) Minus(defence MonsterDefencePower) HunterOffensePower {
	return HunterOffensePower(int(offense) - int(defence))
}

type HunterAttackDamage int

type Hunters []Hunter

func (hunter *Hunter) Attack(monster *Monster, damage HunterOffensePower) (*Monster, *errors2.AppError) {
	return monster.AttackedBy(damage)
}

func (hunter *Hunter) AttackedBy(givenDamage MonsterOffensePower) (*Hunter, *errors2.AppError) {
	var err errors2.AppError
	diff := hunter.Life.Minus(givenDamage)

	if hunter.Life == 0 {
		err = errors2.NewAppError("ハンターは既に倒れています")
		return nil, &err
	}

	if diff >= 0 {
		hunter.Life = diff
	} else {
		hunter.Life = HunterLife(0)
	}

	return hunter, &err
}

func (hunter *Hunter) GetMonsterMaterial(monster *Monster) (*MonsterMaterial, *errors2.AppError) {
	result, err := monster.TakenMaterial()

	return result, &err
}
