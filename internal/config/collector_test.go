package config

import "testing"

func TestInit(t *testing.T) {
	t.Run("Should return empty struct pointer", func(t *testing.T) {
		v := Init()
		if v.DomainName != "" && v.ModulePath != "" {
			t.Errorf("Want %v, Got %v", "", v.DomainName)
		}
	})
}
