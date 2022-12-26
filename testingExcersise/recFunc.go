package recFunc

type Player struct {
	currHealth, maxHealth int
	currEnergy, maxEnergy int
}

func AddPlayer(currHealth, maxHealth, currEnergy, maxEnergy int) Player {
	return Player{
		currEnergy: currEnergy,
		maxEnergy:  maxEnergy,
		currHealth: currHealth,
		maxHealth:  maxHealth,
	}
}

func (p *Player) damage(damageVal int) int {
	if damageVal <= 0 {
		return p.currHealth
	}
	if damageVal > p.currHealth {
		p.currHealth = 0
		return 0
	}
	p.currHealth -= damageVal
	return p.currHealth
}

func (p *Player) recovery(recoveryVal int) int {
	if recoveryVal > p.maxHealth || recoveryVal < 0 {
		return p.currHealth
	}
	if p.currHealth+recoveryVal > p.maxHealth {
		p.currHealth = p.maxHealth
		return p.currHealth
	}
	p.currHealth += recoveryVal
	return p.currHealth
}
