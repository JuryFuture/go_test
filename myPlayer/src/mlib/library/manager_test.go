package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager faild...")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManager faild... not empty")
	}

	m0 := &MusicEntry{"1", "Mw Heart Wil Go On", "Celion Dion", "http://qdox.me/2450123", "MP3"}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() faild...")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() faild...")
	}
	if m.Id != m0.Id || m.Name != m0.Name || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() faild, Found item mismatch.")
	}

	if m1,err := mm.Get(0); err !=nil || m1 == nil {
		t.Error("MusicManager.Get() faild...", err)
	}

	m2 := mm.Remove(0)
	if m2 == nil {
		t.Error("MusicManager.Remove() faild...")
	}
}
