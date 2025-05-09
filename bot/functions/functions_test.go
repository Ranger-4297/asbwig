package functions

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestToInt64(t *testing.T) {
	t.Run("Convert from int", func(t *testing.T) {
		i := ToInt64(1)
		t.Logf("Converted int 1 to int64: %d", i)
		assert.Equal(t, int64(1), i, "Expected int conversion to int64")
	})

	t.Run("Convert from string", func(t *testing.T) {
		i := ToInt64("2")
		t.Logf("Converted string '2' to int64 (expected 2): %d", i)
		assert.Equal(t, int64(2), i, "Expected string '2' conversion to int64")
	})

	t.Run("Convert from float64", func(t *testing.T) {
		i := ToInt64(3.0)
		t.Logf("Converted float64 3.0 to int64 (expected 3): %d", i)
		assert.Equal(t, int64(3), i, "Expected float64 3.0 conversion to int64")
	})

	t.Run("Convert from string float", func(t *testing.T) {
		i := ToInt64("4.0")
		t.Logf("Converted string '4.0' to int64 (expected 4): %d", i)
		assert.Equal(t, int64(4), i, "Expected string '4.0' conversion to int64")
	})

	t.Run("Convert from float64 with decimal", func(t *testing.T) {
		i := ToInt64(5.1)
		t.Logf("Converted float64 5.1 to int64 (expected 0): %d", i)
		assert.Equal(t, int64(0), i, "Expected float64 5.1 to be invalid and return 0")
	})

	t.Run("Convert from string with float decimal", func(t *testing.T) {
		i := ToInt64("6.1")
		t.Logf("Converted string '6.1' to int64 (expected 0): %d", i)
		assert.Equal(t, int64(0), i, "Expected string '6.1' to be invalid and return 0")
	})

	t.Run("Invalid input - Unsupported type", func(t *testing.T) {
		i := ToInt64([]int{1, 2, 3})  // Invalid type (slice), should return 0
		t.Logf("Converted unsupported type []int{1,2,3} to int64 (expected 0): %d", i)
		assert.Equal(t, int64(0), i, "Expected 0 for unsupported types like a slice")
	})
}


func TestIsRoleHigher(t *testing.T) {
	tests := []struct {
		name     string
		higher   *discordgo.Role
		lower    *discordgo.Role
		expected bool
	}{
		{
			name: "Higher position",
			higher: &discordgo.Role{
				ID:       "1",
				Position: 10,
			},
			lower: &discordgo.Role{
				ID:       "2",
				Position: 5,
			},
			expected: true,
		},
		{
			name: "Lower position",
			higher: &discordgo.Role{
				ID:       "1",
				Position: 5,
			},
			lower: &discordgo.Role{
				ID:       "2",
				Position: 10,
			},
			expected: false,
		},
		{
			name: "Same position, lower ID",
			higher: &discordgo.Role{
				ID:       "1",
				Position: 5,
			},
			lower: &discordgo.Role{
				ID:       "2",
				Position: 5,
			},
			expected: true, // Lower ID considered "higher"
		},
		{
			name: "Same position, same ID",
			higher: &discordgo.Role{
				ID:       "1",
				Position: 5,
			},
			lower: &discordgo.Role{
				ID:       "1",
				Position: 5,
			},
			expected: false, // Same role
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsRoleHigher(tt.higher, tt.lower)
			assert.Equal(t, tt.expected, result)
		})
	}
}