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
		return color.RGBA{65, 65, 65, 255}
	case theme.ColorNameHover:
		return color.RGBA{70, 70, 70, 255}
	case theme.ColorNameShadow:
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
	case theme.SizeNameScrollBar:
		return 2
	case theme.SizeNameScrollBarSmall:
		return 2
	default:
		return theme.DefaultTheme().Size(name)
	}
}
