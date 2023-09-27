package actionrpg2d

import "github.com/LuigiVanacore/ebiten_extended"

type ItemType uint

type Item struct {

	name       string
	itemType   ItemType
	text       Text
	font       Font
	textOffset ebiten_extended.Vector2D
}


func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetItemType() ItemType {
	return item.itemType
}

func (item *Item) GetText() Text {
	return item.text
}

func (item *Item) GetFont() Font {
	return item.font
}

func (item *Item) GetTextOffset() ebiten_extended.Vector2D {
	return item.textOffset
}

func (item *Item) SetName(name string) *Item {
	item.name = name
	return item
}

func (item *Item) SetItemType(itemType ItemType) *Item {
	item.itemType = itemType
	return item
}

func (item *Item) SetText(text Text) *Item {
	item.text = text
	return item
}

func (item *Item) SetFont(font Font) *Item {
	item.font = font
	return item
}

func (item *Item) SetTextOffset(textOffset ebiten_extended.Vector2D) *Item {
	item.textOffset = textOffset
	return item
}