package calc

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{2, 1, 3},
		{3, 4, 7},
	}

	for _, c := range cases {
		actual := Add(c.a, c.b)
		if actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{2, 1, 1},
		{3, 4, -1},
	}

	for _, c := range cases {
		actual := Sub(c.a, c.b)
		if actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}

func TestMul(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{2, 1, 2},
		{3, 4, 12},
	}

	for _, c := range cases {
		actual := Mul(c.a, c.b)
		if actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}

func TestDiv(t *testing.T) {
	cases := []struct {
		a, b        int
		expectedVal float64
		expectedErr error
	}{
		{2, 1, 2, nil},
		{3, 4, 0.75, nil},
		{1, 0, 0, ErrDivisionByZero},
	}

	for _, c := range cases {
		actual, err := Div(c.a, c.b)
		if err != c.expectedErr {
			t.Errorf("예상 오류와 일치하지 않습니다. expected: %v, actual: %v", c.expectedErr, err)
		}
		if actual != c.expectedVal {
			t.Errorf("값이 예상과 다릅니다. expected: %f, actual: %f", c.expectedVal, actual)
		}
	}
}
