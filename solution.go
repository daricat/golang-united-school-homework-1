package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	stringWithEmoji := emoji.Sprint("Hello :world_map:!")
	return stringWithEmoji
}
