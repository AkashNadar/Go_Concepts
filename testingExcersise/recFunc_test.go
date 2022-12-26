package recFunc

import "testing"

func TestHealth(t *testing.T) {
	player1 := AddPlayer(70, 100, 50, 100)

	player1.damage(20)
	if player1.currHealth != 50 {
		t.Errorf("after the damage the value should be 50 but was: %v", player1.currHealth)
	}

	player1.damage(-50)
	if player1.currHealth != 50 {
		t.Errorf("if minus damage is given then the value should be the same")
	}

	player1.damage(200)
	if player1.currHealth != 0 {
		t.Errorf("if damage is given above the max health should be 0")
	}

	player1.recovery(200)
	if player1.currHealth != 0 {
		t.Errorf("recovery limit should not be greater than max health")
	}

	player1.recovery(85)
	if player1.currHealth != 85 {
		t.Errorf("recory should be added as the limit is below 100")
	}

	player1.recovery(90)
	if player1.currHealth != 100 {
		t.Errorf("recovery should be full")
	}
}
