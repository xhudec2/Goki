package window

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Theme struct{}

func (m Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameButton:
		return color.RGBA{85, 85, 85, 255}
	case theme.ColorNameHover:
		return color.RGBA{105, 105, 105, 255}
	case theme.ColorNameScrollBar, theme.ColorNameShadow:
		return color.RGBA{0, 0, 0, 0}
	case theme.ColorNameForeground:
		return color.White
	default:
		return color.RGBA{55, 55, 55, 255}
	}
}

func (m Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m Theme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 10
	case theme.SizeNameScrollBar, theme.SizeNameScrollBarSmall:
		return 0
	default:
		return theme.DefaultTheme().Size(name)
	}
}
