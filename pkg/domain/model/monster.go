package model

import (
	errors2 "yu-croco/ddd_on_golang/pkg/errors"
)

type Monster struct {
	Id           MonsterId           `json:"monsterId"`
	Name         MonsterName         `json:"monsterName"`
	Life         MonsterLife         `json:"life"`
	DefencePower MonsterDefencePower `json:"defencePower"`
	OffensePower MonsterOffensePower `json:"offensePower"`
	Materials    MonsterMaterials
	AttackDamage MonsterAttackDamage `json:"attackDamage"`
}

// 完全コンストラクタのための初期化処理サンプル
func NewMonsterId(id int) (*MonsterId, *errors2.AppError) {
	if id <= 0 {
		err := errors2.NewAppError("MonsterIdは1以上の値にしてください")
		return nil, &err
	}

	hunterId := MonsterId(id)
	return &hunterId, nil
}

type MonsterId int
type MonsterName string
type MonsterLife int

func (life MonsterLife) Minus(damage HunterOffensePower) MonsterLife {
	return MonsterLife(int(life) - int(damage))
}

type MonsterDefencePower int

func (defence MonsterDefencePower) BiggerOrSameThan(offense HunterOffensePower) bool {
	return int(defence) >= int(offense)
}

type MonsterOffensePower int

func (offense MonsterOffensePower) Twice() MonsterOffensePower {
	return MonsterOffensePower(int(offense) * 2)
}
func (offense MonsterOffensePower) Minus(defence HunterDefencePower) MonsterOffensePower {
	return MonsterOffensePower(int(offense) - int(defence))
}

type MonsterAttackDamage int

type Monsters []Monster

func (monster *Monster) Attack(hunter *Hunter, damage MonsterOffensePower) (*Hunter, *errors2.AppError) {
	return hunter.AttackedBy(damage)
}

func (monster *Monster) AttackedBy(givenDamage HunterOffensePower) (*Monster, *errors2.AppError) {
	var err errors2.AppError
	diff := monster.Life.Minus(givenDamage)

	if monster.Life == 0 {
		err = errors2.NewAppError("このモンスターはすでに倒しています")
		return nil, &err
	}

	if diff >= 0 {
		monster.Life = diff
	} else {
		monster.Life = MonsterLife(0)
	}

	return monster, &err
}

func (monster *Monster) TakenMaterial() (*MonsterMaterial, errors2.AppError) {
	if monster.Life != 0 {
		return nil, errors2.NewAppError("このモンスターはまだ生きてるので剥ぎ取れません")
	} else {
		return &monster.Materials[0], errors2.AppError{}
	}
}
