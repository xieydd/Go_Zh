package manager

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed!")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty")
	}

	m0 := &Music{"1", "My Heart Will Go On", "Celion Dion", "http://qbox.me/24501234", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager Add() failed")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager Find() failed")
	}
	if m.Name != m0.Name || m.ID != m0.ID || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager Find() failed, Found item mismatch")
	}

	m, err := mm.Get(0)
	if err != nil {
		t.Error("MusicManager Get() failed")
	}

	m = mm.Delete(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager Delete() failed")
	}
}
