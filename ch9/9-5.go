package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

// 方法一
var mu sync.RWMutex

func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icons := icons[name]
		mu.RUnlock()
		return icons
	}
	mu.RUnlock()

	mu.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

// 方法二
var loadIconsOnce sync.Once

func Icon1(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
func loadIcons() {
	icons = map[string]image.Image{
		//"spades.png":   loadIcon("spades.png"),
		//"hearts.png":   loadIcon("hearts.png"),
		//"diamonds.png": loadIcon("diamonds.png"),
		//"clubs.png":    loadIcon("clubs.png"),
	}
}
