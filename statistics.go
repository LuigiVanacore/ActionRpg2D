package actionrpg2d


type Statistics struct {
	healt uint
	maxHealth uint
	mana uint
	maxMana uint
	attack uint
	defense uint
	strenght uint 
	dexterity uint 
	stamina uint 
	speed uint
}


func (statistics *Statistics) GetHealt() uint {
	return statistics.healt
}

func (statistics *Statistics) GetMaxHealth() uint {
	return statistics.maxHealth
}

func (statistics *Statistics) GetMana() uint {
	return statistics.mana
}

func (statistics *Statistics) GetMaxMana() uint {
	return statistics.maxMana
}

func (statistics *Statistics) GetAttack() uint {
	return statistics.attack
}

func (statistics *Statistics) GetDefense() uint {
	return statistics.defense
}

func (statistics *Statistics) GetStrenght() uint {
	return statistics.strenght
}

func (statistics *Statistics) GetDexterity() uint {
	return statistics.dexterity
}

func (statistics *Statistics) GetStamina() uint {
	return statistics.stamina
}

func (statistics *Statistics) GetSpeed() uint {
	return statistics.speed
}

func (statistics *Statistics) SetHealt(healt uint) *Statistics {
	statistics.healt = healt
	return statistics
}

func (statistics *Statistics) SetMaxHealth(maxHealth uint) *Statistics {
	statistics.maxHealth = maxHealth
	return statistics
}

func (statistics *Statistics) SetMana(mana uint) *Statistics {
	statistics.mana = mana
	return statistics
}

func (statistics *Statistics) SetMaxMana(maxMana uint) *Statistics {
	statistics.maxMana = maxMana
	return statistics
}

func (statistics *Statistics) SetAttack(attack uint) *Statistics {
	statistics.attack = attack
	return statistics
}

func (statistics *Statistics) SetDefense(defense uint) *Statistics {
	statistics.defense = defense
	return statistics
}

func (statistics *Statistics) SetStrenght(strenght uint) *Statistics {
	statistics.strenght = strenght
	return statistics
}

func (statistics *Statistics) SetDexterity(dexterity uint) *Statistics {
	statistics.dexterity = dexterity
	return statistics
}

func (statistics *Statistics) SetStamina(stamina uint) *Statistics {
	statistics.stamina = stamina
	return statistics
}

func (statistics *Statistics) SetSpeed(speed uint) *Statistics {
	statistics.speed = speed
	return statistics
}