package nifitui

import "github.com/rivo/tview"

type helpFlex tview.Flex

func (hf *helpFlex) clear() *helpFlex {
	(*tview.Flex)(hf).Clear()
	return hf
}

func (hf *helpFlex) addHelpText(label string) *helpFlex {
	(*tview.Flex)(hf).AddItem(tview.NewTextView().SetText(label), len(label)+1, 0, false)
	return hf
}
