package decimal

import "testing"

func TestClamp_QP_QTGO_12_1(t *testing.T) {
	// [QP-QTGO-12-1] returns the low bound when the value is below the range
	d := NewFromInt(0)
	got := d.Clamp(NewFromInt(1), NewFromInt(3))
	if !got.Equal(NewFromInt(1)) {
		t.Errorf("Clamp(0, 1, 3) = %s, want 1", got)
	}
}

func TestClamp_QP_QTGO_12_2(t *testing.T) {
	// [QP-QTGO-12-2] returns the high bound when the value is above the range
	d := NewFromInt(7)
	got := d.Clamp(NewFromInt(1), NewFromInt(3))
	if !got.Equal(NewFromInt(3)) {
		t.Errorf("Clamp(7, 1, 3) = %s, want 3", got)
	}
}

func TestClamp_QP_QTGO_12_3(t *testing.T) {
	// [QP-QTGO-12-3] returns the value unchanged when it already lies within the range
	d := NewFromInt(2)
	got := d.Clamp(NewFromInt(1), NewFromInt(3))
	if !got.Equal(NewFromInt(2)) {
		t.Errorf("Clamp(2, 1, 3) = %s, want 2", got)
	}
}

func TestClamp_QP_QTGO_12_4(t *testing.T) {
	// [QP-QTGO-12-4] swaps reversed bounds and clamps a value below both to the smaller bound
	d := NewFromInt(0)
	got := d.Clamp(NewFromInt(3), NewFromInt(1))
	if !got.Equal(NewFromInt(1)) {
		t.Errorf("Clamp(0, 3, 1) = %s, want 1", got)
	}
}

func TestClamp_QP_QTGO_12_5(t *testing.T) {
	// [QP-QTGO-12-5] swaps reversed bounds and clamps a value above both to the larger bound
	d := NewFromInt(5)
	got := d.Clamp(NewFromInt(3), NewFromInt(1))
	if !got.Equal(NewFromInt(3)) {
		t.Errorf("Clamp(5, 3, 1) = %s, want 3", got)
	}
}
